package daemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleStaticFiles(absolutePath string, fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *gin.Context) {
		file := c.Param("filepath")

		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open(file)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}

		_ = f.Close()

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
