package miner

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
)

func (miner *Miner) SubscribeFastReceipts(ch chan<- *types.Receipt) event.Subscription {
	log.Trace("new subscription to fast receipt feed")
	return miner.worker.fastReceiptFeed.Subscribe(ch)
}
