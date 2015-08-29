package languages

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Python struct {
	*detectors.Detector
}

func (detector *Python) Message() string {
	return "PYTHON: Version" + detector.Info.Language.Version
}

func (detector *Python) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "server" {
			for _, value := range values {
				regex := regexp.MustCompile(`Python/([\d.]+)`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 1 {
					detector.Info.Language.Name = "Python"
					detector.Info.Language.Version = result[1]
					return true
				}
			}
		}
	}
	return false
}