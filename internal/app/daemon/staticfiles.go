package daemon

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleStaticFiles(absolutePath string, sites []site) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, site := range sites {
			if getDomainName(c.Request.Host) != site.domain {
				continue
			}

			fs := gin.Dir(site.directoryPath, false)
			fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

			file := c.Param("filepath")

			// Check if file exists and/or if we have permission to access it
			f, err := fs.Open(file)
			if err != nil {
				renderNotFound(c, "file not found")
				return
			}

			_ = f.Close()

			fileServer.ServeHTTP(c.Writer, c.Request)

			return
		}

		renderNotFound(c, "domain not found")
	}
}

func renderNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"status":  http.StatusText(http.StatusNotFound),
		"message": message,
	})
}

func getDomainName(host string) string {
	return strings.Split(host, ":")[0]
}
