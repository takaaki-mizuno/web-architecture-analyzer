package detectors

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors/info"
)

type DetectorInterface interface{
	Detect() bool
	Message() string
}

type Detector struct {
	Response *http.Response
	Document *goquery.Document
	Info *info.Info
}
