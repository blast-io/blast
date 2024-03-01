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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blastConfigurationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blastFeeVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"adminClaimGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseGasSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blastConfigurationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blastFeeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ceilClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ceilGasSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasToClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasSecondsToConsume\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minClaimRateBips\",\"type\":\"uint256\"}],\"name\":\"claimGasAtMinClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasSecondsToConsume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasToClaim\",\"type\":\"uint256\"}],\"name\":\"getClaimRateBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_zeroClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilClaimRate\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"readGasParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"etherSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"etherBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"enumGasMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumGasMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"name\":\"setGasMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_zeroClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseClaimRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilGasSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ceilClaimRate\",\"type\":\"uint256\"}],\"name\":\"updateAdminParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeroClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b506040516200169338038062001693833981016040819052620000359162000150565b6001608052600060a081905260c0526001600160a01b0380841660e052828116610100528116610120526200006962000072565b5050506200019a565b600054610100900460ff1615620000df5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000131576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200014b57600080fd5b919050565b6000806000606084860312156200016657600080fd5b620001718462000133565b9250620001816020850162000133565b9150620001916040850162000133565b90509250925092565b60805160a05160c05160e051610100516101205161147f6200021460003960008181610252015281816105df01526106aa0152600081816101d90152818161043f01526109db0152600081816102cb0152818161062001526106e2015260006108970152600061086e01526000610845015261147f6000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806361da985e116100ad578063d4810ba511610071578063d4810ba514610274578063dcbab60814610287578063dde798a41461029a578063f5c5e343146102bd578063f851a440146102c657600080fd5b806361da985e146101d457806388e5f22914610213578063aaa2f6431461023b578063bafe809614610244578063d45e6bdf1461024d57600080fd5b8063445067cb116100f4578063445067cb1461017b578063494588681461019057806354fd4d50146101a35780635767bba5146101b85780635a03838f146101cb57600080fd5b80630951888f14610126578063150111191461014c57806326af4832146101555780632d7a59e114610168575b600080fd5b610139610134366004610f61565b6102ed565b6040519081526020015b60405180910390f35b61013960045481565b610139610163366004610f9d565b610432565b610139610176366004610fdf565b610613565b61018e610189366004610ffa565b6106d7565b005b61013961019e366004611035565b61082f565b6101ab61083e565b6040516101439190611094565b6101396101c6366004611035565b6108e1565b61013960015481565b6101fb7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610143565b6102266102213660046110c7565b61090a565b60408051928352602083019190915201610143565b61013960025481565b61013960055481565b6101fb7f000000000000000000000000000000000000000000000000000000000000000081565b61018e6102823660046110e9565b6109d0565b61018e610295366004610ffa565b610a3c565b6102ad6102a8366004610fdf565b610c32565b604051610143949392919061113a565b61013960035481565b6101fb7f000000000000000000000000000000000000000000000000000000000000000081565b60006005548211156103515760405162461bcd60e51b815260206004820152602260248201527f6465736972656420636c61696d20726174652065786365656473206d6178696d604482015261756d60f01b60648201526084015b60405180910390fd5b60008061035d86610c32565b505091509150600154841161037f5761037686866108e1565b9250505061042b565b60035484101561038f5760035493505b60006003548561039f9190611192565b905060006002546004546103b39190611192565b905060006003546005546103c79190611192565b905060006103de6103d884866111a9565b83610cde565b6002546103eb91906111c8565b905060006103f982886111f6565b9050858111156104065750845b600061041283836111a9565b90506104208c8c8484610432565b985050505050505050505b9392505050565b6000336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461047c5760405162461bcd60e51b81526004016103489061120a565b600080600061048a88610c32565b93505092509250600086116104e15760405162461bcd60e51b815260206004820152601d60248201527f6d757374207769746864726177206e6f6e2d7a65726f20616d6f756e740000006044820152606401610348565b818611156105285760405162461bcd60e51b8152602060048201526014602482015273746f6f206d75636820746f20776974686472617760601b6044820152606401610348565b828511156105715760405162461bcd60e51b81526020600482015260166024820152756e6f7420656e6f75676820676173207365636f6e647360501b6044820152606401610348565b60008061057e878961090a565b90925090506000612710610592848b6111a9565b61059c91906111f6565b905060006105aa828b611192565b90506105ca8c6105ba858a611192565b6105c48d8a611192565b88610d15565b6105d48b83610deb565b8015610604576106047f000000000000000000000000000000000000000000000000000000000000000082610deb565b509a9950505050505050505050565b6000336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106875760405162461bcd60e51b815260206004820152601760248201527621b0b63632b91034b9903737ba103a34329030b236b4b760491b6044820152606401610348565b600061069283610c32565b50509150506106a5836000806000610d15565b6106cf7f000000000000000000000000000000000000000000000000000000000000000082610deb565b90505b919050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107495760405162461bcd60e51b815260206004820152601760248201527621b0b63632b91034b9903737ba103a34329030b236b4b760491b6044820152606401610348565b8285106107685760405162461bcd60e51b815260040161034890611255565b8083106107875760405162461bcd60e51b81526004016103489061129e565b8184106107a65760405162461bcd60e51b8152600401610348906112e7565b600084116107f65760405162461bcd60e51b815260206004820152601c60248201527f6261736520676173207365636f6e6473206d757374206265203e2030000000006044820152606401610348565b6127108111156108185760405162461bcd60e51b815260040161034890611332565b600194909455600292909255600355600455600555565b600061042b83836005546102ed565b60606108697f0000000000000000000000000000000000000000000000000000000000000000610e41565b6108927f0000000000000000000000000000000000000000000000000000000000000000610e41565b6108bb7f0000000000000000000000000000000000000000000000000000000000000000610e41565b6040516020016108cd9392919061138f565b604051602081830303815290604052905090565b60008060006108ef85610c32565b50509150915061090185858385610432565b95945050505050565b6000808061091884866111f6565b9050600254811015610932575050600154905060006109c9565b600454811061095a5760006004548561094b91906111a9565b600554945092506109c9915050565b600060035460055461096c9190611192565b905060006002546004546109809190611192565b90506000600254846109929190611192565b90506000826109a183866111a9565b6109ab91906111f6565b90506000816003546109bd91906111c8565b97508996505050505050505b9250929050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a185760405162461bcd60e51b81526004016103489061120a565b600080610a2484610c32565b505091509150610a3684838386610d15565b50505050565b600054610100900460ff1615808015610a5c5750600054600160ff909116105b80610a765750303b158015610a76575060005460ff166001145b610ad95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610348565b6000805460ff191660011790558015610afc576000805461ff0019166101001790555b838610610b1b5760405162461bcd60e51b815260040161034890611255565b818410610b3a5760405162461bcd60e51b81526004016103489061129e565b828510610b595760405162461bcd60e51b8152600401610348906112e7565b60008511610ba95760405162461bcd60e51b815260206004820152601c60248201527f6261736520676173207365636f6e6473206d757374206265203e2030000000006044820152606401610348565b612710821115610bcb5760405162461bcd60e51b815260040161034890611332565b600186905560028590556003849055600483905560058290558015610c2a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b600080600080600085604051602001610c4b91906113e9565b60408051601f19818403018152919052805160209091012080549091508060001a6001811115610c7d57610c7d611124565b6effffffffffffffffffffffffffffff602083901c1696506bffffffffffffffffffffffff609883901c16955063ffffffff821694509250610cbf8442611192565b610cc990866111a9565b610cd390876111c8565b955050509193509193565b60008215610d0c5781610cf2600185611192565b610cfc91906111f6565b610d079060016111c8565b61042b565b50600092915050565b600160601b82101580610d2c5750600160781b8310155b15610d8a5760405162461bcd60e51b815260206004820152602860248201527f556e6578706563746564207061636b696e672069737375652064756520746f206044820152676f766572666c6f7760c01b6064820152608401610348565b6040514290600090610da09087906020016113e9565b60408051601f1981840301815291905280516020918201209150600090839087901b609887901b60f8876001811115610ddb57610ddb611124565b901b171717909155505050505050565b600080600080600085875af1905080610e3c5760405162461bcd60e51b815260206004820152601360248201527211551217d514905394d1915497d19052531151606a1b6044820152606401610348565b505050565b606081600003610e685750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610e925780610e7c8161142f565b9150610e8b9050600a836111f6565b9150610e6c565b60008167ffffffffffffffff811115610ead57610ead611448565b6040519080825280601f01601f191660200182016040528015610ed7576020820181803683370190505b5090505b8415610f4257610eec600183611192565b9150610ef9600a8661145e565b610f049060306111c8565b60f81b818381518110610f1957610f19611419565b60200101906001600160f81b031916908160001a905350610f3b600a866111f6565b9450610edb565b949350505050565b80356001600160a01b03811681146106d257600080fd5b600080600060608486031215610f7657600080fd5b610f7f84610f4a565b9250610f8d60208501610f4a565b9150604084013590509250925092565b60008060008060808587031215610fb357600080fd5b610fbc85610f4a565b9350610fca60208601610f4a565b93969395505050506040820135916060013590565b600060208284031215610ff157600080fd5b61042b82610f4a565b600080600080600060a0868803121561101257600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b6000806040838503121561104857600080fd5b61105183610f4a565b915061105f60208401610f4a565b90509250929050565b60005b8381101561108357818101518382015260200161106b565b83811115610a365750506000910152565b60208152600082518060208401526110b3816040850160208701611068565b601f01601f19169190910160400192915050565b600080604083850312156110da57600080fd5b50508035926020909101359150565b600080604083850312156110fc57600080fd5b61110583610f4a565b915060208301356002811061111957600080fd5b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b8481526020810184905260408101839052608081016002831061116d57634e487b7160e01b600052602160045260246000fd5b82606083015295945050505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156111a4576111a461117c565b500390565b60008160001904831182151516156111c3576111c361117c565b500290565b600082198211156111db576111db61117c565b500190565b634e487b7160e01b600052601260045260246000fd5b600082611205576112056111e0565b500490565b6020808252602b908201527f43616c6c6572206d75737420626520626c61737420636f6e666967757261746960408201526a1bdb8818dbdb9d1c9858dd60aa1b606082015260800190565b60208082526029908201527f7a65726f20636c61696d2072617465206d757374206265203c206261736520636040820152686c61696d207261746560b81b606082015260800190565b60208082526029908201527f6261736520636c61696d2072617465206d757374206265203c206365696c20636040820152686c61696d207261746560b81b606082015260800190565b6020808252602b908201527f6261736520676173207365636f6e6473206d757374206265203c206365696c2060408201526a676173207365636f6e647360a81b606082015260800190565b60208082526039908201527f6365696c20636c61696d2072617465206d757374206265206c6573732074686160408201527f6e206f7220657175616c20746f2031305f303030206269707300000000000000606082015260800190565b600084516113a1818460208901611068565b8083019050601760f91b80825285516113c1816001850160208a01611068565b600192019182015283516113dc816002840160208801611068565b0160020195945050505050565b60609190911b6bffffffffffffffffffffffff1916815269706172616d657465727360b01b6014820152601e0190565b634e487b7160e01b600052603260045260246000fd5b6000600182016114415761144161117c565b5060010190565b634e487b7160e01b600052604160045260246000fd5b60008261146d5761146d6111e0565b50069056fea164736f6c634300080f000a",
}

// GasABI is the input ABI used to generate the binding from.
// Deprecated: Use GasMetaData.ABI instead.
var GasABI = GasMetaData.ABI

// GasBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasMetaData.Bin instead.
var GasBin = GasMetaData.Bin

// DeployGas deploys a new Ethereum contract, binding an instance of Gas to it.
func DeployGas(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _blastConfigurationContract common.Address, _blastFeeVault common.Address) (common.Address, *types.Transaction, *Gas, error) {
	parsed, err := GasMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasBin), backend, _admin, _blastConfigurationContract, _blastFeeVault)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Gas *GasCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gas.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Gas *GasSession) Version() (string, error) {
	return _Gas.Contract.Version(&_Gas.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Gas *GasCallerSession) Version() (string, error) {
	return _Gas.Contract.Version(&_Gas.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _zeroClaimRate, uint256 _baseGasSeconds, uint256 _baseClaimRate, uint256 _ceilGasSeconds, uint256 _ceilClaimRate) returns()
func (_Gas *GasTransactor) Initialize(opts *bind.TransactOpts, _zeroClaimRate *big.Int, _baseGasSeconds *big.Int, _baseClaimRate *big.Int, _ceilGasSeconds *big.Int, _ceilClaimRate *big.Int) (*types.Transaction, error) {
	return _Gas.contract.Transact(opts, "initialize", _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _zeroClaimRate, uint256 _baseGasSeconds, uint256 _baseClaimRate, uint256 _ceilGasSeconds, uint256 _ceilClaimRate) returns()
func (_Gas *GasSession) Initialize(_zeroClaimRate *big.Int, _baseGasSeconds *big.Int, _baseClaimRate *big.Int, _ceilGasSeconds *big.Int, _ceilClaimRate *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.Initialize(&_Gas.TransactOpts, _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _zeroClaimRate, uint256 _baseGasSeconds, uint256 _baseClaimRate, uint256 _ceilGasSeconds, uint256 _ceilClaimRate) returns()
func (_Gas *GasTransactorSession) Initialize(_zeroClaimRate *big.Int, _baseGasSeconds *big.Int, _baseClaimRate *big.Int, _ceilGasSeconds *big.Int, _ceilClaimRate *big.Int) (*types.Transaction, error) {
	return _Gas.Contract.Initialize(&_Gas.TransactOpts, _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate)
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

// GasInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Gas contract.
type GasInitializedIterator struct {
	Event *GasInitialized // Event containing the contract specifics and raw log

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
func (it *GasInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasInitialized)
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
		it.Event = new(GasInitialized)
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
func (it *GasInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasInitialized represents a Initialized event raised by the Gas contract.
type GasInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gas *GasFilterer) FilterInitialized(opts *bind.FilterOpts) (*GasInitializedIterator, error) {

	logs, sub, err := _Gas.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GasInitializedIterator{contract: _Gas.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gas *GasFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GasInitialized) (event.Subscription, error) {

	logs, sub, err := _Gas.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasInitialized)
				if err := _Gas.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Gas *GasFilterer) ParseInitialized(log types.Log) (*GasInitialized, error) {
	event := new(GasInitialized)
	if err := _Gas.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
