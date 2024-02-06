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

// GasMetaData contains all meta data concerning the Gas contract.
var GasMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blastConfigurationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blastFeeVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_zeroClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilClaimRate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"adminClaimGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseGasSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blastConfigurationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blastFeeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ceilClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ceilGasSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasToClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasSecondsToConsume\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minClaimRateBips\",\"type\":\"uint256\"}],\"name\":\"claimGasAtMinClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasSecondsToConsume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasToClaim\",\"type\":\"uint256\"}],\"name\":\"getClaimRateBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"readGasParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"etherSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"etherBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"enumGasMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumGasMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"name\":\"setGasMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_zeroClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilClaimRate\",\"type\":\"uint256\"}],\"name\":\"updateAdminParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeroClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200116f3803806200116f83398101604081905262000034916200020b565b8285106200009b5760405162461bcd60e51b815260206004820152602960248201527f7a65726f20636c61696d2072617465206d757374206265203c206261736520636044820152686c61696d207261746560b81b60648201526084015b60405180910390fd5b808310620000fe5760405162461bcd60e51b815260206004820152602960248201527f6261736520636c61696d2072617465206d757374206265203c206365696c20636044820152686c61696d207261746560b81b606482015260840162000092565b818410620001635760405162461bcd60e51b815260206004820152602b60248201527f6261736520676173207365636f6e6473206d757374206265203c206365696c2060448201526a676173207365636f6e647360a81b606482015260840162000092565b60008411620001b55760405162461bcd60e51b815260206004820152601c60248201527f6261736520676173207365636f6e6473206d757374206265203e203000000000604482015260640162000092565b6001600160a01b0397881660805295871660a0529390951660c05260009190915560015560029290925560039190915560045562000286565b80516001600160a01b03811681146200020657600080fd5b919050565b600080600080600080600080610100898b0312156200022957600080fd5b6200023489620001ee565b97506200024460208a01620001ee565b96506200025460408a01620001ee565b9550606089015194506080890151935060a0890151925060c0890151915060e089015190509295985092959890939650565b60805160a05160c051610e8f620002e0600039600081816102270152818161057901526106720152600081816101ae015281816103a101526109cb01526000818161028d015281816105e001526106cf0152610e8f6000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806361da985e116100a2578063d45e6bdf11610071578063d45e6bdf14610222578063d4810ba514610249578063dde798a41461025c578063f5c5e3431461027f578063f851a4401461028857600080fd5b806361da985e146101a957806388e5f229146101e8578063aaa2f64314610210578063bafe80961461021957600080fd5b8063445067cb116100de578063445067cb14610165578063494588681461017a5780635767bba51461018d5780635a03838f146101a057600080fd5b80630951888f14610110578063150111191461013657806326af48321461013f5780632d7a59e114610152575b600080fd5b61012361011e366004610bc5565b6102af565b6040519081526020015b60405180910390f35b61012360035481565b61012361014d366004610c01565b610394565b610123610160366004610c43565b6105d3565b610178610173366004610c5e565b6106c4565b005b610123610188366004610c99565b6108c2565b61012361019b366004610c99565b6108d1565b61012360005481565b6101d07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161012d565b6101fb6101f6366004610ccc565b6108fa565b6040805192835260208301919091520161012d565b61012360015481565b61012360045481565b6101d07f000000000000000000000000000000000000000000000000000000000000000081565b610178610257366004610cee565b6109c0565b61026f61026a366004610c43565b610a2c565b60405161012d9493929190610d3f565b61012360025481565b6101d07f000000000000000000000000000000000000000000000000000000000000000081565b60008060006102bd86610a2c565b50509150915060005484116102df576102d686866108d1565b9250505061038d565b6002548410156102ef5760025493505b6000600254856102ff9190610d97565b905060006001546003546103139190610d97565b905060006002546004546103279190610d97565b90506000816103368486610dae565b6103409190610dcd565b60015461034d9190610def565b9050600061035b8288610dcd565b9050858111156103685750845b60006103748383610dae565b90506103828c8c8484610394565b985050505050505050505b9392505050565b6000336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103e75760405162461bcd60e51b81526004016103de90610e07565b60405180910390fd5b60008060006103f588610a2c565b935050925092506000861161044c5760405162461bcd60e51b815260206004820152601d60248201527f6d757374207769746864726177206e6f6e2d7a65726f20616d6f756e7400000060448201526064016103de565b818611156104935760405162461bcd60e51b8152602060048201526014602482015273746f6f206d75636820746f20776974686472617760601b60448201526064016103de565b828511156104dc5760405162461bcd60e51b81526020600482015260166024820152756e6f7420656e6f75676820676173207365636f6e647360501b60448201526064016103de565b6000806104e987896108fa565b909250905060006127106104fd848b610dae565b6105079190610dcd565b90506000610515828b610d97565b90506105358c610525858a610d97565b61052f8d8a610d97565b88610ad8565b6040516001600160a01b038c169083156108fc029084906000818181858888f1935050505015801561056b573d6000803e3d6000fd5b506040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169082156108fc029083906000818181858888f193505050501580156105c2573d6000803e3d6000fd5b50909b9a5050505050505050505050565b6000336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106475760405162461bcd60e51b815260206004820152601760248201527621b0b63632b91034b9903737ba103a34329030b236b4b760491b60448201526064016103de565b600061065283610a2c565b5050915050610665836000806000610ad8565b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169082156108fc029083906000818181858888f193505050501580156106bb573d6000803e3d6000fd5b5090505b919050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107365760405162461bcd60e51b815260206004820152601760248201527621b0b63632b91034b9903737ba103a34329030b236b4b760491b60448201526064016103de565b8285106107975760405162461bcd60e51b815260206004820152602960248201527f7a65726f20636c61696d2072617465206d757374206265203c206261736520636044820152686c61696d207261746560b81b60648201526084016103de565b8083106107f85760405162461bcd60e51b815260206004820152602960248201527f6261736520636c61696d2072617465206d757374206265203c206365696c20636044820152686c61696d207261746560b81b60648201526084016103de565b81841061085b5760405162461bcd60e51b815260206004820152602b60248201527f6261736520676173207365636f6e6473206d757374206265203c206365696c2060448201526a676173207365636f6e647360a81b60648201526084016103de565b600084116108ab5760405162461bcd60e51b815260206004820152601c60248201527f6261736520676173207365636f6e6473206d757374206265203e20300000000060448201526064016103de565b600094909455600192909255600255600355600455565b600061038d83836004546102af565b60008060006108df85610a2c565b5050915091506108f185858385610394565b95945050505050565b600080806109088486610dcd565b90506001548110156109215750506000805491506109b9565b60035481111561094a5760006003548561093b9190610dae565b600454945092506109b9915050565b600060025460045461095c9190610d97565b905060006001546003546109709190610d97565b90506000600154846109829190610d97565b90506000826109918386610dae565b61099b9190610dcd565b90506000816002546109ad9190610def565b97508996505050505050505b9250929050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a085760405162461bcd60e51b81526004016103de90610e07565b600080610a1484610a2c565b505091509150610a2684838386610ad8565b50505050565b600080600080600085604051602001610a459190610e52565b60408051601f19818403018152919052805160209091012080549091508060001a6001811115610a7757610a77610d29565b6effffffffffffffffffffffffffffff602083901c1696506bffffffffffffffffffffffff609883901c16955063ffffffff821694509250610ab98442610d97565b610ac39086610dae565b610acd9087610def565b955050509193509193565b600160601b82101580610aef5750600160781b8310155b15610b4d5760405162461bcd60e51b815260206004820152602860248201527f556e6578706563746564207061636b696e672069737375652064756520746f206044820152676f766572666c6f7760c01b60648201526084016103de565b6040514290600090610b63908790602001610e52565b60408051601f1981840301815291905280516020918201209150600090839087901b609887901b60f8876001811115610b9e57610b9e610d29565b901b171717909155505050505050565b80356001600160a01b03811681146106bf57600080fd5b600080600060608486031215610bda57600080fd5b610be384610bae565b9250610bf160208501610bae565b9150604084013590509250925092565b60008060008060808587031215610c1757600080fd5b610c2085610bae565b9350610c2e60208601610bae565b93969395505050506040820135916060013590565b600060208284031215610c5557600080fd5b61038d82610bae565b600080600080600060a08688031215610c7657600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b60008060408385031215610cac57600080fd5b610cb583610bae565b9150610cc360208401610bae565b90509250929050565b60008060408385031215610cdf57600080fd5b50508035926020909101359150565b60008060408385031215610d0157600080fd5b610d0a83610bae565b9150602083013560028110610d1e57600080fd5b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b84815260208101849052604081018390526080810160028310610d7257634e487b7160e01b600052602160045260246000fd5b82606083015295945050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610da957610da9610d81565b500390565b6000816000190483118215151615610dc857610dc8610d81565b500290565b600082610dea57634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115610e0257610e02610d81565b500190565b6020808252602b908201527f43616c6c6572206d75737420626520626c61737420636f6e666967757261746960408201526a1bdb8818dbdb9d1c9858dd60aa1b606082015260800190565b60609190911b6bffffffffffffffffffffffff1916815269706172616d657465727360b01b6014820152601e019056fea164736f6c634300080f000a",
}

// GasABI is the input ABI used to generate the binding from.
// Deprecated: Use GasMetaData.ABI instead.
var GasABI = GasMetaData.ABI

// GasBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasMetaData.Bin instead.
var GasBin = GasMetaData.Bin

// DeployGas deploys a new Ethereum contract, binding an instance of Gas to it.
func DeployGas(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _blastConfigurationContract common.Address, _blastFeeVault common.Address, _zeroClaimRate *big.Int, _baseGasSeconds *big.Int, _baseClaimRate *big.Int, _ceilGasSeconds *big.Int, _ceilClaimRate *big.Int) (common.Address, *types.Transaction, *Gas, error) {
	parsed, err := GasMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasBin), backend, _admin, _blastConfigurationContract, _blastFeeVault, _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gas{GasCaller: GasCaller{contract: contract}, GasTransactor: GasTransactor{contract: contract}, GasFilterer: GasFilterer{contract: contract}}, nil
}

// Gas is an auto generated Go binding around an Ethereum contract.
type Gas struct {
	GasCaller     // Read-only binding to the contract
	GasTransactor // Write-only binding to the contract
	GasFilterer   // Log filterer for contract events
}

// GasCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasSession struct {
	Contract     *Gas              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasCallerSession struct {
	Contract *GasCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasTransactorSession struct {
	Contract     *GasTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasRaw struct {
	Contract *Gas // Generic contract binding to access the raw methods on
}

// GasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasCallerRaw struct {
	Contract *GasCaller // Generic read-only contract binding to access the raw methods on
}

// GasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasTransactorRaw struct {
	Contract *GasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGas creates a new instance of Gas, bound to a specific deployed contract.
func NewGas(address common.Address, backend bind.ContractBackend) (*Gas, error) {
	contract, err := bindGas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gas{GasCaller: GasCaller{contract: contract}, GasTransactor: GasTransactor{contract: contract}, GasFilterer: GasFilterer{contract: contract}}, nil
}

// NewGasCaller creates a new read-only instance of Gas, bound to a specific deployed contract.
func NewGasCaller(address common.Address, caller bind.ContractCaller) (*GasCaller, error) {
	contract, err := bindGas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasCaller{contract: contract}, nil
}

// NewGasTransactor creates a new write-only instance of Gas, bound to a specific deployed contract.
func NewGasTransactor(address common.Address, transactor bind.ContractTransactor) (*GasTransactor, error) {
	contract, err := bindGas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasTransactor{contract: contract}, nil
}

// NewGasFilterer creates a new log filterer instance of Gas, bound to a specific deployed contract.
func NewGasFilterer(address common.Address, filterer bind.ContractFilterer) (*GasFilterer, error) {
	contract, err := bindGas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasFilterer{contract: contract}, nil
}

// bindGas binds a generic wrapper to an already deployed contract.
func bindGas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gas *GasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gas.Contract.GasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gas *GasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gas.Contract.GasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gas *GasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gas.Contract.GasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gas *GasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gas *GasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gas *GasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gas.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Gas *GasCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Gas *GasSession) Admin() (common.Address, error) {
	return _Gas.Contract.Admin(&_Gas.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Gas *GasCallerSession) Admin() (common.Address, error) {
	return _Gas.Contract.Admin(&_Gas.CallOpts)
}

// BaseClaimRate is a free data retrieval call binding the contract method 0xf5c5e343.
//
// Solidity: function baseClaimRate() view returns(uint256)
func (_Gas *GasCaller) BaseClaimRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "baseClaimRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseClaimRate is a free data retrieval call binding the contract method 0xf5c5e343.
//
// Solidity: function baseClaimRate() view returns(uint256)
func (_Gas *GasSession) BaseClaimRate() (*big.Int, error) {
	return _Gas.Contract.BaseClaimRate(&_Gas.CallOpts)
}

// BaseClaimRate is a free data retrieval call binding the contract method 0xf5c5e343.
//
// Solidity: function baseClaimRate() view returns(uint256)
func (_Gas *GasCallerSession) BaseClaimRate() (*big.Int, error) {
	return _Gas.Contract.BaseClaimRate(&_Gas.CallOpts)
}

// BaseGasSeconds is a free data retrieval call binding the contract method 0xaaa2f643.
//
// Solidity: function baseGasSeconds() view returns(uint256)
func (_Gas *GasCaller) BaseGasSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "baseGasSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseGasSeconds is a free data retrieval call binding the contract method 0xaaa2f643.
//
// Solidity: function baseGasSeconds() view returns(uint256)
func (_Gas *GasSession) BaseGasSeconds() (*big.Int, error) {
	return _Gas.Contract.BaseGasSeconds(&_Gas.CallOpts)
}

// BaseGasSeconds is a free data retrieval call binding the contract method 0xaaa2f643.
//
// Solidity: function baseGasSeconds() view returns(uint256)
func (_Gas *GasCallerSession) BaseGasSeconds() (*big.Int, error) {
	return _Gas.Contract.BaseGasSeconds(&_Gas.CallOpts)
}

// BlastConfigurationContract is a free data retrieval call binding the contract method 0x61da985e.
//
// Solidity: function blastConfigurationContract() view returns(address)
func (_Gas *GasCaller) BlastConfigurationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "blastConfigurationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlastConfigurationContract is a free data retrieval call binding the contract method 0x61da985e.
//
// Solidity: function blastConfigurationContract() view returns(address)
func (_Gas *GasSession) BlastConfigurationContract() (common.Address, error) {
	return _Gas.Contract.BlastConfigurationContract(&_Gas.CallOpts)
}

// BlastConfigurationContract is a free data retrieval call binding the contract method 0x61da985e.
//
// Solidity: function blastConfigurationContract() view returns(address)
func (_Gas *GasCallerSession) BlastConfigurationContract() (common.Address, error) {
	return _Gas.Contract.BlastConfigurationContract(&_Gas.CallOpts)
}

// BlastFeeVault is a free data retrieval call binding the contract method 0xd45e6bdf.
//
// Solidity: function blastFeeVault() view returns(address)
func (_Gas *GasCaller) BlastFeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "blastFeeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlastFeeVault is a free data retrieval call binding the contract method 0xd45e6bdf.
//
// Solidity: function blastFeeVault() view returns(address)
func (_Gas *GasSession) BlastFeeVault() (common.Address, error) {
	return _Gas.Contract.BlastFeeVault(&_Gas.CallOpts)
}

// BlastFeeVault is a free data retrieval call binding the contract method 0xd45e6bdf.
//
// Solidity: function blastFeeVault() view returns(address)
func (_Gas *GasCallerSession) BlastFeeVault() (common.Address, error) {
	return _Gas.Contract.BlastFeeVault(&_Gas.CallOpts)
}

// CeilClaimRate is a free data retrieval call binding the contract method 0xbafe8096.
//
// Solidity: function ceilClaimRate() view returns(uint256)
func (_Gas *GasCaller) CeilClaimRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "ceilClaimRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CeilClaimRate is a free data retrieval call binding the contract method 0xbafe8096.
//
// Solidity: function ceilClaimRate() view returns(uint256)
func (_Gas *GasSession) CeilClaimRate() (*big.Int, error) {
	return _Gas.Contract.CeilClaimRate(&_Gas.CallOpts)
}

// CeilClaimRate is a free data retrieval call binding the contract method 0xbafe8096.
//
// Solidity: function ceilClaimRate() view returns(uint256)
func (_Gas *GasCallerSession) CeilClaimRate() (*big.Int, error) {
	return _Gas.Contract.CeilClaimRate(&_Gas.CallOpts)
}

// CeilGasSeconds is a free data retrieval call binding the contract method 0x15011119.
//
// Solidity: function ceilGasSeconds() view returns(uint256)
func (_Gas *GasCaller) CeilGasSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "ceilGasSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CeilGasSeconds is a free data retrieval call binding the contract method 0x15011119.
//
// Solidity: function ceilGasSeconds() view returns(uint256)
func (_Gas *GasSession) CeilGasSeconds() (*big.Int, error) {
	return _Gas.Contract.CeilGasSeconds(&_Gas.CallOpts)
}

// CeilGasSeconds is a free data retrieval call binding the contract method 0x15011119.
//
// Solidity: function ceilGasSeconds() view returns(uint256)
func (_Gas *GasCallerSession) CeilGasSeconds() (*big.Int, error) {
	return _Gas.Contract.CeilGasSeconds(&_Gas.CallOpts)
}

// GetClaimRateBps is a free data retrieval call binding the contract method 0x88e5f229.
//
// Solidity: function getClaimRateBps(uint256 gasSecondsToConsume, uint256 gasToClaim) view returns(uint256, uint256)
func (_Gas *GasCaller) GetClaimRateBps(opts *bind.CallOpts, gasSecondsToConsume *big.Int, gasToClaim *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "getClaimRateBps", gasSecondsToConsume, gasToClaim)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetClaimRateBps is a free data retrieval call binding the contract method 0x88e5f229.
//
// Solidity: function getClaimRateBps(uint256 gasSecondsToConsume, uint256 gasToClaim) view returns(uint256, uint256)
func (_Gas *GasSession) GetClaimRateBps(gasSecondsToConsume *big.Int, gasToClaim *big.Int) (*big.Int, *big.Int, error) {
	return _Gas.Contract.GetClaimRateBps(&_Gas.CallOpts, gasSecondsToConsume, gasToClaim)
}

// GetClaimRateBps is a free data retrieval call binding the contract method 0x88e5f229.
//
// Solidity: function getClaimRateBps(uint256 gasSecondsToConsume, uint256 gasToClaim) view returns(uint256, uint256)
func (_Gas *GasCallerSession) GetClaimRateBps(gasSecondsToConsume *big.Int, gasToClaim *big.Int) (*big.Int, *big.Int, error) {
	return _Gas.Contract.GetClaimRateBps(&_Gas.CallOpts, gasSecondsToConsume, gasToClaim)
}

// ReadGasParams is a free data retrieval call binding the contract method 0xdde798a4.
//
// Solidity: function readGasParams(address user) view returns(uint256 etherSeconds, uint256 etherBalance, uint256 lastUpdated, uint8 mode)
func (_Gas *GasCaller) ReadGasParams(opts *bind.CallOpts, user common.Address) (struct {
	EtherSeconds *big.Int
	EtherBalance *big.Int
	LastUpdated  *big.Int
	Mode         uint8
}, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "readGasParams", user)

	outstruct := new(struct {
		EtherSeconds *big.Int
		EtherBalance *big.Int
		LastUpdated  *big.Int
		Mode         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EtherSeconds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EtherBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastUpdated = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Mode = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// ReadGasParams is a free data retrieval call binding the contract method 0xdde798a4.
//
// Solidity: function readGasParams(address user) view returns(uint256 etherSeconds, uint256 etherBalance, uint256 lastUpdated, uint8 mode)
func (_Gas *GasSession) ReadGasParams(user common.Address) (struct {
	EtherSeconds *big.Int
	EtherBalance *big.Int
	LastUpdated  *big.Int
	Mode         uint8
}, error) {
	return _Gas.Contract.ReadGasParams(&_Gas.CallOpts, user)
}

// ReadGasParams is a free data retrieval call binding the contract method 0xdde798a4.
//
// Solidity: function readGasParams(address user) view returns(uint256 etherSeconds, uint256 etherBalance, uint256 lastUpdated, uint8 mode)
func (_Gas *GasCallerSession) ReadGasParams(user common.Address) (struct {
	EtherSeconds *big.Int
	EtherBalance *big.Int
	LastUpdated  *big.Int
	Mode         uint8
}, error) {
	return _Gas.Contract.ReadGasParams(&_Gas.CallOpts, user)
}

// ZeroClaimRate is a free data retrieval call binding the contract method 0x5a03838f.
//
// Solidity: function zeroClaimRate() view returns(uint256)
func (_Gas *GasCaller) ZeroClaimRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "zeroClaimRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZeroClaimRate is a free data retrieval call binding the contract method 0x5a03838f.
//
// Solidity: function zeroClaimRate() view returns(uint256)
func (_Gas *GasSession) ZeroClaimRate() (*big.Int, error) {
	return _Gas.Contract.ZeroClaimRate(&_Gas.CallOpts)
}

// ZeroClaimRate is a free data retrieval call binding the contract method 0x5a03838f.
//
// Solidity: function zeroClaimRate() view returns(uint256)
func (_Gas *GasCallerSession) ZeroClaimRate() (*big.Int, error) {
	return _Gas.Contract.ZeroClaimRate(&_Gas.CallOpts)
}

// AdminClaimGas is a paid mutator transaction binding the contract method 0x2d7a59e1.
//
// Solidity: function adminClaimGas(address contractAddress) returns(uint256)
func (_Gas *GasTransactor) AdminClaimGas(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "adminClaimGas", contractAddress)
}

// AdminClaimGas is a paid mutator transaction binding the contract method 0x2d7a59e1.
//
// Solidity: function adminClaimGas(address contractAddress) returns(uint256)
func (_Gas *GasSession) AdminClaimGas(contractAddress common.Address) (*types.Transaction, error) {
	return _Gas.Contract.AdminClaimGas(&_Gas.TransactOpts, contractAddress)
}

// AdminClaimGas is a paid mutator transaction binding the contract method 0x2d7a59e1.
//
// Solidity: function adminClaimGas(address contractAddress) returns(uint256)
func (_Gas *GasTransactorSession) AdminClaimGas(contractAddress common.Address) (*types.Transaction, error) {
	return _Gas.Contract.AdminClaimGas(&_Gas.TransactOpts, contractAddress)
}

// Claim is a paid mutator transaction binding the contract method 0x26af4832.
//
// Solidity: function claim(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) returns(uint256)
func (_Gas *GasTransactor) Claim(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address, gasToClaim *big.Int, gasSecondsToConsume *big.Int) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "claim", contractAddress, recipientOfGas, gasToClaim, gasSecondsToConsume)
}

// Claim is a paid mutator transaction binding the contract method 0x26af4832.
//
// Solidity: function claim(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) returns(uint256)
func (_Gas *GasSession) Claim(contractAddress common.Address, recipientOfGas common.Address, gasToClaim *big.Int, gasSecondsToConsume *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.Claim(&_Gas.TransactOpts, contractAddress, recipientOfGas, gasToClaim, gasSecondsToConsume)
}

// Claim is a paid mutator transaction binding the contract method 0x26af4832.
//
// Solidity: function claim(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) returns(uint256)
func (_Gas *GasTransactorSession) Claim(contractAddress common.Address, recipientOfGas common.Address, gasToClaim *big.Int, gasSecondsToConsume *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.Claim(&_Gas.TransactOpts, contractAddress, recipientOfGas, gasToClaim, gasSecondsToConsume)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x5767bba5.
//
// Solidity: function claimAll(address contractAddress, address recipientOfGas) returns(uint256)
func (_Gas *GasTransactor) ClaimAll(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "claimAll", contractAddress, recipientOfGas)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x5767bba5.
//
// Solidity: function claimAll(address contractAddress, address recipientOfGas) returns(uint256)
func (_Gas *GasSession) ClaimAll(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Gas.Contract.ClaimAll(&_Gas.TransactOpts, contractAddress, recipientOfGas)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x5767bba5.
//
// Solidity: function claimAll(address contractAddress, address recipientOfGas) returns(uint256)
func (_Gas *GasTransactorSession) ClaimAll(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Gas.Contract.ClaimAll(&_Gas.TransactOpts, contractAddress, recipientOfGas)
}

// ClaimGasAtMinClaimRate is a paid mutator transaction binding the contract method 0x0951888f.
//
// Solidity: function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) returns(uint256)
func (_Gas *GasTransactor) ClaimGasAtMinClaimRate(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address, minClaimRateBips *big.Int) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "claimGasAtMinClaimRate", contractAddress, recipientOfGas, minClaimRateBips)
}

// ClaimGasAtMinClaimRate is a paid mutator transaction binding the contract method 0x0951888f.
//
// Solidity: function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) returns(uint256)
func (_Gas *GasSession) ClaimGasAtMinClaimRate(contractAddress common.Address, recipientOfGas common.Address, minClaimRateBips *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.ClaimGasAtMinClaimRate(&_Gas.TransactOpts, contractAddress, recipientOfGas, minClaimRateBips)
}

// ClaimGasAtMinClaimRate is a paid mutator transaction binding the contract method 0x0951888f.
//
// Solidity: function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) returns(uint256)
func (_Gas *GasTransactorSession) ClaimGasAtMinClaimRate(contractAddress common.Address, recipientOfGas common.Address, minClaimRateBips *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.ClaimGasAtMinClaimRate(&_Gas.TransactOpts, contractAddress, recipientOfGas, minClaimRateBips)
}

// ClaimMax is a paid mutator transaction binding the contract method 0x49458868.
//
// Solidity: function claimMax(address contractAddress, address recipientOfGas) returns(uint256)
func (_Gas *GasTransactor) ClaimMax(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "claimMax", contractAddress, recipientOfGas)
}

// ClaimMax is a paid mutator transaction binding the contract method 0x49458868.
//
// Solidity: function claimMax(address contractAddress, address recipientOfGas) returns(uint256)
func (_Gas *GasSession) ClaimMax(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Gas.Contract.ClaimMax(&_Gas.TransactOpts, contractAddress, recipientOfGas)
}

// ClaimMax is a paid mutator transaction binding the contract method 0x49458868.
//
// Solidity: function claimMax(address contractAddress, address recipientOfGas) returns(uint256)
func (_Gas *GasTransactorSession) ClaimMax(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Gas.Contract.ClaimMax(&_Gas.TransactOpts, contractAddress, recipientOfGas)
}

// SetGasMode is a paid mutator transaction binding the contract method 0xd4810ba5.
//
// Solidity: function setGasMode(address contractAddress, uint8 mode) returns()
func (_Gas *GasTransactor) SetGasMode(opts *bind.TransactOpts, contractAddress common.Address, mode uint8) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "setGasMode", contractAddress, mode)
}

// SetGasMode is a paid mutator transaction binding the contract method 0xd4810ba5.
//
// Solidity: function setGasMode(address contractAddress, uint8 mode) returns()
func (_Gas *GasSession) SetGasMode(contractAddress common.Address, mode uint8) (*types.Transaction, error) {
	return _Gas.Contract.SetGasMode(&_Gas.TransactOpts, contractAddress, mode)
}

// SetGasMode is a paid mutator transaction binding the contract method 0xd4810ba5.
//
// Solidity: function setGasMode(address contractAddress, uint8 mode) returns()
func (_Gas *GasTransactorSession) SetGasMode(contractAddress common.Address, mode uint8) (*types.Transaction, error) {
	return _Gas.Contract.SetGasMode(&_Gas.TransactOpts, contractAddress, mode)
}

// UpdateAdminParameters is a paid mutator transaction binding the contract method 0x445067cb.
//
// Solidity: function updateAdminParameters(uint256 _zeroClaimRate, uint256 _baseGasSeconds, uint256 _baseClaimRate, uint256 _ceilGasSeconds, uint256 _ceilClaimRate) returns()
func (_Gas *GasTransactor) UpdateAdminParameters(opts *bind.TransactOpts, _zeroClaimRate *big.Int, _baseGasSeconds *big.Int, _baseClaimRate *big.Int, _ceilGasSeconds *big.Int, _ceilClaimRate *big.Int) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "updateAdminParameters", _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate)
}

// UpdateAdminParameters is a paid mutator transaction binding the contract method 0x445067cb.
//
// Solidity: function updateAdminParameters(uint256 _zeroClaimRate, uint256 _baseGasSeconds, uint256 _baseClaimRate, uint256 _ceilGasSeconds, uint256 _ceilClaimRate) returns()
func (_Gas *GasSession) UpdateAdminParameters(_zeroClaimRate *big.Int, _baseGasSeconds *big.Int, _baseClaimRate *big.Int, _ceilGasSeconds *big.Int, _ceilClaimRate *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.UpdateAdminParameters(&_Gas.TransactOpts, _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate)
}

// UpdateAdminParameters is a paid mutator transaction binding the contract method 0x445067cb.
//
// Solidity: function updateAdminParameters(uint256 _zeroClaimRate, uint256 _baseGasSeconds, uint256 _baseClaimRate, uint256 _ceilGasSeconds, uint256 _ceilClaimRate) returns()
func (_Gas *GasTransactorSession) UpdateAdminParameters(_zeroClaimRate *big.Int, _baseGasSeconds *big.Int, _baseClaimRate *big.Int, _ceilGasSeconds *big.Int, _ceilClaimRate *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.UpdateAdminParameters(&_Gas.TransactOpts, _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate)
}
