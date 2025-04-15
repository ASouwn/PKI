package sharedrpctypes

import (
	"encoding/pem"
)

type CAServer interface {
	HandleCSR(csr *pem.Block, reply *pem.Block) error
}
