package daemon

import (
	"fmt"
	"net/http"
)

const (
	// Move to config config
	port = 9182
)

func (d *Daemon) Start() error {
	e := d.server

	e.Handle(http.MethodGet, "/*filepath", handleStaticFiles("/", d.sites))

	d.Logger.Infof("Server is running on %d port", port)

	return d.server.Run(fmt.Sprintf(":%d", port))
}
