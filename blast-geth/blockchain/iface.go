package blockchain

import (
	"encoding/gob"
	"math/big"
	"net/rpc"
	"time"

	"github.com/hashicorp/go-plugin"
)

func init() {
	gob.Register(plugin.BasicError{})
}

type NewChainStartingArgs struct {
	SerializedGenesis      []byte
	ExtraAllocs            map[string]*big.Int
	AssumeMainnet          bool
	IncludeCatalystAPI     bool
	JWTFilePath            string
	AuthPort               int
	WSPort                 int
	UseDatadir             string
	WhenActivateCancun     *uint64
	WhenActivatePrague     *uint64
	WhenActivateOsaka      *uint64
	WhenActivateBPO1       *uint64
	WhenActivateBPO2       *uint64
	Faucet                 string
	CatalystAuthEnabled    bool
	MinerRecommit          time.Duration
	MinerNewPayloadTimeout time.Duration
}

type Chain interface {
	NewChain(*NewChainStartingArgs) NewChainOrError
	InitExtraConfigs([]byte) error
	StartBlock(timeDelta uint64) error
	WSEndpoint() (string, error)
	AuthEndpoint() (string, error)
	Close() error
	IncludeTxByHash(string) error
	EndBlock() NewBlockOrError
	SetFeeRecipient(string) error
}

type NewBlockOrError struct {
	SerializedBlock []byte
	Err             *plugin.BasicError
}

type NewChainOrError struct {
	SerializedHeader []byte
	Err              *plugin.BasicError
}

type ChainRPC struct {
	client *rpc.Client
}

// Notice how it goes , you hit up these rpcs first
// which then look for the method on the rpc server.
// they always take arg and resp
func (c *ChainRPC) NewChain(startingArgs *NewChainStartingArgs) NewChainOrError {
	var result NewChainOrError

	err := c.client.Call("Plugin.NewChain", startingArgs, &result)
	if err != nil {
		return NewChainOrError{Err: plugin.NewBasicError(err)}
	}
	// if errCall.Message != "" {
	// 	return &errCall
	// }

	return result
}

func (c *ChainRPC) SetFeeRecipient(addr string) error {
	var errCall plugin.BasicError
	err := c.client.Call("Plugin.SetFeeRecipient", addr, &errCall)
	if err != nil {
		return err
	}
	if errCall.Message != "" {
		return &errCall
	}
	return nil
}

func (c *ChainRPC) InitExtraConfigs(configs []byte) error {
	var errCall plugin.BasicError
	err := c.client.Call("Plugin.InitExtraConfigs", configs, &errCall)
	if err != nil {
		return err
	}
	if errCall.Message != "" {
		return &errCall
	}
	return nil
}

func (c *ChainRPC) StartBlock(timeDelta uint64) error {
	var errCall plugin.BasicError
	err := c.client.Call("Plugin.StartBlock", timeDelta, &errCall)
	if err != nil {
		return err
	}
	if errCall.Message != "" {
		return &errCall
	}
	return nil
}

func (c *ChainRPC) WSEndpoint() (string, error) {
	var endpoint string
	err := c.client.Call("Plugin.WSEndpoint", new(any), &endpoint)
	return endpoint, err
}

func (c *ChainRPC) AuthEndpoint() (string, error) {
	var endpoint string
	err := c.client.Call("Plugin.AuthEndpoint", new(any), &endpoint)
	return endpoint, err
}

func (c *ChainRPC) Close() error {
	var errCall plugin.BasicError
	err := c.client.Call("Plugin.Close", new(any), &errCall)
	if err != nil {
		return err
	}
	if errCall.Message != "" {
		return &errCall
	}
	return nil
}

func (c *ChainRPC) IncludeTxByHash(hexHash string) error {
	var errCall plugin.BasicError
	err := c.client.Call("Plugin.IncludeTxByHash", hexHash, &errCall)
	if err != nil {
		return err
	}
	if errCall.Message != "" {
		return &errCall
	}
	return nil

}

// force compliation fail
var (
	_ Chain = (*ChainRPC)(nil)
)

func (c *ChainRPC) EndBlock() NewBlockOrError {
	var result NewBlockOrError

	if err := c.client.Call("Plugin.EndBlock", new(any), &result); err != nil {
		return NewBlockOrError{Err: plugin.NewBasicError(err)}
	}

	return result
}

// NOTE ! One oddity is that things could hang if you don't have the message around to call, so keep the signatures
// exactly as they should be

// And these are in the plugin
// just a little confusing so got to keep in mind, the error returned by the call must
// be assigned to the inout return param which should be a return struct with a
// field for error or the payload, so basically an optional type, or more approipriately
// a result type
type ChainRPCServer struct {
	Impl      Chain
	lastError error
}

func (s *ChainRPCServer) SetFeeRecipient(addr string, err *plugin.BasicError) error {
	errSetFeeRecipient := s.Impl.SetFeeRecipient(addr)
	if errSetFeeRecipient != nil {
		*err = *plugin.NewBasicError(errSetFeeRecipient)
	}
	return nil

}

func (s *ChainRPCServer) NewChain(
	args *NewChainStartingArgs,
	serializedGenBlock *NewChainOrError,
) error {
	*serializedGenBlock = s.Impl.NewChain(args)
	// if errNewChain != nil {
	// 	*err = *plugin.NewBasicError(errNewChain)
	// }
	return nil
}

func (s *ChainRPCServer) InitExtraConfigs(configs []byte, err *plugin.BasicError) error {
	errConfigCall := s.Impl.InitExtraConfigs(configs)
	if errConfigCall != nil {
		*err = *plugin.NewBasicError(errConfigCall)
	}
	return nil
}

func (s *ChainRPCServer) StartBlock(timeDelta uint64, err *plugin.BasicError) error {
	errStartBlock := s.Impl.StartBlock(timeDelta)
	if errStartBlock != nil {
		*err = *plugin.NewBasicError(errStartBlock)
	}
	return nil
}

func (s *ChainRPCServer) WSEndpoint(_ any, reply *string) error {
	*reply, s.lastError = s.Impl.WSEndpoint()
	return nil
}

func (s *ChainRPCServer) AuthEndpoint(_ any, reply *string) error {
	*reply, s.lastError = s.Impl.AuthEndpoint()
	return nil
}

func (s *ChainRPCServer) Close(_ any, err *plugin.BasicError) error {
	errCloseCall := s.Impl.Close()
	if errCloseCall != nil {
		*err = *plugin.NewBasicError(errCloseCall)
	}
	return nil
}

func (s *ChainRPCServer) IncludeTxByHash(hexHash string, err *plugin.BasicError) error {
	errIncludeTx := s.Impl.IncludeTxByHash(hexHash)
	if errIncludeTx != nil {
		*err = *plugin.NewBasicError(errIncludeTx)
	}
	return nil
}

func (s *ChainRPCServer) EndBlock(_ any, serializedBlockOrError *NewBlockOrError) error {
	*serializedBlockOrError = s.Impl.EndBlock()
	return nil
}

type ChainPlugin struct {
	Impl Chain
}

// And these are what make it respectively
func (p *ChainPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ChainRPCServer{Impl: p.Impl}, nil
}

func (ChainPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ChainRPC{client: c}, nil
}
