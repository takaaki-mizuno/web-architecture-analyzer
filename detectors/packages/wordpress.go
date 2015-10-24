package packages

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
	"net/url"
	"strings"
	"github.com/takaaki-mizuno/web-architecture-analyzer/utilities"
)

type WordPress struct {
	*detectors.Detector
}

func (detector *WordPress) Detect() bool {
	found := false
	count := 0
	if !found {
		detector.Document.Find("img").Each(func(i int, s *goquery.Selection) {
			link, err_b := s.Attr("src")
			if !err_b {
				return
			}
			u, err := url.Parse(link)
			if err != nil {
				return
			}
			pathFragments := strings.Split(u.Path, "/")
			if utilities.StringInSlice("wp-content", pathFragments) {
				count++
				if count > 10 {
					found = true
					return
				}
			}
		})
	}
	if found {
		detector.Info.Package.Name = "WordPress"
		if detector.Info.Language.Name != "PHP" {
			detector.Info.Language.Version = ""
			detector.Info.Language.Name = "PHP"
		}
	}
	return found
}