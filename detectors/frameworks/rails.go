package frameworks

import (
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Rails struct {
	*detectors.Detector
}

func (detector *Rails) Detect() bool {
	if cookie, ok := detector.Info.Cookies["_session_id"]; ok {
		regex := regexp.MustCompile(`^[0-9a-f]+`)
		result := regex.FindStringSubmatch(cookie.Value)
		if len(result) > 0 && (detector.Info.Language.Name == "" || detector.Info.Language.Name == "Ruby") {
			detector.Info.Framework.Name = "Rails"
			if detector.Info.Language.Name == "" {
				detector.Info.Language.Version = ""
				detector.Info.Language.Name = "Ruby"
			}
			return true
		}
	}
	return false
}