package daemon

import (
	"fmt"
	"os"
	"path/filepath"
)

type site struct {
	absolutePath string
	domain       string
}

func parseSites() ([]site, error) {
	var sites []site

	// dummy value for testing
	absolutePath, err := filepath.Abs(os.Getenv("TEST_SITE_PATH"))
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	dummySite := site{
		absolutePath: absolutePath,
		domain:       "dummy.site",
	}
	// ------------

	sites = append(sites, dummySite)

	return sites, nil
}
