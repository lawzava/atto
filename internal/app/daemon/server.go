package daemon

import (
	"fmt"
	"net/http"
)

// TODO: Move to config.
const port = 9182

func (d *Daemon) Start() error {
	e := d.server

	e.Handle(http.MethodGet, "/*filepath", handleSites(d.sites))

	d.Logger.Infof("Server is running on %d port", port)

	return d.server.Run(fmt.Sprintf(":%d", port))
}
