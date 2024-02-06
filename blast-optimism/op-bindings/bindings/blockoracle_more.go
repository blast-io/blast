// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const BlockOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/BlockOracle.sol:BlockOracle\",\"label\":\"blocks\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_uint256,t_struct(BlockInfo)1001_storage)\"}],\"types\":{\"t_mapping(t_uint256,t_struct(BlockInfo)1001_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct BlockOracle.BlockInfo)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(BlockInfo)1001_storage\"},\"t_struct(BlockInfo)1001_storage\":{\"encoding\":\"inplace\",\"label\":\"struct BlockOracle.BlockInfo\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Hash)1002\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Timestamp)1003\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var BlockOracleStorageLayout = new(solc.StorageLayout)

var BlockOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100415760003560e01c806354fd4d501461004657806399d548aa14610080578063c2c4c5c1146100b8575b600080fd5b61006a60405180604001604052806005815260200164302e302e3160d81b81525081565b60405161007791906101c8565b60405180910390f35b61009361008e36600461021d565b6100ce565b604080518251815260209283015167ffffffffffffffff169281019290925201610077565b6100c0610134565b604051908152602001610077565b604080518082018252600080825260209182018190528381528082528281208351808501909452805480855260019091015467ffffffffffffffff16928401929092520361012f576040516337cf270560e01b815260040160405180910390fd5b919050565b6000610141600143610236565b60408051808201825282408082524267ffffffffffffffff818116602080860182815260008981529182905287822096518755516001909601805467ffffffffffffffff1916969093169590951790915593519495509093909291849186917fb67ff58b33060fd371a35ae2d9f1c3cdaec9b8197969f6efe2594a1ff4ba68c691a4505090565b600060208083528351808285015260005b818110156101f5578581018301518582016040015282016101d9565b81811115610207576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561022f57600080fd5b5035919050565b60008282101561025657634e487b7160e01b600052601160045260246000fd5b50039056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(BlockOracleStorageLayoutJSON), BlockOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["BlockOracle"] = BlockOracleStorageLayout
	deployedBytecodes["BlockOracle"] = BlockOracleDeployedBin
}
