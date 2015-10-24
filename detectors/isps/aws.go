package isps

import (
	"regexp"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type AWS struct {
	*detectors.Detector
}

func (detector *AWS) Detect() bool {

	regions := map[string]string{
		"us-east-1": "US East (N. Virginia)",
		"us-west-1": "US West (N. California)",
		"us-west-2": "US West (Oregon)",
		"eu-west-1": "EU (Ireland)",
		"eu-central-1": "EU (Frankfurt)",
		"ap-southeast-1": "Asia Pacific (Singapore)",
		"ap-southeast-2": "Asia Pacific (Sydney)",
		"ap-northeast-1": "Asia Pacific (Tokyo)",
		"a-east-1": "South America (Sao Paulo)",
	}

	for _, host := range detector.Info.RealHost {
		regex := regexp.MustCompile(`([a-z0-9\-]+)\.compute\.amazonaws\.com`)
		result := regex.FindStringSubmatch(host)
		if len(result) > 1 {
			detector.Info.ISP.Name = "AWS"
			if region, ok := regions[result[1]]; ok {
				detector.Info.ISP.Region = region
			}
			return true
		}
	}
	return false

}