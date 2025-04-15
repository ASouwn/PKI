package sharedrpctypes

import (
	"encoding/pem"
)

type RAServer interface {
	HandleCSR(csrPem *pem.Block, reply *pem.Block) error
}
