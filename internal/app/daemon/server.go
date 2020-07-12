package daemon

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

const (
	// Move to config config
	port = 9182
)

func (d *Daemon) Start() error {
	e := d.server

	for _, site := range d.sites {
		var (
			relativePath = "/" + site.domain
			absolutePath = joinPaths(e.BasePath(), relativePath)
			fs           = gin.Dir(site.directoryPath, false)
		)

		e.Handle(
			http.MethodGet,
			path.Join(relativePath, "/*filepath"),
			handleStaticFiles(absolutePath, fs),
		)
	}

	e.Handle(http.MethodGet, "/status", statusHandler)

	d.Logger.Infof("Server is running on %d port", port)

	return d.server.Run(fmt.Sprintf(":%d", port))
}

func joinPaths(basePath, relativePath string) string {
	if relativePath == "" {
		return basePath
	}

	finalPath := path.Join(basePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'

	if appendSlash {
		return finalPath + "/"
	}

	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}

	return str[len(str)-1]
}
