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

// L2BlastBridgeMetaData contains all meta data concerning the L2BlastBridge contract.
var L2BlastBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStandardBridge\",\"name\":\"_otherBridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"contractStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETHDirect\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"contractStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200217d3803806200217d833981016040819052620000349162000113565b6001600160a01b0381166080526200004b62000052565b5062000145565b600054610100900460ff1615620000bf5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000111576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012657600080fd5b81516001600160a01b03811681146200013e57600080fd5b9392505050565b608051611ff26200018b6000396000818161022701528181610304015281816103d8015281816104c80152818161074a01528181610c1f015261154d0152611ff26000f3fe6080604052600436106100e15760003560e01c80638129fc1c1161007f578063927ede2d11610059578063927ede2d146102c4578063a47a5c35146102e2578063c89701a2146102f5578063e11013dd1461032857600080fd5b80638129fc1c14610249578063870876231461025e5780638f601f661461027e57600080fd5b80633cb747bf116100bb5780633cb747bf1461017a578063540abf73146101b757806354fd4d50146101d75780637f46ddb21461021557600080fd5b80630166a07a1461013457806309fc8843146101545780631635f5fd1461016757600080fd5b3661012f57333b1561010e5760405162461bcd60e51b8152600401610105906118c6565b60405180910390fd5b61012d33333462030d406040518060200160405280600081525061033b565b005b600080fd5b34801561014057600080fd5b5061012d61014f366004611981565b610496565b61012d610162366004611a32565b6106b1565b61012d610175366004611a85565b610718565b34801561018657600080fd5b5060035461019a906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101c357600080fd5b5061012d6101d2366004611af8565b6109ed565b3480156101e357600080fd5b50610208604051806040016040528060058152602001640312e302e360dc1b81525081565b6040516101ae9190611bc7565b34801561022157600080fd5b5061019a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561025557600080fd5b5061012d610a32565b34801561026a57600080fd5b5061012d610279366004611bda565b610bb1565b34801561028a57600080fd5b506102b6610299366004611c5d565b600260209081526000928352604080842090915290825290205481565b6040519081526020016101ae565b3480156102d057600080fd5b506003546001600160a01b031661019a565b61012d6102f0366004611a85565b610c15565b34801561030157600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061019a565b61012d610336366004611c96565b610ed4565b8234146103b05760405162461bcd60e51b815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c756500006064820152608401610105565b6103bc85858584610f1d565b6003546040516001600160a01b0390911690633dbb202b9085907f000000000000000000000000000000000000000000000000000000000000000090631635f5fd60e01b90610415908b908b9086908a90602401611cf9565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e086901b909216825261045d92918890600401611d36565b6000604051808303818588803b15801561047657600080fd5b505af115801561048a573d6000803e3d6000fd5b50505050505050505050565b6003546001600160a01b031633148015610545575060035460408051636e296e4560e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610516573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053a9190611d70565b6001600160a01b0316145b6105615760405162461bcd60e51b815260040161010590611d8d565b61056a87610f70565b156105fc576105798787610fa0565b6105955760405162461bcd60e51b815260040161010590611df4565b6040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b1580156105df57600080fd5b505af11580156105f3573d6000803e3d6000fd5b50505050610664565b6001600160a01b038088166000908152600260209081526040808320938a168352929052205461062d908490611e7a565b6001600160a01b038089166000818152600260209081526040808320948c1683529390529190912091909155610664908585611073565b6106a8878787878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506110d692505050565b50505050505050565b333b156106d05760405162461bcd60e51b8152600401610105906118c6565b6107133333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061033b92505050565b505050565b6003546001600160a01b0316331480156107c7575060035460408051636e296e4560e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610798573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107bc9190611d70565b6001600160a01b0316145b6107e35760405162461bcd60e51b815260040161010590611d8d565b8234146108585760405162461bcd60e51b815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e742072657175697265640000000000006064820152608401610105565b306001600160a01b038516036108bc5760405162461bcd60e51b815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201526232b63360e91b6064820152608401610105565b6003546001600160a01b039081169085160361092b5760405162461bcd60e51b815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201526732b9b9b2b733b2b960c11b6064820152608401610105565b61096d85858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061113792505050565b600061098a855a866040518060200160405280600081525061117c565b9050806109e55760405162461bcd60e51b815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e73666572206661696044820152621b195960ea1b6064820152608401610105565b505050505050565b6106a887873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061119692505050565b600054610100900460ff1615808015610a525750600054600160ff909116105b80610a6c5750303b158015610a6c575060005460ff166001145b610acf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610105565b6000805460ff191660011790558015610af2576000805461ff0019166101001790555b610b026007602160991b0161129b565b60405163099005e760e31b81526002604360981b0190634c802f3890610b3690309060019060009061dead90600401611ea7565b600060405180830381600087803b158015610b5057600080fd5b505af1158015610b64573d6000803e3d6000fd5b505050508015610bae576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b333b15610bd05760405162461bcd60e51b8152600401610105906118c6565b6109e586863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061119692505050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167311110000000000000000000000000000000011101933016001600160a01b031614610cd5576040805162461bcd60e51b81526020600482015260248101919091527f4c32426c6173744272696467653a2066756e6374696f6e2063616e206f6e6c7960448201527f2062652063616c6c65642066726f6d20746865206f74686572206272696467656064820152608401610105565b823414610d4a5760405162461bcd60e51b815260206004820152603960248201527f4c32426c6173744272696467653a20616d6f756e742073656e7420646f65732060448201527f6e6f74206d6174636820616d6f756e74207265717569726564000000000000006064820152608401610105565b306001600160a01b03851603610dad5760405162461bcd60e51b815260206004820152602260248201527f4c32426c6173744272696467653a2063616e6e6f742073656e6420746f207365604482015261363360f11b6064820152608401610105565b6003546001600160a01b0390811690851603610e1b5760405162461bcd60e51b815260206004820152602760248201527f4c32426c6173744272696467653a2063616e6e6f742073656e6420746f206d6560448201526639b9b2b733b2b960c91b6064820152608401610105565b610e5d85858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061113792505050565b6000610e7a855a866040518060200160405280600081525061117c565b9050806109e55760405162461bcd60e51b815260206004820152602260248201527f4c32426c6173744272696467653a20455448207472616e73666572206661696c604482015261195960f21b6064820152608401610105565b610f173385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061033b92505050565b50505050565b826001600160a01b0316846001600160a01b03167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051610f62929190611ef7565b60405180910390a350505050565b6000610f8382631d1d8b6360e01b611328565b80610f9a5750610f9a8263ec4fc8e360e01b611328565b92915050565b6000610fb383631d1d8b6360e01b611328565b1561103557826001600160a01b031663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ff6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101a9190611d70565b6001600160a01b0316826001600160a01b0316149050610f9a565b826001600160a01b031663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ff6573d6000803e3d6000fd5b6040516001600160a01b03831660248201526044810182905261071390849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261134b565b836001600160a01b0316856001600160a01b0316876001600160a01b03167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd86868660405161112793929190611f10565b60405180910390a4505050505050565b826001600160a01b0316846001600160a01b03167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051610f62929190611ef7565b600080600080845160208601878a8af19695505050505050565b6001600160a01b0387166003604360981b011461121b5760405162461bcd60e51b815260206004820152603b60248201527f4c32426c6173744272696467653a206f6e6c7920555344422063616e2062652060448201527f77697468647261776e2066726f6d2074686973206272696467652e00000000006064820152608401610105565b61122c6003604360981b0187610fa0565b61128c5760405162461bcd60e51b815260206004820152602b60248201527f4c32426c6173744272696467653a2077726f6e672072656d6f746520746f6b6560448201526a37103337b9102aa9a2211760a91b6064820152608401610105565b6106a88787878787878761141d565b600054610100900460ff166113065760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610105565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b600061133383611610565b801561134457506113448383611643565b9392505050565b60006113a0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116cc9092919063ffffffff16565b80519091501561071357808060200190518101906113be9190611f40565b6107135760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610105565b61142687610f70565b156114b8576114358787610fa0565b6114515760405162461bcd60e51b815260040161010590611df4565b604051632770a7eb60e21b81526001600160a01b03868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b15801561149b57600080fd5b505af11580156114af573d6000803e3d6000fd5b50505050611525565b6114cd6001600160a01b0388168630866116e3565b6001600160a01b038088166000908152600260209081526040808320938a16835292905220546114fe908490611f62565b6001600160a01b038089166000908152600260209081526040808320938b16835292905220555b61153387878787878661171b565b6003546040516001600160a01b0390911690633dbb202b907f00000000000000000000000000000000000000000000000000000000000000009062b3503d60e11b9061158d908b908d908c908c908c908b90602401611f7a565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e085901b90921682526115d592918790600401611d36565b600060405180830381600087803b1580156115ef57600080fd5b505af1158015611603573d6000803e3d6000fd5b5050505050505050505050565b6000611623826301ffc9a760e01b611643565b8015610f9a575061163c826001600160e01b0319611643565b1592915050565b604080516001600160e01b03198316602480830191909152825180830390910181526044909101909152602080820180516001600160e01b03166301ffc9a760e01b178152825160009392849283928392918391908a617530fa92503d915060005190508280156116b5575060208210155b80156116c15750600081115b979650505050505050565b60606116db848460008561176c565b949350505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610f179085906323b872dd60e01b9060840161109f565b836001600160a01b0316856001600160a01b0316876001600160a01b03167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf86868660405161112793929190611f10565b6060824710156117cd5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610105565b6001600160a01b0385163b6118245760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610105565b600080866001600160a01b031685876040516118409190611fc9565b60006040518083038185875af1925050503d806000811461187d576040519150601f19603f3d011682016040523d82523d6000602084013e611882565b606091505b50915091506116c18282866060831561189c575081611344565b8251156118ac5782518084602001fd5b8160405162461bcd60e51b81526004016101059190611bc7565b60208082526037908201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60408201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606082015260800190565b6001600160a01b0381168114610bae57600080fd5b60008083601f84011261194a57600080fd5b50813567ffffffffffffffff81111561196257600080fd5b60208301915083602082850101111561197a57600080fd5b9250929050565b600080600080600080600060c0888a03121561199c57600080fd5b87356119a781611923565b965060208801356119b781611923565b955060408801356119c781611923565b945060608801356119d781611923565b93506080880135925060a088013567ffffffffffffffff8111156119fa57600080fd5b611a068a828b01611938565b989b979a50959850939692959293505050565b803563ffffffff81168114611a2d57600080fd5b919050565b600080600060408486031215611a4757600080fd5b611a5084611a19565b9250602084013567ffffffffffffffff811115611a6c57600080fd5b611a7886828701611938565b9497909650939450505050565b600080600080600060808688031215611a9d57600080fd5b8535611aa881611923565b94506020860135611ab881611923565b935060408601359250606086013567ffffffffffffffff811115611adb57600080fd5b611ae788828901611938565b969995985093965092949392505050565b600080600080600080600060c0888a031215611b1357600080fd5b8735611b1e81611923565b96506020880135611b2e81611923565b95506040880135611b3e81611923565b945060608801359350611b5360808901611a19565b925060a088013567ffffffffffffffff8111156119fa57600080fd5b60005b83811015611b8a578181015183820152602001611b72565b83811115610f175750506000910152565b60008151808452611bb3816020860160208601611b6f565b601f01601f19169290920160200192915050565b6020815260006113446020830184611b9b565b60008060008060008060a08789031215611bf357600080fd5b8635611bfe81611923565b95506020870135611c0e81611923565b945060408701359350611c2360608801611a19565b9250608087013567ffffffffffffffff811115611c3f57600080fd5b611c4b89828a01611938565b979a9699509497509295939492505050565b60008060408385031215611c7057600080fd5b8235611c7b81611923565b91506020830135611c8b81611923565b809150509250929050565b60008060008060608587031215611cac57600080fd5b8435611cb781611923565b9350611cc560208601611a19565b9250604085013567ffffffffffffffff811115611ce157600080fd5b611ced87828801611938565b95989497509550505050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611d2c90830184611b9b565b9695505050505050565b6001600160a01b0384168152606060208201819052600090611d5a90830185611b9b565b905063ffffffff83166040830152949350505050565b600060208284031215611d8257600080fd5b815161134481611923565b60208082526041908201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60408201527f792062652063616c6c65642066726f6d20746865206f746865722062726964676060820152606560f81b608082015260a00190565b6020808252604a908201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60408201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60608201526937b1b0b6103a37b5b2b760b11b608082015260a00190565b634e487b7160e01b600052601160045260246000fd5b600082821015611e8c57611e8c611e64565b500390565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b038581168252608082019060038610611ec957611ec9611e91565b85602084015260028510611edf57611edf611e91565b84604084015280841660608401525095945050505050565b8281526040602082015260006116db6040830184611b9b565b60018060a01b0384168152826020820152606060408201526000611f376060830184611b9b565b95945050505050565b600060208284031215611f5257600080fd5b8151801515811461134457600080fd5b60008219821115611f7557611f75611e64565b500190565b6001600160a01b03878116825286811660208301528581166040830152841660608201526080810183905260c060a08201819052600090611fbd90830184611b9b565b98975050505050505050565b60008251611fdb818460208701611b6f565b919091019291505056fea164736f6c634300080f000a",
}

// L2BlastBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L2BlastBridgeMetaData.ABI instead.
var L2BlastBridgeABI = L2BlastBridgeMetaData.ABI

// L2BlastBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2BlastBridgeMetaData.Bin instead.
var L2BlastBridgeBin = L2BlastBridgeMetaData.Bin

// DeployL2BlastBridge deploys a new Ethereum contract, binding an instance of L2BlastBridge to it.
func DeployL2BlastBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _otherBridge common.Address) (common.Address, *types.Transaction, *L2BlastBridge, error) {
	parsed, err := L2BlastBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2BlastBridgeBin), backend, _otherBridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2BlastBridge{L2BlastBridgeCaller: L2BlastBridgeCaller{contract: contract}, L2BlastBridgeTransactor: L2BlastBridgeTransactor{contract: contract}, L2BlastBridgeFilterer: L2BlastBridgeFilterer{contract: contract}}, nil
}

// L2BlastBridge is an auto generated Go binding around an Ethereum contract.
type L2BlastBridge struct {
	L2BlastBridgeCaller     // Read-only binding to the contract
	L2BlastBridgeTransactor // Write-only binding to the contract
	L2BlastBridgeFilterer   // Log filterer for contract events
}

// L2BlastBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2BlastBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BlastBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2BlastBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BlastBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2BlastBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BlastBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2BlastBridgeSession struct {
	Contract     *L2BlastBridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2BlastBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2BlastBridgeCallerSession struct {
	Contract *L2BlastBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2BlastBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2BlastBridgeTransactorSession struct {
	Contract     *L2BlastBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2BlastBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2BlastBridgeRaw struct {
	Contract *L2BlastBridge // Generic contract binding to access the raw methods on
}

// L2BlastBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2BlastBridgeCallerRaw struct {
	Contract *L2BlastBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L2BlastBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2BlastBridgeTransactorRaw struct {
	Contract *L2BlastBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2BlastBridge creates a new instance of L2BlastBridge, bound to a specific deployed contract.
func NewL2BlastBridge(address common.Address, backend bind.ContractBackend) (*L2BlastBridge, error) {
	contract, err := bindL2BlastBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridge{L2BlastBridgeCaller: L2BlastBridgeCaller{contract: contract}, L2BlastBridgeTransactor: L2BlastBridgeTransactor{contract: contract}, L2BlastBridgeFilterer: L2BlastBridgeFilterer{contract: contract}}, nil
}

// NewL2BlastBridgeCaller creates a new read-only instance of L2BlastBridge, bound to a specific deployed contract.
func NewL2BlastBridgeCaller(address common.Address, caller bind.ContractCaller) (*L2BlastBridgeCaller, error) {
	contract, err := bindL2BlastBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeCaller{contract: contract}, nil
}

// NewL2BlastBridgeTransactor creates a new write-only instance of L2BlastBridge, bound to a specific deployed contract.
func NewL2BlastBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L2BlastBridgeTransactor, error) {
	contract, err := bindL2BlastBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeTransactor{contract: contract}, nil
}

// NewL2BlastBridgeFilterer creates a new log filterer instance of L2BlastBridge, bound to a specific deployed contract.
func NewL2BlastBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L2BlastBridgeFilterer, error) {
	contract, err := bindL2BlastBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeFilterer{contract: contract}, nil
}

// bindL2BlastBridge binds a generic wrapper to an already deployed contract.
func bindL2BlastBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2BlastBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2BlastBridge *L2BlastBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2BlastBridge.Contract.L2BlastBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2BlastBridge *L2BlastBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.L2BlastBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2BlastBridge *L2BlastBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.L2BlastBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2BlastBridge *L2BlastBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2BlastBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2BlastBridge *L2BlastBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2BlastBridge *L2BlastBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2BlastBridge.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2BlastBridge *L2BlastBridgeSession) MESSENGER() (common.Address, error) {
	return _L2BlastBridge.Contract.MESSENGER(&_L2BlastBridge.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCallerSession) MESSENGER() (common.Address, error) {
	return _L2BlastBridge.Contract.MESSENGER(&_L2BlastBridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCaller) OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2BlastBridge.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2BlastBridge *L2BlastBridgeSession) OTHERBRIDGE() (common.Address, error) {
	return _L2BlastBridge.Contract.OTHERBRIDGE(&_L2BlastBridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCallerSession) OTHERBRIDGE() (common.Address, error) {
	return _L2BlastBridge.Contract.OTHERBRIDGE(&_L2BlastBridge.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L2BlastBridge *L2BlastBridgeCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2BlastBridge.contract.Call(opts, &out, "deposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L2BlastBridge *L2BlastBridgeSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L2BlastBridge.Contract.Deposits(&_L2BlastBridge.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L2BlastBridge *L2BlastBridgeCallerSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L2BlastBridge.Contract.Deposits(&_L2BlastBridge.CallOpts, arg0, arg1)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2BlastBridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2BlastBridge *L2BlastBridgeSession) Messenger() (common.Address, error) {
	return _L2BlastBridge.Contract.Messenger(&_L2BlastBridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCallerSession) Messenger() (common.Address, error) {
	return _L2BlastBridge.Contract.Messenger(&_L2BlastBridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCaller) OtherBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2BlastBridge.contract.Call(opts, &out, "otherBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2BlastBridge *L2BlastBridgeSession) OtherBridge() (common.Address, error) {
	return _L2BlastBridge.Contract.OtherBridge(&_L2BlastBridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2BlastBridge *L2BlastBridgeCallerSession) OtherBridge() (common.Address, error) {
	return _L2BlastBridge.Contract.OtherBridge(&_L2BlastBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2BlastBridge *L2BlastBridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2BlastBridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2BlastBridge *L2BlastBridgeSession) Version() (string, error) {
	return _L2BlastBridge.Contract.Version(&_L2BlastBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2BlastBridge *L2BlastBridgeCallerSession) Version() (string, error) {
	return _L2BlastBridge.Contract.Version(&_L2BlastBridge.CallOpts)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) BridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "bridgeERC20", _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeERC20(&_L2BlastBridge.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeERC20(&_L2BlastBridge.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) BridgeERC20To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "bridgeERC20To", _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeERC20To(&_L2BlastBridge.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeERC20To(&_L2BlastBridge.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) BridgeETH(opts *bind.TransactOpts, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "bridgeETH", _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeSession) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeETH(&_L2BlastBridge.TransactOpts, _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeETH(&_L2BlastBridge.TransactOpts, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) BridgeETHTo(opts *bind.TransactOpts, _to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "bridgeETHTo", _to, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeSession) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeETHTo(&_L2BlastBridge.TransactOpts, _to, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.BridgeETHTo(&_L2BlastBridge.TransactOpts, _to, _minGasLimit, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) FinalizeBridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "finalizeBridgeERC20", _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.FinalizeBridgeERC20(&_L2BlastBridge.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.FinalizeBridgeERC20(&_L2BlastBridge.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) FinalizeBridgeETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "finalizeBridgeETH", _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeSession) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.FinalizeBridgeETH(&_L2BlastBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.FinalizeBridgeETH(&_L2BlastBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETHDirect is a paid mutator transaction binding the contract method 0xa47a5c35.
//
// Solidity: function finalizeBridgeETHDirect(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) FinalizeBridgeETHDirect(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "finalizeBridgeETHDirect", _from, _to, _amount, _extraData)
}

// FinalizeBridgeETHDirect is a paid mutator transaction binding the contract method 0xa47a5c35.
//
// Solidity: function finalizeBridgeETHDirect(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeSession) FinalizeBridgeETHDirect(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.FinalizeBridgeETHDirect(&_L2BlastBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETHDirect is a paid mutator transaction binding the contract method 0xa47a5c35.
//
// Solidity: function finalizeBridgeETHDirect(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) FinalizeBridgeETHDirect(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L2BlastBridge.Contract.FinalizeBridgeETHDirect(&_L2BlastBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2BlastBridge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2BlastBridge *L2BlastBridgeSession) Initialize() (*types.Transaction, error) {
	return _L2BlastBridge.Contract.Initialize(&_L2BlastBridge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) Initialize() (*types.Transaction, error) {
	return _L2BlastBridge.Contract.Initialize(&_L2BlastBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2BlastBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2BlastBridge *L2BlastBridgeSession) Receive() (*types.Transaction, error) {
	return _L2BlastBridge.Contract.Receive(&_L2BlastBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2BlastBridge *L2BlastBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _L2BlastBridge.Contract.Receive(&_L2BlastBridge.TransactOpts)
}

// L2BlastBridgeERC20BridgeFinalizedIterator is returned from FilterERC20BridgeFinalized and is used to iterate over the raw logs and unpacked data for ERC20BridgeFinalized events raised by the L2BlastBridge contract.
type L2BlastBridgeERC20BridgeFinalizedIterator struct {
	Event *L2BlastBridgeERC20BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L2BlastBridgeERC20BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BlastBridgeERC20BridgeFinalized)
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
		it.Event = new(L2BlastBridgeERC20BridgeFinalized)
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
func (it *L2BlastBridgeERC20BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BlastBridgeERC20BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BlastBridgeERC20BridgeFinalized represents a ERC20BridgeFinalized event raised by the L2BlastBridge contract.
type L2BlastBridgeERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeFinalized is a free log retrieval operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L2BlastBridgeERC20BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2BlastBridge.contract.FilterLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeERC20BridgeFinalizedIterator{contract: _L2BlastBridge.contract, event: "ERC20BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeFinalized is a free log subscription operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *L2BlastBridgeERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2BlastBridge.contract.WatchLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BlastBridgeERC20BridgeFinalized)
				if err := _L2BlastBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
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

// ParseERC20BridgeFinalized is a log parse operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) ParseERC20BridgeFinalized(log types.Log) (*L2BlastBridgeERC20BridgeFinalized, error) {
	event := new(L2BlastBridgeERC20BridgeFinalized)
	if err := _L2BlastBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BlastBridgeERC20BridgeInitiatedIterator is returned from FilterERC20BridgeInitiated and is used to iterate over the raw logs and unpacked data for ERC20BridgeInitiated events raised by the L2BlastBridge contract.
type L2BlastBridgeERC20BridgeInitiatedIterator struct {
	Event *L2BlastBridgeERC20BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L2BlastBridgeERC20BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BlastBridgeERC20BridgeInitiated)
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
		it.Event = new(L2BlastBridgeERC20BridgeInitiated)
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
func (it *L2BlastBridgeERC20BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BlastBridgeERC20BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BlastBridgeERC20BridgeInitiated represents a ERC20BridgeInitiated event raised by the L2BlastBridge contract.
type L2BlastBridgeERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeInitiated is a free log retrieval operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) FilterERC20BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L2BlastBridgeERC20BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2BlastBridge.contract.FilterLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeERC20BridgeInitiatedIterator{contract: _L2BlastBridge.contract, event: "ERC20BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeInitiated is a free log subscription operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) WatchERC20BridgeInitiated(opts *bind.WatchOpts, sink chan<- *L2BlastBridgeERC20BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2BlastBridge.contract.WatchLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BlastBridgeERC20BridgeInitiated)
				if err := _L2BlastBridge.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
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

// ParseERC20BridgeInitiated is a log parse operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) ParseERC20BridgeInitiated(log types.Log) (*L2BlastBridgeERC20BridgeInitiated, error) {
	event := new(L2BlastBridgeERC20BridgeInitiated)
	if err := _L2BlastBridge.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BlastBridgeETHBridgeFinalizedIterator is returned from FilterETHBridgeFinalized and is used to iterate over the raw logs and unpacked data for ETHBridgeFinalized events raised by the L2BlastBridge contract.
type L2BlastBridgeETHBridgeFinalizedIterator struct {
	Event *L2BlastBridgeETHBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L2BlastBridgeETHBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BlastBridgeETHBridgeFinalized)
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
		it.Event = new(L2BlastBridgeETHBridgeFinalized)
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
func (it *L2BlastBridgeETHBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BlastBridgeETHBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BlastBridgeETHBridgeFinalized represents a ETHBridgeFinalized event raised by the L2BlastBridge contract.
type L2BlastBridgeETHBridgeFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHBridgeFinalized is a free log retrieval operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) FilterETHBridgeFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2BlastBridgeETHBridgeFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2BlastBridge.contract.FilterLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeETHBridgeFinalizedIterator{contract: _L2BlastBridge.contract, event: "ETHBridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchETHBridgeFinalized is a free log subscription operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) WatchETHBridgeFinalized(opts *bind.WatchOpts, sink chan<- *L2BlastBridgeETHBridgeFinalized, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2BlastBridge.contract.WatchLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BlastBridgeETHBridgeFinalized)
				if err := _L2BlastBridge.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
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

// ParseETHBridgeFinalized is a log parse operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) ParseETHBridgeFinalized(log types.Log) (*L2BlastBridgeETHBridgeFinalized, error) {
	event := new(L2BlastBridgeETHBridgeFinalized)
	if err := _L2BlastBridge.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BlastBridgeETHBridgeInitiatedIterator is returned from FilterETHBridgeInitiated and is used to iterate over the raw logs and unpacked data for ETHBridgeInitiated events raised by the L2BlastBridge contract.
type L2BlastBridgeETHBridgeInitiatedIterator struct {
	Event *L2BlastBridgeETHBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L2BlastBridgeETHBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BlastBridgeETHBridgeInitiated)
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
		it.Event = new(L2BlastBridgeETHBridgeInitiated)
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
func (it *L2BlastBridgeETHBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BlastBridgeETHBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BlastBridgeETHBridgeInitiated represents a ETHBridgeInitiated event raised by the L2BlastBridge contract.
type L2BlastBridgeETHBridgeInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHBridgeInitiated is a free log retrieval operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) FilterETHBridgeInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2BlastBridgeETHBridgeInitiatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2BlastBridge.contract.FilterLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeETHBridgeInitiatedIterator{contract: _L2BlastBridge.contract, event: "ETHBridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchETHBridgeInitiated is a free log subscription operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) WatchETHBridgeInitiated(opts *bind.WatchOpts, sink chan<- *L2BlastBridgeETHBridgeInitiated, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2BlastBridge.contract.WatchLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BlastBridgeETHBridgeInitiated)
				if err := _L2BlastBridge.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
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

// ParseETHBridgeInitiated is a log parse operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L2BlastBridge *L2BlastBridgeFilterer) ParseETHBridgeInitiated(log types.Log) (*L2BlastBridgeETHBridgeInitiated, error) {
	event := new(L2BlastBridgeETHBridgeInitiated)
	if err := _L2BlastBridge.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BlastBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2BlastBridge contract.
type L2BlastBridgeInitializedIterator struct {
	Event *L2BlastBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L2BlastBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BlastBridgeInitialized)
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
		it.Event = new(L2BlastBridgeInitialized)
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
func (it *L2BlastBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BlastBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BlastBridgeInitialized represents a Initialized event raised by the L2BlastBridge contract.
type L2BlastBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2BlastBridge *L2BlastBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2BlastBridgeInitializedIterator, error) {

	logs, sub, err := _L2BlastBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2BlastBridgeInitializedIterator{contract: _L2BlastBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2BlastBridge *L2BlastBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2BlastBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L2BlastBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BlastBridgeInitialized)
				if err := _L2BlastBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2BlastBridge *L2BlastBridgeFilterer) ParseInitialized(log types.Log) (*L2BlastBridgeInitialized, error) {
	event := new(L2BlastBridgeInitialized)
	if err := _L2BlastBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
