package frameworks

import (
	"strings"
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Servlet struct {
	*detectors.Detector
}

func (detector *Servlet) Detect() bool {
	for key, values := range detector.Response.Header {
		if strings.ToLower(key) == "x-powered-by" {
			for _, value := range values {
				regex := regexp.MustCompile(`Servlet(\s+([\d.]+))?`)
				result := regex.FindStringSubmatch(value)
				if len(result) > 0 {
					detector.Info.Framework.Name = "Servlet"
					detector.Info.Framework.Version = ""
					if len(result) > 2 {
						detector.Info.Framework.Version = result[2]
					}
					if detector.Info.Language.Name != "Java" {
						detector.Info.Language.Version = ""
						detector.Info.Language.Name = "Java"
					}
					return true
				}
			}
		}
	}
	if _, ok := detector.Info.Cookies["JSESSIONID"]; ok {
		detector.Info.Framework.Name = "Servlet"
		if detector.Info.Language.Name != "Java" {
			detector.Info.Language.Version = ""
			detector.Info.Language.Name = "Java"
		}
		return true
	}
	return false

}