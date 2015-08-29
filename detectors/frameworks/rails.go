package frameworks

import (
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type Rails struct {
	*detectors.Detector
}

func (detector *Default) Rails() bool {
	return false
}