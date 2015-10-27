package frameworks

import (
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type FuelPHP struct {
	*detectors.Detector
}

func (detector *FuelPHP) Detect() bool {
	if _, ok := detector.Info.Cookies["fuelcid"]; ok {
		detector.Info.Framework.Name = "FuelPHP"
		if detector.Info.Language.Name != "PHP" {
			detector.Info.Language.Version = ""
			detector.Info.Language.Name = "PHP"
		}
		return true
	}
	return false
}