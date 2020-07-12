package server

import (
	"atto/pkg/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r   *gin.Engine
	log *logger.Log
}

func New(l *logger.Log) *Server {
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

	return &Server{r, l}
}

func (e Server) Engine() *gin.Engine {
	return e.r
}

const (
	maxHeaderBytes = 1 << 20

	// Move to shared config
	port         = 9182
	readTimeout  = 5 * time.Second
	writeTimeout = 5 * time.Second
)

func (e Server) Serve() error {
	serverAddr := fmt.Sprintf(":%d", port)

	srv := &http.Server{
		Addr:           serverAddr,
		Handler:        e.r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	e.log.Infof("server is running on %s port", serverAddr)

	return srv.ListenAndServe()
}
