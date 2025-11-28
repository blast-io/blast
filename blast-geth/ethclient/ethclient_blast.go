package ethclient

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

func (ec *Client) SubscribeFastReceipt(
	ctx context.Context, ch chan<- *types.Receipt,
) (ethereum.Subscription, error) {
	sub, err := ec.c.Subscribe(ctx, "blast", ch, "fastReceipt")
	if err != nil {
		// Defensively prefer returning nil interface explicitly on error-path, instead
		// of letting default golang behavior wrap it with non-nil interface that stores
		// nil concrete type value.
		return nil, err
	}
	return sub, nil
}

// BlastTransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (ec *Client) BlastTransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := ec.c.CallContext(ctx, &r, "blast_getTransactionReceipt", txHash)
	if err == nil && r == nil {
		return nil, ethereum.NotFound
	}

	return r, err
}

// BlastTransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (ec *Client) BlastSendTransaction(
	ctx context.Context, tx *types.Transaction,
) error {

	marshaled, err := tx.MarshalBinary()
	if err != nil {
		return err
	}

	return ec.c.CallContext(ctx, nil, "blast_sendRawTransaction", hexutil.Encode(marshaled))
}
