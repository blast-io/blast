package eth

import (
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

func (s *Ethereum) NewBlastAPI() any {
	return ethapi.NewBlastAPI(s.blockchain, s.blastAuthPool, s.miner)
}
