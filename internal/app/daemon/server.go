package daemon

import (
	"net/http"
)

// TODO: Move to config.
const addr = ":9182"

func (d *Daemon) Start() error {
	http.HandleFunc("/", handleSites(d.sites))

	d.Logger.Infof("Server is running on %s port", addr)

	return http.ListenAndServe(addr, nil)
}
