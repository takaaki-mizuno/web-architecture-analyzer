package frameworks

import (
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Laravel struct {
*detectors.Detector
}

func (detector *Laravel) Detect() bool {
	if _, ok := detector.Info.Cookies["laravel_session"]; ok {
		detector.Info.Framework.Name = "Lalavel"
		if detector.Info.Language.Name != "PHP" {
			detector.Info.Language.Version = ""
			detector.Info.Language.Name = "PHP"
		}
		return true
	}
	return false
}