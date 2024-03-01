// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1BlockStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"number\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1001,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"timestamp\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1002,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"basefee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"hash\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1004,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"sequenceNumber\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint64\"},{\"astId\":1005,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bytes32\"},{\"astId\":1006,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeOverhead\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeScalar\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_uint256\"}],\"types\":{\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"}}}"

var L1BlockStorageLayout = new(solc.StorageLayout)

var L1BlockDeployedBin = "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c80638381f58a116100715780638381f58a146101465780638b239f731461015a5780639e8c496614610163578063b80777ea1461016c578063e591b28214610187578063e81b2c6d146101ba57600080fd5b8063015d8eb9146100ae57806309bd5a60146100c357806354fd4d50146100df5780635cf249691461011057806364ca23ef14610119575b600080fd5b6100c16100bc3660046102da565b6101c3565b005b6100cc60025481565b6040519081526020015b60405180910390f35b610103604051806040016040528060058152602001640312e312e360dc1b81525081565b6040516100d6919061034c565b6100cc60015481565b60035461012d9067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100d6565b60005461012d9067ffffffffffffffff1681565b6100cc60055481565b6100cc60065481565b60005461012d90600160401b900467ffffffffffffffff1681565b6101a273deaddeaddeaddeaddeaddeaddeaddeaddead000181565b6040516001600160a01b0390911681526020016100d6565b6100cc60045481565b3373deaddeaddeaddeaddeaddeaddeaddeaddead0001146102505760405162461bcd60e51b815260206004820152603b60248201527f4c31426c6f636b3a206f6e6c7920746865206465706f7369746f72206163636f60448201527f756e742063616e20736574204c3120626c6f636b2076616c7565730000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff988916600160401b026fffffffffffffffffffffffffffffffff199091169989169990991798909817909755600194909455600292909255600380549190941667ffffffffffffffff199190911617909255600491909155600555600655565b803567ffffffffffffffff811681146102d557600080fd5b919050565b600080600080600080600080610100898b0312156102f757600080fd5b610300896102bd565b975061030e60208a016102bd565b9650604089013595506060890135945061032a60808a016102bd565b979a969950949793969560a0850135955060c08501359460e001359350915050565b600060208083528351808285015260005b818110156103795785810183015185820160400152820161035d565b8181111561038b576000604083870101525b50601f01601f191692909201604001939250505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1BlockStorageLayoutJSON), L1BlockStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1Block"] = L1BlockStorageLayout
	deployedBytecodes["L1Block"] = L1BlockDeployedBin
}
