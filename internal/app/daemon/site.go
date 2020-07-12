package daemon

import (
	"fmt"
	"path/filepath"
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
