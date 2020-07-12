package main

import (
	"atto/internal/app/daemon"
	"log"
	"time"
)

// Inherit from build variables.
var currentVersion = "UNKNOWN" // nolint:gochecknoglobals // used for version injection on build

func main() {
	time.Local = time.UTC

	d := daemon.New()

	d.Logger.
		With("Version", currentVersion).
		Info("Atto Daemon started!")

	log.Fatal(d.Start())
}
