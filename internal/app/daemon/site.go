package daemon

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type site struct {
	directoryPath string
	domain        string
	tls           bool
}

func handleSites(sites []site) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, s := range sites {
			if r.Host != s.domain {
				continue
			}

			fs := http.Dir(s.directoryPath)
			renderFile(w, r, fs)

			return
		}

		renderNotFound(w, "domain not found")
	}
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
		domain:        "dummy.site" + addr,
		tls:           true,
	}
	// ------------

	sites = append(sites, dummySite)

	return sites, nil
}
