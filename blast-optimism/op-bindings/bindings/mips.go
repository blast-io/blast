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

// MIPSMetaData contains all meta data concerning the MIPS contract.
var MIPSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPreimageOracle\",\"name\":\"_oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BRK_START\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIPreimageOracle\",\"name\":\"oracle_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_stateData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_localContext\",\"type\":\"uint256\"}],\"name\":\"step\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611cb2380380611cb283398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051611c21610091600039600081816078015261149f0152611c216000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063155633fe146100465780637dc0d1d01461006b578063836e7b32146100a2575b600080fd5b610051634000000081565b60405163ffffffff90911681526020015b60405180910390f35b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152602001610062565b6100b56100b0366004611b50565b6100c3565b604051908152602001610062565b60006100cd611a7d565b608081146100da57600080fd5b604051610600146100ea57600080fd5b608487146100f757600080fd5b6101a4851461010557600080fd5b8635608052602087013560a052604087013560e090811c60c09081526044890135821c82526048890135821c61010052604c890135821c610120526050890135821c61014052605489013590911c61016052605888013560f890811c610180526059890135901c6101a052605a880135901c6101c0526102006101e0819052606288019060005b60208110156101b057823560e01c825260049092019160209091019060010161018c565b505050806101200151156101ce576101c661060e565b915050610605565b6101408101805160010167ffffffffffffffff16905260608101516000906101f69082610712565b9050603f601a82901c16600281148061021557508063ffffffff166003145b1561026a5760006002836303ffffff1663ffffffff16901b846080015163f00000001617905061025f8263ffffffff1660021461025357601f610256565b60005b60ff16826107ce565b945050505050610605565b6101608301516000908190601f601086901c81169190601587901c166020811061029657610296611bc4565b602002015192508063ffffffff851615806102b757508463ffffffff16601c145b156102ee578661016001518263ffffffff16602081106102d9576102d9611bc4565b6020020151925050601f600b86901c166103aa565b60208563ffffffff161015610350578463ffffffff16600c148061031857508463ffffffff16600d145b8061032957508463ffffffff16600e145b1561033a578561ffff1692506103aa565b6103498661ffff16601061089a565b92506103aa565b60288563ffffffff1610158061036c57508463ffffffff166022145b8061037d57508463ffffffff166026145b156103aa578661016001518263ffffffff166020811061039f5761039f611bc4565b602002015192508190505b60048563ffffffff16101580156103c7575060088563ffffffff16105b806103d857508463ffffffff166001145b156103f7576103e9858784876108ef565b975050505050505050610605565b63ffffffff600060208783161061045c576104178861ffff16601061089a565b9095019463fffffffc861661042d816001610712565b915060288863ffffffff161015801561044d57508763ffffffff16603014155b1561045a57809250600093505b505b600061046a89888885610adc565b63ffffffff9081169150603f8a1690891615801561048f575060088163ffffffff1610155b80156104a15750601c8163ffffffff16105b1561057e578063ffffffff16600814806104c157508063ffffffff166009145b156104f8576104e68163ffffffff166008146104dd57856104e0565b60005b896107ce565b9b505050505050505050505050610605565b8063ffffffff16600a03610518576104e6858963ffffffff8a1615611248565b8063ffffffff16600b03610539576104e6858963ffffffff8a161515611248565b8063ffffffff16600c03610550576104e68d611305565b60108163ffffffff161015801561056d5750601c8163ffffffff16105b1561057e576104e68189898861176a565b8863ffffffff166038148015610599575063ffffffff861615155b156105ce5760018b61016001518763ffffffff16602081106105bd576105bd611bc4565b63ffffffff90921660209290920201525b8363ffffffff1663ffffffff146105eb576105eb84600184611964565b6105f785836001611248565b9b5050505050505050505050505b95945050505050565b60408051608051815260a051602082015260dc519181019190915260fc51604482015261011c51604882015261013c51604c82015261015c51605082015261017c5160548201526101805161019f5160588301526101a0516101bf5160598401526101d851605a840152600092610200929091606283019190855b60208110156106ad57601c8601518452602090950194600490930192600101610689565b506000835283830384a06000945080600181146106cd57600395506106f5565b8280156106e557600181146106ee57600296506106f3565b600096506106f3565b600196505b505b50505081900390206001600160f81b031660f89190911b17919050565b60008061071e83611a08565b9050600384161561072e57600080fd5b6020810190358460051c8160005b601b8110156107945760208501943583821c600116801561076457600181146107795761078a565b6000848152602083905260409020935061078a565b600082815260208590526040902093505b505060010161073c565b5060805191508181146107af57630badf00d60005260206000fd5b5050601f94909416601c0360031b9390931c63ffffffff169392505050565b60006107d8611a7d565b60809050806060015160040163ffffffff16816080015163ffffffff161461083c5760405162461bcd60e51b81526020600482015260126024820152711a9d5b5c081a5b8819195b185e481cdb1bdd60721b60448201526064015b60405180910390fd5b60608101805160808301805163ffffffff90811690935285831690529085161561089257806008018261016001518663ffffffff166020811061088157610881611bc4565b63ffffffff90921660209290920201525b61060561060e565b600063ffffffff83811660001980850183169190911c821615159160016020869003821681901b830191861691821b92911b01826108d95760006108db565b815b90861663ffffffff16179250505092915050565b60006108f9611a7d565b608090506000816060015160040163ffffffff16826080015163ffffffff161461095c5760405162461bcd60e51b8152602060048201526014602482015273189c985b98da081a5b8819195b185e481cdb1bdd60621b6044820152606401610833565b8663ffffffff166004148061097757508663ffffffff166005145b156109f35760008261016001518663ffffffff166020811061099b5761099b611bc4565b602002015190508063ffffffff168563ffffffff161480156109c357508763ffffffff166004145b806109eb57508063ffffffff168563ffffffff16141580156109eb57508763ffffffff166005145b915050610a70565b8663ffffffff16600603610a105760008460030b13159050610a70565b8663ffffffff16600703610a2c5760008460030b139050610a70565b8663ffffffff16600103610a7057601f601087901c166000819003610a555760008560030b1291505b8063ffffffff16600103610a6e5760008560030b121591505b505b606082018051608084015163ffffffff169091528115610ab6576002610a9b8861ffff16601061089a565b63ffffffff90811690911b8201600401166080840152610ac8565b60808301805160040163ffffffff1690525b610ad061060e565b98975050505050505050565b6000603f601a86901c16801580610b0b575060088163ffffffff1610158015610b0b5750600f8163ffffffff16105b15610f3d57603f86168160088114610b525760098114610b5b57600a8114610b6457600b8114610b6d57600c8114610b7657600d8114610b7f57600e8114610b8857610b8d565b60209150610b8d565b60219150610b8d565b602a9150610b8d565b602b9150610b8d565b60249150610b8d565b60259150610b8d565b602691505b508063ffffffff16600003610bb45750505063ffffffff8216601f600686901c161b611240565b8063ffffffff16600203610bda5750505063ffffffff8216601f600686901c161c611240565b8063ffffffff16600303610c1057601f600688901c16610c0663ffffffff8716821c602083900361089a565b9350505050611240565b8063ffffffff16600403610c325750505063ffffffff8216601f84161b611240565b8063ffffffff16600603610c545750505063ffffffff8216601f84161c611240565b8063ffffffff16600703610c8757610c7e8663ffffffff168663ffffffff16901c8760200361089a565b92505050611240565b8063ffffffff16600803610c9f578592505050611240565b8063ffffffff16600903610cb7578592505050611240565b8063ffffffff16600a03610ccf578592505050611240565b8063ffffffff16600b03610ce7578592505050611240565b8063ffffffff16600c03610cff578592505050611240565b8063ffffffff16600f03610d17578592505050611240565b8063ffffffff16601003610d2f578592505050611240565b8063ffffffff16601103610d47578592505050611240565b8063ffffffff16601203610d5f578592505050611240565b8063ffffffff16601303610d77578592505050611240565b8063ffffffff16601803610d8f578592505050611240565b8063ffffffff16601903610da7578592505050611240565b8063ffffffff16601a03610dbf578592505050611240565b8063ffffffff16601b03610dd7578592505050611240565b8063ffffffff16602003610df057505050828201611240565b8063ffffffff16602103610e0957505050828201611240565b8063ffffffff16602203610e2257505050818303611240565b8063ffffffff16602303610e3b57505050818303611240565b8063ffffffff16602403610e5457505050828216611240565b8063ffffffff16602503610e6d57505050828217611240565b8063ffffffff16602603610e8657505050828218611240565b8063ffffffff16602703610ea05750505082821719611240565b8063ffffffff16602a03610ed1578460030b8660030b12610ec2576000610ec5565b60015b60ff1692505050611240565b8063ffffffff16602b03610ef9578463ffffffff168663ffffffff1610610ec2576000610ec5565b60405162461bcd60e51b815260206004820152601360248201527234b73b30b634b21034b739ba393ab1ba34b7b760691b6044820152606401610833565b50610ef9565b8063ffffffff16601c03610fc157603f86166002819003610f6357505050828202611240565b8063ffffffff1660201480610f7e57508063ffffffff166021145b15610f37578063ffffffff16602003610f95579419945b60005b6380000000871615610fb7576401fffffffe600197881b169601610f98565b9250611240915050565b8063ffffffff16600f03610fe357505065ffffffff0000601083901b16611240565b8063ffffffff1660200361101f576110178560031660080260180363ffffffff168463ffffffff16901c60ff16600861089a565b915050611240565b8063ffffffff16602103611054576110178560021660080260100363ffffffff168463ffffffff16901c61ffff16601061089a565b8063ffffffff1660220361108357505063ffffffff60086003851602811681811b198416918316901b17611240565b8063ffffffff1660230361109a5782915050611240565b8063ffffffff166024036110cc578460031660080260180363ffffffff168363ffffffff16901c60ff16915050611240565b8063ffffffff166025036110ff578460021660080260100363ffffffff168363ffffffff16901c61ffff16915050611240565b8063ffffffff1660260361113157505063ffffffff60086003851602601803811681811c198416918316901c17611240565b8063ffffffff1660280361116757505060ff63ffffffff60086003861602601803811682811b9091188316918416901b17611240565b8063ffffffff1660290361119e57505061ffff63ffffffff60086002861602601003811682811b9091188316918416901b17611240565b8063ffffffff16602a036111cd57505063ffffffff60086003851602811681811c198316918416901c17611240565b8063ffffffff16602b036111e45783915050611240565b8063ffffffff16602e0361121657505063ffffffff60086003851602601803811681811b198316918416901b17611240565b8063ffffffff1660300361122d5782915050611240565b8063ffffffff16603803610ef957839150505b949350505050565b6000611252611a7d565b506080602063ffffffff86161061129c5760405162461bcd60e51b815260206004820152600e60248201526d3b30b634b2103932b3b4b9ba32b960911b6044820152606401610833565b63ffffffff8516158015906112ae5750825b156112e257838161016001518663ffffffff16602081106112d1576112d1611bc4565b63ffffffff90921660209290920201525b60808101805163ffffffff8082166060850152600490910116905261060561060e565b600061130f611a7d565b506101e051604081015160808083015160a084015160c09094015191936000928392919063ffffffff8616610ffa036113895781610fff81161561135857610fff811661100003015b8363ffffffff1660000361137f5760e08801805163ffffffff838201169091529550611383565b8395505b50611729565b8563ffffffff16610fcd036113a45763400000009450611729565b8563ffffffff16611018036113bc5760019450611729565b8563ffffffff16611096036113f257600161012088015260ff83166101008801526113e561060e565b9998505050505050505050565b8563ffffffff16610fa3036115c85763ffffffff8316156117295760041963ffffffff8416016115a057600061142f8363fffffffc166001610712565b60208901519091508060001a60010361146a57604080516000838152336020528d83526060902091526001600160f81b0316600160f81b1790505b6040808a0151905163e03110e160e01b81526004810183905263ffffffff909116602482015260009081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063e03110e1906044016040805180830381865afa1580156114e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115099190611bda565b91509150600386168060040382811015611521578092505b508186101561152e578591505b8260088302610100031c9250826008828460040303021b9250600180600883600403021b036001806008858560040303021b039150811981169050838119871617955050506115858663fffffffc16600186611964565b60408b018051820163ffffffff16905297506115c392505050565b60021963ffffffff8416016115b757809450611729565b63ffffffff9450600993505b611729565b8563ffffffff16610fa40361167d5763ffffffff8316600114806115f2575063ffffffff83166002145b80611603575063ffffffff83166004145b1561161057809450611729565b60051963ffffffff8416016115b75760006116328363fffffffc166001610712565b6020890151909150600384166004038381101561164d578093505b83900360089081029290921c600019600193850293841b0116911b17602088015260006040880152935083611729565b8563ffffffff16610fd703611729578163ffffffff1660030361171d5763ffffffff831615806116b3575063ffffffff83166005145b806116c4575063ffffffff83166003145b156116d25760009450611729565b63ffffffff8316600114806116ed575063ffffffff83166002145b806116fe575063ffffffff83166006145b8061170f575063ffffffff83166004145b156115b75760019450611729565b63ffffffff9450601693505b6101608701805163ffffffff808816604090920191909152905185821660e09091015260808801805180831660608b015260040190911690526113e561060e565b6000611774611a7d565b506080600063ffffffff8716601003611792575060c08101516118fb565b8663ffffffff166011036117b15763ffffffff861660c08301526118fb565b8663ffffffff166012036117ca575060a08101516118fb565b8663ffffffff166013036117e95763ffffffff861660a08301526118fb565b8663ffffffff1660180361181d5763ffffffff600387810b9087900b02602081901c821660c08501521660a08301526118fb565b8663ffffffff1660190361184e5763ffffffff86811681871602602081901c821660c08501521660a08301526118fb565b8663ffffffff16601a036118a4578460030b8660030b8161187157611871611bfe565b0763ffffffff1660c0830152600385810b9087900b8161189357611893611bfe565b0563ffffffff1660a08301526118fb565b8663ffffffff16601b036118fb578463ffffffff168663ffffffff16816118cd576118cd611bfe565b0663ffffffff90811660c0840152858116908716816118ee576118ee611bfe565b0463ffffffff1660a08301525b63ffffffff84161561193657808261016001518563ffffffff166020811061192557611925611bc4565b63ffffffff90921660209290920201525b60808201805163ffffffff8082166060860152600490910116905261195961060e565b979650505050505050565b600061196f83611a08565b9050600384161561197f57600080fd5b6020810190601f8516601c0360031b83811b913563ffffffff90911b1916178460051c60005b601b8110156119fd5760208401933582821c60011680156119cd57600181146119e2576119f3565b600085815260208390526040902094506119f3565b600082815260208690526040902094505b50506001016119a5565b505060805250505050565b60ff8116610380026101a4810190369061052401811015611a775760405162461bcd60e51b815260206004820152602360248201527f636865636b207468617420746865726520697320656e6f7567682063616c6c6460448201526261746160e81b6064820152608401610833565b50919050565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101919091526101608101611ae3611ae8565b905290565b6040518061040001604052806020906020820280368337509192915050565b60008083601f840112611b1957600080fd5b50813567ffffffffffffffff811115611b3157600080fd5b602083019150836020828501011115611b4957600080fd5b9250929050565b600080600080600060608688031215611b6857600080fd5b853567ffffffffffffffff80821115611b8057600080fd5b611b8c89838a01611b07565b90975095506020880135915080821115611ba557600080fd5b50611bb288828901611b07565b96999598509660400135949350505050565b634e487b7160e01b600052603260045260246000fd5b60008060408385031215611bed57600080fd5b505080516020909101519092909150565b634e487b7160e01b600052601260045260246000fdfea164736f6c634300080f000a",
}

// MIPSABI is the input ABI used to generate the binding from.
// Deprecated: Use MIPSMetaData.ABI instead.
var MIPSABI = MIPSMetaData.ABI

// MIPSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MIPSMetaData.Bin instead.
var MIPSBin = MIPSMetaData.Bin

// DeployMIPS deploys a new Ethereum contract, binding an instance of MIPS to it.
func DeployMIPS(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address) (common.Address, *types.Transaction, *MIPS, error) {
	parsed, err := MIPSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MIPSBin), backend, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MIPS{MIPSCaller: MIPSCaller{contract: contract}, MIPSTransactor: MIPSTransactor{contract: contract}, MIPSFilterer: MIPSFilterer{contract: contract}}, nil
}

// MIPS is an auto generated Go binding around an Ethereum contract.
type MIPS struct {
	MIPSCaller     // Read-only binding to the contract
	MIPSTransactor // Write-only binding to the contract
	MIPSFilterer   // Log filterer for contract events
}

// MIPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type MIPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MIPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MIPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MIPSSession struct {
	Contract     *MIPS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MIPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MIPSCallerSession struct {
	Contract *MIPSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MIPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MIPSTransactorSession struct {
	Contract     *MIPSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MIPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type MIPSRaw struct {
	Contract *MIPS // Generic contract binding to access the raw methods on
}

// MIPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MIPSCallerRaw struct {
	Contract *MIPSCaller // Generic read-only contract binding to access the raw methods on
}

// MIPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MIPSTransactorRaw struct {
	Contract *MIPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMIPS creates a new instance of MIPS, bound to a specific deployed contract.
func NewMIPS(address common.Address, backend bind.ContractBackend) (*MIPS, error) {
	contract, err := bindMIPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MIPS{MIPSCaller: MIPSCaller{contract: contract}, MIPSTransactor: MIPSTransactor{contract: contract}, MIPSFilterer: MIPSFilterer{contract: contract}}, nil
}

// NewMIPSCaller creates a new read-only instance of MIPS, bound to a specific deployed contract.
func NewMIPSCaller(address common.Address, caller bind.ContractCaller) (*MIPSCaller, error) {
	contract, err := bindMIPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MIPSCaller{contract: contract}, nil
}

// NewMIPSTransactor creates a new write-only instance of MIPS, bound to a specific deployed contract.
func NewMIPSTransactor(address common.Address, transactor bind.ContractTransactor) (*MIPSTransactor, error) {
	contract, err := bindMIPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MIPSTransactor{contract: contract}, nil
}

// NewMIPSFilterer creates a new log filterer instance of MIPS, bound to a specific deployed contract.
func NewMIPSFilterer(address common.Address, filterer bind.ContractFilterer) (*MIPSFilterer, error) {
	contract, err := bindMIPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MIPSFilterer{contract: contract}, nil
}

// bindMIPS binds a generic wrapper to an already deployed contract.
func bindMIPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MIPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MIPS *MIPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MIPS.Contract.MIPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MIPS *MIPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MIPS.Contract.MIPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MIPS *MIPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MIPS.Contract.MIPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MIPS *MIPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MIPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MIPS *MIPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MIPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MIPS *MIPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MIPS.Contract.contract.Transact(opts, method, params...)
}

// BRKSTART is a free data retrieval call binding the contract method 0x155633fe.
//
// Solidity: function BRK_START() view returns(uint32)
func (_MIPS *MIPSCaller) BRKSTART(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MIPS.contract.Call(opts, &out, "BRK_START")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BRKSTART is a free data retrieval call binding the contract method 0x155633fe.
//
// Solidity: function BRK_START() view returns(uint32)
func (_MIPS *MIPSSession) BRKSTART() (uint32, error) {
	return _MIPS.Contract.BRKSTART(&_MIPS.CallOpts)
}

// BRKSTART is a free data retrieval call binding the contract method 0x155633fe.
//
// Solidity: function BRK_START() view returns(uint32)
func (_MIPS *MIPSCallerSession) BRKSTART() (uint32, error) {
	return _MIPS.Contract.BRKSTART(&_MIPS.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address oracle_)
func (_MIPS *MIPSCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MIPS.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address oracle_)
func (_MIPS *MIPSSession) Oracle() (common.Address, error) {
	return _MIPS.Contract.Oracle(&_MIPS.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address oracle_)
func (_MIPS *MIPSCallerSession) Oracle() (common.Address, error) {
	return _MIPS.Contract.Oracle(&_MIPS.CallOpts)
}

// Step is a paid mutator transaction binding the contract method 0x836e7b32.
//
// Solidity: function step(bytes _stateData, bytes _proof, uint256 _localContext) returns(bytes32)
func (_MIPS *MIPSTransactor) Step(opts *bind.TransactOpts, _stateData []byte, _proof []byte, _localContext *big.Int) (*types.Transaction, error) {
	return _MIPS.contract.Transact(opts, "step", _stateData, _proof, _localContext)
}

// Step is a paid mutator transaction binding the contract method 0x836e7b32.
//
// Solidity: function step(bytes _stateData, bytes _proof, uint256 _localContext) returns(bytes32)
func (_MIPS *MIPSSession) Step(_stateData []byte, _proof []byte, _localContext *big.Int) (*types.Transaction, error) {
	return _MIPS.Contract.Step(&_MIPS.TransactOpts, _stateData, _proof, _localContext)
}

// Step is a paid mutator transaction binding the contract method 0x836e7b32.
//
// Solidity: function step(bytes _stateData, bytes _proof, uint256 _localContext) returns(bytes32)
func (_MIPS *MIPSTransactorSession) Step(_stateData []byte, _proof []byte, _localContext *big.Int) (*types.Transaction, error) {
	return _MIPS.Contract.Step(&_MIPS.TransactOpts, _stateData, _proof, _localContext)
}
