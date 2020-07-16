package daemon

import (
	"atto/pkg/logger"
	"fmt"
)

type Daemon struct {
	Logger *logger.Log

	sites *[]site
}

func New(l *logger.Log) (*Daemon, error) {
	daemon := Daemon{Logger: l}

	sites, err := parseSites()
	if err != nil {
		return nil, fmt.Errorf("failed to parse sites: %w", err)
	}

	daemon.sites = &sites

	return &daemon, nil
}
