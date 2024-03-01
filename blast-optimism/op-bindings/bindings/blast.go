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

// BlastMetaData contains all meta data concerning the Blast contract.
var BlastMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yieldContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GAS_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YIELD_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimAllGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfYield\",\"type\":\"address\"}],\"name\":\"claimAllYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasToClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasSecondsToConsume\",\"type\":\"uint256\"}],\"name\":\"claimGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minClaimRateBips\",\"type\":\"uint256\"}],\"name\":\"claimGasAtMinClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimMaxGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfYield\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"_yieldMode\",\"type\":\"uint8\"},{\"internalType\":\"enumGasMode\",\"name\":\"_gasMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"configure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureAutomaticYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureAutomaticYieldOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureClaimableGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureClaimableGasOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureClaimableYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureClaimableYieldOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumYieldMode\",\"name\":\"_yieldMode\",\"type\":\"uint8\"},{\"internalType\":\"enumGasMode\",\"name\":\"_gasMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"configureContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"configureGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureGovernorOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureVoidGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureVoidGasOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureVoidYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureVoidYieldOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governorMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"readClaimableYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"readGasParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumGasMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"readYieldConfiguration\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001f0938038062001f09833981016040819052620000359162000148565b6001608052600060a081905260c0526001600160a01b0380831661010052811660e052620000626200006a565b505062000180565b600054610100900460ff1615620000d75760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000129576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200014357600080fd5b919050565b600080604083850312156200015c57600080fd5b62000167836200012b565b915062000177602084016200012b565b90509250929050565b60805160a05160c05160e05161010051611c9c6200026d600039600081816103410152818161052e015281816106370152818161082c0152818161098e01528181610ae701528181610ef401528181610fd801528181611121015281816112670152818161131d01526115da01526000818161031a015281816106d2015281816107a00152818161089901528181610b8401528181610da301528181610e4101528181611036015281816110920152818161118d015281816113cb0152818161146d0152818161151f015261167201526000610a1f015260006109f6015260006109cd0152611c9c6000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c8063a0278c9011610104578063eb59acdc116100a2578063f971966211610071578063f971966214610433578063fafce39e14610446578063fd8c4b9d14610459578063fe9fbb801461047e57600080fd5b8063eb59acdc146103f2578063eb86469814610405578063ec3278e814610418578063f098767a1461042b57600080fd5b8063b71d6dd4116100de578063b71d6dd41461036b578063c8992e611461037e578063dde798a414610391578063e43581b8146103b457600080fd5b8063a0278c9014610315578063a70d11bd1461033c578063aa857d981461036357600080fd5b806354fd4d501161017c57806385ebdc271161014b57806385ebdc271461029b578063860043b6146102dc578063908c8502146102ef578063954fa5ee1461030257600080fd5b806354fd4d5014610263578063662aa11d146102785780637114177a1461028b5780638129fc1c1461029357600080fd5b806337ebe3a8116101b857806337ebe3a8146102225780633ba5713e146102355780634c802f38146102485780634e606c471461025b57600080fd5b80630951888f146101df5780630ca12c4b146102055780632210dfb11461021a575b600080fd5b6101f26101ed366004611853565b610491565b6040519081526020015b60405180910390f35b61021861021336600461188f565b6105a6565b005b6102186105fb565b6102186102303660046118c2565b6106a3565b6102186102433660046118c2565b610771565b6102186102563660046118f9565b6107cd565b610218610952565b61026b6109c6565b6040516101fc919061197b565b6101f261028636600461188f565b610a69565b610218610b55565b610218610c22565b6102c46102a93660046118c2565b6001602052600090815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016101fc565b6101f26102ea36600461188f565b610d2a565b6102186102fd3660046118c2565b610eb8565b6101f261031036600461188f565b610f5a565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b610218611007565b6102186103793660046118c2565b611063565b61021861038c3660046119ae565b6110bf565b6103a461039f3660046118c2565b61123f565b6040516101fc9493929190611a2b565b6103e26103c23660046118c2565b6001600160a01b0390811660009081526001602052604090205416331490565b60405190151581526020016101fc565b6102186104003660046118c2565b6112e1565b6102186104133660046118c2565b611355565b6101f26104263660046118c2565b6113a9565b61021861143e565b6101f2610441366004611853565b61149a565b6101f2610454366004611a4d565b61154e565b61046c6104673660046118c2565b611650565b60405160ff90911681526020016101fc565b6103e261048c3660046118c2565b6116df565b600061049c846116df565b6105005760405162461bcd60e51b815260206004820152602a60248201527f4e6f7420616c6c6f77656420746f20636c61696d20676173206174206d696e20604482015269636c61696d207261746560b01b60648201526084015b60405180910390fd5b604051630951888f60e01b81526001600160a01b0385811660048301528481166024830152604482018490527f00000000000000000000000000000000000000000000000000000000000000001690630951888f906064015b6020604051808303816000875af1158015610578573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059c9190611a8f565b90505b9392505050565b6105af816116df565b6105cb5760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b0390811660009081526001602052604090208054919092166001600160a01b0319909116179055565b610604336116df565b6106205760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061066f903390600090600401611aec565b600060405180830381600087803b15801561068957600080fd5b505af115801561069d573d6000803e3d6000fd5b50505050565b6106ac816116df565b6106c85760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260025b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076d9190611a8f565b5050565b61077a816116df565b6107965760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260006106fb565b6107d6846116df565b6107f25760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b038481166000908152600160205260409081902080546001600160a01b0319168484161790555163d4810ba560e01b81527f00000000000000000000000000000000000000000000000000000000000000009091169063d4810ba5906108659087908690600401611aec565b600060405180830381600087803b15801561087f57600080fd5b505af1158015610893573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633bdbe9a5858560028111156108d9576108d96119f3565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610927573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094b9190611a8f565b5050505050565b61095b336116df565b6109775760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061066f903390600190600401611aec565b60606109f17f0000000000000000000000000000000000000000000000000000000000000000611736565b610a1a7f0000000000000000000000000000000000000000000000000000000000000000611736565b610a437f0000000000000000000000000000000000000000000000000000000000000000611736565b604051602001610a5593929190611b09565b604051602081830303815290604052905090565b6000610a74836116df565b610ac05760405162461bcd60e51b815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d206d6178206761730000000060448201526064016104f7565b604051630928b10d60e31b81526001600160a01b03848116600483015283811660248301527f000000000000000000000000000000000000000000000000000000000000000016906349458868906044015b6020604051808303816000875af1158015610b31573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059f9190611a8f565b610b5e336116df565b610b7a5760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a53360005b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610bfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1f9190611a8f565b50565b600054610100900460ff1615808015610c425750600054600160ff909116105b80610c5c5750303b158015610c5c575060005460ff166001145b610cbf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104f7565b6000805460ff191660011790558015610ce2576000805461ff0019166101001790555b8015610c1f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6000610d35836116df565b610d815760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c6400000060448201526064016104f7565b60405163e12f3a6160e01b81526001600160a01b0384811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015610dec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e109190611a8f565b60405163132d974d60e31b81526001600160a01b0386811660048301528581166024830152604482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063996cba68906064016020604051808303816000875af1158015610e8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb09190611a8f565b949350505050565b610ec1816116df565b610edd5760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba590610f2c908490600190600401611aec565b600060405180830381600087803b158015610f4657600080fd5b505af115801561094b573d6000803e3d6000fd5b6000610f65836116df565b610fb15760405162461bcd60e51b815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d20616c6c206761730000000060448201526064016104f7565b604051635767bba560e01b81526001600160a01b03848116600483015283811660248301527f00000000000000000000000000000000000000000000000000000000000000001690635767bba590604401610b12565b611010336116df565b61102c5760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336001610bad565b61106c816116df565b6110885760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260016106fb565b6110c8336116df565b6110e45760405162461bcd60e51b81526004016104f790611aa8565b336000818152600160205260409081902080546001600160a01b0319166001600160a01b0385811691909117909155905163d4810ba560e01b81527f00000000000000000000000000000000000000000000000000000000000000009091169163d4810ba59161115991908690600401611aec565b600060405180830381600087803b15801561117357600080fd5b505af1158015611187573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633bdbe9a5338560028111156111cd576111cd6119f3565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af115801561121b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069d9190611a8f565b604051633779e62960e21b81526001600160a01b0382811660048301526000918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063dde798a490602401608060405180830381865afa1580156112ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d29190611b63565b93509350935093509193509193565b6112ea816116df565b6113065760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba590610f2c908490600090600401611aec565b61135e336116df565b61137a5760405162461bcd60e51b81526004016104f790611aa8565b33600090815260016020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b60405163e12f3a6160e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015611414573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114389190611a8f565b92915050565b611447336116df565b6114635760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336002610bad565b60006114a5846116df565b6114f15760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c6400000060448201526064016104f7565b60405163132d974d60e31b81526001600160a01b0385811660048301528481166024830152604482018490527f0000000000000000000000000000000000000000000000000000000000000000169063996cba6890606401610559565b6000611559856116df565b6115a55760405162461bcd60e51b815260206004820152601860248201527f4e6f7420616c6c6f77656420746f20636c61696d20676173000000000000000060448201526064016104f7565b604051631357a41960e11b81526001600160a01b038681166004830152858116602483015260448201859052606482018490527f000000000000000000000000000000000000000000000000000000000000000016906326af4832906084016020604051808303816000875af1158015611623573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116479190611a8f565b95945050505050565b60405163c44b11f760e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063c44b11f790602401602060405180830381865afa1580156116bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114389190611ba4565b6001600160a01b0380821660009081526001602052604081205490911633148061143857506001600160a01b03808316600090815260016020526040902054161580156114385750506001600160a01b0316331490565b60608160000361175d5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115611787578061177181611bdd565b91506117809050600a83611c0c565b9150611761565b60008167ffffffffffffffff8111156117a2576117a2611c20565b6040519080825280601f01601f1916602001820160405280156117cc576020820181803683370190505b5090505b8415610eb0576117e1600183611c36565b91506117ee600a86611c4d565b6117f9906030611c61565b60f81b81838151811061180e5761180e611c79565b60200101906001600160f81b031916908160001a905350611830600a86611c0c565b94506117d0565b80356001600160a01b038116811461184e57600080fd5b919050565b60008060006060848603121561186857600080fd5b61187184611837565b925061187f60208501611837565b9150604084013590509250925092565b600080604083850312156118a257600080fd5b6118ab83611837565b91506118b960208401611837565b90509250929050565b6000602082840312156118d457600080fd5b61059f82611837565b80356003811061184e57600080fd5b60028110610c1f57600080fd5b6000806000806080858703121561190f57600080fd5b61191885611837565b9350611926602086016118dd565b92506040850135611936816118ec565b915061194460608601611837565b905092959194509250565b60005b8381101561196a578181015183820152602001611952565b8381111561069d5750506000910152565b602081526000825180602084015261199a81604085016020870161194f565b601f01601f19169190910160400192915050565b6000806000606084860312156119c357600080fd5b6119cc846118dd565b925060208401356119dc816118ec565b91506119ea60408501611837565b90509250925092565b634e487b7160e01b600052602160045260246000fd5b60028110611a2757634e487b7160e01b600052602160045260246000fd5b9052565b8481526020810184905260408101839052608081016116476060830184611a09565b60008060008060808587031215611a6357600080fd5b611a6c85611837565b9350611a7a60208601611837565b93969395505050506040820135916060013590565b600060208284031215611aa157600080fd5b5051919050565b60208082526024908201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e746040820152631c9858dd60e21b606082015260800190565b6001600160a01b03831681526040810161059f6020830184611a09565b60008451611b1b81846020890161194f565b8083019050601760f91b8082528551611b3b816001850160208a0161194f565b60019201918201528351611b5681600284016020880161194f565b0160020195945050505050565b60008060008060808587031215611b7957600080fd5b8451935060208501519250604085015191506060850151611b99816118ec565b939692955090935050565b600060208284031215611bb657600080fd5b815160ff8116811461059f57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611bef57611bef611bc7565b5060010190565b634e487b7160e01b600052601260045260246000fd5b600082611c1b57611c1b611bf6565b500490565b634e487b7160e01b600052604160045260246000fd5b600082821015611c4857611c48611bc7565b500390565b600082611c5c57611c5c611bf6565b500690565b60008219821115611c7457611c74611bc7565b500190565b634e487b7160e01b600052603260045260246000fdfea164736f6c634300080f000a",
}

// BlastABI is the input ABI used to generate the binding from.
// Deprecated: Use BlastMetaData.ABI instead.
var BlastABI = BlastMetaData.ABI

// BlastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlastMetaData.Bin instead.
var BlastBin = BlastMetaData.Bin

// DeployBlast deploys a new Ethereum contract, binding an instance of Blast to it.
func DeployBlast(auth *bind.TransactOpts, backend bind.ContractBackend, _gasContract common.Address, _yieldContract common.Address) (common.Address, *types.Transaction, *Blast, error) {
	parsed, err := BlastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlastBin), backend, _gasContract, _yieldContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blast{BlastCaller: BlastCaller{contract: contract}, BlastTransactor: BlastTransactor{contract: contract}, BlastFilterer: BlastFilterer{contract: contract}}, nil
}

// Blast is an auto generated Go binding around an Ethereum contract.
type Blast struct {
	BlastCaller     // Read-only binding to the contract
	BlastTransactor // Write-only binding to the contract
	BlastFilterer   // Log filterer for contract events
}

// BlastCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlastSession struct {
	Contract     *Blast            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlastCallerSession struct {
	Contract *BlastCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BlastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlastTransactorSession struct {
	Contract     *BlastTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlastRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlastRaw struct {
	Contract *Blast // Generic contract binding to access the raw methods on
}

// BlastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlastCallerRaw struct {
	Contract *BlastCaller // Generic read-only contract binding to access the raw methods on
}

// BlastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlastTransactorRaw struct {
	Contract *BlastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlast creates a new instance of Blast, bound to a specific deployed contract.
func NewBlast(address common.Address, backend bind.ContractBackend) (*Blast, error) {
	contract, err := bindBlast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blast{BlastCaller: BlastCaller{contract: contract}, BlastTransactor: BlastTransactor{contract: contract}, BlastFilterer: BlastFilterer{contract: contract}}, nil
}

// NewBlastCaller creates a new read-only instance of Blast, bound to a specific deployed contract.
func NewBlastCaller(address common.Address, caller bind.ContractCaller) (*BlastCaller, error) {
	contract, err := bindBlast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlastCaller{contract: contract}, nil
}

// NewBlastTransactor creates a new write-only instance of Blast, bound to a specific deployed contract.
func NewBlastTransactor(address common.Address, transactor bind.ContractTransactor) (*BlastTransactor, error) {
	contract, err := bindBlast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlastTransactor{contract: contract}, nil
}

// NewBlastFilterer creates a new log filterer instance of Blast, bound to a specific deployed contract.
func NewBlastFilterer(address common.Address, filterer bind.ContractFilterer) (*BlastFilterer, error) {
	contract, err := bindBlast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlastFilterer{contract: contract}, nil
}

// bindBlast binds a generic wrapper to an already deployed contract.
func bindBlast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blast *BlastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blast.Contract.BlastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blast *BlastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.Contract.BlastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blast *BlastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blast.Contract.BlastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blast *BlastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blast *BlastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blast *BlastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blast.Contract.contract.Transact(opts, method, params...)
}

// GASCONTRACT is a free data retrieval call binding the contract method 0xa70d11bd.
//
// Solidity: function GAS_CONTRACT() view returns(address)
func (_Blast *BlastCaller) GASCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "GAS_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GASCONTRACT is a free data retrieval call binding the contract method 0xa70d11bd.
//
// Solidity: function GAS_CONTRACT() view returns(address)
func (_Blast *BlastSession) GASCONTRACT() (common.Address, error) {
	return _Blast.Contract.GASCONTRACT(&_Blast.CallOpts)
}

// GASCONTRACT is a free data retrieval call binding the contract method 0xa70d11bd.
//
// Solidity: function GAS_CONTRACT() view returns(address)
func (_Blast *BlastCallerSession) GASCONTRACT() (common.Address, error) {
	return _Blast.Contract.GASCONTRACT(&_Blast.CallOpts)
}

// YIELDCONTRACT is a free data retrieval call binding the contract method 0xa0278c90.
//
// Solidity: function YIELD_CONTRACT() view returns(address)
func (_Blast *BlastCaller) YIELDCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "YIELD_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YIELDCONTRACT is a free data retrieval call binding the contract method 0xa0278c90.
//
// Solidity: function YIELD_CONTRACT() view returns(address)
func (_Blast *BlastSession) YIELDCONTRACT() (common.Address, error) {
	return _Blast.Contract.YIELDCONTRACT(&_Blast.CallOpts)
}

// YIELDCONTRACT is a free data retrieval call binding the contract method 0xa0278c90.
//
// Solidity: function YIELD_CONTRACT() view returns(address)
func (_Blast *BlastCallerSession) YIELDCONTRACT() (common.Address, error) {
	return _Blast.Contract.YIELDCONTRACT(&_Blast.CallOpts)
}

// GovernorMap is a free data retrieval call binding the contract method 0x85ebdc27.
//
// Solidity: function governorMap(address ) view returns(address)
func (_Blast *BlastCaller) GovernorMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "governorMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernorMap is a free data retrieval call binding the contract method 0x85ebdc27.
//
// Solidity: function governorMap(address ) view returns(address)
func (_Blast *BlastSession) GovernorMap(arg0 common.Address) (common.Address, error) {
	return _Blast.Contract.GovernorMap(&_Blast.CallOpts, arg0)
}

// GovernorMap is a free data retrieval call binding the contract method 0x85ebdc27.
//
// Solidity: function governorMap(address ) view returns(address)
func (_Blast *BlastCallerSession) GovernorMap(arg0 common.Address) (common.Address, error) {
	return _Blast.Contract.GovernorMap(&_Blast.CallOpts, arg0)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address contractAddress) view returns(bool)
func (_Blast *BlastCaller) IsAuthorized(opts *bind.CallOpts, contractAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "isAuthorized", contractAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address contractAddress) view returns(bool)
func (_Blast *BlastSession) IsAuthorized(contractAddress common.Address) (bool, error) {
	return _Blast.Contract.IsAuthorized(&_Blast.CallOpts, contractAddress)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address contractAddress) view returns(bool)
func (_Blast *BlastCallerSession) IsAuthorized(contractAddress common.Address) (bool, error) {
	return _Blast.Contract.IsAuthorized(&_Blast.CallOpts, contractAddress)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address contractAddress) view returns(bool)
func (_Blast *BlastCaller) IsGovernor(opts *bind.CallOpts, contractAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "isGovernor", contractAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address contractAddress) view returns(bool)
func (_Blast *BlastSession) IsGovernor(contractAddress common.Address) (bool, error) {
	return _Blast.Contract.IsGovernor(&_Blast.CallOpts, contractAddress)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address contractAddress) view returns(bool)
func (_Blast *BlastCallerSession) IsGovernor(contractAddress common.Address) (bool, error) {
	return _Blast.Contract.IsGovernor(&_Blast.CallOpts, contractAddress)
}

// ReadClaimableYield is a free data retrieval call binding the contract method 0xec3278e8.
//
// Solidity: function readClaimableYield(address contractAddress) view returns(uint256)
func (_Blast *BlastCaller) ReadClaimableYield(opts *bind.CallOpts, contractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "readClaimableYield", contractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadClaimableYield is a free data retrieval call binding the contract method 0xec3278e8.
//
// Solidity: function readClaimableYield(address contractAddress) view returns(uint256)
func (_Blast *BlastSession) ReadClaimableYield(contractAddress common.Address) (*big.Int, error) {
	return _Blast.Contract.ReadClaimableYield(&_Blast.CallOpts, contractAddress)
}

// ReadClaimableYield is a free data retrieval call binding the contract method 0xec3278e8.
//
// Solidity: function readClaimableYield(address contractAddress) view returns(uint256)
func (_Blast *BlastCallerSession) ReadClaimableYield(contractAddress common.Address) (*big.Int, error) {
	return _Blast.Contract.ReadClaimableYield(&_Blast.CallOpts, contractAddress)
}

// ReadGasParams is a free data retrieval call binding the contract method 0xdde798a4.
//
// Solidity: function readGasParams(address contractAddress) view returns(uint256, uint256, uint256, uint8)
func (_Blast *BlastCaller) ReadGasParams(opts *bind.CallOpts, contractAddress common.Address) (*big.Int, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "readGasParams", contractAddress)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// ReadGasParams is a free data retrieval call binding the contract method 0xdde798a4.
//
// Solidity: function readGasParams(address contractAddress) view returns(uint256, uint256, uint256, uint8)
func (_Blast *BlastSession) ReadGasParams(contractAddress common.Address) (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Blast.Contract.ReadGasParams(&_Blast.CallOpts, contractAddress)
}

// ReadGasParams is a free data retrieval call binding the contract method 0xdde798a4.
//
// Solidity: function readGasParams(address contractAddress) view returns(uint256, uint256, uint256, uint8)
func (_Blast *BlastCallerSession) ReadGasParams(contractAddress common.Address) (*big.Int, *big.Int, *big.Int, uint8, error) {
	return _Blast.Contract.ReadGasParams(&_Blast.CallOpts, contractAddress)
}

// ReadYieldConfiguration is a free data retrieval call binding the contract method 0xfd8c4b9d.
//
// Solidity: function readYieldConfiguration(address contractAddress) view returns(uint8)
func (_Blast *BlastCaller) ReadYieldConfiguration(opts *bind.CallOpts, contractAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "readYieldConfiguration", contractAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ReadYieldConfiguration is a free data retrieval call binding the contract method 0xfd8c4b9d.
//
// Solidity: function readYieldConfiguration(address contractAddress) view returns(uint8)
func (_Blast *BlastSession) ReadYieldConfiguration(contractAddress common.Address) (uint8, error) {
	return _Blast.Contract.ReadYieldConfiguration(&_Blast.CallOpts, contractAddress)
}

// ReadYieldConfiguration is a free data retrieval call binding the contract method 0xfd8c4b9d.
//
// Solidity: function readYieldConfiguration(address contractAddress) view returns(uint8)
func (_Blast *BlastCallerSession) ReadYieldConfiguration(contractAddress common.Address) (uint8, error) {
	return _Blast.Contract.ReadYieldConfiguration(&_Blast.CallOpts, contractAddress)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Blast *BlastCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Blast.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Blast *BlastSession) Version() (string, error) {
	return _Blast.Contract.Version(&_Blast.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Blast *BlastCallerSession) Version() (string, error) {
	return _Blast.Contract.Version(&_Blast.CallOpts)
}

// ClaimAllGas is a paid mutator transaction binding the contract method 0x954fa5ee.
//
// Solidity: function claimAllGas(address contractAddress, address recipientOfGas) returns(uint256)
func (_Blast *BlastTransactor) ClaimAllGas(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "claimAllGas", contractAddress, recipientOfGas)
}

// ClaimAllGas is a paid mutator transaction binding the contract method 0x954fa5ee.
//
// Solidity: function claimAllGas(address contractAddress, address recipientOfGas) returns(uint256)
func (_Blast *BlastSession) ClaimAllGas(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ClaimAllGas(&_Blast.TransactOpts, contractAddress, recipientOfGas)
}

// ClaimAllGas is a paid mutator transaction binding the contract method 0x954fa5ee.
//
// Solidity: function claimAllGas(address contractAddress, address recipientOfGas) returns(uint256)
func (_Blast *BlastTransactorSession) ClaimAllGas(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ClaimAllGas(&_Blast.TransactOpts, contractAddress, recipientOfGas)
}

// ClaimAllYield is a paid mutator transaction binding the contract method 0x860043b6.
//
// Solidity: function claimAllYield(address contractAddress, address recipientOfYield) returns(uint256)
func (_Blast *BlastTransactor) ClaimAllYield(opts *bind.TransactOpts, contractAddress common.Address, recipientOfYield common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "claimAllYield", contractAddress, recipientOfYield)
}

// ClaimAllYield is a paid mutator transaction binding the contract method 0x860043b6.
//
// Solidity: function claimAllYield(address contractAddress, address recipientOfYield) returns(uint256)
func (_Blast *BlastSession) ClaimAllYield(contractAddress common.Address, recipientOfYield common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ClaimAllYield(&_Blast.TransactOpts, contractAddress, recipientOfYield)
}

// ClaimAllYield is a paid mutator transaction binding the contract method 0x860043b6.
//
// Solidity: function claimAllYield(address contractAddress, address recipientOfYield) returns(uint256)
func (_Blast *BlastTransactorSession) ClaimAllYield(contractAddress common.Address, recipientOfYield common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ClaimAllYield(&_Blast.TransactOpts, contractAddress, recipientOfYield)
}

// ClaimGas is a paid mutator transaction binding the contract method 0xfafce39e.
//
// Solidity: function claimGas(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) returns(uint256)
func (_Blast *BlastTransactor) ClaimGas(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address, gasToClaim *big.Int, gasSecondsToConsume *big.Int) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "claimGas", contractAddress, recipientOfGas, gasToClaim, gasSecondsToConsume)
}

// ClaimGas is a paid mutator transaction binding the contract method 0xfafce39e.
//
// Solidity: function claimGas(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) returns(uint256)
func (_Blast *BlastSession) ClaimGas(contractAddress common.Address, recipientOfGas common.Address, gasToClaim *big.Int, gasSecondsToConsume *big.Int) (*types.Transaction, error) {
	return _Blast.Contract.ClaimGas(&_Blast.TransactOpts, contractAddress, recipientOfGas, gasToClaim, gasSecondsToConsume)
}

// ClaimGas is a paid mutator transaction binding the contract method 0xfafce39e.
//
// Solidity: function claimGas(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) returns(uint256)
func (_Blast *BlastTransactorSession) ClaimGas(contractAddress common.Address, recipientOfGas common.Address, gasToClaim *big.Int, gasSecondsToConsume *big.Int) (*types.Transaction, error) {
	return _Blast.Contract.ClaimGas(&_Blast.TransactOpts, contractAddress, recipientOfGas, gasToClaim, gasSecondsToConsume)
}

// ClaimGasAtMinClaimRate is a paid mutator transaction binding the contract method 0x0951888f.
//
// Solidity: function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) returns(uint256)
func (_Blast *BlastTransactor) ClaimGasAtMinClaimRate(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address, minClaimRateBips *big.Int) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "claimGasAtMinClaimRate", contractAddress, recipientOfGas, minClaimRateBips)
}

// ClaimGasAtMinClaimRate is a paid mutator transaction binding the contract method 0x0951888f.
//
// Solidity: function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) returns(uint256)
func (_Blast *BlastSession) ClaimGasAtMinClaimRate(contractAddress common.Address, recipientOfGas common.Address, minClaimRateBips *big.Int) (*types.Transaction, error) {
	return _Blast.Contract.ClaimGasAtMinClaimRate(&_Blast.TransactOpts, contractAddress, recipientOfGas, minClaimRateBips)
}

// ClaimGasAtMinClaimRate is a paid mutator transaction binding the contract method 0x0951888f.
//
// Solidity: function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) returns(uint256)
func (_Blast *BlastTransactorSession) ClaimGasAtMinClaimRate(contractAddress common.Address, recipientOfGas common.Address, minClaimRateBips *big.Int) (*types.Transaction, error) {
	return _Blast.Contract.ClaimGasAtMinClaimRate(&_Blast.TransactOpts, contractAddress, recipientOfGas, minClaimRateBips)
}

// ClaimMaxGas is a paid mutator transaction binding the contract method 0x662aa11d.
//
// Solidity: function claimMaxGas(address contractAddress, address recipientOfGas) returns(uint256)
func (_Blast *BlastTransactor) ClaimMaxGas(opts *bind.TransactOpts, contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "claimMaxGas", contractAddress, recipientOfGas)
}

// ClaimMaxGas is a paid mutator transaction binding the contract method 0x662aa11d.
//
// Solidity: function claimMaxGas(address contractAddress, address recipientOfGas) returns(uint256)
func (_Blast *BlastSession) ClaimMaxGas(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ClaimMaxGas(&_Blast.TransactOpts, contractAddress, recipientOfGas)
}

// ClaimMaxGas is a paid mutator transaction binding the contract method 0x662aa11d.
//
// Solidity: function claimMaxGas(address contractAddress, address recipientOfGas) returns(uint256)
func (_Blast *BlastTransactorSession) ClaimMaxGas(contractAddress common.Address, recipientOfGas common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ClaimMaxGas(&_Blast.TransactOpts, contractAddress, recipientOfGas)
}

// ClaimYield is a paid mutator transaction binding the contract method 0xf9719662.
//
// Solidity: function claimYield(address contractAddress, address recipientOfYield, uint256 amount) returns(uint256)
func (_Blast *BlastTransactor) ClaimYield(opts *bind.TransactOpts, contractAddress common.Address, recipientOfYield common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "claimYield", contractAddress, recipientOfYield, amount)
}

// ClaimYield is a paid mutator transaction binding the contract method 0xf9719662.
//
// Solidity: function claimYield(address contractAddress, address recipientOfYield, uint256 amount) returns(uint256)
func (_Blast *BlastSession) ClaimYield(contractAddress common.Address, recipientOfYield common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blast.Contract.ClaimYield(&_Blast.TransactOpts, contractAddress, recipientOfYield, amount)
}

// ClaimYield is a paid mutator transaction binding the contract method 0xf9719662.
//
// Solidity: function claimYield(address contractAddress, address recipientOfYield, uint256 amount) returns(uint256)
func (_Blast *BlastTransactorSession) ClaimYield(contractAddress common.Address, recipientOfYield common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Blast.Contract.ClaimYield(&_Blast.TransactOpts, contractAddress, recipientOfYield, amount)
}

// Configure is a paid mutator transaction binding the contract method 0xc8992e61.
//
// Solidity: function configure(uint8 _yieldMode, uint8 _gasMode, address governor) returns()
func (_Blast *BlastTransactor) Configure(opts *bind.TransactOpts, _yieldMode uint8, _gasMode uint8, governor common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configure", _yieldMode, _gasMode, governor)
}

// Configure is a paid mutator transaction binding the contract method 0xc8992e61.
//
// Solidity: function configure(uint8 _yieldMode, uint8 _gasMode, address governor) returns()
func (_Blast *BlastSession) Configure(_yieldMode uint8, _gasMode uint8, governor common.Address) (*types.Transaction, error) {
	return _Blast.Contract.Configure(&_Blast.TransactOpts, _yieldMode, _gasMode, governor)
}

// Configure is a paid mutator transaction binding the contract method 0xc8992e61.
//
// Solidity: function configure(uint8 _yieldMode, uint8 _gasMode, address governor) returns()
func (_Blast *BlastTransactorSession) Configure(_yieldMode uint8, _gasMode uint8, governor common.Address) (*types.Transaction, error) {
	return _Blast.Contract.Configure(&_Blast.TransactOpts, _yieldMode, _gasMode, governor)
}

// ConfigureAutomaticYield is a paid mutator transaction binding the contract method 0x7114177a.
//
// Solidity: function configureAutomaticYield() returns()
func (_Blast *BlastTransactor) ConfigureAutomaticYield(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureAutomaticYield")
}

// ConfigureAutomaticYield is a paid mutator transaction binding the contract method 0x7114177a.
//
// Solidity: function configureAutomaticYield() returns()
func (_Blast *BlastSession) ConfigureAutomaticYield() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureAutomaticYield(&_Blast.TransactOpts)
}

// ConfigureAutomaticYield is a paid mutator transaction binding the contract method 0x7114177a.
//
// Solidity: function configureAutomaticYield() returns()
func (_Blast *BlastTransactorSession) ConfigureAutomaticYield() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureAutomaticYield(&_Blast.TransactOpts)
}

// ConfigureAutomaticYieldOnBehalf is a paid mutator transaction binding the contract method 0x3ba5713e.
//
// Solidity: function configureAutomaticYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactor) ConfigureAutomaticYieldOnBehalf(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureAutomaticYieldOnBehalf", contractAddress)
}

// ConfigureAutomaticYieldOnBehalf is a paid mutator transaction binding the contract method 0x3ba5713e.
//
// Solidity: function configureAutomaticYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastSession) ConfigureAutomaticYieldOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureAutomaticYieldOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureAutomaticYieldOnBehalf is a paid mutator transaction binding the contract method 0x3ba5713e.
//
// Solidity: function configureAutomaticYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactorSession) ConfigureAutomaticYieldOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureAutomaticYieldOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureClaimableGas is a paid mutator transaction binding the contract method 0x4e606c47.
//
// Solidity: function configureClaimableGas() returns()
func (_Blast *BlastTransactor) ConfigureClaimableGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureClaimableGas")
}

// ConfigureClaimableGas is a paid mutator transaction binding the contract method 0x4e606c47.
//
// Solidity: function configureClaimableGas() returns()
func (_Blast *BlastSession) ConfigureClaimableGas() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableGas(&_Blast.TransactOpts)
}

// ConfigureClaimableGas is a paid mutator transaction binding the contract method 0x4e606c47.
//
// Solidity: function configureClaimableGas() returns()
func (_Blast *BlastTransactorSession) ConfigureClaimableGas() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableGas(&_Blast.TransactOpts)
}

// ConfigureClaimableGasOnBehalf is a paid mutator transaction binding the contract method 0x908c8502.
//
// Solidity: function configureClaimableGasOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactor) ConfigureClaimableGasOnBehalf(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureClaimableGasOnBehalf", contractAddress)
}

// ConfigureClaimableGasOnBehalf is a paid mutator transaction binding the contract method 0x908c8502.
//
// Solidity: function configureClaimableGasOnBehalf(address contractAddress) returns()
func (_Blast *BlastSession) ConfigureClaimableGasOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableGasOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureClaimableGasOnBehalf is a paid mutator transaction binding the contract method 0x908c8502.
//
// Solidity: function configureClaimableGasOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactorSession) ConfigureClaimableGasOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableGasOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureClaimableYield is a paid mutator transaction binding the contract method 0xf098767a.
//
// Solidity: function configureClaimableYield() returns()
func (_Blast *BlastTransactor) ConfigureClaimableYield(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureClaimableYield")
}

// ConfigureClaimableYield is a paid mutator transaction binding the contract method 0xf098767a.
//
// Solidity: function configureClaimableYield() returns()
func (_Blast *BlastSession) ConfigureClaimableYield() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableYield(&_Blast.TransactOpts)
}

// ConfigureClaimableYield is a paid mutator transaction binding the contract method 0xf098767a.
//
// Solidity: function configureClaimableYield() returns()
func (_Blast *BlastTransactorSession) ConfigureClaimableYield() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableYield(&_Blast.TransactOpts)
}

// ConfigureClaimableYieldOnBehalf is a paid mutator transaction binding the contract method 0x37ebe3a8.
//
// Solidity: function configureClaimableYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactor) ConfigureClaimableYieldOnBehalf(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureClaimableYieldOnBehalf", contractAddress)
}

// ConfigureClaimableYieldOnBehalf is a paid mutator transaction binding the contract method 0x37ebe3a8.
//
// Solidity: function configureClaimableYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastSession) ConfigureClaimableYieldOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableYieldOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureClaimableYieldOnBehalf is a paid mutator transaction binding the contract method 0x37ebe3a8.
//
// Solidity: function configureClaimableYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactorSession) ConfigureClaimableYieldOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureClaimableYieldOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureContract is a paid mutator transaction binding the contract method 0x4c802f38.
//
// Solidity: function configureContract(address contractAddress, uint8 _yieldMode, uint8 _gasMode, address _newGovernor) returns()
func (_Blast *BlastTransactor) ConfigureContract(opts *bind.TransactOpts, contractAddress common.Address, _yieldMode uint8, _gasMode uint8, _newGovernor common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureContract", contractAddress, _yieldMode, _gasMode, _newGovernor)
}

// ConfigureContract is a paid mutator transaction binding the contract method 0x4c802f38.
//
// Solidity: function configureContract(address contractAddress, uint8 _yieldMode, uint8 _gasMode, address _newGovernor) returns()
func (_Blast *BlastSession) ConfigureContract(contractAddress common.Address, _yieldMode uint8, _gasMode uint8, _newGovernor common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureContract(&_Blast.TransactOpts, contractAddress, _yieldMode, _gasMode, _newGovernor)
}

// ConfigureContract is a paid mutator transaction binding the contract method 0x4c802f38.
//
// Solidity: function configureContract(address contractAddress, uint8 _yieldMode, uint8 _gasMode, address _newGovernor) returns()
func (_Blast *BlastTransactorSession) ConfigureContract(contractAddress common.Address, _yieldMode uint8, _gasMode uint8, _newGovernor common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureContract(&_Blast.TransactOpts, contractAddress, _yieldMode, _gasMode, _newGovernor)
}

// ConfigureGovernor is a paid mutator transaction binding the contract method 0xeb864698.
//
// Solidity: function configureGovernor(address _governor) returns()
func (_Blast *BlastTransactor) ConfigureGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureGovernor", _governor)
}

// ConfigureGovernor is a paid mutator transaction binding the contract method 0xeb864698.
//
// Solidity: function configureGovernor(address _governor) returns()
func (_Blast *BlastSession) ConfigureGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureGovernor(&_Blast.TransactOpts, _governor)
}

// ConfigureGovernor is a paid mutator transaction binding the contract method 0xeb864698.
//
// Solidity: function configureGovernor(address _governor) returns()
func (_Blast *BlastTransactorSession) ConfigureGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureGovernor(&_Blast.TransactOpts, _governor)
}

// ConfigureGovernorOnBehalf is a paid mutator transaction binding the contract method 0x0ca12c4b.
//
// Solidity: function configureGovernorOnBehalf(address _newGovernor, address contractAddress) returns()
func (_Blast *BlastTransactor) ConfigureGovernorOnBehalf(opts *bind.TransactOpts, _newGovernor common.Address, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureGovernorOnBehalf", _newGovernor, contractAddress)
}

// ConfigureGovernorOnBehalf is a paid mutator transaction binding the contract method 0x0ca12c4b.
//
// Solidity: function configureGovernorOnBehalf(address _newGovernor, address contractAddress) returns()
func (_Blast *BlastSession) ConfigureGovernorOnBehalf(_newGovernor common.Address, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureGovernorOnBehalf(&_Blast.TransactOpts, _newGovernor, contractAddress)
}

// ConfigureGovernorOnBehalf is a paid mutator transaction binding the contract method 0x0ca12c4b.
//
// Solidity: function configureGovernorOnBehalf(address _newGovernor, address contractAddress) returns()
func (_Blast *BlastTransactorSession) ConfigureGovernorOnBehalf(_newGovernor common.Address, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureGovernorOnBehalf(&_Blast.TransactOpts, _newGovernor, contractAddress)
}

// ConfigureVoidGas is a paid mutator transaction binding the contract method 0x2210dfb1.
//
// Solidity: function configureVoidGas() returns()
func (_Blast *BlastTransactor) ConfigureVoidGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureVoidGas")
}

// ConfigureVoidGas is a paid mutator transaction binding the contract method 0x2210dfb1.
//
// Solidity: function configureVoidGas() returns()
func (_Blast *BlastSession) ConfigureVoidGas() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidGas(&_Blast.TransactOpts)
}

// ConfigureVoidGas is a paid mutator transaction binding the contract method 0x2210dfb1.
//
// Solidity: function configureVoidGas() returns()
func (_Blast *BlastTransactorSession) ConfigureVoidGas() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidGas(&_Blast.TransactOpts)
}

// ConfigureVoidGasOnBehalf is a paid mutator transaction binding the contract method 0xeb59acdc.
//
// Solidity: function configureVoidGasOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactor) ConfigureVoidGasOnBehalf(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureVoidGasOnBehalf", contractAddress)
}

// ConfigureVoidGasOnBehalf is a paid mutator transaction binding the contract method 0xeb59acdc.
//
// Solidity: function configureVoidGasOnBehalf(address contractAddress) returns()
func (_Blast *BlastSession) ConfigureVoidGasOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidGasOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureVoidGasOnBehalf is a paid mutator transaction binding the contract method 0xeb59acdc.
//
// Solidity: function configureVoidGasOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactorSession) ConfigureVoidGasOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidGasOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureVoidYield is a paid mutator transaction binding the contract method 0xaa857d98.
//
// Solidity: function configureVoidYield() returns()
func (_Blast *BlastTransactor) ConfigureVoidYield(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureVoidYield")
}

// ConfigureVoidYield is a paid mutator transaction binding the contract method 0xaa857d98.
//
// Solidity: function configureVoidYield() returns()
func (_Blast *BlastSession) ConfigureVoidYield() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidYield(&_Blast.TransactOpts)
}

// ConfigureVoidYield is a paid mutator transaction binding the contract method 0xaa857d98.
//
// Solidity: function configureVoidYield() returns()
func (_Blast *BlastTransactorSession) ConfigureVoidYield() (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidYield(&_Blast.TransactOpts)
}

// ConfigureVoidYieldOnBehalf is a paid mutator transaction binding the contract method 0xb71d6dd4.
//
// Solidity: function configureVoidYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactor) ConfigureVoidYieldOnBehalf(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "configureVoidYieldOnBehalf", contractAddress)
}

// ConfigureVoidYieldOnBehalf is a paid mutator transaction binding the contract method 0xb71d6dd4.
//
// Solidity: function configureVoidYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastSession) ConfigureVoidYieldOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidYieldOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// ConfigureVoidYieldOnBehalf is a paid mutator transaction binding the contract method 0xb71d6dd4.
//
// Solidity: function configureVoidYieldOnBehalf(address contractAddress) returns()
func (_Blast *BlastTransactorSession) ConfigureVoidYieldOnBehalf(contractAddress common.Address) (*types.Transaction, error) {
	return _Blast.Contract.ConfigureVoidYieldOnBehalf(&_Blast.TransactOpts, contractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Blast *BlastTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blast.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Blast *BlastSession) Initialize() (*types.Transaction, error) {
	return _Blast.Contract.Initialize(&_Blast.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Blast *BlastTransactorSession) Initialize() (*types.Transaction, error) {
	return _Blast.Contract.Initialize(&_Blast.TransactOpts)
}

// BlastInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Blast contract.
type BlastInitializedIterator struct {
	Event *BlastInitialized // Event containing the contract specifics and raw log

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
func (it *BlastInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlastInitialized)
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
		it.Event = new(BlastInitialized)
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
func (it *BlastInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlastInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlastInitialized represents a Initialized event raised by the Blast contract.
type BlastInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Blast *BlastFilterer) FilterInitialized(opts *bind.FilterOpts) (*BlastInitializedIterator, error) {

	logs, sub, err := _Blast.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BlastInitializedIterator{contract: _Blast.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Blast *BlastFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BlastInitialized) (event.Subscription, error) {

	logs, sub, err := _Blast.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlastInitialized)
				if err := _Blast.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Blast *BlastFilterer) ParseInitialized(log types.Log) (*BlastInitialized, error) {
	event := new(BlastInitialized)
	if err := _Blast.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
