package servers

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Apache struct {
	*detectors.Detector
}

func (detector *Apache) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "server" {
			for _, value := range values {
				regex := regexp.MustCompile(`Apache`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 0 {
					detector.Info.Server.Name = "Apache"
					detector.Info.Server.Version = ""
					return true
				}
			}
		}
	}
	return false
}