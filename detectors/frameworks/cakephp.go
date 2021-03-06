package frameworks

import (
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type CakePHP struct {
	*detectors.Detector
}

func (detector *CakePHP) Detect() bool {
	if _, ok := detector.Info.Cookies["CAKEPHP"]; ok {
		detector.Info.Framework.Name = "CakePHP"
		if detector.Info.Language.Name != "PHP" {
			detector.Info.Language.Version = ""
			detector.Info.Language.Name = "PHP"
		}
		return true
	}
	return false
}