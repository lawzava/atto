package daemon

import "atto/pkg/logger"

type Daemon struct {
	Logger *logger.Log
}

func New() *Daemon {
	l := logger.New(logger.Info)

	return &Daemon{l}
}
