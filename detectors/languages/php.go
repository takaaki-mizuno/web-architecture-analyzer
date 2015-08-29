package languages

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type PHP struct {
	*detectors.Detector
}

func (detector *PHP) Message() string {
	return "PHP: Version" + detector.Info.Language.Version
}

func (detector *PHP) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "x-powered-by" {
			for _, value := range values {
				regex := regexp.MustCompile(`PHP/([\d.]+)`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 1 {
					detector.Info.Language.Name = "PHP"
					detector.Info.Language.Version = result[1]
					return true
				}
			}
		}
	}
	return false
}