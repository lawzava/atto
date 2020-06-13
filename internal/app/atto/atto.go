package atto

import "atto/pkg/logger"

type Atto struct {
	Logger *logger.Log
}

func New() *Atto {
	l := logger.New(logger.Info)

	return &Atto{l}
}
