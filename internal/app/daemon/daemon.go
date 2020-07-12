package daemon

import (
	"atto/pkg/logger"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Daemon struct {
	Logger *logger.Log

	server *gin.Engine
	sites  []site
}

func New(l *logger.Log) (*Daemon, error) {
	daemon := Daemon{Logger: l}

	sites, err := parseSites()
	if err != nil {
		return nil, fmt.Errorf("failed to parse sites: %w", err)
	}

	daemon.sites = sites

	r := gin.New()

	// CORS
	if gin.Mode() != gin.ReleaseMode {
		defaultConfig := cors.DefaultConfig()
		defaultConfig.AllowAllOrigins = true
		defaultConfig.AllowHeaders = []string{"*"}
		r.Use(cors.New(defaultConfig))
	}

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.RedirectFixedPath = true

	daemon.server = r

	return &daemon, nil
}
