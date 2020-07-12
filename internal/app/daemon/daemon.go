package daemon

import (
	"atto/pkg/logger"
	"atto/pkg/server"
	"fmt"
)

type Daemon struct {
	Logger *logger.Log

	server *server.Server
	sites  []site
}

func New(l *logger.Log) (*Daemon, error) {
	daemon := Daemon{Logger: l}

	sites, err := parseSites()
	if err != nil {
		return nil, fmt.Errorf("failed to parse sites: %w", err)
	}

	daemon.sites = sites

	daemon.server = server.New(l)

	return &daemon, nil
}
