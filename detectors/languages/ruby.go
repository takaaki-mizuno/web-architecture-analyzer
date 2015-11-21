package languages

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Ruby struct {
	*detectors.Detector
}

func (detector *Ruby) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "x-powered-by" {
			for _, value := range values {
				regex := regexp.MustCompile(`Phusion\s+Passenger`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 0 {
					detector.Info.Language.Name = "Ruby"
					detector.Info.Language.Version = ""
					return true
				}
			}
		}
	}
	return false
}
