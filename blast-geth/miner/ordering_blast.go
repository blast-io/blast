package miner

import (
	"container/heap"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
)

// txByTime implements both the sort and the heap interface, making it useful
// for all at once sorting as well as individually adding and removing elements.
type txByTime []*txWithMinerFee

func (s txByTime) Len() int { return len(s) }
func (s txByTime) Less(i, j int) bool {
	// If the prices are equal, use the time the transaction was first seen for
	// deterministic sorting
	// cmp := s[i].fees.Cmp(s[j].fees)
	// if cmp == 0 {
	return s[i].tx.Time.Before(s[j].tx.Time)
	// }
	// return cmp > 0
}
func (s txByTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s *txByTime) Push(x any) {
	*s = append(*s, x.(*txWithMinerFee))
}

func (s *txByTime) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*s = old[0 : n-1]
	return x
}

// transactionsByTimeAndNonce represents a set of transactions that can return
// transactions in a time received sorted order, while supporting removing
// entire batches of transactions for non-executable accounts.
type transactionsByTimeAndNonce struct {
	txs     map[common.Address][]*txpool.LazyTransaction // Per account nonce-sorted list of transactions
	heads   txByTime                                     // Next transaction for each unique account (time heap)
	signer  types.Signer                                 // Signer for the set of transactions
	baseFee *big.Int                                     // Current base fee
}

// newTransactionsByPriceAndNonce creates a transaction set that can retrieve
// price sorted transactions in a nonce-honouring way.
//
// Note, the input map is reowned so the caller should not interact any more with
// if after providing it to the constructor.
func newTransactionsByTimeAndNonce(signer types.Signer, txs map[common.Address][]*txpool.LazyTransaction, baseFee *big.Int) *transactionsByTimeAndNonce {
	// Initialize a received time based heap with the head transactions
	heads := make(txByTime, 0, len(txs))
	for from, accTxs := range txs {
		if accTxs[0] == nil {
			continue
		}
		wrapped, err := newTxWithMinerFee(accTxs[0], from, baseFee)
		if err != nil {
			delete(txs, from)
			continue
		}
		if wrapped == nil {
			continue
		}
		heads = append(heads, wrapped)
		txs[from] = accTxs[1:]
	}
	heap.Init(&heads)

	// Assemble and return the transaction set
	return &transactionsByTimeAndNonce{
		txs:     txs,
		heads:   heads,
		signer:  signer,
		baseFee: baseFee,
	}
}

// Peek returns the next transaction by time.
func (t *transactionsByTimeAndNonce) Peek() *txpool.LazyTransaction {
	if len(t.heads) == 0 {
		return nil
	}
	return t.heads[0].tx
}

// Shift replaces the current best head with the next one from the same account.
func (t *transactionsByTimeAndNonce) Shift() {
	acc := t.heads[0].from
	if txs, ok := t.txs[acc]; ok && len(txs) > 0 {
		if wrapped, err := newTxWithMinerFee(txs[0], acc, t.baseFee); err == nil {
			t.heads[0], t.txs[acc] = wrapped, txs[1:]
			heap.Fix(&t.heads, 0)
			return
		}
	}
	heap.Pop(&t.heads)
}

// Pop removes the best transaction, *not* replacing it with the next one from
// the same account. This should be used when a transaction cannot be executed
// and hence all subsequent ones should be discarded from the same account.
func (t *transactionsByTimeAndNonce) Pop() {
	heap.Pop(&t.heads)
}
