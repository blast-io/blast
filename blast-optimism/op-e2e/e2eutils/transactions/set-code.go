package transactions

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

func CreateSetCodeTx(pk *ecdsa.PrivateKey, chainID uint64, to common.Address) *types.SetCodeTx {
	signed, _ := types.SignSetCode(pk, types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(chainID),
		Address: common.Address{0x05, 0x05, 0x05},
	},
	)

	setCodeTx := &types.SetCodeTx{
		ChainID:   uint256.NewInt(chainID),
		Nonce:     0,
		GasTipCap: uint256.NewInt(2200000000000),
		GasFeeCap: uint256.NewInt(5000000000000),
		Gas:       25000,
		To:        to,
		Value:     uint256.NewInt(99),
		Data:      make([]byte, 50),
		AuthList:  []types.SetCodeAuthorization{signed},
	}
	return setCodeTx
}
