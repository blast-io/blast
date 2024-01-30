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
	Bin: "0x60c06040523480156200001157600080fd5b506040516200249838038062002498833981016040819052620000349162000069565b6001600160a01b0391821660a05216608052620000a1565b80516001600160a01b03811681146200006457600080fd5b919050565b600080604083850312156200007d57600080fd5b62000088836200004c565b915062000098602084016200004c565b90509250929050565b60805160a0516123296200016f6000396000818161032801528181610575015281816107a701528181610b6601528181610d8f01528181610e85015281816113260152818161144a01528181611738015281816118d601528181611a210152611e94015260008181610301015281816108be01528181610a2d01528181610bd301528181610f9e0152818161111a015281816111de01528181611524015281816115fc015281816117a401528181611b8901528181611ca701528181611d990152611f5201526123296000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c8063a70d11bd116100f9578063eb86469811610097578063f971966211610071578063f971966214610427578063fafce39e1461043a578063fd8c4b9d1461044d578063fe9fbb801461047257600080fd5b8063eb864698146103f9578063ec3278e81461040c578063f098767a1461041f57600080fd5b8063c8992e61116100d3578063c8992e6114610365578063dde798a414610378578063e43581b81461039b578063eb59acdc146103e657600080fd5b8063a70d11bd14610323578063aa857d981461034a578063b71d6dd41461035257600080fd5b8063662aa11d11610166578063860043b611610140578063860043b6146102c3578063908c8502146102d6578063954fa5ee146102e9578063a0278c90146102fc57600080fd5b8063662aa11d1461024d5780637114177a1461026057806385ebdc271461026857600080fd5b806337ebe3a8116101a257806337ebe3a81461020c5780633ba5713e1461021f5780634c802f38146102325780634e606c471461024557600080fd5b80630951888f146101c95780630ca12c4b146101ef5780632210dfb114610204575b600080fd5b6101dc6101d7366004612066565b610485565b6040519081526020015b60405180910390f35b6102026101fd3660046120a2565b6105ed565b005b6102026106d6565b61020261021a3660046120d5565b610813565b61020261022d3660046120d5565b610982565b61020261024036600461210c565b610a5a565b610202610cbe565b6101dc61025b3660046120a2565b610dc7565b610202610ef3565b61029e6102763660046120d5565b60006020819052908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e6565b6101dc6102d13660046120a2565b611061565b6102026102e43660046120d5565b611255565b6101dc6102f73660046120a2565b61138c565b61029e7f000000000000000000000000000000000000000000000000000000000000000081565b61029e7f000000000000000000000000000000000000000000000000000000000000000081565b610202611479565b6102026103603660046120d5565b611551565b610202610373366004612162565b611629565b61038b6103863660046120d5565b611888565b6040516101e69493929190612211565b6103d66103a93660046120d5565b73ffffffffffffffffffffffffffffffffffffffff90811660009081526020819052604090205416331490565b60405190151581526020016101e6565b6102026103f43660046120d5565b611950565b6102026104073660046120d5565b611a59565b6101dc61041a3660046120d5565b611b41565b610202611bfc565b6101dc610435366004612066565b611cd4565b6101dc610448366004612233565b611dc8565b61046061045b3660046120d5565b611f0a565b60405160ff90911681526020016101e6565b6103d66104803660046120d5565b611fbf565b600061049084611fbf565b610521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4e6f7420616c6c6f77656420746f20636c61696d20676173206174206d696e2060448201527f636c61696d20726174650000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6040517f0951888f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301528481166024830152604482018490527f00000000000000000000000000000000000000000000000000000000000000001690630951888f906064015b6020604051808303816000875af11580156105bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e39190612275565b90505b9392505050565b6105f681611fbf565b610681576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff90811660009081526020819052604090208054919092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055565b6106df33611fbf565b61076a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b6040517fd4810ba500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba5906107df90339060009060040161228e565b600060405180830381600087803b1580156107f957600080fd5b505af115801561080d573d6000803e3d6000fd5b50505050565b61081c81611fbf565b6108a7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260025b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260ff1660248201526044016020604051808303816000875af115801561095a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097e9190612275565b5050565b61098b81611fbf565b610a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260006108e7565b610a6384611fbf565b610aee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff8481166000908152602081905260409081902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016848416179055517fd4810ba50000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000009091169063d4810ba590610b9f908790869060040161228e565b600060405180830381600087803b158015610bb957600080fd5b505af1158015610bcd573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633bdbe9a585856002811115610c2057610c206121a7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260ff1660248201526044016020604051808303816000875af1158015610c93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb79190612275565b5050505050565b610cc733611fbf565b610d52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b6040517fd4810ba500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba5906107df90339060019060040161228e565b6000610dd283611fbf565b610e38576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d206d617820676173000000006044820152606401610518565b6040517f4945886800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015283811660248301527f000000000000000000000000000000000000000000000000000000000000000016906349458868906044015b6020604051808303816000875af1158015610ecf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e69190612275565b610efc33611fbf565b610f87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a53360005b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260ff1660248201526044016020604051808303816000875af115801561103a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105e9190612275565b50565b600061106c83611fbf565b6110d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c640000006044820152606401610518565b6040517fe12f3a6100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015611163573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111879190612275565b6040517f996cba6800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301528581166024830152604482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063996cba68906064016020604051808303816000875af1158015611229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124d9190612275565b949350505050565b61125e81611fbf565b6112e9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b6040517fd4810ba500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061135e90849060019060040161228e565b600060405180830381600087803b15801561137857600080fd5b505af1158015610cb7573d6000803e3d6000fd5b600061139783611fbf565b6113fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d20616c6c20676173000000006044820152606401610518565b6040517f5767bba500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015283811660248301527f00000000000000000000000000000000000000000000000000000000000000001690635767bba590604401610eb0565b61148233611fbf565b61150d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336001610fc7565b61155a81611fbf565b6115e5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260016108e7565b61163233611fbf565b6116bd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b336000818152602081905260409081902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8581169190911790915590517fd4810ba50000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000009091169163d4810ba5916117709190869060040161228e565b600060405180830381600087803b15801561178a57600080fd5b505af115801561179e573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633bdbe9a5338560028111156117f1576117f16121a7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260ff1660248201526044016020604051808303816000875af1158015611864573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080d9190612275565b6040517fdde798a400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063dde798a490602401608060405180830381865afa15801561191d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061194191906122b8565b93509350935093509193509193565b61195981611fbf565b6119e4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b6040517fd4810ba500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061135e90849060009060040161228e565b611a6233611fbf565b611aed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b33600090815260208190526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6040517fe12f3a6100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015611bd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bf69190612275565b92915050565b611c0533611fbf565b611c90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e7460448201527f72616374000000000000000000000000000000000000000000000000000000006064820152608401610518565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336002610fc7565b6000611cdf84611fbf565b611d45576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c640000006044820152606401610518565b6040517f996cba6800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301528481166024830152604482018490527f0000000000000000000000000000000000000000000000000000000000000000169063996cba68906064016105a0565b6000611dd385611fbf565b611e39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4e6f7420616c6c6f77656420746f20636c61696d2067617300000000000000006044820152606401610518565b6040517f26af483200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152858116602483015260448201859052606482018490527f000000000000000000000000000000000000000000000000000000000000000016906326af4832906084016020604051808303816000875af1158015611edd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f019190612275565b95945050505050565b6040517fc44b11f700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063c44b11f790602401602060405180830381865afa158015611f9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bf691906122f9565b73ffffffffffffffffffffffffffffffffffffffff808216600090815260208190526040812054909116331480611bf6575073ffffffffffffffffffffffffffffffffffffffff80831660009081526020819052604090205416158015611bf657505073ffffffffffffffffffffffffffffffffffffffff16331490565b803573ffffffffffffffffffffffffffffffffffffffff8116811461206157600080fd5b919050565b60008060006060848603121561207b57600080fd5b6120848461203d565b92506120926020850161203d565b9150604084013590509250925092565b600080604083850312156120b557600080fd5b6120be8361203d565b91506120cc6020840161203d565b90509250929050565b6000602082840312156120e757600080fd5b6105e68261203d565b80356003811061206157600080fd5b6002811061105e57600080fd5b6000806000806080858703121561212257600080fd5b61212b8561203d565b9350612139602086016120f0565b92506040850135612149816120ff565b91506121576060860161203d565b905092959194509250565b60008060006060848603121561217757600080fd5b612180846120f0565b92506020840135612190816120ff565b915061219e6040850161203d565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811061220d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b848152602081018490526040810183905260808101611f0160608301846121d6565b6000806000806080858703121561224957600080fd5b6122528561203d565b93506122606020860161203d565b93969395505050506040820135916060013590565b60006020828403121561228757600080fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff83168152604081016105e660208301846121d6565b600080600080608085870312156122ce57600080fd5b84519350602085015192506040850151915060608501516122ee816120ff565b939692955090935050565b60006020828403121561230b57600080fd5b815160ff811681146105e657600080fdfea164736f6c634300080f000a",
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
