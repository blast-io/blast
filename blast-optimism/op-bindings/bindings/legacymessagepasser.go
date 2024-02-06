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

// LegacyMessagePasserMetaData contains all meta data concerning the LegacyMessagePasser contract.
var LegacyMessagePasserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"passMessageToL1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sentMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061029e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806354fd4d501461004657806382e3702d14610080578063cafa81dc146100b3575b600080fd5b61006a604051806040016040528060058152602001640312e312e360dc1b81525081565b6040516100779190610147565b60405180910390f35b6100a361008e36600461017a565b60006020819052908152604090205460ff1681565b6040519015158152602001610077565b6100c66100c13660046101a9565b6100c8565b005b600160008083336040516020016100e092919061025a565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff191691151591909117905550565b60005b8381101561013257818101518382015260200161011a565b83811115610141576000848401525b50505050565b6020815260008251806020840152610166816040850160208701610117565b601f01601f19169190910160400192915050565b60006020828403121561018c57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156101bb57600080fd5b813567ffffffffffffffff808211156101d357600080fd5b818401915084601f8301126101e757600080fd5b8135818111156101f9576101f9610193565b604051601f8201601f19908116603f0116810190838211818310171561022157610221610193565b8160405282815287602084870101111561023a57600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000835161026c818460208801610117565b60609390931b6bffffffffffffffffffffffff1916919092019081526014019291505056fea164736f6c634300080f000a",
}

// LegacyMessagePasserABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyMessagePasserMetaData.ABI instead.
var LegacyMessagePasserABI = LegacyMessagePasserMetaData.ABI

// LegacyMessagePasserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LegacyMessagePasserMetaData.Bin instead.
var LegacyMessagePasserBin = LegacyMessagePasserMetaData.Bin

// DeployLegacyMessagePasser deploys a new Ethereum contract, binding an instance of LegacyMessagePasser to it.
func DeployLegacyMessagePasser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LegacyMessagePasser, error) {
	parsed, err := LegacyMessagePasserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LegacyMessagePasserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LegacyMessagePasser{LegacyMessagePasserCaller: LegacyMessagePasserCaller{contract: contract}, LegacyMessagePasserTransactor: LegacyMessagePasserTransactor{contract: contract}, LegacyMessagePasserFilterer: LegacyMessagePasserFilterer{contract: contract}}, nil
}

// LegacyMessagePasser is an auto generated Go binding around an Ethereum contract.
type LegacyMessagePasser struct {
	LegacyMessagePasserCaller     // Read-only binding to the contract
	LegacyMessagePasserTransactor // Write-only binding to the contract
	LegacyMessagePasserFilterer   // Log filterer for contract events
}

// LegacyMessagePasserCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyMessagePasserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMessagePasserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyMessagePasserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMessagePasserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyMessagePasserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMessagePasserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyMessagePasserSession struct {
	Contract     *LegacyMessagePasser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LegacyMessagePasserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyMessagePasserCallerSession struct {
	Contract *LegacyMessagePasserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LegacyMessagePasserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyMessagePasserTransactorSession struct {
	Contract     *LegacyMessagePasserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LegacyMessagePasserRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyMessagePasserRaw struct {
	Contract *LegacyMessagePasser // Generic contract binding to access the raw methods on
}

// LegacyMessagePasserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyMessagePasserCallerRaw struct {
	Contract *LegacyMessagePasserCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyMessagePasserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyMessagePasserTransactorRaw struct {
	Contract *LegacyMessagePasserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyMessagePasser creates a new instance of LegacyMessagePasser, bound to a specific deployed contract.
func NewLegacyMessagePasser(address common.Address, backend bind.ContractBackend) (*LegacyMessagePasser, error) {
	contract, err := bindLegacyMessagePasser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyMessagePasser{LegacyMessagePasserCaller: LegacyMessagePasserCaller{contract: contract}, LegacyMessagePasserTransactor: LegacyMessagePasserTransactor{contract: contract}, LegacyMessagePasserFilterer: LegacyMessagePasserFilterer{contract: contract}}, nil
}

// NewLegacyMessagePasserCaller creates a new read-only instance of LegacyMessagePasser, bound to a specific deployed contract.
func NewLegacyMessagePasserCaller(address common.Address, caller bind.ContractCaller) (*LegacyMessagePasserCaller, error) {
	contract, err := bindLegacyMessagePasser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyMessagePasserCaller{contract: contract}, nil
}

// NewLegacyMessagePasserTransactor creates a new write-only instance of LegacyMessagePasser, bound to a specific deployed contract.
func NewLegacyMessagePasserTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyMessagePasserTransactor, error) {
	contract, err := bindLegacyMessagePasser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyMessagePasserTransactor{contract: contract}, nil
}

// NewLegacyMessagePasserFilterer creates a new log filterer instance of LegacyMessagePasser, bound to a specific deployed contract.
func NewLegacyMessagePasserFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyMessagePasserFilterer, error) {
	contract, err := bindLegacyMessagePasser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyMessagePasserFilterer{contract: contract}, nil
}

// bindLegacyMessagePasser binds a generic wrapper to an already deployed contract.
func bindLegacyMessagePasser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyMessagePasserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyMessagePasser *LegacyMessagePasserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyMessagePasser.Contract.LegacyMessagePasserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyMessagePasser *LegacyMessagePasserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMessagePasser.Contract.LegacyMessagePasserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyMessagePasser *LegacyMessagePasserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyMessagePasser.Contract.LegacyMessagePasserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyMessagePasser *LegacyMessagePasserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyMessagePasser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyMessagePasser *LegacyMessagePasserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMessagePasser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyMessagePasser *LegacyMessagePasserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyMessagePasser.Contract.contract.Transact(opts, method, params...)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_LegacyMessagePasser *LegacyMessagePasserCaller) SentMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _LegacyMessagePasser.contract.Call(opts, &out, "sentMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_LegacyMessagePasser *LegacyMessagePasserSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _LegacyMessagePasser.Contract.SentMessages(&_LegacyMessagePasser.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_LegacyMessagePasser *LegacyMessagePasserCallerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _LegacyMessagePasser.Contract.SentMessages(&_LegacyMessagePasser.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyMessagePasser *LegacyMessagePasserCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyMessagePasser.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyMessagePasser *LegacyMessagePasserSession) Version() (string, error) {
	return _LegacyMessagePasser.Contract.Version(&_LegacyMessagePasser.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyMessagePasser *LegacyMessagePasserCallerSession) Version() (string, error) {
	return _LegacyMessagePasser.Contract.Version(&_LegacyMessagePasser.CallOpts)
}

// PassMessageToL1 is a paid mutator transaction binding the contract method 0xcafa81dc.
//
// Solidity: function passMessageToL1(bytes _message) returns()
func (_LegacyMessagePasser *LegacyMessagePasserTransactor) PassMessageToL1(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _LegacyMessagePasser.contract.Transact(opts, "passMessageToL1", _message)
}

// PassMessageToL1 is a paid mutator transaction binding the contract method 0xcafa81dc.
//
// Solidity: function passMessageToL1(bytes _message) returns()
func (_LegacyMessagePasser *LegacyMessagePasserSession) PassMessageToL1(_message []byte) (*types.Transaction, error) {
	return _LegacyMessagePasser.Contract.PassMessageToL1(&_LegacyMessagePasser.TransactOpts, _message)
}

// PassMessageToL1 is a paid mutator transaction binding the contract method 0xcafa81dc.
//
// Solidity: function passMessageToL1(bytes _message) returns()
func (_LegacyMessagePasser *LegacyMessagePasserTransactorSession) PassMessageToL1(_message []byte) (*types.Transaction, error) {
	return _LegacyMessagePasser.Contract.PassMessageToL1(&_LegacyMessagePasser.TransactOpts, _message)
}
