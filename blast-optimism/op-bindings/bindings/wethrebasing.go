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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApproveFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"name\":\"DistributeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReporter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimableAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceIsInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumYieldMode\",\"name\":\"yieldMode\",\"type\":\"uint8\"}],\"name\":\"Configure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"yieldMode\",\"type\":\"uint8\"}],\"name\":\"configure\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61012060405234801561001157600080fd5b50604360981b608052601260a052600160c052600060e08190526101005261003761003c565b6100fb565b600054610100900460ff16156100a85760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100f9576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160a05160c05160e05161010051612276620001476000396000610827015260006107fe015260006107d5015260006102e5015260008181610395015261148001526122766000f3fe6080604052600436106101a05760003560e01c80637ae556b5116100ec578063aad3ec961161008a578063d505accf11610064578063d505accf146104cc578063dd62ed3e146104ec578063e12f3a6114610532578063e20ccec31461055257600080fd5b8063aad3ec9614610477578063c44b11f714610497578063d0e30db0146104c457600080fd5b806384b0196e116100c657806384b0196e1461040457806395d89b411461042c578063a035b1fe14610441578063a9059cbb1461045757600080fd5b80637ae556b5146103835780637ecebe00146103cf5780638129fc1c146103ef57600080fd5b80632e1a7d4d116101595780633644e515116101335780633644e5151461031957806354fd4d501461032e5780635b9af12b1461034357806370a082311461036357600080fd5b80632e1a7d4d1461027f57806330adf81f1461029f578063313ce567146102d357600080fd5b806306661abd146101b457806306fdde03146101d8578063095ea7b3146101fa57806318160ddd1461022a5780631a33757d1461023f57806323b872dd1461025f57600080fd5b366101af576101ad610568565b005b600080fd5b3480156101c057600080fd5b50609d545b6040519081526020015b60405180910390f35b3480156101e457600080fd5b506101ed6105b9565b6040516101cf9190611cdc565b34801561020657600080fd5b5061021a610215366004611d0b565b610647565b60405190151581526020016101cf565b34801561023657600080fd5b506101c5610661565b34801561024b57600080fd5b506101c561025a366004611d35565b610685565b34801561026b57600080fd5b5061021a61027a366004611d56565b6106db565b34801561028b57600080fd5b506101ad61029a366004611d92565b6106fd565b3480156102ab57600080fd5b506101c57f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b3480156102df57600080fd5b506103077f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101cf565b34801561032557600080fd5b506101c56107c4565b34801561033a57600080fd5b506101ed6107ce565b34801561034f57600080fd5b506101ad61035e366004611d92565b610871565b34801561036f57600080fd5b506101c561037e366004611dab565b61087d565b34801561038f57600080fd5b506103b77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101cf565b3480156103db57600080fd5b506101c56103ea366004611dab565b610909565b3480156103fb57600080fd5b506101ad610927565b34801561041057600080fd5b50610419610b4a565b6040516101cf9796959493929190611dc6565b34801561043857600080fd5b506101ed610be8565b34801561044d57600080fd5b506101c560685481565b34801561046357600080fd5b5061021a610472366004611d0b565b610bf5565b34801561048357600080fd5b506101c5610492366004611d0b565b610c0b565b3480156104a357600080fd5b506104b76104b2366004611dab565b610d87565b6040516101cf9190611e86565b6101ad610568565b3480156104d857600080fd5b506101ad6104e7366004611e94565b610da5565b3480156104f857600080fd5b506101c5610507366004611f07565b6001600160a01b03918216600090815260a26020908152604080832093909416825291909152205490565b34801561053e57600080fd5b506101c561054d366004611dab565b610f09565b34801561055e57600080fd5b506101c560695481565b336105738134610fa0565b806001600160a01b03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516105ae91815260200190565b60405180910390a250565b609a80546105c690611f3a565b80601f01602080910402602001604051908101604052809291908181526020018280546105f290611f3a565b801561063f5780601f106106145761010080835404028352916020019161063f565b820191906000526020600020905b81548152906001019060200180831161062257829003601f168201915b505050505081565b600033610655818585611008565b60019150505b92915050565b600060a054609d546068546106769190611f84565b6106809190611fa3565b905090565b600061069133836110b8565b336001600160a01b03167fcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36836040516106ca9190611e86565b60405180910390a261065b3361087d565b60006106e88433846111ea565b6106f3848484611247565b5060019392505050565b3361070881836112ee565b6000816001600160a01b03168360405160006040518083038185875af1925050503d8060008114610755576040519150601f19603f3d011682016040523d82523d6000602084013e61075a565b606091505b505090508061077c5760405163b12d13eb60e01b815260040160405180910390fd5b816001600160a01b03167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65846040516107b791815260200190565b60405180910390a2505050565b6000610680611362565b60606107f97f000000000000000000000000000000000000000000000000000000000000000061136c565b6108227f000000000000000000000000000000000000000000000000000000000000000061136c565b61084b7f000000000000000000000000000000000000000000000000000000000000000061136c565b60405160200161085d93929190611fbb565b604051602081830303815290604052905090565b61087a81611475565b50565b6001600160a01b038116600090815260a1602052604081205460ff16818160028111156108ac576108ac611e5c565b036108e7576001600160a01b0383166000908152609c6020908152604080832054609e909252909120546108e09190611531565b9150610903565b6001600160a01b0383166000908152609f602052604090205491505b50919050565b6001600160a01b03811660009081526035602052604081205461065b565b600054610100900460ff16158080156109475750600054600160ff909116105b806109615750303b158015610961575060005460ff166001145b6109c95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156109ec576000805461ff0019166101001790555b610a9d6040518060400160405280600d81526020016c2bb930b83832b21022ba3432b960991b815250604051806040016040528060048152602001630ae8aa8960e31b815250604360981b6001600160a01b031663a035b1fe6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a989190612015565b61154c565b60405163099005e760e31b81526002604360981b0190634c802f3890610ad0903090600090819061dead9060040161202e565b600060405180830381600087803b158015610aea57600080fd5b505af1158015610afe573d6000803e3d6000fd5b50505050801561087a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6000606080600080600060606001546000801b148015610b6a5750600254155b610bae5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064016109c0565b610bb661159e565b610bbe611630565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b609b80546105c690611f3a565b6000610c02338484611247565b50600192915050565b6000336001600160a01b038416610c355760405163088d6c0d60e31b815260040160405180910390fd5b6002610c4082610d87565b6002811115610c5157610c51611e5c565b14610c6f5760405163ebf953c760e01b815260040160405180910390fd5b6001600160a01b0381166000908152609c6020908152604080832054609e909252822054610c9d9190611531565b6001600160a01b0383166000908152609f602052604081205491925090610cc4908361208b565b905080851115610ce757604051631e9acf1760e31b815260040160405180910390fd5b600080610cfc610cf7888661208b565b61163f565b91509150610d31858383609f60008a6001600160a01b03166001600160a01b0316815260200190815260200160002054611674565b610d3b8888610fa0565b6040518781526001600160a01b0389169033907f70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd9870689060200160405180910390a350949695505050505050565b6001600160a01b0316600090815260a1602052604090205460ff1690565b83421115610df55760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016109c0565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610e248c611710565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610e7f82611736565b90506000610e8f82878787611763565b9050896001600160a01b0316816001600160a01b031614610ef25760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016109c0565b610efd8a8a8a611008565b50505050505050505050565b60006002610f1683610d87565b6002811115610f2757610f27611e5c565b14610f455760405163ebf953c760e01b815260040160405180910390fd5b6001600160a01b0382166000908152609c6020908152604080832054609e909252822054610f739190611531565b6001600160a01b0384166000908152609f6020526040902054909150610f99908261208b565b9392505050565b600081610fac8461087d565b610fb69190611fa3565b9050610fc48382600061178b565b6000610fcf84610d87565b90506001816002811115610fe557610fe5611e5c565b03611002578260a06000828254610ffc9190611fa3565b90915550505b50505050565b6001600160a01b03831661102f5760405163eb3b083560e01b815260040160405180910390fd5b6001600160a01b0382166110565760405163076e33c360e31b815260040160405180910390fd5b6001600160a01b03838116600081815260a2602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006110c383610d87565b9050600060028260028111156110db576110db611e5c565b03611116576001600160a01b0384166000908152609c6020908152604080832054609e9092529091205461110f9190611531565b9050611122565b61111f8461087d565b90505b6001600160a01b038416600090815260a160205260409020805484919060ff1916600183600281111561115757611157611e5c565b02179055506001600160a01b0384166000908152609f60205260409020546111818583600161178b565b600183600281111561119557611195611e5c565b036111b2578060a060008282546111ac919061208b565b90915550505b60018460028111156111c6576111c6611e5c565b036111e3578160a060008282546111dd9190611fa3565b90915550505b5050505050565b6001600160a01b03838116600090815260a260209081526040808320938616835292905220546000198114611002578082111561123a576040516313be252b60e01b815260040160405180910390fd5b6110028484848403611008565b6001600160a01b03831661126e57604051630b07e54560e11b815260040160405180910390fd5b6001600160a01b03821661129557604051633a954ecd60e21b815260040160405180910390fd5b61129f83826112ee565b6112a98282610fa0565b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516110ab91815260200190565b60006112f98361087d565b90508082111561131c57604051631e9acf1760e31b815260040160405180910390fd5b61132a83838303600061178b565b600061133584610d87565b9050600181600281111561134b5761134b611e5c565b03611002578260a06000828254610ffc919061208b565b6000610680611892565b6060816000036113935750506040805180820190915260018152600360fc1b602082015290565b8160005b81156113bd57806113a7816120a2565b91506113b69050600a836120d1565b9150611397565b60008167ffffffffffffffff8111156113d8576113d8612075565b6040519080825280601f01601f191660200182016040528015611402576020820181803683370190505b5090505b841561146d5761141760018361208b565b9150611424600a866120e5565b61142f906030611fa3565b60f81b818381518110611444576114446120f9565b60200101906001600160f81b031916908160001a905350611466600a866120d1565b9450611406565b949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146114be57604051631d73770560e11b815260040160405180910390fd5b6000609d546068546114d09190611f84565b9050600060a05482476114e3919061208b565b6114ed919061208b565b9050609d548110806114ff5750609d54155b1561150957505050565b609d5461151690826120d1565b606860008282546115279190611fa3565b9091555050505050565b600081836068546115429190611f84565b610f999190611fa3565b600054610100900460ff166115735760405162461bcd60e51b81526004016109c09061210f565b61157c83611906565b61158581611950565b609a61159184826121a9565b50609b61100283826121a9565b6060600380546115ad90611f3a565b80601f01602080910402602001604051908101604052809291908181526020018280546115d990611f3a565b80156116265780601f106115fb57610100808354040283529160200191611626565b820191906000526020600020905b81548152906001019060200180831161160957829003601f168201915b5050505050905090565b6060600480546115ad90611f3a565b600080606854600003611650575091565b60685461165d90846120d1565b91506068548361166d91906120e5565b9050915091565b6001600160a01b0384166000908152609c6020526040902054609d5461169b908590611fa3565b6116a5919061208b565b609d556001600160a01b0384166000908152609e602052604090205460a0546116cf908490611fa3565b6116d9919061208b565b60a0556001600160a01b039093166000908152609c6020908152604080832094909455609e815283822092909255609f9091522055565b6001600160a01b0381166000908152603560205260409020805460018101825590610903565b600061065b611743611362565b8360405161190160f01b8152600281019290925260228201526042902090565b60008060006117748787878761199d565b9150915061178181611a61565b5095945050505050565b60008060008061179a87610d87565b905060008160028111156117b0576117b0611e5c565b036117c8576117be8661163f565b909450925061187d565b60018160028111156117dc576117dc611e5c565b036117e95785915061187d565b60028160028111156117fd576117fd611e5c565b0361187d57859150818561186d576001600160a01b0388166000908152609c6020908152604080832054609e9092529091205461183a9190611531565b6001600160a01b0389166000908152609f60205260409020549091506118608883611fa3565b61186a919061208b565b90505b6118768161163f565b9095509350505b61188987858585611674565b50505050505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6118bd611bab565b6118c5611c04565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b600054610100900460ff1661192d5760405162461bcd60e51b81526004016109c09061210f565b61087a81604051806040016040528060018152602001603160f81b815250611c35565b600054610100900460ff166119775760405162461bcd60e51b81526004016109c09061210f565b606854156119985760405163131cb46d60e21b815260040160405180910390fd5b606855565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156119d45750600090506003611a58565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611a28573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611a5157600060019250925050611a58565b9150600090505b94509492505050565b6000816004811115611a7557611a75611e5c565b03611a7d5750565b6001816004811115611a9157611a91611e5c565b03611ade5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016109c0565b6002816004811115611af257611af2611e5c565b03611b3f5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016109c0565b6003816004811115611b5357611b53611e5c565b0361087a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016109c0565b600080611bb661159e565b805190915015611bcd578051602090910120919050565b6001548015611bdc5792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b600080611c0f611630565b805190915015611c26578051602090910120919050565b6002548015611bdc5792915050565b600054610100900460ff16611c5c5760405162461bcd60e51b81526004016109c09061210f565b6003611c6883826121a9565b506004611c7582826121a9565b50506000600181905560025550565b60005b83811015611c9f578181015183820152602001611c87565b838111156110025750506000910152565b60008151808452611cc8816020860160208601611c84565b601f01601f19169290920160200192915050565b602081526000610f996020830184611cb0565b80356001600160a01b0381168114611d0657600080fd5b919050565b60008060408385031215611d1e57600080fd5b611d2783611cef565b946020939093013593505050565b600060208284031215611d4757600080fd5b813560038110610f9957600080fd5b600080600060608486031215611d6b57600080fd5b611d7484611cef565b9250611d8260208501611cef565b9150604084013590509250925092565b600060208284031215611da457600080fd5b5035919050565b600060208284031215611dbd57600080fd5b610f9982611cef565b60ff60f81b881681526000602060e081840152611de660e084018a611cb0565b8381036040850152611df8818a611cb0565b606085018990526001600160a01b038816608086015260a0850187905284810360c0860152855180825283870192509083019060005b81811015611e4a57835183529284019291840191600101611e2e565b50909c9b505050505050505050505050565b634e487b7160e01b600052602160045260246000fd5b60038110611e8257611e82611e5c565b9052565b6020810161065b8284611e72565b600080600080600080600060e0888a031215611eaf57600080fd5b611eb888611cef565b9650611ec660208901611cef565b95506040880135945060608801359350608088013560ff81168114611eea57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611f1a57600080fd5b611f2383611cef565b9150611f3160208401611cef565b90509250929050565b600181811c90821680611f4e57607f821691505b60208210810361090357634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615611f9e57611f9e611f6e565b500290565b60008219821115611fb657611fb6611f6e565b500190565b60008451611fcd818460208901611c84565b8083019050601760f91b8082528551611fed816001850160208a01611c84565b60019201918201528351612008816002840160208801611c84565b0160020195945050505050565b60006020828403121561202757600080fd5b5051919050565b6001600160a01b038581168252608082019061204d6020840187611e72565b6002851061205d5761205d611e5c565b84604084015280841660608401525095945050505050565b634e487b7160e01b600052604160045260246000fd5b60008282101561209d5761209d611f6e565b500390565b6000600182016120b4576120b4611f6e565b5060010190565b634e487b7160e01b600052601260045260246000fd5b6000826120e0576120e06120bb565b500490565b6000826120f4576120f46120bb565b500690565b634e487b7160e01b600052603260045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b601f8211156121a457600081815260208120601f850160051c810160208610156121815750805b601f850160051c820191505b818110156121a05782815560010161218d565b5050505b505050565b815167ffffffffffffffff8111156121c3576121c3612075565b6121d7816121d18454611f3a565b8461215a565b602080601f83116001811461220c57600084156121f45750858301515b600019600386901b1c1916600185901b1785556121a0565b600085815260208120601f198616915b8281101561223b5788860151825594840194600190910190840161221c565b50858210156122595787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
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

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_WETHRebasing *WETHRebasingCaller) REPORTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "REPORTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_WETHRebasing *WETHRebasingSession) REPORTER() (common.Address, error) {
	return _WETHRebasing.Contract.REPORTER(&_WETHRebasing.CallOpts)
}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_WETHRebasing *WETHRebasingCallerSession) REPORTER() (common.Address, error) {
	return _WETHRebasing.Contract.REPORTER(&_WETHRebasing.CallOpts)
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

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WETHRebasing *WETHRebasingCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WETHRebasing *WETHRebasingSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WETHRebasing.Contract.Eip712Domain(&_WETHRebasing.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WETHRebasing *WETHRebasingCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WETHRebasing.Contract.Eip712Domain(&_WETHRebasing.CallOpts)
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

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) Pending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "pending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) Pending() (*big.Int, error) {
	return _WETHRebasing.Contract.Pending(&_WETHRebasing.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) Pending() (*big.Int, error) {
	return _WETHRebasing.Contract.Pending(&_WETHRebasing.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_WETHRebasing *WETHRebasingCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETHRebasing.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_WETHRebasing *WETHRebasingSession) Price() (*big.Int, error) {
	return _WETHRebasing.Contract.Price(&_WETHRebasing.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_WETHRebasing *WETHRebasingCallerSession) Price() (*big.Int, error) {
	return _WETHRebasing.Contract.Price(&_WETHRebasing.CallOpts)
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

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_WETHRebasing *WETHRebasingTransactor) AddValue(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "addValue", value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_WETHRebasing *WETHRebasingSession) AddValue(value *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.AddValue(&_WETHRebasing.TransactOpts, value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_WETHRebasing *WETHRebasingTransactorSession) AddValue(value *big.Int) (*types.Transaction, error) {
	return _WETHRebasing.Contract.AddValue(&_WETHRebasing.TransactOpts, value)
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
// Solidity: function initialize() returns()
func (_WETHRebasing *WETHRebasingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHRebasing.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_WETHRebasing *WETHRebasingSession) Initialize() (*types.Transaction, error) {
	return _WETHRebasing.Contract.Initialize(&_WETHRebasing.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
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

// WETHRebasingEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the WETHRebasing contract.
type WETHRebasingEIP712DomainChangedIterator struct {
	Event *WETHRebasingEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *WETHRebasingEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingEIP712DomainChanged)
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
		it.Event = new(WETHRebasingEIP712DomainChanged)
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
func (it *WETHRebasingEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingEIP712DomainChanged represents a EIP712DomainChanged event raised by the WETHRebasing contract.
type WETHRebasingEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WETHRebasing *WETHRebasingFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*WETHRebasingEIP712DomainChangedIterator, error) {

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &WETHRebasingEIP712DomainChangedIterator{contract: _WETHRebasing.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WETHRebasing *WETHRebasingFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *WETHRebasingEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingEIP712DomainChanged)
				if err := _WETHRebasing.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WETHRebasing *WETHRebasingFilterer) ParseEIP712DomainChanged(log types.Log) (*WETHRebasingEIP712DomainChanged, error) {
	event := new(WETHRebasingEIP712DomainChanged)
	if err := _WETHRebasing.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// WETHRebasingNewPriceIterator is returned from FilterNewPrice and is used to iterate over the raw logs and unpacked data for NewPrice events raised by the WETHRebasing contract.
type WETHRebasingNewPriceIterator struct {
	Event *WETHRebasingNewPrice // Event containing the contract specifics and raw log

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
func (it *WETHRebasingNewPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHRebasingNewPrice)
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
		it.Event = new(WETHRebasingNewPrice)
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
func (it *WETHRebasingNewPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHRebasingNewPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHRebasingNewPrice represents a NewPrice event raised by the WETHRebasing contract.
type WETHRebasingNewPrice struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewPrice is a free log retrieval operation binding the contract event 0x270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b7.
//
// Solidity: event NewPrice(uint256 price)
func (_WETHRebasing *WETHRebasingFilterer) FilterNewPrice(opts *bind.FilterOpts) (*WETHRebasingNewPriceIterator, error) {

	logs, sub, err := _WETHRebasing.contract.FilterLogs(opts, "NewPrice")
	if err != nil {
		return nil, err
	}
	return &WETHRebasingNewPriceIterator{contract: _WETHRebasing.contract, event: "NewPrice", logs: logs, sub: sub}, nil
}

// WatchNewPrice is a free log subscription operation binding the contract event 0x270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b7.
//
// Solidity: event NewPrice(uint256 price)
func (_WETHRebasing *WETHRebasingFilterer) WatchNewPrice(opts *bind.WatchOpts, sink chan<- *WETHRebasingNewPrice) (event.Subscription, error) {

	logs, sub, err := _WETHRebasing.contract.WatchLogs(opts, "NewPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHRebasingNewPrice)
				if err := _WETHRebasing.contract.UnpackLog(event, "NewPrice", log); err != nil {
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
func (_WETHRebasing *WETHRebasingFilterer) ParseNewPrice(log types.Log) (*WETHRebasingNewPrice, error) {
	event := new(WETHRebasingNewPrice)
	if err := _WETHRebasing.contract.UnpackLog(event, "NewPrice", log); err != nil {
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
