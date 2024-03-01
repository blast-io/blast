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

// USDBMetaData contains all meta data concerning the USDB contract.
var USDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdYieldManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApproveFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerIsNotBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"name\":\"DistributeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReporter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimableAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceIsInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumYieldMode\",\"name\":\"yieldMode\",\"type\":\"uint8\"}],\"name\":\"Configure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REMOTE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"yieldMode\",\"type\":\"uint8\"}],\"name\":\"configure\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b50604051620024f1380380620024f1833981016040819052620000359162000156565b6001600160a01b03808416608052601260a052600160c052600060e081905261010052828116610140528116610120526200006f62000078565b505050620001a0565b600054610100900460ff1615620000e55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000137576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200015157600080fd5b919050565b6000806000606084860312156200016c57600080fd5b620001778462000139565b9250620001876020850162000139565b9150620001976040850162000139565b90509250925092565b60805160a05160c05160e0516101005161012051610140516122d36200021e600039600081816104c6015281816104ef015281816106a80152610b26015260008181610222015261044b015260006107b60152600061078d01526000610764015260006102f001526000818161037401526113e601526122d36000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c80637ecebe001161010f578063c44b11f7116100a2578063e12f3a6111610071578063e12f3a61146104a8578063e20ccec3146104bb578063e78cea92146104c4578063ee9a31a2146104ea57600080fd5b8063c44b11f714610416578063d505accf14610436578063d6c0b2c414610449578063dd62ed3e1461046f57600080fd5b80639dc29fac116100de5780639dc29fac146103d4578063a035b1fe146103e7578063a9059cbb146103f0578063aad3ec961461040357600080fd5b80637ecebe00146103965780638129fc1c146103a957806384b0196e146103b157806395d89b41146103cc57600080fd5b806330adf81f1161018757806354fd4d501161015657806354fd4d50146103415780635b9af12b1461034957806370a082311461035c5780637ae556b51461036f57600080fd5b806330adf81f146102c4578063313ce567146102eb5780633644e5151461032457806340c10f191461032c57600080fd5b8063095ea7b3116101c3578063095ea7b31461028357806318160ddd146102965780631a33757d1461029e57806323b872dd146102b157600080fd5b806301ffc9a7146101f5578063033964be1461021d57806306661abd1461025c57806306fdde031461026e575b600080fd5b610208610203366004611cd0565b610511565b60405190151581526020015b60405180910390f35b6102447f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610214565b609d545b604051908152602001610214565b61027661054f565b6040516102149190611d52565b610208610291366004611d81565b6105dd565b6102606105f7565b6102606102ac366004611dab565b61061b565b6102086102bf366004611dcc565b610671565b6102607f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b6103127f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff9091168152602001610214565b610260610693565b61033f61033a366004611d81565b61069d565b005b61027661075d565b61033f610357366004611e08565b610800565b61026061036a366004611e21565b61080c565b6102447f000000000000000000000000000000000000000000000000000000000000000081565b6102606103a4366004611e21565b610898565b61033f6108b6565b6103b9610a70565b6040516102149796959493929190611e3c565b610276610b0e565b61033f6103e2366004611d81565b610b1b565b61026060685481565b6102086103fe366004611d81565b610bd3565b610260610411366004611d81565b610be9565b610429610424366004611e21565b610d65565b6040516102149190611efc565b61033f610444366004611f0a565b610d83565b7f0000000000000000000000000000000000000000000000000000000000000000610244565b61026061047d366004611f7d565b6001600160a01b03918216600090815260a26020908152604080832093909416825291909152205490565b6102606104b6366004611e21565b610ee7565b61026060695481565b7f0000000000000000000000000000000000000000000000000000000000000000610244565b6102447f000000000000000000000000000000000000000000000000000000000000000081565b60006301ffc9a760e01b63ec4fc8e360e01b6001600160e01b0319841682148061054757506001600160e01b0319848116908216145b949350505050565b609a805461055c90611fb0565b80601f016020809104026020016040519081016040528092919081815260200182805461058890611fb0565b80156105d55780601f106105aa576101008083540402835291602001916105d5565b820191906000526020600020905b8154815290600101906020018083116105b857829003601f168201915b505050505081565b6000336105eb818585610f7e565b60019150505b92915050565b600060a054609d5460685461060c9190611ffa565b6106169190612019565b905090565b6000610627338361102e565b336001600160a01b03167fcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36836040516106609190611efc565b60405180910390a26105f13361080c565b600061067e843384611160565b6106898484846111c3565b5060019392505050565b600061061661126a565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106e65760405163ea0e1ccb60e01b815260040160405180910390fd5b6001600160a01b03821661070d57604051633a954ecd60e21b815260040160405180910390fd5b6107178282611274565b6040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a35050565b60606107887f00000000000000000000000000000000000000000000000000000000000000006112db565b6107b17f00000000000000000000000000000000000000000000000000000000000000006112db565b6107da7f00000000000000000000000000000000000000000000000000000000000000006112db565b6040516020016107ec93929190612031565b604051602081830303815290604052905090565b610809816113dc565b50565b6001600160a01b038116600090815260a1602052604081205460ff168181600281111561083b5761083b611ed2565b03610876576001600160a01b0383166000908152609c6020908152604080832054609e9092529091205461086f919061146f565b9150610892565b6001600160a01b0383166000908152609f602052604090205491505b50919050565b6001600160a01b0381166000908152603560205260408120546105f1565b600054610100900460ff16158080156108d65750600054600160ff909116105b806108f05750303b1580156108f0575060005460ff166001145b6109585760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801561097b576000805461ff0019166101001790555b6109c2604051806040016040528060048152602001632aa9a22160e11b815250604051806040016040528060048152602001632aa9a22160e11b815250633b9aca0061148a565b60405163099005e760e31b81526002604360981b0190634c802f38906109f690309060019060009061dead9060040161208b565b600060405180830381600087803b158015610a1057600080fd5b505af1158015610a24573d6000803e3d6000fd5b505050508015610809576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6000606080600080600060606001546000801b148015610a905750600254155b610ad45760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b604482015260640161094f565b610adc6114dc565b610ae461156e565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b609b805461055c90611fb0565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b645760405163ea0e1ccb60e01b815260040160405180910390fd5b6001600160a01b038216610b8b57604051630b07e54560e11b815260040160405180910390fd5b610b95828261157d565b6040518181526000906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610751565b6000610be03384846111c3565b50600192915050565b6000336001600160a01b038416610c135760405163088d6c0d60e31b815260040160405180910390fd5b6002610c1e82610d65565b6002811115610c2f57610c2f611ed2565b14610c4d5760405163ebf953c760e01b815260040160405180910390fd5b6001600160a01b0381166000908152609c6020908152604080832054609e909252822054610c7b919061146f565b6001600160a01b0383166000908152609f602052604081205491925090610ca290836120e8565b905080851115610cc557604051631e9acf1760e31b815260040160405180910390fd5b600080610cda610cd588866120e8565b6115f1565b91509150610d0f858383609f60008a6001600160a01b03166001600160a01b0316815260200190815260200160002054611626565b610d198888611274565b6040518781526001600160a01b0389169033907f70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd9870689060200160405180910390a350949695505050505050565b6001600160a01b0316600090815260a1602052604090205460ff1690565b83421115610dd35760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e65000000604482015260640161094f565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610e028c6116c2565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610e5d826116e8565b90506000610e6d82878787611715565b9050896001600160a01b0316816001600160a01b031614610ed05760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e61747572650000604482015260640161094f565b610edb8a8a8a610f7e565b50505050505050505050565b60006002610ef483610d65565b6002811115610f0557610f05611ed2565b14610f235760405163ebf953c760e01b815260040160405180910390fd5b6001600160a01b0382166000908152609c6020908152604080832054609e909252822054610f51919061146f565b6001600160a01b0384166000908152609f6020526040902054909150610f7790826120e8565b9392505050565b6001600160a01b038316610fa55760405163eb3b083560e01b815260040160405180910390fd5b6001600160a01b038216610fcc5760405163076e33c360e31b815260040160405180910390fd5b6001600160a01b03838116600081815260a2602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b600061103983610d65565b90506000600282600281111561105157611051611ed2565b0361108c576001600160a01b0384166000908152609c6020908152604080832054609e90925290912054611085919061146f565b9050611098565b6110958461080c565b90505b6001600160a01b038416600090815260a160205260409020805484919060ff191660018360028111156110cd576110cd611ed2565b02179055506001600160a01b0384166000908152609f60205260409020546110f78583600161173d565b600183600281111561110b5761110b611ed2565b03611128578060a0600082825461112291906120e8565b90915550505b600184600281111561113c5761113c611ed2565b03611159578160a060008282546111539190612019565b90915550505b5050505050565b6001600160a01b03838116600090815260a2602090815260408083209386168352929052205460001981146111bd57808211156111b0576040516313be252b60e01b815260040160405180910390fd5b6111bd8484848403610f7e565b50505050565b6001600160a01b0383166111ea57604051630b07e54560e11b815260040160405180910390fd5b6001600160a01b03821661121157604051633a954ecd60e21b815260040160405180910390fd5b61121b838261157d565b6112258282611274565b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161102191815260200190565b6000610616611844565b6000816112808461080c565b61128a9190612019565b90506112988382600061173d565b60006112a384610d65565b905060018160028111156112b9576112b9611ed2565b036111bd578260a060008282546112d09190612019565b909155505050505050565b6060816000036113025750506040805180820190915260018152600360fc1b602082015290565b8160005b811561132c5780611316816120ff565b91506113259050600a8361212e565b9150611306565b60008167ffffffffffffffff811115611347576113476120d2565b6040519080825280601f01601f191660200182016040528015611371576020820181803683370190505b5090505b8415610547576113866001836120e8565b9150611393600a86612142565b61139e906030612019565b60f81b8183815181106113b3576113b3612156565b60200101906001600160f81b031916908160001a9053506113d5600a8661212e565b9450611375565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167311110000000000000000000000000000000011101933016001600160a01b03161461144557604051631d73770560e11b815260040160405180910390fd5b801561146357806069600082825461145d9190612019565b90915550505b61146b6118b8565b5050565b600081836068546114809190611ffa565b610f779190612019565b600054610100900460ff166114b15760405162461bcd60e51b815260040161094f9061216c565b6114ba83611952565b6114c38161199c565b609a6114cf8482612206565b50609b6111bd8382612206565b6060600380546114eb90611fb0565b80601f016020809104026020016040519081016040528092919081815260200182805461151790611fb0565b80156115645780601f1061153957610100808354040283529160200191611564565b820191906000526020600020905b81548152906001019060200180831161154757829003601f168201915b5050505050905090565b6060600480546114eb90611fb0565b60006115888361080c565b9050808211156115ab57604051631e9acf1760e31b815260040160405180910390fd5b6115b983838303600061173d565b60006115c484610d65565b905060018160028111156115da576115da611ed2565b036111bd578260a060008282546112d091906120e8565b600080606854600003611602575091565b60685461160f908461212e565b91506068548361161f9190612142565b9050915091565b6001600160a01b0384166000908152609c6020526040902054609d5461164d908590612019565b61165791906120e8565b609d556001600160a01b0384166000908152609e602052604090205460a054611681908490612019565b61168b91906120e8565b60a0556001600160a01b039093166000908152609c6020908152604080832094909455609e815283822092909255609f9091522055565b6001600160a01b0381166000908152603560205260409020805460018101825590610892565b60006105f16116f561126a565b8360405161190160f01b8152600281019290925260228201526042902090565b6000806000611726878787876119e9565b9150915061173381611aad565b5095945050505050565b60008060008061174c87610d65565b9050600081600281111561176257611762611ed2565b0361177a57611770866115f1565b909450925061182f565b600181600281111561178e5761178e611ed2565b0361179b5785915061182f565b60028160028111156117af576117af611ed2565b0361182f57859150818561181f576001600160a01b0388166000908152609c6020908152604080832054609e909252909120546117ec919061146f565b6001600160a01b0389166000908152609f60205260409020549091506118128883612019565b61181c91906120e8565b90505b611828816115f1565b9095509350505b61183b87858585611626565b50505050505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61186f611bf7565b611877611c50565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b60006118c3609d5490565b60695410806118d25750609d54155b156118dd5750600090565b609d546069546118ed919061212e565b606860008282546118fe9190612019565b9091555050609d546069546119139190612142565b6069556068546040519081527f270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b79060200160405180910390a150600190565b600054610100900460ff166119795760405162461bcd60e51b815260040161094f9061216c565b61080981604051806040016040528060018152602001603160f81b815250611c81565b600054610100900460ff166119c35760405162461bcd60e51b815260040161094f9061216c565b606854156119e45760405163131cb46d60e21b815260040160405180910390fd5b606855565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611a205750600090506003611aa4565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611a74573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611a9d57600060019250925050611aa4565b9150600090505b94509492505050565b6000816004811115611ac157611ac1611ed2565b03611ac95750565b6001816004811115611add57611add611ed2565b03611b2a5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161094f565b6002816004811115611b3e57611b3e611ed2565b03611b8b5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161094f565b6003816004811115611b9f57611b9f611ed2565b036108095760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161094f565b600080611c026114dc565b805190915015611c19578051602090910120919050565b6001548015611c285792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b600080611c5b61156e565b805190915015611c72578051602090910120919050565b6002548015611c285792915050565b600054610100900460ff16611ca85760405162461bcd60e51b815260040161094f9061216c565b6003611cb48382612206565b506004611cc18282612206565b50506000600181905560025550565b600060208284031215611ce257600080fd5b81356001600160e01b031981168114610f7757600080fd5b60005b83811015611d15578181015183820152602001611cfd565b838111156111bd5750506000910152565b60008151808452611d3e816020860160208601611cfa565b601f01601f19169290920160200192915050565b602081526000610f776020830184611d26565b80356001600160a01b0381168114611d7c57600080fd5b919050565b60008060408385031215611d9457600080fd5b611d9d83611d65565b946020939093013593505050565b600060208284031215611dbd57600080fd5b813560038110610f7757600080fd5b600080600060608486031215611de157600080fd5b611dea84611d65565b9250611df860208501611d65565b9150604084013590509250925092565b600060208284031215611e1a57600080fd5b5035919050565b600060208284031215611e3357600080fd5b610f7782611d65565b60ff60f81b881681526000602060e081840152611e5c60e084018a611d26565b8381036040850152611e6e818a611d26565b606085018990526001600160a01b038816608086015260a0850187905284810360c0860152855180825283870192509083019060005b81811015611ec057835183529284019291840191600101611ea4565b50909c9b505050505050505050505050565b634e487b7160e01b600052602160045260246000fd5b60038110611ef857611ef8611ed2565b9052565b602081016105f18284611ee8565b600080600080600080600060e0888a031215611f2557600080fd5b611f2e88611d65565b9650611f3c60208901611d65565b95506040880135945060608801359350608088013560ff81168114611f6057600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611f9057600080fd5b611f9983611d65565b9150611fa760208401611d65565b90509250929050565b600181811c90821680611fc457607f821691505b60208210810361089257634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561201457612014611fe4565b500290565b6000821982111561202c5761202c611fe4565b500190565b60008451612043818460208901611cfa565b8083019050601760f91b8082528551612063816001850160208a01611cfa565b6001920191820152835161207e816002840160208801611cfa565b0160020195945050505050565b6001600160a01b03858116825260808201906120aa6020840187611ee8565b600285106120ba576120ba611ed2565b84604084015280841660608401525095945050505050565b634e487b7160e01b600052604160045260246000fd5b6000828210156120fa576120fa611fe4565b500390565b60006001820161211157612111611fe4565b5060010190565b634e487b7160e01b600052601260045260246000fd5b60008261213d5761213d612118565b500490565b60008261215157612151612118565b500690565b634e487b7160e01b600052603260045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b601f82111561220157600081815260208120601f850160051c810160208610156121de5750805b601f850160051c820191505b818110156121fd578281556001016121ea565b5050505b505050565b815167ffffffffffffffff811115612220576122206120d2565b6122348161222e8454611fb0565b846121b7565b602080601f83116001811461226957600084156122515750858301515b600019600386901b1c1916600185901b1785556121fd565b600085815260208120601f198616915b8281101561229857888601518255948401946001909101908401612279565b50858210156122b65787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
}

// USDBABI is the input ABI used to generate the binding from.
// Deprecated: Use USDBMetaData.ABI instead.
var USDBABI = USDBMetaData.ABI

// USDBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use USDBMetaData.Bin instead.
var USDBBin = USDBMetaData.Bin

// DeployUSDB deploys a new Ethereum contract, binding an instance of USDB to it.
func DeployUSDB(auth *bind.TransactOpts, backend bind.ContractBackend, _usdYieldManager common.Address, _l2Bridge common.Address, _remoteToken common.Address) (common.Address, *types.Transaction, *USDB, error) {
	parsed, err := USDBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDBBin), backend, _usdYieldManager, _l2Bridge, _remoteToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDB{USDBCaller: USDBCaller{contract: contract}, USDBTransactor: USDBTransactor{contract: contract}, USDBFilterer: USDBFilterer{contract: contract}}, nil
}

// USDB is an auto generated Go binding around an Ethereum contract.
type USDB struct {
	USDBCaller     // Read-only binding to the contract
	USDBTransactor // Write-only binding to the contract
	USDBFilterer   // Log filterer for contract events
}

// USDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDBSession struct {
	Contract     *USDB             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDBCallerSession struct {
	Contract *USDBCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// USDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDBTransactorSession struct {
	Contract     *USDBTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDBRaw struct {
	Contract *USDB // Generic contract binding to access the raw methods on
}

// USDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDBCallerRaw struct {
	Contract *USDBCaller // Generic read-only contract binding to access the raw methods on
}

// USDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDBTransactorRaw struct {
	Contract *USDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDB creates a new instance of USDB, bound to a specific deployed contract.
func NewUSDB(address common.Address, backend bind.ContractBackend) (*USDB, error) {
	contract, err := bindUSDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDB{USDBCaller: USDBCaller{contract: contract}, USDBTransactor: USDBTransactor{contract: contract}, USDBFilterer: USDBFilterer{contract: contract}}, nil
}

// NewUSDBCaller creates a new read-only instance of USDB, bound to a specific deployed contract.
func NewUSDBCaller(address common.Address, caller bind.ContractCaller) (*USDBCaller, error) {
	contract, err := bindUSDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDBCaller{contract: contract}, nil
}

// NewUSDBTransactor creates a new write-only instance of USDB, bound to a specific deployed contract.
func NewUSDBTransactor(address common.Address, transactor bind.ContractTransactor) (*USDBTransactor, error) {
	contract, err := bindUSDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDBTransactor{contract: contract}, nil
}

// NewUSDBFilterer creates a new log filterer instance of USDB, bound to a specific deployed contract.
func NewUSDBFilterer(address common.Address, filterer bind.ContractFilterer) (*USDBFilterer, error) {
	contract, err := bindUSDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDBFilterer{contract: contract}, nil
}

// bindUSDB binds a generic wrapper to an already deployed contract.
func bindUSDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDB *USDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDB.Contract.USDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDB *USDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDB.Contract.USDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDB *USDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDB.Contract.USDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDB *USDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDB *USDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDB *USDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDB.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_USDB *USDBCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_USDB *USDBSession) BRIDGE() (common.Address, error) {
	return _USDB.Contract.BRIDGE(&_USDB.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_USDB *USDBCallerSession) BRIDGE() (common.Address, error) {
	return _USDB.Contract.BRIDGE(&_USDB.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_USDB *USDBCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_USDB *USDBSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _USDB.Contract.DOMAINSEPARATOR(&_USDB.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_USDB *USDBCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _USDB.Contract.DOMAINSEPARATOR(&_USDB.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_USDB *USDBCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_USDB *USDBSession) PERMITTYPEHASH() ([32]byte, error) {
	return _USDB.Contract.PERMITTYPEHASH(&_USDB.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_USDB *USDBCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _USDB.Contract.PERMITTYPEHASH(&_USDB.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_USDB *USDBCaller) REMOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "REMOTE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_USDB *USDBSession) REMOTETOKEN() (common.Address, error) {
	return _USDB.Contract.REMOTETOKEN(&_USDB.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_USDB *USDBCallerSession) REMOTETOKEN() (common.Address, error) {
	return _USDB.Contract.REMOTETOKEN(&_USDB.CallOpts)
}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_USDB *USDBCaller) REPORTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "REPORTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_USDB *USDBSession) REPORTER() (common.Address, error) {
	return _USDB.Contract.REPORTER(&_USDB.CallOpts)
}

// REPORTER is a free data retrieval call binding the contract method 0x7ae556b5.
//
// Solidity: function REPORTER() view returns(address)
func (_USDB *USDBCallerSession) REPORTER() (common.Address, error) {
	return _USDB.Contract.REPORTER(&_USDB.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDB *USDBCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDB *USDBSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDB.Contract.Allowance(&_USDB.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDB *USDBCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDB.Contract.Allowance(&_USDB.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 value)
func (_USDB *USDBCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 value)
func (_USDB *USDBSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDB.Contract.BalanceOf(&_USDB.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 value)
func (_USDB *USDBCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDB.Contract.BalanceOf(&_USDB.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_USDB *USDBCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_USDB *USDBSession) Bridge() (common.Address, error) {
	return _USDB.Contract.Bridge(&_USDB.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_USDB *USDBCallerSession) Bridge() (common.Address, error) {
	return _USDB.Contract.Bridge(&_USDB.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_USDB *USDBCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_USDB *USDBSession) Count() (*big.Int, error) {
	return _USDB.Contract.Count(&_USDB.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_USDB *USDBCallerSession) Count() (*big.Int, error) {
	return _USDB.Contract.Count(&_USDB.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDB *USDBCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDB *USDBSession) Decimals() (uint8, error) {
	return _USDB.Contract.Decimals(&_USDB.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDB *USDBCallerSession) Decimals() (uint8, error) {
	return _USDB.Contract.Decimals(&_USDB.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_USDB *USDBCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "eip712Domain")

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
func (_USDB *USDBSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _USDB.Contract.Eip712Domain(&_USDB.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_USDB *USDBCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _USDB.Contract.Eip712Domain(&_USDB.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_USDB *USDBCaller) GetClaimableAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "getClaimableAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_USDB *USDBSession) GetClaimableAmount(account common.Address) (*big.Int, error) {
	return _USDB.Contract.GetClaimableAmount(&_USDB.CallOpts, account)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_USDB *USDBCallerSession) GetClaimableAmount(account common.Address) (*big.Int, error) {
	return _USDB.Contract.GetClaimableAmount(&_USDB.CallOpts, account)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address account) view returns(uint8)
func (_USDB *USDBCaller) GetConfiguration(opts *bind.CallOpts, account common.Address) (uint8, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "getConfiguration", account)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address account) view returns(uint8)
func (_USDB *USDBSession) GetConfiguration(account common.Address) (uint8, error) {
	return _USDB.Contract.GetConfiguration(&_USDB.CallOpts, account)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address account) view returns(uint8)
func (_USDB *USDBCallerSession) GetConfiguration(account common.Address) (uint8, error) {
	return _USDB.Contract.GetConfiguration(&_USDB.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDB *USDBCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDB *USDBSession) Name() (string, error) {
	return _USDB.Contract.Name(&_USDB.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDB *USDBCallerSession) Name() (string, error) {
	return _USDB.Contract.Name(&_USDB.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_USDB *USDBCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_USDB *USDBSession) Nonces(owner common.Address) (*big.Int, error) {
	return _USDB.Contract.Nonces(&_USDB.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_USDB *USDBCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _USDB.Contract.Nonces(&_USDB.CallOpts, owner)
}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_USDB *USDBCaller) Pending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "pending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_USDB *USDBSession) Pending() (*big.Int, error) {
	return _USDB.Contract.Pending(&_USDB.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe20ccec3.
//
// Solidity: function pending() view returns(uint256)
func (_USDB *USDBCallerSession) Pending() (*big.Int, error) {
	return _USDB.Contract.Pending(&_USDB.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_USDB *USDBCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_USDB *USDBSession) Price() (*big.Int, error) {
	return _USDB.Contract.Price(&_USDB.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_USDB *USDBCallerSession) Price() (*big.Int, error) {
	return _USDB.Contract.Price(&_USDB.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_USDB *USDBCaller) RemoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "remoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_USDB *USDBSession) RemoteToken() (common.Address, error) {
	return _USDB.Contract.RemoteToken(&_USDB.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_USDB *USDBCallerSession) RemoteToken() (common.Address, error) {
	return _USDB.Contract.RemoteToken(&_USDB.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_USDB *USDBCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_USDB *USDBSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _USDB.Contract.SupportsInterface(&_USDB.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_USDB *USDBCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _USDB.Contract.SupportsInterface(&_USDB.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDB *USDBCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDB *USDBSession) Symbol() (string, error) {
	return _USDB.Contract.Symbol(&_USDB.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDB *USDBCallerSession) Symbol() (string, error) {
	return _USDB.Contract.Symbol(&_USDB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDB *USDBCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDB *USDBSession) TotalSupply() (*big.Int, error) {
	return _USDB.Contract.TotalSupply(&_USDB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDB *USDBCallerSession) TotalSupply() (*big.Int, error) {
	return _USDB.Contract.TotalSupply(&_USDB.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_USDB *USDBCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDB.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_USDB *USDBSession) Version() (string, error) {
	return _USDB.Contract.Version(&_USDB.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_USDB *USDBCallerSession) Version() (string, error) {
	return _USDB.Contract.Version(&_USDB.CallOpts)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_USDB *USDBTransactor) AddValue(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "addValue", value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_USDB *USDBSession) AddValue(value *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.AddValue(&_USDB.TransactOpts, value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 value) returns()
func (_USDB *USDBTransactorSession) AddValue(value *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.AddValue(&_USDB.TransactOpts, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDB *USDBTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDB *USDBSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Approve(&_USDB.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDB *USDBTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Approve(&_USDB.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_USDB *USDBTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_USDB *USDBSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Burn(&_USDB.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_USDB *USDBTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Burn(&_USDB.TransactOpts, _from, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 amount) returns(uint256)
func (_USDB *USDBTransactor) Claim(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "claim", recipient, amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 amount) returns(uint256)
func (_USDB *USDBSession) Claim(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Claim(&_USDB.TransactOpts, recipient, amount)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 amount) returns(uint256)
func (_USDB *USDBTransactorSession) Claim(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Claim(&_USDB.TransactOpts, recipient, amount)
}

// Configure is a paid mutator transaction binding the contract method 0x1a33757d.
//
// Solidity: function configure(uint8 yieldMode) returns(uint256)
func (_USDB *USDBTransactor) Configure(opts *bind.TransactOpts, yieldMode uint8) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "configure", yieldMode)
}

// Configure is a paid mutator transaction binding the contract method 0x1a33757d.
//
// Solidity: function configure(uint8 yieldMode) returns(uint256)
func (_USDB *USDBSession) Configure(yieldMode uint8) (*types.Transaction, error) {
	return _USDB.Contract.Configure(&_USDB.TransactOpts, yieldMode)
}

// Configure is a paid mutator transaction binding the contract method 0x1a33757d.
//
// Solidity: function configure(uint8 yieldMode) returns(uint256)
func (_USDB *USDBTransactorSession) Configure(yieldMode uint8) (*types.Transaction, error) {
	return _USDB.Contract.Configure(&_USDB.TransactOpts, yieldMode)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_USDB *USDBTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_USDB *USDBSession) Initialize() (*types.Transaction, error) {
	return _USDB.Contract.Initialize(&_USDB.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_USDB *USDBTransactorSession) Initialize() (*types.Transaction, error) {
	return _USDB.Contract.Initialize(&_USDB.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_USDB *USDBTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_USDB *USDBSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Mint(&_USDB.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_USDB *USDBTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Mint(&_USDB.TransactOpts, _to, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_USDB *USDBTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_USDB *USDBSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _USDB.Contract.Permit(&_USDB.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_USDB *USDBTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _USDB.Contract.Permit(&_USDB.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_USDB *USDBTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_USDB *USDBSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Transfer(&_USDB.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_USDB *USDBTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.Transfer(&_USDB.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_USDB *USDBTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_USDB *USDBSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.TransferFrom(&_USDB.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_USDB *USDBTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDB.Contract.TransferFrom(&_USDB.TransactOpts, from, to, amount)
}

// USDBApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the USDB contract.
type USDBApprovalIterator struct {
	Event *USDBApproval // Event containing the contract specifics and raw log

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
func (it *USDBApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDBApproval)
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
		it.Event = new(USDBApproval)
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
func (it *USDBApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDBApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDBApproval represents a Approval event raised by the USDB contract.
type USDBApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDB *USDBFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*USDBApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDB.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &USDBApprovalIterator{contract: _USDB.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDB *USDBFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *USDBApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDB.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDBApproval)
				if err := _USDB.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_USDB *USDBFilterer) ParseApproval(log types.Log) (*USDBApproval, error) {
	event := new(USDBApproval)
	if err := _USDB.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDBClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the USDB contract.
type USDBClaimIterator struct {
	Event *USDBClaim // Event containing the contract specifics and raw log

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
func (it *USDBClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDBClaim)
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
		it.Event = new(USDBClaim)
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
func (it *USDBClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDBClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDBClaim represents a Claim event raised by the USDB contract.
type USDBClaim struct {
	Account   common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd987068.
//
// Solidity: event Claim(address indexed account, address indexed recipient, uint256 amount)
func (_USDB *USDBFilterer) FilterClaim(opts *bind.FilterOpts, account []common.Address, recipient []common.Address) (*USDBClaimIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDB.contract.FilterLogs(opts, "Claim", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &USDBClaimIterator{contract: _USDB.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd987068.
//
// Solidity: event Claim(address indexed account, address indexed recipient, uint256 amount)
func (_USDB *USDBFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *USDBClaim, account []common.Address, recipient []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDB.contract.WatchLogs(opts, "Claim", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDBClaim)
				if err := _USDB.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_USDB *USDBFilterer) ParseClaim(log types.Log) (*USDBClaim, error) {
	event := new(USDBClaim)
	if err := _USDB.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDBConfigureIterator is returned from FilterConfigure and is used to iterate over the raw logs and unpacked data for Configure events raised by the USDB contract.
type USDBConfigureIterator struct {
	Event *USDBConfigure // Event containing the contract specifics and raw log

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
func (it *USDBConfigureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDBConfigure)
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
		it.Event = new(USDBConfigure)
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
func (it *USDBConfigureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDBConfigureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDBConfigure represents a Configure event raised by the USDB contract.
type USDBConfigure struct {
	Account   common.Address
	YieldMode uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfigure is a free log retrieval operation binding the contract event 0xcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36.
//
// Solidity: event Configure(address indexed account, uint8 yieldMode)
func (_USDB *USDBFilterer) FilterConfigure(opts *bind.FilterOpts, account []common.Address) (*USDBConfigureIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDB.contract.FilterLogs(opts, "Configure", accountRule)
	if err != nil {
		return nil, err
	}
	return &USDBConfigureIterator{contract: _USDB.contract, event: "Configure", logs: logs, sub: sub}, nil
}

// WatchConfigure is a free log subscription operation binding the contract event 0xcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36.
//
// Solidity: event Configure(address indexed account, uint8 yieldMode)
func (_USDB *USDBFilterer) WatchConfigure(opts *bind.WatchOpts, sink chan<- *USDBConfigure, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDB.contract.WatchLogs(opts, "Configure", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDBConfigure)
				if err := _USDB.contract.UnpackLog(event, "Configure", log); err != nil {
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
func (_USDB *USDBFilterer) ParseConfigure(log types.Log) (*USDBConfigure, error) {
	event := new(USDBConfigure)
	if err := _USDB.contract.UnpackLog(event, "Configure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDBEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the USDB contract.
type USDBEIP712DomainChangedIterator struct {
	Event *USDBEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *USDBEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDBEIP712DomainChanged)
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
		it.Event = new(USDBEIP712DomainChanged)
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
func (it *USDBEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDBEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDBEIP712DomainChanged represents a EIP712DomainChanged event raised by the USDB contract.
type USDBEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_USDB *USDBFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*USDBEIP712DomainChangedIterator, error) {

	logs, sub, err := _USDB.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &USDBEIP712DomainChangedIterator{contract: _USDB.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_USDB *USDBFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *USDBEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _USDB.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDBEIP712DomainChanged)
				if err := _USDB.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_USDB *USDBFilterer) ParseEIP712DomainChanged(log types.Log) (*USDBEIP712DomainChanged, error) {
	event := new(USDBEIP712DomainChanged)
	if err := _USDB.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDBInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the USDB contract.
type USDBInitializedIterator struct {
	Event *USDBInitialized // Event containing the contract specifics and raw log

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
func (it *USDBInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDBInitialized)
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
		it.Event = new(USDBInitialized)
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
func (it *USDBInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDBInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDBInitialized represents a Initialized event raised by the USDB contract.
type USDBInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_USDB *USDBFilterer) FilterInitialized(opts *bind.FilterOpts) (*USDBInitializedIterator, error) {

	logs, sub, err := _USDB.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &USDBInitializedIterator{contract: _USDB.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_USDB *USDBFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *USDBInitialized) (event.Subscription, error) {

	logs, sub, err := _USDB.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDBInitialized)
				if err := _USDB.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_USDB *USDBFilterer) ParseInitialized(log types.Log) (*USDBInitialized, error) {
	event := new(USDBInitialized)
	if err := _USDB.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDBNewPriceIterator is returned from FilterNewPrice and is used to iterate over the raw logs and unpacked data for NewPrice events raised by the USDB contract.
type USDBNewPriceIterator struct {
	Event *USDBNewPrice // Event containing the contract specifics and raw log

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
func (it *USDBNewPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDBNewPrice)
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
		it.Event = new(USDBNewPrice)
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
func (it *USDBNewPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDBNewPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDBNewPrice represents a NewPrice event raised by the USDB contract.
type USDBNewPrice struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewPrice is a free log retrieval operation binding the contract event 0x270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b7.
//
// Solidity: event NewPrice(uint256 price)
func (_USDB *USDBFilterer) FilterNewPrice(opts *bind.FilterOpts) (*USDBNewPriceIterator, error) {

	logs, sub, err := _USDB.contract.FilterLogs(opts, "NewPrice")
	if err != nil {
		return nil, err
	}
	return &USDBNewPriceIterator{contract: _USDB.contract, event: "NewPrice", logs: logs, sub: sub}, nil
}

// WatchNewPrice is a free log subscription operation binding the contract event 0x270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b7.
//
// Solidity: event NewPrice(uint256 price)
func (_USDB *USDBFilterer) WatchNewPrice(opts *bind.WatchOpts, sink chan<- *USDBNewPrice) (event.Subscription, error) {

	logs, sub, err := _USDB.contract.WatchLogs(opts, "NewPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDBNewPrice)
				if err := _USDB.contract.UnpackLog(event, "NewPrice", log); err != nil {
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
func (_USDB *USDBFilterer) ParseNewPrice(log types.Log) (*USDBNewPrice, error) {
	event := new(USDBNewPrice)
	if err := _USDB.contract.UnpackLog(event, "NewPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDBTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the USDB contract.
type USDBTransferIterator struct {
	Event *USDBTransfer // Event containing the contract specifics and raw log

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
func (it *USDBTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDBTransfer)
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
		it.Event = new(USDBTransfer)
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
func (it *USDBTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDBTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDBTransfer represents a Transfer event raised by the USDB contract.
type USDBTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDB *USDBFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDBTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDB.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDBTransferIterator{contract: _USDB.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDB *USDBFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *USDBTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDB.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDBTransfer)
				if err := _USDB.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_USDB *USDBFilterer) ParseTransfer(log types.Log) (*USDBTransfer, error) {
	event := new(USDBTransfer)
	if err := _USDB.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
