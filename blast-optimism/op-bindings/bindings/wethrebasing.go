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

// WETHRebasingMetaData contains all meta data concerning the WETHRebasing contract.
var WETHRebasingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApproveFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotClaimToSameAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimableAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumYieldMode\",\"name\":\"yieldMode\",\"type\":\"uint8\"}],\"name\":\"Configure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"yieldMode\",\"type\":\"uint8\"}],\"name\":\"configure\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101c060405234801561001157600080fd5b50604080518082018252600d81526c2bb930b83832b21022ba3432b960991b60208083019182528351808501855260048152630ae8aa8960e31b90820152835180850185526001808252603160f81b91830191909152925190912060e08181527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66101008181524660a081815288517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818901819052818b019790975260608101949094526080808501929092523084820181905289518086038301815260c0958601909a52895199909701989098209081905291859052610120849052601261014081905261016087905260006101808190526101a0819052975193519151929793969194929390919080611f7461019382396000610718015260006106ef015260006106c601526000610289015260006112e1015260006113300152600061130b015260006112640152600061128e015260006112b80152611f746000f3fe6080604052600436106101445760003560e01c806370a08231116100b6578063aad3ec961161006f578063aad3ec9614610379578063c44b11f714610399578063d0e30db0146103c6578063d505accf146103ce578063dd62ed3e146103ee578063e12f3a611461043457600080fd5b806370a08231146102e75780637ecebe00146103075780638129fc1c14610327578063872697291461032f57806395d89b4114610344578063a9059cbb1461035957600080fd5b806323b872dd1161010857806323b872dd146102035780632e1a7d4d1461022357806330adf81f14610243578063313ce567146102775780633644e515146102bd57806354fd4d50146102d257600080fd5b806306661abd1461015857806306fdde031461017c578063095ea7b31461019e57806318160ddd146101ce5780631a33757d146101e357600080fd5b3661015357610151610454565b005b600080fd5b34801561016457600080fd5b506006545b6040519081526020015b60405180910390f35b34801561018857600080fd5b506101916104a5565b6040516101739190611aa5565b3480156101aa57600080fd5b506101be6101b9366004611af4565b610533565b6040519015158152602001610173565b3480156101da57600080fd5b5061016961054d565b3480156101ef57600080fd5b506101696101fe366004611b1e565b610576565b34801561020f57600080fd5b506101be61021e366004611b3f565b6105cc565b34801561022f57600080fd5b5061015161023e366004611b7b565b6105ee565b34801561024f57600080fd5b506101697f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b34801561028357600080fd5b506102ab7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff9091168152602001610173565b3480156102c957600080fd5b506101696106b5565b3480156102de57600080fd5b506101916106bf565b3480156102f357600080fd5b50610169610302366004611b94565b610762565b34801561031357600080fd5b50610169610322366004611b94565b6107f4565b610151610812565b34801561033b57600080fd5b50610169610a44565b34801561035057600080fd5b50610191610a4e565b34801561036557600080fd5b506101be610374366004611af4565b610a5b565b34801561038557600080fd5b50610169610394366004611af4565b610a71565b3480156103a557600080fd5b506103b96103b4366004611b94565b610bf8565b6040516101739190611bd9565b610151610454565b3480156103da57600080fd5b506101516103e9366004611be7565b610c16565b3480156103fa57600080fd5b50610169610409366004611c5a565b6001600160a01b039182166000908152600b6020908152604080832093909416825291909152205490565b34801561044057600080fd5b5061016961044f366004611b94565b610d7a565b3361045f8134610df0565b806001600160a01b03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3460405161049a91815260200190565b60405180910390a250565b600380546104b290611c84565b80601f01602080910402602001604051908101604052809291908181526020018280546104de90611c84565b801561052b5780601f106105005761010080835404028352916020019161052b565b820191906000526020600020905b81548152906001019060200180831161050e57829003601f168201915b505050505081565b600033610541818585610e60565b60019150505b92915050565b600060095460065461055d610a44565b6105679190611cce565b6105719190611ced565b905090565b60006105823383610f0f565b336001600160a01b03167fcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36836040516105bb9190611bd9565b60405180910390a261054733610762565b60006105d9843384610ff9565b6105e4848484611056565b5060019392505050565b336105f981836111e2565b6000816001600160a01b03168360405160006040518083038185875af1925050503d8060008114610646576040519150601f19603f3d011682016040523d82523d6000602084013e61064b565b606091505b505090508061066d5760405163b12d13eb60e01b815260040160405180910390fd5b816001600160a01b03167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65846040516106a891815260200190565b60405180910390a2505050565b6000610571611257565b60606106ea7f000000000000000000000000000000000000000000000000000000000000000061137e565b6107137f000000000000000000000000000000000000000000000000000000000000000061137e565b61073c7f000000000000000000000000000000000000000000000000000000000000000061137e565b60405160200161074e93929190611d05565b604051602081830303815290604052905090565b6001600160a01b0381166000908152600a602052604081205460ff168181600281111561079157610791611baf565b036107d2576107cb6107a1610a44565b6001600160a01b038516600090815260056020908152604080832054600790925290912054611487565b91506107ee565b6001600160a01b03831660009081526008602052604090205491505b50919050565b6001600160a01b038116600090815260016020526040812054610547565b600054610100900460ff16158080156108325750600054600160ff909116105b8061084c5750303b15801561084c575060005460ff166001145b6108b45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156108d7576000805461ff0019166101001790555b6109226040518060400160405280600d81526020016c2bb930b83832b21022ba3432b960991b815250604051806040016040528060048152602001630ae8aa8960e31b81525061149e565b60405163099005e760e31b81526002604360981b0190634c802f389061095390309060009081908190600401611d5f565b600060405180830381600087803b15801561096d57600080fd5b505af1158015610981573d6000803e3d6000fd5b50505050604360981b6001600160a01b031663a035b1fe6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109eb9190611da6565b34146109f657600080fd5b60016006558015610a41576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000610571611527565b600480546104b290611c84565b6000610a68338484611056565b50600192915050565b6000336001600160a01b0384168103610a9d57604051637520beed60e01b815260040160405180910390fd5b6002610aa882610bf8565b6002811115610ab957610ab9611baf565b14610ad75760405163ebf953c760e01b815260040160405180910390fd5b6000610ae1610a44565b6001600160a01b03831660009081526005602090815260408083205460079092528220549293509091610b15918491611487565b6001600160a01b03841660009081526008602052604081205491925090610b3c9083611dbf565b905080861115610b5f57604051631e9acf1760e31b815260040160405180910390fd5b600080610b7585610b708a87611dbf565b61155f565b91509150610b838989610df0565b6001600160a01b038616600090815260086020526040902054610bab90879084908490611591565b6040518881526001600160a01b038a169033907f70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd9870689060200160405180910390a35095979650505050505050565b6001600160a01b03166000908152600a602052604090205460ff1690565b83421115610c665760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016108ab565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610c958c61162d565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610cf082611655565b90506000610d00828787876116a3565b9050896001600160a01b0316816001600160a01b031614610d635760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016108ab565b610d6e8a8a8a610e60565b50505050505050505050565b60006002610d8783610bf8565b6002811115610d9857610d98611baf565b14610db65760405163ebf953c760e01b815260040160405180910390fd5b6000610dc36107a1610a44565b6001600160a01b038416600090815260086020526040902054909150610de99082611dbf565b9392505050565b600081610dfc84610762565b610e069190611ced565b9050610e1c8382610e15610a44565b60006116cb565b6000610e2784610bf8565b90506001816002811115610e3d57610e3d611baf565b03610e5a578260096000828254610e549190611ced565b90915550505b50505050565b6001600160a01b038316610e875760405163eb3b083560e01b815260040160405180910390fd5b6001600160a01b038216610eae5760405163076e33c360e31b815260040160405180910390fd5b6001600160a01b038381166000818152600b602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6000610f1a83610762565b90506000610f2784610bf8565b6001600160a01b0385166000908152600a602052604090208054919250849160ff19166001836002811115610f5e57610f5e611baf565b02179055506001600160a01b038416600090815260086020526040902054610f908584610f89610a44565b60016116cb565b6001826002811115610fa457610fa4611baf565b03610fc1578060096000828254610fbb9190611dbf565b90915550505b6001846002811115610fd557610fd5611baf565b03610ff2578260096000828254610fec9190611ced565b90915550505b5050505050565b6001600160a01b038381166000908152600b60209081526040808320938616835292905220546000198114610e5a5780821115611049576040516313be252b60e01b815260040160405180910390fd5b610e5a8484848403610e60565b6001600160a01b03831661107d57604051630b07e54560e11b815260040160405180910390fd5b6001600160a01b0382166110a457604051633a954ecd60e21b815260040160405180910390fd5b60006110ae610a44565b905060006110bb85610762565b9050808311156110de57604051631e9acf1760e31b815260040160405180910390fd5b60006110e985610762565b9050611101866110f98685611dbf565b8560006116cb565b61110f856110f98684611ced565b600061111a87610bf8565b9050600181600281111561113057611130611baf565b0361114d5784600960008282546111479190611dbf565b90915550505b600061115887610bf8565b9050600181600281111561116e5761116e611baf565b0361118b5785600960008282546111859190611ced565b90915550505b866001600160a01b0316886001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef886040516111d091815260200190565b60405180910390a35050505050505050565b60006111ed83610762565b90508082111561121057604051631e9acf1760e31b815260040160405180910390fd5b61121f83838303610e15610a44565b600061122a84610bf8565b9050600181600281111561124057611240611baf565b03610e5a578260096000828254610e549190611dbf565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156112b057507f000000000000000000000000000000000000000000000000000000000000000046145b156112da57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b6060816000036113a55750506040805180820190915260018152600360fc1b602082015290565b8160005b81156113cf57806113b981611dd6565b91506113c89050600a83611e05565b91506113a9565b60008167ffffffffffffffff8111156113ea576113ea611e19565b6040519080825280601f01601f191660200182016040528015611414576020820181803683370190505b5090505b841561147f57611429600183611dbf565b9150611436600a86611e2f565b611441906030611ced565b60f81b81838151811061145657611456611e43565b60200101906001600160f81b031916908160001a905350611478600a86611e05565b9450611418565b949350505050565b6000816114948486611cce565b61147f9190611ced565b600054610100900460ff166115095760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016108ab565b60036115158382611ea7565b5060046115228282611ea7565b505050565b60006006546000036115395750600090565b600654346009544761154b9190611dbf565b6115559190611dbf565b6105719190611e05565b6000808360000361157157508161158a565b61157b8484611e05565b91506115878484611e2f565b90505b9250929050565b6001600160a01b0384166000908152600560205260409020546006546115b8908590611ced565b6115c29190611dbf565b6006556001600160a01b0384166000908152600760205260409020546009546115ec908490611ced565b6115f69190611dbf565b6009556001600160a01b03909316600090815260056020908152604080832094909455600781528382209290925560089091522055565b6001600160a01b038116600090815260016020526040812080548154600101825591506107ee565b6000610547611662611257565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b60008060006116b4878787876117d6565b915091506116c1816118c3565b5095945050505050565b6000806000806116da88610bf8565b905060008160028111156116f0576116f0611baf565b03611709576116ff868861155f565b90945092506117c0565b600181600281111561171d5761171d611baf565b0361172a578691506117c0565b600281600281111561173e5761173e611baf565b036117c05786915081856117af576001600160a01b03891660009081526005602090815260408083205460079092529091205461177c918991611487565b6001600160a01b038a166000908152600860205260409020549091506117a28983611ced565b6117ac9190611dbf565b90505b6117b9878261155f565b9095509350505b6117cc88858585611591565b5050505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561180d57506000905060036118ba565b8460ff16601b1415801561182557508460ff16601c14155b1561183657506000905060046118ba565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561188a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166118b3576000600192509250506118ba565b9150600090505b94509492505050565b60008160048111156118d7576118d7611baf565b036118df5750565b60018160048111156118f3576118f3611baf565b036119405760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016108ab565b600281600481111561195457611954611baf565b036119a15760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016108ab565b60038160048111156119b5576119b5611baf565b03611a0d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016108ab565b6004816004811115611a2157611a21611baf565b03610a415760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016108ab565b60005b83811015611a94578181015183820152602001611a7c565b83811115610e5a5750506000910152565b6020815260008251806020840152611ac4816040850160208701611a79565b601f01601f19169190910160400192915050565b80356001600160a01b0381168114611aef57600080fd5b919050565b60008060408385031215611b0757600080fd5b611b1083611ad8565b946020939093013593505050565b600060208284031215611b3057600080fd5b813560038110610de957600080fd5b600080600060608486031215611b5457600080fd5b611b5d84611ad8565b9250611b6b60208501611ad8565b9150604084013590509250925092565b600060208284031215611b8d57600080fd5b5035919050565b600060208284031215611ba657600080fd5b610de982611ad8565b634e487b7160e01b600052602160045260246000fd5b60038110611bd557611bd5611baf565b9052565b602081016105478284611bc5565b600080600080600080600060e0888a031215611c0257600080fd5b611c0b88611ad8565b9650611c1960208901611ad8565b95506040880135945060608801359350608088013560ff81168114611c3d57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611c6d57600080fd5b611c7683611ad8565b915061158760208401611ad8565b600181811c90821680611c9857607f821691505b6020821081036107ee57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615611ce857611ce8611cb8565b500290565b60008219821115611d0057611d00611cb8565b500190565b60008451611d17818460208901611a79565b8083019050601760f91b8082528551611d37816001850160208a01611a79565b60019201918201528351611d52816002840160208801611a79565b0160020195945050505050565b6001600160a01b0385811682526080820190611d7e6020840187611bc5565b60028510611d8e57611d8e611baf565b84604084015280841660608401525095945050505050565b600060208284031215611db857600080fd5b5051919050565b600082821015611dd157611dd1611cb8565b500390565b600060018201611de857611de8611cb8565b5060010190565b634e487b7160e01b600052601260045260246000fd5b600082611e1457611e14611def565b500490565b634e487b7160e01b600052604160045260246000fd5b600082611e3e57611e3e611def565b500690565b634e487b7160e01b600052603260045260246000fd5b601f82111561152257600081815260208120601f850160051c81016020861015611e805750805b601f850160051c820191505b81811015611e9f57828155600101611e8c565b505050505050565b815167ffffffffffffffff811115611ec157611ec1611e19565b611ed581611ecf8454611c84565b84611e59565b602080601f831160018114611f0a5760008415611ef25750858301515b600019600386901b1c1916600185901b178555611e9f565b600085815260208120601f198616915b82811015611f3957888601518255948401946001909101908401611f1a565b5085821015611f575787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
}

// WETHRebasingABI is the input ABI used to generate the binding from.
// Deprecated: Use WETHRebasingMetaData.ABI instead.
var WETHRebasingABI = WETHRebasingMetaData.ABI

// WETHRebasingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WETHRebasingMetaData.Bin instead.
var WETHRebasingBin = WETHRebasingMetaData.Bin

// DeployWETHRebasing deploys a new Ethereum contract, binding an instance of WETHRebasing to it.
func DeployWETHRebasing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WETHRebasing, error) {
	parsed, err := WETHRebasingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WETHRebasingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WETHRebasing{WETHRebasingCaller: WETHRebasingCaller{contract: contract}, WETHRebasingTransactor: WETHRebasingTransactor{contract: contract}, WETHRebasingFilterer: WETHRebasingFilterer{contract: contract}}, nil
}

// WETHRebasing is an auto generated Go binding around an Ethereum contract.
type WETHRebasing struct {
	WETHRebasingCaller     // Read-only binding to the contract
	WETHRebasingTransactor // Write-only binding to the contract
	WETHRebasingFilterer   // Log filterer for contract events
}

// WETHRebasingCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHRebasingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHRebasingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHRebasingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHRebasingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHRebasingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHRebasingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHRebasingSession struct {
	Contract     *WETHRebasing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHRebasingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHRebasingCallerSession struct {
	Contract *WETHRebasingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WETHRebasingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHRebasingTransactorSession struct {
	Contract     *WETHRebasingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WETHRebasingRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHRebasingRaw struct {
	Contract *WETHRebasing // Generic contract binding to access the raw methods on
}

// WETHRebasingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHRebasingCallerRaw struct {
	Contract *WETHRebasingCaller // Generic read-only contract binding to access the raw methods on
}

// WETHRebasingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHRebasingTransactorRaw struct {
	Contract *WETHRebasingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETHRebasing creates a new instance of WETHRebasing, bound to a specific deployed contract.
func NewWETHRebasing(address common.Address, backend bind.ContractBackend) (*WETHRebasing, error) {
	contract, err := bindWETHRebasing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETHRebasing{WETHRebasingCaller: WETHRebasingCaller{contract: contract}, WETHRebasingTransactor: WETHRebasingTransactor{contract: contract}, WETHRebasingFilterer: WETHRebasingFilterer{contract: contract}}, nil
}

// NewWETHRebasingCaller creates a new read-only instance of WETHRebasing, bound to a specific deployed contract.
func NewWETHRebasingCaller(address common.Address, caller bind.ContractCaller) (*WETHRebasingCaller, error) {
	contract, err := bindWETHRebasing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingCaller{contract: contract}, nil
}

// NewWETHRebasingTransactor creates a new write-only instance of WETHRebasing, bound to a specific deployed contract.
func NewWETHRebasingTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHRebasingTransactor, error) {
	contract, err := bindWETHRebasing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingTransactor{contract: contract}, nil
}

// NewWETHRebasingFilterer creates a new log filterer instance of WETHRebasing, bound to a specific deployed contract.
func NewWETHRebasingFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHRebasingFilterer, error) {
	contract, err := bindWETHRebasing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingFilterer{contract: contract}, nil
}

// bindWETHRebasing binds a generic wrapper to an already deployed contract.
func bindWETHRebasing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WETHRebasingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHRebasing *WETHRebasingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHRebasing.Contract.WETHRebasingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHRebasing *WETHRebasingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHRebasing.Contract.WETHRebasingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHRebasing *WETHRebasingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHRebasing.Contract.WETHRebasingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHRebasing *WETHRebasingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHRebasing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHRebasing *WETHRebasingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHRebasing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHRebasing *WETHRebasingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHRebasing.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WETHRebasing *WETHRebasingCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WETHRebasing *WETHRebasingSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WETHRebasing.Contract.DOMAINSEPARATOR(&_WETHRebasing.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WETHRebasing *WETHRebasingCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WETHRebasing.Contract.DOMAINSEPARATOR(&_WETHRebasing.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_WETHRebasing *WETHRebasingCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_WETHRebasing *WETHRebasingSession) PERMITTYPEHASH() ([32]byte, error) {
	return _WETHRebasing.Contract.PERMITTYPEHASH(&_WETHRebasing.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_WETHRebasing *WETHRebasingCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _WETHRebasing.Contract.PERMITTYPEHASH(&_WETHRebasing.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.Allowance(&_WETHRebasing.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.Allowance(&_WETHRebasing.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 value)
func (_WETHRebasing *WETHRebasingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 value)
func (_WETHRebasing *WETHRebasingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.BalanceOf(&_WETHRebasing.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 value)
func (_WETHRebasing *WETHRebasingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.BalanceOf(&_WETHRebasing.CallOpts, account)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) Count() (*big.Int, error) {
	return _WETHRebasing.Contract.Count(&_WETHRebasing.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) Count() (*big.Int, error) {
	return _WETHRebasing.Contract.Count(&_WETHRebasing.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETHRebasing *WETHRebasingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETHRebasing *WETHRebasingSession) Decimals() (uint8, error) {
	return _WETHRebasing.Contract.Decimals(&_WETHRebasing.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETHRebasing *WETHRebasingCallerSession) Decimals() (uint8, error) {
	return _WETHRebasing.Contract.Decimals(&_WETHRebasing.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) GetClaimableAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "getClaimableAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) GetClaimableAmount(account common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.GetClaimableAmount(&_WETHRebasing.CallOpts, account)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) GetClaimableAmount(account common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.GetClaimableAmount(&_WETHRebasing.CallOpts, account)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address account) view returns(uint8)
func (_WETHRebasing *WETHRebasingCaller) GetConfiguration(opts *bind.CallOpts, account common.Address) (uint8, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "getConfiguration", account)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address account) view returns(uint8)
func (_WETHRebasing *WETHRebasingSession) GetConfiguration(account common.Address) (uint8, error) {
	return _WETHRebasing.Contract.GetConfiguration(&_WETHRebasing.CallOpts, account)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address account) view returns(uint8)
func (_WETHRebasing *WETHRebasingCallerSession) GetConfiguration(account common.Address) (uint8, error) {
	return _WETHRebasing.Contract.GetConfiguration(&_WETHRebasing.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETHRebasing *WETHRebasingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETHRebasing *WETHRebasingSession) Name() (string, error) {
	return _WETHRebasing.Contract.Name(&_WETHRebasing.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETHRebasing *WETHRebasingCallerSession) Name() (string, error) {
	return _WETHRebasing.Contract.Name(&_WETHRebasing.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.Nonces(&_WETHRebasing.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WETHRebasing.Contract.Nonces(&_WETHRebasing.CallOpts, owner)
}

// SharePrice is a free data retrieval call binding the contract method 0x87269729.
//
// Solidity: function sharePrice() view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) SharePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "sharePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharePrice is a free data retrieval call binding the contract method 0x87269729.
//
// Solidity: function sharePrice() view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) SharePrice() (*big.Int, error) {
	return _WETHRebasing.Contract.SharePrice(&_WETHRebasing.CallOpts)
}

// SharePrice is a free data retrieval call binding the contract method 0x87269729.
//
// Solidity: function sharePrice() view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) SharePrice() (*big.Int, error) {
	return _WETHRebasing.Contract.SharePrice(&_WETHRebasing.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETHRebasing *WETHRebasingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETHRebasing *WETHRebasingSession) Symbol() (string, error) {
	return _WETHRebasing.Contract.Symbol(&_WETHRebasing.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETHRebasing *WETHRebasingCallerSession) Symbol() (string, error) {
	return _WETHRebasing.Contract.Symbol(&_WETHRebasing.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) TotalSupply() (*big.Int, error) {
	return _WETHRebasing.Contract.TotalSupply(&_WETHRebasing.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) TotalSupply() (*big.Int, error) {
	return _WETHRebasing.Contract.TotalSupply(&_WETHRebasing.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_WETHRebasing *WETHRebasingCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_WETHRebasing *WETHRebasingSession) Version() (string, error) {
	return _WETHRebasing.Contract.Version(&_WETHRebasing.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_WETHRebasing *WETHRebasingCallerSession) Version() (string, error) {
	return _WETHRebasing.Contract.Version(&_WETHRebasing.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Approve(&_WETHRebasing.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Approve(&_WETHRebasing.TransactOpts, spender, amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 amount) returns(uint256)
func (_WETHRebasing *WETHRebasingTransactor) Claim(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "claim", recipient, amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 amount) returns(uint256)
func (_WETHRebasing *WETHRebasingSession) Claim(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Claim(&_WETHRebasing.TransactOpts, recipient, amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 amount) returns(uint256)
func (_WETHRebasing *WETHRebasingTransactorSession) Claim(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Claim(&_WETHRebasing.TransactOpts, recipient, amount)
}

// Configure is a paid mutator transaction binding the contract method 0x1a33757d.
//
// Solidity: function configure(uint8 yieldMode) returns(uint256)
func (_WETHRebasing *WETHRebasingTransactor) Configure(opts *bind.TransactOpts, yieldMode uint8) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "configure", yieldMode)
}

// Configure is a paid mutator transaction binding the contract method 0x1a33757d.
//
// Solidity: function configure(uint8 yieldMode) returns(uint256)
func (_WETHRebasing *WETHRebasingSession) Configure(yieldMode uint8) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Configure(&_WETHRebasing.TransactOpts, yieldMode)
}

// Configure is a paid mutator transaction binding the contract method 0x1a33757d.
//
// Solidity: function configure(uint8 yieldMode) returns(uint256)
func (_WETHRebasing *WETHRebasingTransactorSession) Configure(yieldMode uint8) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Configure(&_WETHRebasing.TransactOpts, yieldMode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETHRebasing *WETHRebasingTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETHRebasing *WETHRebasingSession) Deposit() (*types.Transaction, error) {
	return _WETHRebasing.Contract.Deposit(&_WETHRebasing.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETHRebasing *WETHRebasingTransactorSession) Deposit() (*types.Transaction, error) {
	return _WETHRebasing.Contract.Deposit(&_WETHRebasing.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_WETHRebasing *WETHRebasingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_WETHRebasing *WETHRebasingSession) Initialize() (*types.Transaction, error) {
	return _WETHRebasing.Contract.Initialize(&_WETHRebasing.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_WETHRebasing *WETHRebasingTransactorSession) Initialize() (*types.Transaction, error) {
	return _WETHRebasing.Contract.Initialize(&_WETHRebasing.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WETHRebasing *WETHRebasingTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WETHRebasing *WETHRebasingSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Permit(&_WETHRebasing.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WETHRebasing *WETHRebasingTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Permit(&_WETHRebasing.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Transfer(&_WETHRebasing.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Transfer(&_WETHRebasing.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.TransferFrom(&_WETHRebasing.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WETHRebasing *WETHRebasingTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.TransferFrom(&_WETHRebasing.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETHRebasing *WETHRebasingTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETHRebasing *WETHRebasingSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Withdraw(&_WETHRebasing.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETHRebasing *WETHRebasingTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.Withdraw(&_WETHRebasing.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHRebasing *WETHRebasingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHRebasing.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHRebasing *WETHRebasingSession) Receive() (*types.Transaction, error) {
	return _WETHRebasing.Contract.Receive(&_WETHRebasing.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHRebasing *WETHRebasingTransactorSession) Receive() (*types.Transaction, error) {
	return _WETHRebasing.Contract.Receive(&_WETHRebasing.TransactOpts)
}

// WETHRebasingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WETHRebasing contract.
type WETHRebasingApprovalIterator struct {
	Event *WETHRebasingApproval // Event containing the contract specifics and raw log

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
func (it *WETHRebasingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingApproval)
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
		it.Event = new(WETHRebasingApproval)
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
func (it *WETHRebasingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingApproval represents a Approval event raised by the WETHRebasing contract.
type WETHRebasingApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETHRebasing *WETHRebasingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WETHRebasingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingApprovalIterator{contract: _WETHRebasing.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETHRebasing *WETHRebasingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETHRebasingApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingApproval)
				if err := _WETHRebasing.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETHRebasing *WETHRebasingFilterer) ParseApproval(log types.Log) (*WETHRebasingApproval, error) {
	event := new(WETHRebasingApproval)
	if err := _WETHRebasing.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETHRebasingClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the WETHRebasing contract.
type WETHRebasingClaimIterator struct {
	Event *WETHRebasingClaim // Event containing the contract specifics and raw log

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
func (it *WETHRebasingClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingClaim)
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
		it.Event = new(WETHRebasingClaim)
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
func (it *WETHRebasingClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingClaim represents a Claim event raised by the WETHRebasing contract.
type WETHRebasingClaim struct {
	Account   common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd987068.
//
// Solidity: event Claim(address indexed account, address indexed recipient, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) FilterClaim(opts *bind.FilterOpts, account []common.Address, recipient []common.Address) (*WETHRebasingClaimIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "Claim", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingClaimIterator{contract: _WETHRebasing.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd987068.
//
// Solidity: event Claim(address indexed account, address indexed recipient, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *WETHRebasingClaim, account []common.Address, recipient []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "Claim", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingClaim)
				if err := _WETHRebasing.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd987068.
//
// Solidity: event Claim(address indexed account, address indexed recipient, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) ParseClaim(log types.Log) (*WETHRebasingClaim, error) {
	event := new(WETHRebasingClaim)
	if err := _WETHRebasing.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETHRebasingConfigureIterator is returned from FilterConfigure and is used to iterate over the raw logs and unpacked data for Configure events raised by the WETHRebasing contract.
type WETHRebasingConfigureIterator struct {
	Event *WETHRebasingConfigure // Event containing the contract specifics and raw log

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
func (it *WETHRebasingConfigureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingConfigure)
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
		it.Event = new(WETHRebasingConfigure)
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
func (it *WETHRebasingConfigureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingConfigureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingConfigure represents a Configure event raised by the WETHRebasing contract.
type WETHRebasingConfigure struct {
	Account   common.Address
	YieldMode uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfigure is a free log retrieval operation binding the contract event 0xcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36.
//
// Solidity: event Configure(address indexed account, uint8 yieldMode)
func (_WETHRebasing *WETHRebasingFilterer) FilterConfigure(opts *bind.FilterOpts, account []common.Address) (*WETHRebasingConfigureIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "Configure", accountRule)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingConfigureIterator{contract: _WETHRebasing.contract, event: "Configure", logs: logs, sub: sub}, nil
}

// WatchConfigure is a free log subscription operation binding the contract event 0xcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36.
//
// Solidity: event Configure(address indexed account, uint8 yieldMode)
func (_WETHRebasing *WETHRebasingFilterer) WatchConfigure(opts *bind.WatchOpts, sink chan<- *WETHRebasingConfigure, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "Configure", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingConfigure)
				if err := _WETHRebasing.contract.UnpackLog(event, "Configure", log); err != nil {
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

// ParseConfigure is a log parse operation binding the contract event 0xcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36.
//
// Solidity: event Configure(address indexed account, uint8 yieldMode)
func (_WETHRebasing *WETHRebasingFilterer) ParseConfigure(log types.Log) (*WETHRebasingConfigure, error) {
	event := new(WETHRebasingConfigure)
	if err := _WETHRebasing.contract.UnpackLog(event, "Configure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETHRebasingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WETHRebasing contract.
type WETHRebasingDepositIterator struct {
	Event *WETHRebasingDeposit // Event containing the contract specifics and raw log

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
func (it *WETHRebasingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingDeposit)
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
		it.Event = new(WETHRebasingDeposit)
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
func (it *WETHRebasingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingDeposit represents a Deposit event raised by the WETHRebasing contract.
type WETHRebasingDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*WETHRebasingDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingDepositIterator{contract: _WETHRebasing.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WETHRebasingDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingDeposit)
				if err := _WETHRebasing.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) ParseDeposit(log types.Log) (*WETHRebasingDeposit, error) {
	event := new(WETHRebasingDeposit)
	if err := _WETHRebasing.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETHRebasingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WETHRebasing contract.
type WETHRebasingInitializedIterator struct {
	Event *WETHRebasingInitialized // Event containing the contract specifics and raw log

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
func (it *WETHRebasingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingInitialized)
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
		it.Event = new(WETHRebasingInitialized)
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
func (it *WETHRebasingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingInitialized represents a Initialized event raised by the WETHRebasing contract.
type WETHRebasingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WETHRebasing *WETHRebasingFilterer) FilterInitialized(opts *bind.FilterOpts) (*WETHRebasingInitializedIterator, error) {

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WETHRebasingInitializedIterator{contract: _WETHRebasing.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WETHRebasing *WETHRebasingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WETHRebasingInitialized) (event.Subscription, error) {

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingInitialized)
				if err := _WETHRebasing.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WETHRebasing *WETHRebasingFilterer) ParseInitialized(log types.Log) (*WETHRebasingInitialized, error) {
	event := new(WETHRebasingInitialized)
	if err := _WETHRebasing.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETHRebasingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WETHRebasing contract.
type WETHRebasingTransferIterator struct {
	Event *WETHRebasingTransfer // Event containing the contract specifics and raw log

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
func (it *WETHRebasingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingTransfer)
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
		it.Event = new(WETHRebasingTransfer)
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
func (it *WETHRebasingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingTransfer represents a Transfer event raised by the WETHRebasing contract.
type WETHRebasingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETHRebasing *WETHRebasingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WETHRebasingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingTransferIterator{contract: _WETHRebasing.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETHRebasing *WETHRebasingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETHRebasingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingTransfer)
				if err := _WETHRebasing.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETHRebasing *WETHRebasingFilterer) ParseTransfer(log types.Log) (*WETHRebasingTransfer, error) {
	event := new(WETHRebasingTransfer)
	if err := _WETHRebasing.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETHRebasingWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WETHRebasing contract.
type WETHRebasingWithdrawalIterator struct {
	Event *WETHRebasingWithdrawal // Event containing the contract specifics and raw log

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
func (it *WETHRebasingWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingWithdrawal)
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
		it.Event = new(WETHRebasingWithdrawal)
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
func (it *WETHRebasingWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingWithdrawal represents a Withdrawal event raised by the WETHRebasing contract.
type WETHRebasingWithdrawal struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed account, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) FilterWithdrawal(opts *bind.FilterOpts, account []common.Address) (*WETHRebasingWithdrawalIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "Withdrawal", accountRule)
	if err != nil {
		return nil, err
	}
	return &WETHRebasingWithdrawalIterator{contract: _WETHRebasing.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed account, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WETHRebasingWithdrawal, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "Withdrawal", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingWithdrawal)
				if err := _WETHRebasing.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed account, uint256 amount)
func (_WETHRebasing *WETHRebasingFilterer) ParseWithdrawal(log types.Log) (*WETHRebasingWithdrawal, error) {
	event := new(WETHRebasingWithdrawal)
	if err := _WETHRebasing.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
