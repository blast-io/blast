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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"name\":\"DistributeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReporter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceIsInitialized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewPrice\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REPORTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b50604051610a0f380380610a0f83398101604081905261003091610119565b6001600160a01b038116608052600160a052600060c081905260e05261005461005a565b50610149565b600054610100900460ff16156100c65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614610117576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561012b57600080fd5b81516001600160a01b038116811461014257600080fd5b9392505050565b60805160a05160c05160e0516108876101886000396000610180015260006101570152600061012e01526000818160c8015261055c01526108876000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637ae556b51161005b5780637ae556b5146100c3578063a035b1fe14610102578063e20ccec31461010b578063fe4b84df1461011457600080fd5b806306661abd1461008257806354fd4d50146100995780635b9af12b146100ae575b600080fd5b6033545b6040519081526020015b60405180910390f35b6100a1610127565b60405161009091906106a6565b6100c16100bc3660046106d9565b6101ca565b005b6100ea7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610090565b61008660015481565b61008660025481565b6100c16101223660046106d9565b6101d6565b60606101527f0000000000000000000000000000000000000000000000000000000000000000610354565b61017b7f0000000000000000000000000000000000000000000000000000000000000000610354565b6101a47f0000000000000000000000000000000000000000000000000000000000000000610354565b6040516020016101b6939291906106f2565b604051602081830303815290604052905090565b6101d38161045d565b50565b600054610100900460ff16158080156101f65750600054600160ff909116105b806102105750303b158015610210575060005460ff166001145b6102785760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801561029b576000805461ff0019166101001790555b6102a4826104c1565b60405163099005e760e31b81526002604360981b0190634c802f38906102d890309060019060009061dead90600401610762565b600060405180830381600087803b1580156102f257600080fd5b505af1158015610306573d6000803e3d6000fd5b505050508015610350576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b60608160000361037b5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156103a5578061038f816107c8565b915061039e9050600a836107f7565b915061037f565b60008167ffffffffffffffff8111156103c0576103c061080b565b6040519080825280601f01601f1916602001820160405280156103ea576020820181803683370190505b5090505b8415610455576103ff600183610821565b915061040c600a86610838565b61041790603061084c565b60f81b81838151811061042c5761042c610864565b60200101906001600160f81b031916908160001a90535061044e600a866107f7565b94506103ee565b949350505050565b61046681610552565b604051635b9af12b60e01b8152600481018290526004604360981b0190635b9af12b90602401600060405180830381600087803b1580156104a657600080fd5b505af11580156104ba573d6000803e3d6000fd5b5050505050565b600054610100900460ff1661052c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840161026f565b6001541561054d5760405163131cb46d60e21b815260040160405180910390fd5b600155565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167311110000000000000000000000000000000011101933016001600160a01b0316146105bb57604051631d73770560e11b815260040160405180910390fd5b80156105d95780600260008282546105d3919061084c565b90915550505b61035060006105e760335490565b60025410806105f65750603354155b156106015750600090565b60335460025461061191906107f7565b60016000828254610622919061084c565b90915550506033546002546106379190610838565b6002556001546040519081527f270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b79060200160405180910390a150600190565b60005b83811015610691578181015183820152602001610679565b838111156106a0576000848401525b50505050565b60208152600082518060208401526106c5816040850160208701610676565b601f01601f19169190910160400192915050565b6000602082840312156106eb57600080fd5b5035919050565b60008451610704818460208901610676565b8083019050601760f91b8082528551610724816001850160208a01610676565b6001920191820152835161073f816002840160208801610676565b0160020195945050505050565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b0385811682526080820190600386106107845761078461074c565b8560208401526002851061079a5761079a61074c565b84604084015280841660608401525095945050505050565b634e487b7160e01b600052601160045260246000fd5b6000600182016107da576107da6107b2565b5060010190565b634e487b7160e01b600052601260045260246000fd5b600082610806576108066107e1565b500490565b634e487b7160e01b600052604160045260246000fd5b600082821015610833576108336107b2565b500390565b600082610847576108476107e1565b500690565b6000821982111561085f5761085f6107b2565b500190565b634e487b7160e01b600052603260045260246000fdfea164736f6c634300080f000a",
}

// SharesABI is the input ABI used to generate the binding from.
// Deprecated: Use SharesMetaData.ABI instead.
var SharesABI = SharesMetaData.ABI

// SharesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SharesMetaData.Bin instead.
var SharesBin = SharesMetaData.Bin

// DeployShares deploys a new Ethereum contract, binding an instance of Shares to it.
func DeployShares(auth *bind.TransactOpts, backend bind.ContractBackend, _reporter common.Address) (common.Address, *types.Transaction, *Shares, error) {
	parsed, err := SharesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SharesBin), backend, _reporter)
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
