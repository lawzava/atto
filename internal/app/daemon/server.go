package daemon

import (
	"atto/pkg/handling"
	"net/http"
)

func (d *Daemon) Start() error {
	e := d.server.Engine()
	//
	//for _, site := range d.sites {
	//	http.Handle("/"+site.domain, http.FileServer(http.Dir(site.absolutePath)))
	//	d.Logger.
	//		With("path", "/"+site.domain).
	//		With("directory", site.absolutePath).
	//		Info("Serving site " + site.domain)
	//}

	e.Handle(http.MethodGet, "/status", handling.HandlerFunc(statusHandler))

	return d.server.Serve()
}
