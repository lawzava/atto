package daemon

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type site struct {
	directoryPath string
	domain        string
	tls           bool
}

func parseSites() ([]site, error) {
	var sites []site

	// dummy dir for testing
	directoryPath, err := filepath.Abs("dummysite")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	dummySite := site{
		directoryPath: directoryPath,
		domain:        "dummy.site",
		tls:           true,
	}
	// ------------

	sites = append(sites, dummySite)

	return sites, nil
}

func handleSites(sites []site) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, s := range sites {
			if getDomainName(c.Request.Host) != s.domain {
				continue
			}

			fs := gin.Dir(s.directoryPath, false)
			renderFile(c, fs)

			return
		}

		renderNotFound(c, "domain not found")
	}
}

func getDomainName(host string) string {
	return strings.Split(host, ":")[0]
}
