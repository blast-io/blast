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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yieldContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GAS_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YIELD_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimAllGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfYield\",\"type\":\"address\"}],\"name\":\"claimAllYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasToClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasSecondsToConsume\",\"type\":\"uint256\"}],\"name\":\"claimGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minClaimRateBips\",\"type\":\"uint256\"}],\"name\":\"claimGasAtMinClaimRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfGas\",\"type\":\"address\"}],\"name\":\"claimMaxGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientOfYield\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumYieldMode\",\"name\":\"_yieldMode\",\"type\":\"uint8\"},{\"internalType\":\"enumGasMode\",\"name\":\"_gasMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"configure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureAutomaticYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureAutomaticYieldOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureClaimableGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureClaimableGasOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureClaimableYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureClaimableYieldOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"enumYieldMode\",\"name\":\"_yieldMode\",\"type\":\"uint8\"},{\"internalType\":\"enumGasMode\",\"name\":\"_gasMode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"configureContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"configureGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureGovernorOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureVoidGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureVoidGasOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configureVoidYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"configureVoidYieldOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governorMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"readClaimableYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"readGasParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumGasMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"readYieldConfiguration\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620019ab380380620019ab833981016040819052620000349162000069565b6001600160a01b0391821660a05216608052620000a1565b80516001600160a01b03811681146200006457600080fd5b919050565b600080604083850312156200007d57600080fd5b62000088836200004c565b915062000098602084016200004c565b90509250929050565b60805160a05161183c6200016f6000396000818161030e015281816104fb01528181610604015281816107f90152818161095b01528181610a1101528181610d1601528181610dfa01528181610f43015281816110890152818161113f01526113fc0152600081816102e70152818161069f0152818161076d0152818161086601528181610aae01528181610bc501528181610c6301528181610e5801528181610eb401528181610faf015281816111ed0152818161128f015281816113410152611494015261183c6000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c8063a70d11bd116100f9578063eb86469811610097578063f971966211610071578063f971966214610400578063fafce39e14610413578063fd8c4b9d14610426578063fe9fbb801461044b57600080fd5b8063eb864698146103d2578063ec3278e8146103e5578063f098767a146103f857600080fd5b8063c8992e61116100d3578063c8992e611461034b578063dde798a41461035e578063e43581b814610381578063eb59acdc146103bf57600080fd5b8063a70d11bd14610309578063aa857d9814610330578063b71d6dd41461033857600080fd5b8063662aa11d11610166578063860043b611610140578063860043b6146102a9578063908c8502146102bc578063954fa5ee146102cf578063a0278c90146102e257600080fd5b8063662aa11d1461024d5780637114177a1461026057806385ebdc271461026857600080fd5b806337ebe3a8116101a257806337ebe3a81461020c5780633ba5713e1461021f5780634c802f38146102325780634e606c471461024557600080fd5b80630951888f146101c95780630ca12c4b146101ef5780632210dfb114610204575b600080fd5b6101dc6101d7366004611574565b61045e565b6040519081526020015b60405180910390f35b6102026101fd3660046115b0565b610573565b005b6102026105c8565b61020261021a3660046115e3565b610670565b61020261022d3660046115e3565b61073e565b61020261024036600461161a565b61079a565b61020261091f565b6101dc61025b3660046115b0565b610993565b610202610a7f565b6102916102763660046115e3565b6000602081905290815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016101e6565b6101dc6102b73660046115b0565b610b4c565b6102026102ca3660046115e3565b610cda565b6101dc6102dd3660046115b0565b610d7c565b6102917f000000000000000000000000000000000000000000000000000000000000000081565b6102917f000000000000000000000000000000000000000000000000000000000000000081565b610202610e29565b6102026103463660046115e3565b610e85565b610202610359366004611670565b610ee1565b61037161036c3660046115e3565b611061565b6040516101e694939291906116ed565b6103af61038f3660046115e3565b6001600160a01b0390811660009081526020819052604090205416331490565b60405190151581526020016101e6565b6102026103cd3660046115e3565b611103565b6102026103e03660046115e3565b611177565b6101dc6103f33660046115e3565b6111cb565b610202611260565b6101dc61040e366004611574565b6112bc565b6101dc61042136600461170f565b611370565b6104396104343660046115e3565b611472565b60405160ff90911681526020016101e6565b6103af6104593660046115e3565b611501565b600061046984611501565b6104cd5760405162461bcd60e51b815260206004820152602a60248201527f4e6f7420616c6c6f77656420746f20636c61696d20676173206174206d696e20604482015269636c61696d207261746560b01b60648201526084015b60405180910390fd5b604051630951888f60e01b81526001600160a01b0385811660048301528481166024830152604482018490527f00000000000000000000000000000000000000000000000000000000000000001690630951888f906064015b6020604051808303816000875af1158015610545573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105699190611751565b90505b9392505050565b61057c81611501565b6105985760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b0390811660009081526020819052604090208054919092166001600160a01b0319909116179055565b6105d133611501565b6105ed5760405162461bcd60e51b81526004016104c49061176a565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061063c9033906000906004016117ae565b600060405180830381600087803b15801561065657600080fd5b505af115801561066a573d6000803e3d6000fd5b50505050565b61067981611501565b6106955760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260025b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610716573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073a9190611751565b5050565b61074781611501565b6107635760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260006106c8565b6107a384611501565b6107bf5760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b038481166000908152602081905260409081902080546001600160a01b0319168484161790555163d4810ba560e01b81527f00000000000000000000000000000000000000000000000000000000000000009091169063d4810ba59061083290879086906004016117ae565b600060405180830381600087803b15801561084c57600080fd5b505af1158015610860573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633bdbe9a5858560028111156108a6576108a66116b5565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af11580156108f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109189190611751565b5050505050565b61092833611501565b6109445760405162461bcd60e51b81526004016104c49061176a565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061063c9033906001906004016117ae565b600061099e83611501565b6109ea5760405162461bcd60e51b815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d206d6178206761730000000060448201526064016104c4565b604051630928b10d60e31b81526001600160a01b03848116600483015283811660248301527f000000000000000000000000000000000000000000000000000000000000000016906349458868906044015b6020604051808303816000875af1158015610a5b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c9190611751565b610a8833611501565b610aa45760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a53360005b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610b25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b499190611751565b50565b6000610b5783611501565b610ba35760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c6400000060448201526064016104c4565b60405163e12f3a6160e01b81526001600160a01b0384811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015610c0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c329190611751565b60405163132d974d60e31b81526001600160a01b0386811660048301528581166024830152604482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063996cba68906064016020604051808303816000875af1158015610cae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd29190611751565b949350505050565b610ce381611501565b610cff5760405162461bcd60e51b81526004016104c49061176a565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba590610d4e9084906001906004016117ae565b600060405180830381600087803b158015610d6857600080fd5b505af1158015610918573d6000803e3d6000fd5b6000610d8783611501565b610dd35760405162461bcd60e51b815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d20616c6c206761730000000060448201526064016104c4565b604051635767bba560e01b81526001600160a01b03848116600483015283811660248301527f00000000000000000000000000000000000000000000000000000000000000001690635767bba590604401610a3c565b610e3233611501565b610e4e5760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336001610ad7565b610e8e81611501565b610eaa5760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260016106c8565b610eea33611501565b610f065760405162461bcd60e51b81526004016104c49061176a565b336000818152602081905260409081902080546001600160a01b0319166001600160a01b0385811691909117909155905163d4810ba560e01b81527f00000000000000000000000000000000000000000000000000000000000000009091169163d4810ba591610f7b919086906004016117ae565b600060405180830381600087803b158015610f9557600080fd5b505af1158015610fa9573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633bdbe9a533856002811115610fef57610fef6116b5565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af115801561103d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066a9190611751565b604051633779e62960e21b81526001600160a01b0382811660048301526000918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063dde798a490602401608060405180830381865afa1580156110d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f491906117cb565b93509350935093509193509193565b61110c81611501565b6111285760405162461bcd60e51b81526004016104c49061176a565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba590610d4e9084906000906004016117ae565b61118033611501565b61119c5760405162461bcd60e51b81526004016104c49061176a565b33600090815260208190526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b60405163e12f3a6160e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015611236573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125a9190611751565b92915050565b61126933611501565b6112855760405162461bcd60e51b81526004016104c49061176a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336002610ad7565b60006112c784611501565b6113135760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c6400000060448201526064016104c4565b60405163132d974d60e31b81526001600160a01b0385811660048301528481166024830152604482018490527f0000000000000000000000000000000000000000000000000000000000000000169063996cba6890606401610526565b600061137b85611501565b6113c75760405162461bcd60e51b815260206004820152601860248201527f4e6f7420616c6c6f77656420746f20636c61696d20676173000000000000000060448201526064016104c4565b604051631357a41960e11b81526001600160a01b038681166004830152858116602483015260448201859052606482018490527f000000000000000000000000000000000000000000000000000000000000000016906326af4832906084016020604051808303816000875af1158015611445573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114699190611751565b95945050505050565b60405163c44b11f760e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063c44b11f790602401602060405180830381865afa1580156114dd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125a919061180c565b6001600160a01b0380821660009081526020819052604081205490911633148061125a57506001600160a01b038083166000908152602081905260409020541615801561125a5750506001600160a01b0316331490565b80356001600160a01b038116811461156f57600080fd5b919050565b60008060006060848603121561158957600080fd5b61159284611558565b92506115a060208501611558565b9150604084013590509250925092565b600080604083850312156115c357600080fd5b6115cc83611558565b91506115da60208401611558565b90509250929050565b6000602082840312156115f557600080fd5b61056c82611558565b80356003811061156f57600080fd5b60028110610b4957600080fd5b6000806000806080858703121561163057600080fd5b61163985611558565b9350611647602086016115fe565b925060408501356116578161160d565b915061166560608601611558565b905092959194509250565b60008060006060848603121561168557600080fd5b61168e846115fe565b9250602084013561169e8161160d565b91506116ac60408501611558565b90509250925092565b634e487b7160e01b600052602160045260246000fd5b600281106116e957634e487b7160e01b600052602160045260246000fd5b9052565b84815260208101849052604081018390526080810161146960608301846116cb565b6000806000806080858703121561172557600080fd5b61172e85611558565b935061173c60208601611558565b93969395505050506040820135916060013590565b60006020828403121561176357600080fd5b5051919050565b60208082526024908201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e746040820152631c9858dd60e21b606082015260800190565b6001600160a01b03831681526040810161056c60208301846116cb565b600080600080608085870312156117e157600080fd5b84519350602085015192506040850151915060608501516118018161160d565b939692955090935050565b60006020828403121561181e57600080fd5b815160ff8116811461056c57600080fdfea164736f6c634300080f000a",
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
