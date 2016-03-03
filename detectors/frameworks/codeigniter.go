package frameworks

import (
    "github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Codeigniter struct {
    *detectors.Detector
}

func (detector *Codeigniter) Detect() bool {
    if _, ok := detector.Info.Cookies["ci_session"]; ok {
        detector.Info.Framework.Name = "Codeigniter"
        if detector.Info.Language.Name != "PHP" {
            detector.Info.Language.Version = ""
            detector.Info.Language.Name = "PHP"
        }
        return true
    }
    return false
}