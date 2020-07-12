package daemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func renderFile(c *gin.Context, fs http.FileSystem) {
	file := c.Param("filepath")

	// Check if file exists and/or if we have permission to access it
	f, err := fs.Open(file)
	if err != nil {
		renderNotFound(c, "file not found")
		return
	}

	_ = f.Close()

	fileServer := http.StripPrefix("/", http.FileServer(fs))
	fileServer.ServeHTTP(c.Writer, c.Request)
}

func renderNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"status":  http.StatusText(http.StatusNotFound),
		"message": message,
	})
}
