package daemon

import (
	"net/http"
)

// TODO: Move to config.
const addr = ":9182"

func (d *Daemon) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleSites(d.sites))

	d.Logger.Infof("Server is running on %s port", addr)

	// Serve TLS here as well
	return http.ListenAndServe(addr, mux)
}
