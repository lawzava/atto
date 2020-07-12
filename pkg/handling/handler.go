package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerFunc(f http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}
