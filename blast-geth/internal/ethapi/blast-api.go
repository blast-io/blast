package ethapi

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type FastMiner interface {
	SubscribeFastReceipts(ch chan<- *types.Receipt) event.Subscription
}

type BlastAPI struct {
	chain *core.BlockChain
	pool  *legacypool.LegacyPool
	miner FastMiner
}

func NewBlastAPI(b *core.BlockChain, authedPool *legacypool.LegacyPool, m FastMiner) *BlastAPI {
	return &BlastAPI{chain: b, pool: authedPool, miner: m}
}

func (b *BlastAPI) SendRawTransaction(
	ctx context.Context, input hexutil.Bytes,
) (common.Hash, error) {
	tx := new(types.Transaction)
	if err := tx.UnmarshalBinary(input); err != nil {
		return common.Hash{}, err
	}
	tx.AuthOnly.Store(true)
	return tx.Hash(), b.pool.Add([]*types.Transaction{tx}, true, false)[0]
}

func (b *BlastAPI) receiptToMap(receipt *types.Receipt, hsh common.Hash) map[string]any {
	fields := map[string]interface{}{
		// "blockHash":         blockHash,
		// "blockNumber":       hexutil.Uint64(blockNumber),
		"transactionHash": hsh,
		// "transactionIndex":  hexutil.Uint64(txIndex),
		// "from":              from,
		// "to":                tx.To(),
		"gasUsed":           hexutil.Uint64(receipt.GasUsed),
		"cumulativeGasUsed": hexutil.Uint64(receipt.CumulativeGasUsed),
		// "contractAddress":   nil,
		"logs":      receipt.Logs,
		"logsBloom": receipt.Bloom,
		// "type":              hexutil.Uint(tx.Type()),
		"effectiveGasPrice": (*hexutil.Big)(receipt.EffectiveGasPrice),
	}

	// Assign receipt status or post state.
	if len(receipt.PostState) > 0 {
		fields["root"] = hexutil.Bytes(receipt.PostState)
	} else {
		fields["status"] = hexutil.Uint(receipt.Status)
	}
	// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
	if receipt.ContractAddress != (common.Address{}) {
		fields["contractAddress"] = receipt.ContractAddress
	}
	if receipt.Logs == nil {
		fields["logs"] = []*types.Log{}
	}

	return fields

}

func (b *BlastAPI) GetTransactionReceipt(ctx context.Context, hsh common.Hash) (map[string]any, error) {
	receipt := b.chain.FastReceiptCache(hsh)
	if receipt == nil {
		return nil, ethereum.NotFound
	}

	return b.receiptToMap(receipt, hsh), nil
}

func (a *BlastAPI) FastReceipt(ctx context.Context) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	rpcSub := notifier.CreateSubscription()

	go func() {
		incoming := make(chan *types.Receipt)
		sub := a.miner.SubscribeFastReceipts(incoming)
		defer sub.Unsubscribe()
		log.Trace("created fast receipt subscription to miner")

		for {
			select {
			case arb := <-incoming:
				log.Trace("received incoming so now sending to notify", "rpc-sub-id", rpcSub.ID, "arb", arb)
				notifier.Notify(rpcSub.ID, arb)
			case err := <-rpcSub.Err():
				if err != nil {
					log.Trace("problem with fast receipt subscription", "err", err)
				}
				return
			}
		}
	}()

	return rpcSub, nil
}
