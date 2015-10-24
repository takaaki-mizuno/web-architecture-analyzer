package frameworks

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Play struct {
	*detectors.Detector
}

func (detector *Play) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "server" {
			for _, value := range values {
				regex := regexp.MustCompile(`Play(/([\d.]+))?`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 0 {
					detector.Info.Framework.Name = "Play"
					detector.Info.Framework.Version = ""
					if len(result) > 2 {
						detector.Info.Framework.Version = result[2]
					}
					return true
				}
			}
		}
	}
	if _, ok := detector.Info.Cookies["PLAY_SESSION"]; ok {
		detector.Info.Framework.Name = "Play"
		return true
	}
	return false
}