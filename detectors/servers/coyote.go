package servers

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Coyote struct {
	*detectors.Detector
}

func (detector *Coyote) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "server" {
			for _, value := range values {
				regex := regexp.MustCompile(`Apache-Coyote(/([\d.]+))?`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 0 {
					detector.Info.Server.Name = "Coyote"
					detector.Info.Server.Version = ""
					if len(result) > 2 {
						detector.Info.Server.Version = result[2]
					}
					return true
				}
			}
		}
	}
	return false
}