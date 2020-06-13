package main

import (
	"atto/internal/app/atto"
	"time"
)

const currentVersion = "v0.1"

func main() {
	time.Local = time.UTC

	a := atto.New()

	a.Logger.
		With("Version", currentVersion).
		Info("Atto started!")
}
