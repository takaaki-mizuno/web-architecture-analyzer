package frameworks

import (
	"github.com/takaaki-mizuno/web-architecture-analyzer/detectors"
)

type CakePHP struct {
	*detectors.Detector
}

func (detector *Default) CakePHP() bool {
	return false
}