package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/rpc"
)

type BatcherDriver interface {
	StartBatchSubmitting() error
	StopBatchSubmitting(ctx context.Context) error
}

type ReorgSender func(context.Context, uint64) error

type adminAPI struct {
	*rpc.CommonAdminAPI
	b           BatcherDriver
	reorgSender ReorgSender
}

func NewAdminAPI(dr BatcherDriver, m metrics.RPCMetricer, log log.Logger, sender ReorgSender) *adminAPI {
	return &adminAPI{
		CommonAdminAPI: rpc.NewCommonAdminAPI(m, log),
		b:              dr,
		reorgSender:    sender,
	}
}

func (a *adminAPI) SendReorgBatch(ctx context.Context, l2BlockNumber uint64) error {
	return a.reorgSender(ctx, l2BlockNumber)
}

func GetAdminAPI(api *adminAPI) gethrpc.API {
	return gethrpc.API{
		Namespace: "admin",
		Service:   api,
	}
}

func (a *adminAPI) StartBatcher(_ context.Context) error {
	return a.b.StartBatchSubmitting()
}

func (a *adminAPI) StopBatcher(ctx context.Context) error {
	return a.b.StopBatchSubmitting(ctx)
}
