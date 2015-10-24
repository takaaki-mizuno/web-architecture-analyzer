package packages

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Drupal struct {
	*detectors.Detector
}

func (detector *Drupal) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "x-generator" {
			for _, value := range values {
				regex := regexp.MustCompile(`Drupal(\s+([\d.]+))?`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 0 {
					detector.Info.Package.Name = "Drupal"
					detector.Info.Package.Version = ""
					if len(result) > 2 {
						detector.Info.Package.Version = result[2]
					}
					return true
				}
			}
		}
	}
	return false
}