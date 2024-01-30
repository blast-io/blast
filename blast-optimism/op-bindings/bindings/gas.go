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
	Bin: "0x60e06040523480156200001157600080fd5b50604051620016d2380380620016d2833981016040819052620000349162000285565b8285106200009b5760405162461bcd60e51b815260206004820152602960248201527f7a65726f20636c61696d2072617465206d757374206265203c206261736520636044820152686c61696d207261746560b81b60648201526084015b60405180910390fd5b808310620000fe5760405162461bcd60e51b815260206004820152602960248201527f6261736520636c61696d2072617465206d757374206265203c206365696c20636044820152686c61696d207261746560b81b606482015260840162000092565b818410620001635760405162461bcd60e51b815260206004820152602b60248201527f6261736520676173207365636f6e6473206d757374206265203c206365696c2060448201526a676173207365636f6e647360a81b606482015260840162000092565b60008411620001b55760405162461bcd60e51b815260206004820152601c60248201527f6261736520676173207365636f6e6473206d757374206265203e203000000000604482015260640162000092565b6127108111156200022f5760405162461bcd60e51b815260206004820152603960248201527f6365696c20636c61696d2072617465206d757374206265206c6573732074686160448201527f6e206f7220657175616c20746f2031305f303030206269707300000000000000606482015260840162000092565b6001600160a01b0397881660805295871660a0529390951660c05260009190915560015560029290925560039190915560045562000300565b80516001600160a01b03811681146200028057600080fd5b919050565b600080600080600080600080610100898b031215620002a357600080fd5b620002ae8962000268565b9750620002be60208a0162000268565b9650620002ce60408a0162000268565b9550606089015194506080890151935060a0890151925060c0890151915060e089015190509295985092959890939650565b60805160a05160c0516113786200035a60003960008181610244015281816106ce01526107c60152600081816101be015281816104600152610c670152600081816102aa0152818161071c015261080b01526113786000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806361da985e116100b2578063d45e6bdf11610081578063dde798a411610066578063dde798a414610279578063f5c5e3431461029c578063f851a440146102a557600080fd5b8063d45e6bdf1461023f578063d4810ba51461026657600080fd5b806361da985e146101b957806388e5f22914610205578063aaa2f6431461022d578063bafe80961461023657600080fd5b8063445067cb116100ee578063445067cb14610175578063494588681461018a5780635767bba51461019d5780635a03838f146101b057600080fd5b80630951888f14610120578063150111191461014657806326af48321461014f5780632d7a59e114610162575b600080fd5b61013361012e3660046110a7565b6102cc565b6040519081526020015b60405180910390f35b61013360035481565b61013361015d3660046110e3565b610446565b610133610170366004611125565b610702565b610188610183366004611140565b6107f3565b005b61013361019836600461117b565b610b52565b6101336101ab36600461117b565b610b61565b61013360005481565b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013d565b6102186102133660046111ae565b610b8a565b6040805192835260208301919091520161013d565b61013360015481565b61013360045481565b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b6101886102743660046111d0565b610c4f565b61028c610287366004611125565b610d38565b60405161013d949392919061123a565b61013360025481565b6101e07f000000000000000000000000000000000000000000000000000000000000000081565b6000600454821115610365576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f6465736972656420636c61696d20726174652065786365656473206d6178696d60448201527f756d00000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60008061037186610d38565b50509150915060005484116103935761038a8686610b61565b9250505061043f565b6002548410156103a35760025493505b6000600254856103b391906112c4565b905060006001546003546103c791906112c4565b905060006002546004546103db91906112c4565b905060006103f26103ec84866112db565b83610e53565b6001546103ff9190611318565b9050600061040d8288611330565b90508581111561041a5750845b600061042683836112db565b90506104348c8c8484610446565b985050505050505050505b9392505050565b60003373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461050d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f43616c6c6572206d75737420626520626c61737420636f6e666967757261746960448201527f6f6e20636f6e7472616374000000000000000000000000000000000000000000606482015260840161035c565b600080600061051b88610d38565b935050925092506000861161058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6d757374207769746864726177206e6f6e2d7a65726f20616d6f756e74000000604482015260640161035c565b818611156105f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f746f6f206d75636820746f207769746864726177000000000000000000000000604482015260640161035c565b82851115610660576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f6e6f7420656e6f75676820676173207365636f6e647300000000000000000000604482015260640161035c565b60008061066d8789610b8a565b90925090506000612710610681848b6112db565b61068b9190611330565b90506000610699828b6112c4565b90506106b98c6106a9858a6112c4565b6106b38d8a6112c4565b88610e8a565b6106c38b83611009565b80156106f3576106f37f000000000000000000000000000000000000000000000000000000000000000082611009565b509a9950505050505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146107a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207468652061646d696e000000000000000000604482015260640161035c565b60006107ae83610d38565b50509150506107c1836000806000610e8a565b6107eb7f000000000000000000000000000000000000000000000000000000000000000082611009565b90505b919050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610892576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207468652061646d696e000000000000000000604482015260640161035c565b828510610921576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f7a65726f20636c61696d2072617465206d757374206265203c2062617365206360448201527f6c61696d20726174650000000000000000000000000000000000000000000000606482015260840161035c565b8083106109b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f6261736520636c61696d2072617465206d757374206265203c206365696c206360448201527f6c61696d20726174650000000000000000000000000000000000000000000000606482015260840161035c565b818410610a3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f6261736520676173207365636f6e6473206d757374206265203c206365696c2060448201527f676173207365636f6e6473000000000000000000000000000000000000000000606482015260840161035c565b60008411610aa9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f6261736520676173207365636f6e6473206d757374206265203e203000000000604482015260640161035c565b612710811115610b3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f6365696c20636c61696d2072617465206d757374206265206c6573732074686160448201527f6e206f7220657175616c20746f2031305f303030206269707300000000000000606482015260840161035c565b600094909455600192909255600255600355600455565b600061043f83836004546102cc565b6000806000610b6f85610d38565b505091509150610b8185858385610446565b95945050505050565b60008080610b988486611330565b9050600154811015610bb1575050600080549150610c48565b6003548110610bd957600060035485610bca91906112db565b60045494509250610c48915050565b6000600254600454610beb91906112c4565b90506000600154600354610bff91906112c4565b9050600060015484610c1191906112c4565b9050600082610c2083866112db565b610c2a9190611330565b9050600081600254610c3c9190611318565b97508996505050505050505b9250929050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f43616c6c6572206d75737420626520626c61737420636f6e666967757261746960448201527f6f6e20636f6e7472616374000000000000000000000000000000000000000000606482015260840161035c565b600080610d2084610d38565b505091509150610d3284838386610e8a565b50505050565b600080600080600085604051602001610da2919060609190911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001681527f706172616d6574657273000000000000000000000000000000000000000000006014820152601e0190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012080549091508060001a6001811115610df257610df261120b565b6effffffffffffffffffffffffffffff602083901c1696506bffffffffffffffffffffffff609883901c16955063ffffffff821694509250610e3484426112c4565b610e3e90866112db565b610e489087611318565b955050509193509193565b60008215610e815781610e676001856112c4565b610e719190611330565b610e7c906001611318565b61043f565b50600092915050565b6c0100000000000000000000000082101580610eb657506f010000000000000000000000000000008310155b15610f43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f556e6578706563746564207061636b696e672069737375652064756520746f2060448201527f6f766572666c6f77000000000000000000000000000000000000000000000000606482015260840161035c565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b1660208201527f706172616d65746572730000000000000000000000000000000000000000000060348201524290600090603e01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020918201209150600090839087901b609887901b60f8876001811115610ff957610ff961120b565b901b171717909155505050505050565b600080600080600085875af190508061107e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4554485f5452414e534645525f4641494c454400000000000000000000000000604482015260640161035c565b505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107ee57600080fd5b6000806000606084860312156110bc57600080fd5b6110c584611083565b92506110d360208501611083565b9150604084013590509250925092565b600080600080608085870312156110f957600080fd5b61110285611083565b935061111060208601611083565b93969395505050506040820135916060013590565b60006020828403121561113757600080fd5b61043f82611083565b600080600080600060a0868803121561115857600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b6000806040838503121561118e57600080fd5b61119783611083565b91506111a560208401611083565b90509250929050565b600080604083850312156111c157600080fd5b50508035926020909101359150565b600080604083850312156111e357600080fd5b6111ec83611083565b915060208301356002811061120057600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b84815260208101849052604081018390526080810160028310611286577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82606083015295945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156112d6576112d6611295565b500390565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561131357611313611295565b500290565b6000821982111561132b5761132b611295565b500190565b600082611366577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c634300080f000a",
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
