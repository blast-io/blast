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

// ProxyAdminMetaData contains all meta data concerning the ProxyAdmin contract.
var ProxyAdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"contractAddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"implementationName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUpgrading\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxyType\",\"outputs\":[{\"internalType\":\"enumProxyAdmin.ProxyType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAddressManager\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setImplementationName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"enumProxyAdmin.ProxyType\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"setProxyType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_upgrading\",\"type\":\"bool\"}],\"name\":\"setUpgrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516115da3803806115da83398101604081905261002f91610097565b61003833610047565b61004181610047565b506100c7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100a957600080fd5b81516001600160a01b03811681146100c057600080fd5b9392505050565b611504806100d66000396000f3fe6080604052600436106101095760003560e01c8063860f7cda1161009557806399a88ec41161006457806399a88ec4146102d05780639b2ea4bd146102f0578063b794726214610310578063f2fde38b1461033a578063f3b7dead1461035a57600080fd5b8063860f7cda1461025f5780638d52d4a01461027f5780638da5cb5b1461029f5780639623609d146102bd57600080fd5b8063238181ae116100dc578063238181ae146101a05780633ab76e9f146101cd5780636bd9f516146101ed578063715018a61461022a5780637eff275e1461023f57600080fd5b80630652b57a1461010e57806307c8f7b0146101305780631b8b921d14610150578063204e1c7a14610163575b600080fd5b34801561011a57600080fd5b5061012e610129366004610f05565b61037a565b005b34801561013c57600080fd5b5061012e61014b366004610f22565b6103a4565b61012e61015e366004611009565b6103ca565b34801561016f57600080fd5b5061018361017e366004610f05565b61048a565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101ac57600080fd5b506101c06101bb366004610f05565b610644565b60405161019791906110b1565b3480156101d957600080fd5b50600354610183906001600160a01b031681565b3480156101f957600080fd5b5061021d610208366004610f05565b60016020526000908152604090205460ff1681565b60405161019791906110da565b34801561023657600080fd5b5061012e6106de565b34801561024b57600080fd5b5061012e61025a366004611102565b6106f2565b34801561026b57600080fd5b5061012e61027a366004611009565b610821565b34801561028b57600080fd5b5061012e61029a36600461113b565b61084b565b3480156102ab57600080fd5b506000546001600160a01b0316610183565b61012e6102cb36600461116d565b610894565b3480156102dc57600080fd5b5061012e6102eb366004611102565b610a24565b3480156102fc57600080fd5b5061012e61030b3660046111cf565b610c28565b34801561031c57600080fd5b50600354600160a01b900460ff166040519015158152602001610197565b34801561034657600080fd5b5061012e610355366004610f05565b610c98565b34801561036657600080fd5b50610183610375366004610f05565b610d11565b610382610e46565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6103ac610e46565b60038054911515600160a01b0260ff60a01b19909216919091179055565b6103d2610e46565b6000826001600160a01b031634836040516103ed9190611216565b60006040518083038185875af1925050503d806000811461042a576040519150601f19603f3d011682016040523d82523d6000602084013e61042f565b606091505b50509050806104855760405162461bcd60e51b815260206004820181905260248201527f50726f787941646d696e3a2063616c6c20746f2070726f7879206661696c656460448201526064015b60405180910390fd5b505050565b6001600160a01b03811660009081526001602052604081205460ff16818160028111156104b9576104b96110c4565b0361052757826001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105209190611232565b9392505050565b600181600281111561053b5761053b6110c4565b0361057e57826001600160a01b031663aaf10f426040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104fc573d6000803e3d6000fd5b6002816002811115610592576105926110c4565b036105f6576003546001600160a01b0384811660009081526002602052604090819020905163bf40fac160e01b8152919092169163bf40fac1916105d99190600401611283565b602060405180830381865afa1580156104fc573d6000803e3d6000fd5b60405162461bcd60e51b815260206004820152601e60248201527f50726f787941646d696e3a20756e6b6e6f776e2070726f787920747970650000604482015260640161047c565b50919050565b6002602052600090815260409020805461065d9061124f565b80601f01602080910402602001604051908101604052809291908181526020018280546106899061124f565b80156106d65780601f106106ab576101008083540402835291602001916106d6565b820191906000526020600020905b8154815290600101906020018083116106b957829003601f168201915b505050505081565b6106e6610e46565b6106f06000610ea0565b565b6106fa610e46565b6001600160a01b03821660009081526001602052604081205460ff1690816002811115610729576107296110c4565b0361078f576040516308f2839760e41b81526001600160a01b038381166004830152841690638f283970906024015b600060405180830381600087803b15801561077257600080fd5b505af1158015610786573d6000803e3d6000fd5b50505050505050565b60018160028111156107a3576107a36110c4565b036107d6576040516313af403560e01b81526001600160a01b0383811660048301528416906313af403590602401610758565b60028160028111156107ea576107ea6110c4565b036105f65760035460405163f2fde38b60e01b81526001600160a01b0384811660048301529091169063f2fde38b90602401610758565b610829610e46565b6001600160a01b03821660009081526002602052604090206104858282611354565b610853610e46565b6001600160a01b03821660009081526001602081905260409091208054839260ff199091169083600281111561088b5761088b6110c4565b02179055505050565b61089c610e46565b6001600160a01b03831660009081526001602052604081205460ff16908160028111156108cb576108cb6110c4565b0361094d5760405163278f794360e11b81526001600160a01b03851690634f1ef2869034906109009087908790600401611414565b60006040518083038185885af115801561091e573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f191682016040526109479190810190611440565b50610a1e565b6109578484610a24565b6000846001600160a01b031634846040516109729190611216565b60006040518083038185875af1925050503d80600081146109af576040519150601f19603f3d011682016040523d82523d6000602084013e6109b4565b606091505b5050905080610a1c5760405162461bcd60e51b815260206004820152602e60248201527f50726f787941646d696e3a2063616c6c20746f2070726f78792061667465722060448201526d1d5c19dc9859194819985a5b195960921b606482015260840161047c565b505b50505050565b610a2c610e46565b6001600160a01b03821660009081526001602052604081205460ff1690816002811115610a5b57610a5b6110c4565b03610a8e57604051631b2ce7f360e11b81526001600160a01b038381166004830152841690633659cfe690602401610758565b6001816002811115610aa257610aa26110c4565b03610afb57604051634d8587ed60e11b81527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60048201526001600160a01b038381166024830152841690639b0b0fda90604401610758565b6002816002811115610b0f57610b0f6110c4565b03610c20576001600160a01b03831660009081526002602052604081208054610b379061124f565b80601f0160208091040260200160405190810160405280929190818152602001828054610b639061124f565b8015610bb05780601f10610b8557610100808354040283529160200191610bb0565b820191906000526020600020905b815481529060010190602001808311610b9357829003601f168201915b5050600354604051639b2ea4bd60e01b81529495506001600160a01b031693639b2ea4bd9350610be8925085915087906004016114b7565b600060405180830381600087803b158015610c0257600080fd5b505af1158015610c16573d6000803e3d6000fd5b5050505050505050565b6104856114e1565b610c30610e46565b600354604051639b2ea4bd60e01b81526001600160a01b0390911690639b2ea4bd90610c6290859085906004016114b7565b600060405180830381600087803b158015610c7c57600080fd5b505af1158015610c90573d6000803e3d6000fd5b505050505050565b610ca0610e46565b6001600160a01b038116610d055760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161047c565b610d0e81610ea0565b50565b6001600160a01b03811660009081526001602052604081205460ff1681816002811115610d4057610d406110c4565b03610d8357826001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104fc573d6000803e3d6000fd5b6001816002811115610d9757610d976110c4565b03610dda57826001600160a01b031663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104fc573d6000803e3d6000fd5b6002816002811115610dee57610dee6110c4565b036105f657600360009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104fc573d6000803e3d6000fd5b6000546001600160a01b031633146106f05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161047c565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114610d0e57600080fd5b600060208284031215610f1757600080fd5b813561052081610ef0565b600060208284031215610f3457600080fd5b8135801515811461052057600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610f8357610f83610f44565b604052919050565b600067ffffffffffffffff821115610fa557610fa5610f44565b50601f01601f191660200190565b600082601f830112610fc457600080fd5b8135610fd7610fd282610f8b565b610f5a565b818152846020838601011115610fec57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806040838503121561101c57600080fd5b823561102781610ef0565b9150602083013567ffffffffffffffff81111561104357600080fd5b61104f85828601610fb3565b9150509250929050565b60005b8381101561107457818101518382015260200161105c565b83811115610a1e5750506000910152565b6000815180845261109d816020860160208601611059565b601f01601f19169290920160200192915050565b6020815260006105206020830184611085565b634e487b7160e01b600052602160045260246000fd5b60208101600383106110fc57634e487b7160e01b600052602160045260246000fd5b91905290565b6000806040838503121561111557600080fd5b823561112081610ef0565b9150602083013561113081610ef0565b809150509250929050565b6000806040838503121561114e57600080fd5b823561115981610ef0565b915060208301356003811061113057600080fd5b60008060006060848603121561118257600080fd5b833561118d81610ef0565b9250602084013561119d81610ef0565b9150604084013567ffffffffffffffff8111156111b957600080fd5b6111c586828701610fb3565b9150509250925092565b600080604083850312156111e257600080fd5b823567ffffffffffffffff8111156111f957600080fd5b61120585828601610fb3565b925050602083013561113081610ef0565b60008251611228818460208701611059565b9190910192915050565b60006020828403121561124457600080fd5b815161052081610ef0565b600181811c9082168061126357607f821691505b60208210810361063e57634e487b7160e01b600052602260045260246000fd5b60006020808352600084546112978161124f565b808487015260406001808416600081146112b857600181146112d257611300565b60ff1985168984015283151560051b890183019550611300565b896000528660002060005b858110156112f85781548b82018601529083019088016112dd565b8a0184019650505b509398975050505050505050565b601f82111561048557600081815260208120601f850160051c810160208610156113355750805b601f850160051c820191505b81811015610c9057828155600101611341565b815167ffffffffffffffff81111561136e5761136e610f44565b6113828161137c845461124f565b8461130e565b602080601f8311600181146113b7576000841561139f5750858301515b600019600386901b1c1916600185901b178555610c90565b600085815260208120601f198616915b828110156113e6578886015182559484019460019091019084016113c7565b50858210156114045787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b038316815260406020820181905260009061143890830184611085565b949350505050565b60006020828403121561145257600080fd5b815167ffffffffffffffff81111561146957600080fd5b8201601f8101841361147a57600080fd5b8051611488610fd282610f8b565b81815285602083850101111561149d57600080fd5b6114ae826020830160208601611059565b95945050505050565b6040815260006114ca6040830185611085565b905060018060a01b03831660208301529392505050565b634e487b7160e01b600052600160045260246000fdfea164736f6c634300080f000a",
}

// ProxyAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyAdminMetaData.ABI instead.
var ProxyAdminABI = ProxyAdminMetaData.ABI

// ProxyAdminBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProxyAdminMetaData.Bin instead.
var ProxyAdminBin = ProxyAdminMetaData.Bin

// DeployProxyAdmin deploys a new Ethereum contract, binding an instance of ProxyAdmin to it.
func DeployProxyAdmin(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *ProxyAdmin, error) {
	parsed, err := ProxyAdminMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyAdminBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// ProxyAdmin is an auto generated Go binding around an Ethereum contract.
type ProxyAdmin struct {
	ProxyAdminCaller     // Read-only binding to the contract
	ProxyAdminTransactor // Write-only binding to the contract
	ProxyAdminFilterer   // Log filterer for contract events
}

// ProxyAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyAdminSession struct {
	Contract     *ProxyAdmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyAdminCallerSession struct {
	Contract *ProxyAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProxyAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyAdminTransactorSession struct {
	Contract     *ProxyAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyAdminRaw struct {
	Contract *ProxyAdmin // Generic contract binding to access the raw methods on
}

// ProxyAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyAdminCallerRaw struct {
	Contract *ProxyAdminCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyAdminTransactorRaw struct {
	Contract *ProxyAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyAdmin creates a new instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdmin(address common.Address, backend bind.ContractBackend) (*ProxyAdmin, error) {
	contract, err := bindProxyAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyAdmin{ProxyAdminCaller: ProxyAdminCaller{contract: contract}, ProxyAdminTransactor: ProxyAdminTransactor{contract: contract}, ProxyAdminFilterer: ProxyAdminFilterer{contract: contract}}, nil
}

// NewProxyAdminCaller creates a new read-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminCaller(address common.Address, caller bind.ContractCaller) (*ProxyAdminCaller, error) {
	contract, err := bindProxyAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminCaller{contract: contract}, nil
}

// NewProxyAdminTransactor creates a new write-only instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyAdminTransactor, error) {
	contract, err := bindProxyAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminTransactor{contract: contract}, nil
}

// NewProxyAdminFilterer creates a new log filterer instance of ProxyAdmin, bound to a specific deployed contract.
func NewProxyAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyAdminFilterer, error) {
	contract, err := bindProxyAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminFilterer{contract: contract}, nil
}

// bindProxyAdmin binds a generic wrapper to an already deployed contract.
func bindProxyAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyAdminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.ProxyAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ProxyAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdmin *ProxyAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdmin *ProxyAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ProxyAdmin *ProxyAdminSession) AddressManager() (common.Address, error) {
	return _ProxyAdmin.Contract.AddressManager(&_ProxyAdmin.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) AddressManager() (common.Address, error) {
	return _ProxyAdmin.Contract.AddressManager(&_ProxyAdmin.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyAdmin(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyAdmin", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, _proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyAdmin(&_ProxyAdmin.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) GetProxyImplementation(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "getProxyImplementation", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _ProxyAdmin.Contract.GetProxyImplementation(&_ProxyAdmin.CallOpts, _proxy)
}

// ImplementationName is a free data retrieval call binding the contract method 0x238181ae.
//
// Solidity: function implementationName(address ) view returns(string)
func (_ProxyAdmin *ProxyAdminCaller) ImplementationName(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "implementationName", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImplementationName is a free data retrieval call binding the contract method 0x238181ae.
//
// Solidity: function implementationName(address ) view returns(string)
func (_ProxyAdmin *ProxyAdminSession) ImplementationName(arg0 common.Address) (string, error) {
	return _ProxyAdmin.Contract.ImplementationName(&_ProxyAdmin.CallOpts, arg0)
}

// ImplementationName is a free data retrieval call binding the contract method 0x238181ae.
//
// Solidity: function implementationName(address ) view returns(string)
func (_ProxyAdmin *ProxyAdminCallerSession) ImplementationName(arg0 common.Address) (string, error) {
	return _ProxyAdmin.Contract.ImplementationName(&_ProxyAdmin.CallOpts, arg0)
}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ProxyAdmin *ProxyAdminCaller) IsUpgrading(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "isUpgrading")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ProxyAdmin *ProxyAdminSession) IsUpgrading() (bool, error) {
	return _ProxyAdmin.Contract.IsUpgrading(&_ProxyAdmin.CallOpts)
}

// IsUpgrading is a free data retrieval call binding the contract method 0xb7947262.
//
// Solidity: function isUpgrading() view returns(bool)
func (_ProxyAdmin *ProxyAdminCallerSession) IsUpgrading() (bool, error) {
	return _ProxyAdmin.Contract.IsUpgrading(&_ProxyAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdmin *ProxyAdminCallerSession) Owner() (common.Address, error) {
	return _ProxyAdmin.Contract.Owner(&_ProxyAdmin.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x6bd9f516.
//
// Solidity: function proxyType(address ) view returns(uint8)
func (_ProxyAdmin *ProxyAdminCaller) ProxyType(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _ProxyAdmin.contract.Call(opts, &out, "proxyType", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProxyType is a free data retrieval call binding the contract method 0x6bd9f516.
//
// Solidity: function proxyType(address ) view returns(uint8)
func (_ProxyAdmin *ProxyAdminSession) ProxyType(arg0 common.Address) (uint8, error) {
	return _ProxyAdmin.Contract.ProxyType(&_ProxyAdmin.CallOpts, arg0)
}

// ProxyType is a free data retrieval call binding the contract method 0x6bd9f516.
//
// Solidity: function proxyType(address ) view returns(uint8)
func (_ProxyAdmin *ProxyAdminCallerSession) ProxyType(arg0 common.Address) (uint8, error) {
	return _ProxyAdmin.Contract.ProxyType(&_ProxyAdmin.CallOpts, arg0)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address _proxy, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactor) Call(opts *bind.TransactOpts, _proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "call", _proxy, _data)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address _proxy, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminSession) Call(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Call(&_ProxyAdmin.TransactOpts, _proxy, _data)
}

// Call is a paid mutator transaction binding the contract method 0x1b8b921d.
//
// Solidity: function call(address _proxy, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) Call(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Call(&_ProxyAdmin.TransactOpts, _proxy, _data)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, _proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "changeProxyAdmin", _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_ProxyAdmin *ProxyAdminSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.ChangeProxyAdmin(&_ProxyAdmin.TransactOpts, _proxy, _newAdmin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdmin.Contract.RenounceOwnership(&_ProxyAdmin.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetAddress(opts *bind.TransactOpts, _name string, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setAddress", _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_ProxyAdmin *ProxyAdminSession) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddress(&_ProxyAdmin.TransactOpts, _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddress(&_ProxyAdmin.TransactOpts, _name, _address)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetAddressManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setAddressManager", _address)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address _address) returns()
func (_ProxyAdmin *ProxyAdminSession) SetAddressManager(_address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddressManager(&_ProxyAdmin.TransactOpts, _address)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address _address) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetAddressManager(_address common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetAddressManager(&_ProxyAdmin.TransactOpts, _address)
}

// SetImplementationName is a paid mutator transaction binding the contract method 0x860f7cda.
//
// Solidity: function setImplementationName(address _address, string _name) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetImplementationName(opts *bind.TransactOpts, _address common.Address, _name string) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setImplementationName", _address, _name)
}

// SetImplementationName is a paid mutator transaction binding the contract method 0x860f7cda.
//
// Solidity: function setImplementationName(address _address, string _name) returns()
func (_ProxyAdmin *ProxyAdminSession) SetImplementationName(_address common.Address, _name string) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetImplementationName(&_ProxyAdmin.TransactOpts, _address, _name)
}

// SetImplementationName is a paid mutator transaction binding the contract method 0x860f7cda.
//
// Solidity: function setImplementationName(address _address, string _name) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetImplementationName(_address common.Address, _name string) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetImplementationName(&_ProxyAdmin.TransactOpts, _address, _name)
}

// SetProxyType is a paid mutator transaction binding the contract method 0x8d52d4a0.
//
// Solidity: function setProxyType(address _address, uint8 _type) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetProxyType(opts *bind.TransactOpts, _address common.Address, _type uint8) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setProxyType", _address, _type)
}

// SetProxyType is a paid mutator transaction binding the contract method 0x8d52d4a0.
//
// Solidity: function setProxyType(address _address, uint8 _type) returns()
func (_ProxyAdmin *ProxyAdminSession) SetProxyType(_address common.Address, _type uint8) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetProxyType(&_ProxyAdmin.TransactOpts, _address, _type)
}

// SetProxyType is a paid mutator transaction binding the contract method 0x8d52d4a0.
//
// Solidity: function setProxyType(address _address, uint8 _type) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetProxyType(_address common.Address, _type uint8) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetProxyType(&_ProxyAdmin.TransactOpts, _address, _type)
}

// SetUpgrading is a paid mutator transaction binding the contract method 0x07c8f7b0.
//
// Solidity: function setUpgrading(bool _upgrading) returns()
func (_ProxyAdmin *ProxyAdminTransactor) SetUpgrading(opts *bind.TransactOpts, _upgrading bool) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "setUpgrading", _upgrading)
}

// SetUpgrading is a paid mutator transaction binding the contract method 0x07c8f7b0.
//
// Solidity: function setUpgrading(bool _upgrading) returns()
func (_ProxyAdmin *ProxyAdminSession) SetUpgrading(_upgrading bool) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetUpgrading(&_ProxyAdmin.TransactOpts, _upgrading)
}

// SetUpgrading is a paid mutator transaction binding the contract method 0x07c8f7b0.
//
// Solidity: function setUpgrading(bool _upgrading) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) SetUpgrading(_upgrading bool) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.SetUpgrading(&_ProxyAdmin.TransactOpts, _upgrading)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.TransferOwnership(&_ProxyAdmin.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address _proxy, address _implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactor) Upgrade(opts *bind.TransactOpts, _proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgrade", _proxy, _implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address _proxy, address _implementation) returns()
func (_ProxyAdmin *ProxyAdminSession) Upgrade(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, _proxy, _implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address _proxy, address _implementation) returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) Upgrade(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.Upgrade(&_ProxyAdmin.TransactOpts, _proxy, _implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactor) UpgradeAndCall(opts *bind.TransactOpts, _proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.contract.Transact(opts, "upgradeAndCall", _proxy, _implementation, _data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminSession) UpgradeAndCall(_proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, _proxy, _implementation, _data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_ProxyAdmin *ProxyAdminTransactorSession) UpgradeAndCall(_proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _ProxyAdmin.Contract.UpgradeAndCall(&_ProxyAdmin.TransactOpts, _proxy, _implementation, _data)
}

// ProxyAdminOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferredIterator struct {
	Event *ProxyAdminOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProxyAdminOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAdminOwnershipTransferred)
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
		it.Event = new(ProxyAdminOwnershipTransferred)
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
func (it *ProxyAdminOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAdminOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAdminOwnershipTransferred represents a OwnershipTransferred event raised by the ProxyAdmin contract.
type ProxyAdminOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyAdminOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminOwnershipTransferredIterator{contract: _ProxyAdmin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyAdminOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdmin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAdminOwnershipTransferred)
				if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdmin *ProxyAdminFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyAdminOwnershipTransferred, error) {
	event := new(ProxyAdminOwnershipTransferred)
	if err := _ProxyAdmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
