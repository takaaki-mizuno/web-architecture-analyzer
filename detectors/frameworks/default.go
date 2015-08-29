package frameworks

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Default struct {
	*detectors.Detector
}

func (detector *Default) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "x-framework" {
			for _, value := range values {
				regex := regexp.MustCompile(`^([^/]+)/([\d.]+)`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 2 {
					detector.Info.Framework.Name = result[1]
					detector.Info.Framework.Version = result[2]
					return true
				} else {
					detector.Info.Framework.Name = value
					detector.Info.Framework.Version = ""
				}
			}
		}
	}
	return false
}