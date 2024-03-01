// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DelayedVetoableMetaData contains all meta data concerning the DelayedVetoable contract.
var DelayedVetoableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vetoer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatingDelay_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyDelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwardingEarly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetUnitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"DelayActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Forwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Initiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Vetoed\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initiator_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"}],\"name\":\"queuedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"queuedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vetoer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vetoer_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b506040516107cf3803806107cf8339810160408190526100309161006e565b6001600160a01b0393841660a05291831660c05290911660805260e0526100b9565b80516001600160a01b038116811461006957600080fd5b919050565b6000806000806080858703121561008457600080fd5b61008d85610052565b935061009b60208601610052565b92506100a960408601610052565b6060959095015193969295505050565b60805160a05160c05160e0516106ac61012360003960006101ca01526000818161011d015281816101900152818161024b015281816103a30152610437015260008181610151015281816102f501526104cf0152600081816104a3015261053101526106ac6000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806354fd4d501461006c5780635c39fcc1146100a65780636a42b8f8146100c6578063b912de5d146100dc578063d4b83992146100ef578063d8bff440146100f7575b61006a6100ff565b005b610090604051806040016040528060058152602001640312e302e360dc1b81525081565b60405161009d91906105cc565b60405180910390f35b6100ae61042d565b6040516001600160a01b03909116815260200161009d565b6100ce610464565b60405190815260200161009d565b6100ce6100ea366004610621565b610472565b6100ae610499565b6100ae6104c5565b3615801561010d5750600054155b1561022357336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015906101745750336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614155b156101c85760405163295a81c160e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660048201523360248201526044015b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000060008190556040519081527febf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a89060200160405180910390a1565b6000803660405161023592919061063a565b6040519081900390209050336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156102845750600081815260016020526040902054155b156102ea5760005460000361029c5761029c816104f1565b6000818152600160205260408082204290555182917f87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13916102df9190369061064a565b60405180910390a250565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561032f575060008181526001602052604090205415155b15610377576000818152600160205260408082208290555182917fbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040916102df9190369061064a565b60008181526001602052604081205490036103d65760405163295a81c160e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660048201523360248201526044016101bf565b60008054828252600160205260409091205442916103f391610679565b1015610412576040516343dc986d60e01b815260040160405180910390fd5b60008181526001602052604081205561042a816104f1565b50565b60003361045957507f000000000000000000000000000000000000000000000000000000000000000090565b6104616100ff565b90565b600033610459575060005490565b60003361048c575060009081526001602052604090205490565b6104946100ff565b919050565b60003361045957507f000000000000000000000000000000000000000000000000000000000000000090565b60003361045957507f000000000000000000000000000000000000000000000000000000000000000090565b807f4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e122960003660405161052492919061064a565b60405180910390a26000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031660003660405161056a92919061063a565b6000604051808303816000865af19150503d80600081146105a7576040519150601f19603f3d011682016040523d82523d6000602084013e6105ac565b606091505b5090925090508115156001036105c457805160208201f35b805160208201fd5b600060208083528351808285015260005b818110156105f9578581018301518582016040015282016105dd565b8181111561060b576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561063357600080fd5b5035919050565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000821982111561069a57634e487b7160e01b600052601160045260246000fd5b50019056fea164736f6c634300080f000a",
}

// DelayedVetoableABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayedVetoableMetaData.ABI instead.
var DelayedVetoableABI = DelayedVetoableMetaData.ABI

// DelayedVetoableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelayedVetoableMetaData.Bin instead.
var DelayedVetoableBin = DelayedVetoableMetaData.Bin

// DeployDelayedVetoable deploys a new Ethereum contract, binding an instance of DelayedVetoable to it.
func DeployDelayedVetoable(auth *bind.TransactOpts, backend bind.ContractBackend, vetoer_ common.Address, initiator_ common.Address, target_ common.Address, operatingDelay_ *big.Int) (common.Address, *types.Transaction, *DelayedVetoable, error) {
	parsed, err := DelayedVetoableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelayedVetoableBin), backend, vetoer_, initiator_, target_, operatingDelay_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelayedVetoable{DelayedVetoableCaller: DelayedVetoableCaller{contract: contract}, DelayedVetoableTransactor: DelayedVetoableTransactor{contract: contract}, DelayedVetoableFilterer: DelayedVetoableFilterer{contract: contract}}, nil
}

// DelayedVetoable is an auto generated Go binding around an Ethereum contract.
type DelayedVetoable struct {
	DelayedVetoableCaller     // Read-only binding to the contract
	DelayedVetoableTransactor // Write-only binding to the contract
	DelayedVetoableFilterer   // Log filterer for contract events
}

// DelayedVetoableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedVetoableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedVetoableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedVetoableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedVetoableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayedVetoableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedVetoableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedVetoableSession struct {
	Contract     *DelayedVetoable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelayedVetoableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedVetoableCallerSession struct {
	Contract *DelayedVetoableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DelayedVetoableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedVetoableTransactorSession struct {
	Contract     *DelayedVetoableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DelayedVetoableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedVetoableRaw struct {
	Contract *DelayedVetoable // Generic contract binding to access the raw methods on
}

// DelayedVetoableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedVetoableCallerRaw struct {
	Contract *DelayedVetoableCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedVetoableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedVetoableTransactorRaw struct {
	Contract *DelayedVetoableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedVetoable creates a new instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoable(address common.Address, backend bind.ContractBackend) (*DelayedVetoable, error) {
	contract, err := bindDelayedVetoable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoable{DelayedVetoableCaller: DelayedVetoableCaller{contract: contract}, DelayedVetoableTransactor: DelayedVetoableTransactor{contract: contract}, DelayedVetoableFilterer: DelayedVetoableFilterer{contract: contract}}, nil
}

// NewDelayedVetoableCaller creates a new read-only instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoableCaller(address common.Address, caller bind.ContractCaller) (*DelayedVetoableCaller, error) {
	contract, err := bindDelayedVetoable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableCaller{contract: contract}, nil
}

// NewDelayedVetoableTransactor creates a new write-only instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoableTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedVetoableTransactor, error) {
	contract, err := bindDelayedVetoable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableTransactor{contract: contract}, nil
}

// NewDelayedVetoableFilterer creates a new log filterer instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoableFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayedVetoableFilterer, error) {
	contract, err := bindDelayedVetoable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableFilterer{contract: contract}, nil
}

// bindDelayedVetoable binds a generic wrapper to an already deployed contract.
func bindDelayedVetoable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelayedVetoableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedVetoable *DelayedVetoableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedVetoable.Contract.DelayedVetoableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedVetoable *DelayedVetoableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.DelayedVetoableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedVetoable *DelayedVetoableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.DelayedVetoableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedVetoable *DelayedVetoableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedVetoable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedVetoable *DelayedVetoableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedVetoable *DelayedVetoableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedVetoable *DelayedVetoableCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelayedVetoable.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedVetoable *DelayedVetoableSession) Version() (string, error) {
	return _DelayedVetoable.Contract.Version(&_DelayedVetoable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedVetoable *DelayedVetoableCallerSession) Version() (string, error) {
	return _DelayedVetoable.Contract.Version(&_DelayedVetoable.CallOpts)
}

// Delay is a paid mutator transaction binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() returns(uint256 delay_)
func (_DelayedVetoable *DelayedVetoableTransactor) Delay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "delay")
}

// Delay is a paid mutator transaction binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() returns(uint256 delay_)
func (_DelayedVetoable *DelayedVetoableSession) Delay() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Delay(&_DelayedVetoable.TransactOpts)
}

// Delay is a paid mutator transaction binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() returns(uint256 delay_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Delay() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Delay(&_DelayedVetoable.TransactOpts)
}

// Initiator is a paid mutator transaction binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() returns(address initiator_)
func (_DelayedVetoable *DelayedVetoableTransactor) Initiator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "initiator")
}

// Initiator is a paid mutator transaction binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() returns(address initiator_)
func (_DelayedVetoable *DelayedVetoableSession) Initiator() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Initiator(&_DelayedVetoable.TransactOpts)
}

// Initiator is a paid mutator transaction binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() returns(address initiator_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Initiator() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Initiator(&_DelayedVetoable.TransactOpts)
}

// QueuedAt is a paid mutator transaction binding the contract method 0xb912de5d.
//
// Solidity: function queuedAt(bytes32 callHash) returns(uint256 queuedAt_)
func (_DelayedVetoable *DelayedVetoableTransactor) QueuedAt(opts *bind.TransactOpts, callHash [32]byte) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "queuedAt", callHash)
}

// QueuedAt is a paid mutator transaction binding the contract method 0xb912de5d.
//
// Solidity: function queuedAt(bytes32 callHash) returns(uint256 queuedAt_)
func (_DelayedVetoable *DelayedVetoableSession) QueuedAt(callHash [32]byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.QueuedAt(&_DelayedVetoable.TransactOpts, callHash)
}

// QueuedAt is a paid mutator transaction binding the contract method 0xb912de5d.
//
// Solidity: function queuedAt(bytes32 callHash) returns(uint256 queuedAt_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) QueuedAt(callHash [32]byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.QueuedAt(&_DelayedVetoable.TransactOpts, callHash)
}

// Target is a paid mutator transaction binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address target_)
func (_DelayedVetoable *DelayedVetoableTransactor) Target(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "target")
}

// Target is a paid mutator transaction binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address target_)
func (_DelayedVetoable *DelayedVetoableSession) Target() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Target(&_DelayedVetoable.TransactOpts)
}

// Target is a paid mutator transaction binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address target_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Target() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Target(&_DelayedVetoable.TransactOpts)
}

// Vetoer is a paid mutator transaction binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() returns(address vetoer_)
func (_DelayedVetoable *DelayedVetoableTransactor) Vetoer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "vetoer")
}

// Vetoer is a paid mutator transaction binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() returns(address vetoer_)
func (_DelayedVetoable *DelayedVetoableSession) Vetoer() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Vetoer(&_DelayedVetoable.TransactOpts)
}

// Vetoer is a paid mutator transaction binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() returns(address vetoer_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Vetoer() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Vetoer(&_DelayedVetoable.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DelayedVetoable *DelayedVetoableTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DelayedVetoable.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DelayedVetoable *DelayedVetoableSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Fallback(&_DelayedVetoable.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DelayedVetoable *DelayedVetoableTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Fallback(&_DelayedVetoable.TransactOpts, calldata)
}

// DelayedVetoableDelayActivatedIterator is returned from FilterDelayActivated and is used to iterate over the raw logs and unpacked data for DelayActivated events raised by the DelayedVetoable contract.
type DelayedVetoableDelayActivatedIterator struct {
	Event *DelayedVetoableDelayActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableDelayActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableDelayActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableDelayActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableDelayActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableDelayActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableDelayActivated represents a DelayActivated event raised by the DelayedVetoable contract.
type DelayedVetoableDelayActivated struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelayActivated is a free log retrieval operation binding the contract event 0xebf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a8.
//
// Solidity: event DelayActivated(uint256 delay)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterDelayActivated(opts *bind.FilterOpts) (*DelayedVetoableDelayActivatedIterator, error) {

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "DelayActivated")
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableDelayActivatedIterator{contract: _DelayedVetoable.contract, event: "DelayActivated", logs: logs, sub: sub}, nil
}

// WatchDelayActivated is a free log subscription operation binding the contract event 0xebf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a8.
//
// Solidity: event DelayActivated(uint256 delay)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchDelayActivated(opts *bind.WatchOpts, sink chan<- *DelayedVetoableDelayActivated) (event.Subscription, error) {

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "DelayActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableDelayActivated)
				if err := _DelayedVetoable.contract.UnpackLog(event, "DelayActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayActivated is a log parse operation binding the contract event 0xebf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a8.
//
// Solidity: event DelayActivated(uint256 delay)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseDelayActivated(log types.Log) (*DelayedVetoableDelayActivated, error) {
	event := new(DelayedVetoableDelayActivated)
	if err := _DelayedVetoable.contract.UnpackLog(event, "DelayActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedVetoableForwardedIterator is returned from FilterForwarded and is used to iterate over the raw logs and unpacked data for Forwarded events raised by the DelayedVetoable contract.
type DelayedVetoableForwardedIterator struct {
	Event *DelayedVetoableForwarded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableForwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableForwarded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableForwarded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableForwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableForwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableForwarded represents a Forwarded event raised by the DelayedVetoable contract.
type DelayedVetoableForwarded struct {
	CallHash [32]byte
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterForwarded is a free log retrieval operation binding the contract event 0x4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e1229.
//
// Solidity: event Forwarded(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterForwarded(opts *bind.FilterOpts, callHash [][32]byte) (*DelayedVetoableForwardedIterator, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "Forwarded", callHashRule)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableForwardedIterator{contract: _DelayedVetoable.contract, event: "Forwarded", logs: logs, sub: sub}, nil
}

// WatchForwarded is a free log subscription operation binding the contract event 0x4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e1229.
//
// Solidity: event Forwarded(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchForwarded(opts *bind.WatchOpts, sink chan<- *DelayedVetoableForwarded, callHash [][32]byte) (event.Subscription, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "Forwarded", callHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableForwarded)
				if err := _DelayedVetoable.contract.UnpackLog(event, "Forwarded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForwarded is a log parse operation binding the contract event 0x4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e1229.
//
// Solidity: event Forwarded(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseForwarded(log types.Log) (*DelayedVetoableForwarded, error) {
	event := new(DelayedVetoableForwarded)
	if err := _DelayedVetoable.contract.UnpackLog(event, "Forwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedVetoableInitiatedIterator is returned from FilterInitiated and is used to iterate over the raw logs and unpacked data for Initiated events raised by the DelayedVetoable contract.
type DelayedVetoableInitiatedIterator struct {
	Event *DelayedVetoableInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableInitiated represents a Initiated event raised by the DelayedVetoable contract.
type DelayedVetoableInitiated struct {
	CallHash [32]byte
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitiated is a free log retrieval operation binding the contract event 0x87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13.
//
// Solidity: event Initiated(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterInitiated(opts *bind.FilterOpts, callHash [][32]byte) (*DelayedVetoableInitiatedIterator, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "Initiated", callHashRule)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableInitiatedIterator{contract: _DelayedVetoable.contract, event: "Initiated", logs: logs, sub: sub}, nil
}

// WatchInitiated is a free log subscription operation binding the contract event 0x87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13.
//
// Solidity: event Initiated(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchInitiated(opts *bind.WatchOpts, sink chan<- *DelayedVetoableInitiated, callHash [][32]byte) (event.Subscription, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "Initiated", callHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableInitiated)
				if err := _DelayedVetoable.contract.UnpackLog(event, "Initiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiated is a log parse operation binding the contract event 0x87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13.
//
// Solidity: event Initiated(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseInitiated(log types.Log) (*DelayedVetoableInitiated, error) {
	event := new(DelayedVetoableInitiated)
	if err := _DelayedVetoable.contract.UnpackLog(event, "Initiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedVetoableVetoedIterator is returned from FilterVetoed and is used to iterate over the raw logs and unpacked data for Vetoed events raised by the DelayedVetoable contract.
type DelayedVetoableVetoedIterator struct {
	Event *DelayedVetoableVetoed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableVetoedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableVetoed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableVetoed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableVetoedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableVetoedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableVetoed represents a Vetoed event raised by the DelayedVetoable contract.
type DelayedVetoableVetoed struct {
	CallHash [32]byte
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVetoed is a free log retrieval operation binding the contract event 0xbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040.
//
// Solidity: event Vetoed(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterVetoed(opts *bind.FilterOpts, callHash [][32]byte) (*DelayedVetoableVetoedIterator, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "Vetoed", callHashRule)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableVetoedIterator{contract: _DelayedVetoable.contract, event: "Vetoed", logs: logs, sub: sub}, nil
}

// WatchVetoed is a free log subscription operation binding the contract event 0xbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040.
//
// Solidity: event Vetoed(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchVetoed(opts *bind.WatchOpts, sink chan<- *DelayedVetoableVetoed, callHash [][32]byte) (event.Subscription, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "Vetoed", callHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableVetoed)
				if err := _DelayedVetoable.contract.UnpackLog(event, "Vetoed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVetoed is a log parse operation binding the contract event 0xbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040.
//
// Solidity: event Vetoed(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseVetoed(log types.Log) (*DelayedVetoableVetoed, error) {
	event := new(DelayedVetoableVetoed)
	if err := _DelayedVetoable.contract.UnpackLog(event, "Vetoed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
