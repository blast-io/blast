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

// SharesMetaData contains all meta data concerning the Shares contract.
var SharesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"name\":\"DistributeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReporter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceIsInitialized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewPrice\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REPORTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b50604051610a4c380380610a4c83398101604081905261003091610225565b6001600160a01b038116608052600160a052600060c081905260e0526100558261005c565b5050610262565b600054610100900460ff161580801561007c5750600054600160ff909116105b806100a757506100953061018560201b6103bd1760201c565b1580156100a7575060005460ff166001145b61010f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610132576000805461ff0019166101001790555b61013b82610194565b8015610181576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6001600160a01b03163b151590565b600054610100900460ff166101ff5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610106565b600154156102205760405163131cb46d60e21b815260040160405180910390fd5b600155565b6000806040838503121561023857600080fd5b825160208401519092506001600160a01b038116811461025757600080fd5b809150509250929050565b60805160a05160c05160e0516107ab6102a160003960006101ce015260006101a50152600061017c01526000818160db015261022201526107ab6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80637ae556b51161005b5780637ae556b5146100d6578063a035b1fe14610115578063e20ccec31461011e578063fe4b84df1461012757600080fd5b806306661abd1461008d5780634291cd11146100a457806354fd4d50146100ae5780635b9af12b146100c3575b600080fd5b6033545b6040519081526020015b60405180910390f35b6100ac61013a565b005b6100b6610175565b60405161009b9190610630565b6100ac6100d1366004610663565b610218565b6100fd7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161009b565b61009160015481565b61009160025481565b6100ac610135366004610663565b6102ab565b6101426103cc565b610173576033546002546040516307b4790760e11b8152600481019290925260248201526044015b60405180910390fd5b565b60606101a07f0000000000000000000000000000000000000000000000000000000000000000610466565b6101c97f0000000000000000000000000000000000000000000000000000000000000000610466565b6101f27f0000000000000000000000000000000000000000000000000000000000000000610466565b6040516020016102049392919061067c565b604051602081830303815290604052905090565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167311110000000000000000000000000000000011101933016001600160a01b03161461028157604051631d73770560e11b815260040160405180910390fd5b801561029f57806002600082825461029991906106ec565b90915550505b6102a76103cc565b5050565b600054610100900460ff16158080156102cb5750600054600160ff909116105b806102e55750303b1580156102e5575060005460ff166001145b6103485760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161016a565b6000805460ff19166001179055801561036b576000805461ff0019166101001790555b6103748261056f565b80156102a7576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6001600160a01b03163b151590565b60006103d760335490565b60025410806103e65750603354155b156103f15750600090565b603354600254610401919061071a565b6001600082825461041291906106ec565b9091555050603354600254610427919061072e565b6002556001546040519081527f270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b79060200160405180910390a150600190565b60608160000361048d5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156104b757806104a181610742565b91506104b09050600a8361071a565b9150610491565b60008167ffffffffffffffff8111156104d2576104d261075b565b6040519080825280601f01601f1916602001820160405280156104fc576020820181803683370190505b5090505b841561056757610511600183610771565b915061051e600a8661072e565b6105299060306106ec565b60f81b81838151811061053e5761053e610788565b60200101906001600160f81b031916908160001a905350610560600a8661071a565b9450610500565b949350505050565b600054610100900460ff166105da5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840161016a565b600154156105fb5760405163131cb46d60e21b815260040160405180910390fd5b600155565b60005b8381101561061b578181015183820152602001610603565b8381111561062a576000848401525b50505050565b602081526000825180602084015261064f816040850160208701610600565b601f01601f19169190910160400192915050565b60006020828403121561067557600080fd5b5035919050565b6000845161068e818460208901610600565b8083019050601760f91b80825285516106ae816001850160208a01610600565b600192019182015283516106c9816002840160208801610600565b0160020195945050505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156106ff576106ff6106d6565b500190565b634e487b7160e01b600052601260045260246000fd5b60008261072957610729610704565b500490565b60008261073d5761073d610704565b500690565b600060018201610754576107546106d6565b5060010190565b634e487b7160e01b600052604160045260246000fd5b600082821015610783576107836106d6565b500390565b634e487b7160e01b600052603260045260246000fdfea164736f6c634300080f000a",
}

// SharesABI is the input ABI used to generate the binding from.
// Deprecated: Use SharesMetaData.ABI instead.
var SharesABI = SharesMetaData.ABI

// SharesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SharesMetaData.Bin instead.
var SharesBin = SharesMetaData.Bin

// DeployShares deploys a new Ethereum contract, binding an instance of Shares to it.
func DeployShares(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _reporter common.Address) (common.Address, *types.Transaction, *Shares, error) {
	parsed, err := SharesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SharesBin), backend, _price, _reporter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Shares{SharesCaller: SharesCaller{contract: contract}, SharesTransactor: SharesTransactor{contract: contract}, SharesFilterer: SharesFilterer{contract: contract}}, nil
}

// Shares is an auto generated Go binding around an Ethereum contract.
type Shares struct {
	SharesCaller     // Read-only binding to the contract
	SharesTransactor // Write-only binding to the contract
	SharesFilterer   // Log filterer for contract events
}

// SharesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SharesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SharesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SharesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SharesSession struct {
	Contract     *Shares           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SharesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SharesCallerSession struct {
	Contract *SharesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SharesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SharesTransactorSession struct {
	Contract     *SharesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SharesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SharesRaw struct {
	Contract *Shares // Generic contract binding to access the raw methods on
}

// SharesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SharesCallerRaw struct {
	Contract *SharesCaller // Generic read-only contract binding to access the raw methods on
}

// SharesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SharesTransactorRaw struct {
	Contract *SharesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShares creates a new instance of Shares, bound to a specific deployed contract.
func NewShares(address common.Address, backend bind.ContractBackend) (*Shares, error) {
	contract, err := bindShares(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shares{SharesCaller: SharesCaller{contract: contract}, SharesTransactor: SharesTransactor{contract: contract}, SharesFilterer: SharesFilterer{contract: contract}}, nil
}

// NewSharesCaller creates a new read-only instance of Shares, bound to a specific deployed contract.
func NewSharesCaller(address common.Address, caller bind.ContractCaller) (*SharesCaller, error) {
	contract, err := bindShares(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SharesCaller{contract: contract}, nil
}

// NewSharesTransactor creates a new write-only instance of Shares, bound to a specific deployed contract.
func NewSharesTransactor(address common.Address, transactor bind.ContractTransactor) (*SharesTransactor, error) {
	contract, err := bindShares(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SharesTransactor{contract: contract}, nil
}

// NewSharesFilterer creates a new log filterer instance of Shares, bound to a specific deployed contract.
func NewSharesFilterer(address common.Address, filterer bind.ContractFilterer) (*SharesFilterer, error) {
	contract, err := bindShares(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SharesFilterer{contract: contract}, nil
}

// bindShares binds a generic wrapper to an already deployed contract.
func bindShares(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SharesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shares *SharesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shares.Contract.SharesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shares *SharesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shares.Contract.SharesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shares *SharesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shares.Contract.SharesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shares *SharesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shares.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shares *SharesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shares.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shares *SharesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shares.Contract.contract.Transact(opts, method, params...)
}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_Shares *SharesCaller) REPORTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Shares.contract.Call(opts, &out, "REPORTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_Shares *SharesSession) REPORTER() (common.Address, error) {
	return _Shares.Contract.REPORTER(&_Shares.CallOpts)
}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_Shares *SharesCallerSession) REPORTER() (common.Address, error) {
	return _Shares.Contract.REPORTER(&_Shares.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Shares *SharesCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Shares.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Shares *SharesSession) Count() (*big.Int, error) {
	return _Shares.Contract.Count(&_Shares.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Shares *SharesCallerSession) Count() (*big.Int, error) {
	return _Shares.Contract.Count(&_Shares.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_Shares *SharesCaller) Pending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Shares.contract.Call(opts, &out, "pending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_Shares *SharesSession) Pending() (*big.Int, error) {
	return _Shares.Contract.Pending(&_Shares.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_Shares *SharesCallerSession) Pending() (*big.Int, error) {
	return _Shares.Contract.Pending(&_Shares.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Shares *SharesCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Shares.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Shares *SharesSession) Price() (*big.Int, error) {
	return _Shares.Contract.Price(&_Shares.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Shares *SharesCallerSession) Price() (*big.Int, error) {
	return _Shares.Contract.Price(&_Shares.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Shares *SharesCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Shares.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Shares *SharesSession) Version() (string, error) {
	return _Shares.Contract.Version(&_Shares.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Shares *SharesCallerSession) Version() (string, error) {
	return _Shares.Contract.Version(&_Shares.CallOpts)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_Shares *SharesTransactor) AddValue(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Shares.contract.Transact(opts, "addValue", value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_Shares *SharesSession) AddValue(value *big.Int) (*types.Transaction, error) {
	return _Shares.Contract.AddValue(&_Shares.TransactOpts, value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_Shares *SharesTransactorSession) AddValue(value *big.Int) (*types.Transaction, error) {
	return _Shares.Contract.AddValue(&_Shares.TransactOpts, value)
}

// DistributePending is a paid mutator transaction binding the contract method 0x4291cd11.
//
// Solidity: function distributePending() returns()
func (_Shares *SharesTransactor) DistributePending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shares.contract.Transact(opts, "distributePending")
}

// DistributePending is a paid mutator transaction binding the contract method 0x4291cd11.
//
// Solidity: function distributePending() returns()
func (_Shares *SharesSession) DistributePending() (*types.Transaction, error) {
	return _Shares.Contract.DistributePending(&_Shares.TransactOpts)
}

// DistributePending is a paid mutator transaction binding the contract method 0x4291cd11.
//
// Solidity: function distributePending() returns()
func (_Shares *SharesTransactorSession) DistributePending() (*types.Transaction, error) {
	return _Shares.Contract.DistributePending(&_Shares.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _price) returns()
func (_Shares *SharesTransactor) Initialize(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Shares.contract.Transact(opts, "initialize", _price)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _price) returns()
func (_Shares *SharesSession) Initialize(_price *big.Int) (*types.Transaction, error) {
	return _Shares.Contract.Initialize(&_Shares.TransactOpts, _price)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _price) returns()
func (_Shares *SharesTransactorSession) Initialize(_price *big.Int) (*types.Transaction, error) {
	return _Shares.Contract.Initialize(&_Shares.TransactOpts, _price)
}

// SharesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Shares contract.
type SharesInitializedIterator struct {
	Event *SharesInitialized // Event containing the contract specifics and raw log

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
func (it *SharesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SharesInitialized)
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
		it.Event = new(SharesInitialized)
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
func (it *SharesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SharesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SharesInitialized represents a Initialized event raised by the Shares contract.
type SharesInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Shares *SharesFilterer) FilterInitialized(opts *bind.FilterOpts) (*SharesInitializedIterator, error) {

	logs, sub, err := _Shares.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SharesInitializedIterator{contract: _Shares.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Shares *SharesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SharesInitialized) (event.Subscription, error) {

	logs, sub, err := _Shares.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SharesInitialized)
				if err := _Shares.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Shares *SharesFilterer) ParseInitialized(log types.Log) (*SharesInitialized, error) {
	event := new(SharesInitialized)
	if err := _Shares.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SharesNewPriceIterator is returned from FilterNewPrice and is used to iterate over the raw logs and unpacked data for NewPrice events raised by the Shares contract.
type SharesNewPriceIterator struct {
	Event *SharesNewPrice // Event containing the contract specifics and raw log

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
func (it *SharesNewPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SharesNewPrice)
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
		it.Event = new(SharesNewPrice)
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
func (it *SharesNewPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SharesNewPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SharesNewPrice represents a NewPrice event raised by the Shares contract.
type SharesNewPrice struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewPrice is a free log retrieval operation binding the contract event 0x270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b7.
//
// Solidity: event NewPrice(uint256 price)
func (_Shares *SharesFilterer) FilterNewPrice(opts *bind.FilterOpts) (*SharesNewPriceIterator, error) {

	logs, sub, err := _Shares.contract.FilterLogs(opts, "NewPrice")
	if err != nil {
		return nil, err
	}
	return &SharesNewPriceIterator{contract: _Shares.contract, event: "NewPrice", logs: logs, sub: sub}, nil
}

// WatchNewPrice is a free log subscription operation binding the contract event 0x270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b7.
//
// Solidity: event NewPrice(uint256 price)
func (_Shares *SharesFilterer) WatchNewPrice(opts *bind.WatchOpts, sink chan<- *SharesNewPrice) (event.Subscription, error) {

	logs, sub, err := _Shares.contract.WatchLogs(opts, "NewPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SharesNewPrice)
				if err := _Shares.contract.UnpackLog(event, "NewPrice", log); err != nil {
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

// ParseNewPrice is a log parse operation binding the contract event 0x270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b7.
//
// Solidity: event NewPrice(uint256 price)
func (_Shares *SharesFilterer) ParseNewPrice(log types.Log) (*SharesNewPrice, error) {
	event := new(SharesNewPrice)
	if err := _Shares.contract.UnpackLog(event, "NewPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
