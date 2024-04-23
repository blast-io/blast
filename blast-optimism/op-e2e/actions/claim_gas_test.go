package actions

import (
	"bytes"
	"encoding/json"
	"math/big"
	"os"
	"os/exec"
	"testing"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func loadContract(t *testing.T) ([]byte, abi.ABI, abi.ABI) {
	cwd, _ := os.Getwd()
	t.Logf("Current working directory %v", cwd)
	cmd := exec.Command(
		"solc", "--combined-json", "bin,abi", "--no-cbor-metadata",
		"gas_tester.sol",
	)
	output := bytes.Buffer{}
	oops := bytes.Buffer{}
	cmd.Stdout = &output
	cmd.Stderr = &oops
	if err := cmd.Run(); err != nil {
		t.Fatalf("solc died %v because %v", err, oops.String())
	}
	type loadIt struct {
		Contracts struct {
			Wrapper struct {
				ABI      json.RawMessage `json:"abi"`
				Bytecode string          `json:"bin"`
			} `json:"gas_tester.sol:Worker"`
			IBlast struct {
				ABI json.RawMessage `json:"abi"`
			} `json:"gas_tester.sol:IBlast"`
		} `json:"contracts"`
	}
	var l loadIt
	// t.Logf("output of solc %v", string(output.Bytes()))
	if err := json.Unmarshal(output.Bytes(), &l); err != nil {
		t.Fatalf("couldn't unmarshal %v", err)
	}
	//	t.Logf("load it %v", l)
	if len(l.Contracts.Wrapper.Bytecode) == 0 {
		t.Fatalf("contract bytecode non existent")
	}
	abiHandle, err := abi.JSON(bytes.NewReader(l.Contracts.Wrapper.ABI))
	require.Nil(t, err, "couldn't load up abi %v", err)
	iblastABI, err := abi.JSON(bytes.NewReader(l.Contracts.IBlast.ABI))

	t.Logf("worker abi looks like %v", abiHandle)
	t.Logf("iblast abi looks like %v", iblastABI)

	return common.Hex2Bytes(l.Contracts.Wrapper.Bytecode), abiHandle, iblastABI
}

func TestGasTrackerClaim(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	sd, dp, miner, sequencer, seqEngine, _, _ := setupEIP4844Test(t, log)
	contractByteCode, contractABI, iBlastABI := loadContract(gt)

	cnstr, err := contractABI.Constructor.Inputs.Pack()
	require.Nil(t, err, "constructor cant die")
	payload := append(contractByteCode, cnstr...)
	l2Client := seqEngine.EthClient()
	sayBlock := func(ctx string) {
		blkNum, err := l2Client.BlockNumber(t.Ctx())
		require.NoError(t, err, "block num cant die")
		t.Logf("%v block number is %v", ctx, blkNum)
	}

	sequencer.ActL2PipelineFull(t)
	miner.ActEmptyBlock(t)

	amt, err := l2Client.BalanceAt(t.Ctx(), dp.Addresses.Alice, nil)
	require.Nil(t, err, "l2client cant die")
	t.Logf("amt alice %v", amt)

	signer := types.LatestSigner(sd.L2Cfg.Config)
	n, err := l2Client.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
	require.NoError(t, err, "couldnt get pending nonce")
	newContractAddr := crypto.CreateAddress(dp.Addresses.Alice, n)

	tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
		ChainID:   sd.L2Cfg.Config.ChainID,
		Nonce:     n,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
		Gas:       2_000_000,
		To:        nil,
		Value:     common.Big0,
		Data:      payload,
	})

	sayBlock("before contract creation ")
	require.NoError(gt, l2Client.SendTransaction(t.Ctx(), tx))
	sequencer.ActL2StartBlock(t)
	seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
	sequencer.ActL2EndBlock(t)
	sayBlock("cranked the l2 chain ")

	runtimeCode, err := l2Client.CodeAt(t.Ctx(), newContractAddr, nil)
	require.NoError(t, err, "missing runtime contract code after deployment")
	require.Greater(t, len(runtimeCode), 0, "Zero length contract code")

	readGas := func() (*big.Int, *big.Int, *big.Int, uint8) {
		readGasEncoded, err := iBlastABI.Pack("readGasParams", newContractAddr)
		require.NoError(t, err, "not possible couldnt pack read gas params")
		readGasResult, err := l2Client.CallContract(t.Ctx(), ethereum.CallMsg{
			To:   &predeploys.BlastAddr,
			Data: readGasEncoded,
		}, nil)
		require.NoError(t, err, "call contract died")
		unpacked, err := iBlastABI.Unpack("readGasParams", readGasResult)
		require.NoError(t, err, "unpack died")
		etherSeconds, etherBalance, lastUpdated, mode :=
			unpacked[0].(*big.Int), unpacked[1].(*big.Int), unpacked[2].(*big.Int), unpacked[3].(uint8)
		return etherSeconds, etherBalance, lastUpdated, mode
	}

	wasteGas := func() {
		wasteEncoded, err := contractABI.Pack("burn_gas")
		require.NoError(t, err, "couldnt pack burn_gas")
		n, err := l2Client.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
		require.NoError(t, err, "couldnt get pending nonce")
		tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
			ChainID:   sd.L2Cfg.Config.ChainID,
			Nonce:     n,
			GasTipCap: big.NewInt(2 * params.GWei),
			GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
			Gas:       50_000,
			To:        &newContractAddr,
			Value:     common.Big0,
			Data:      wasteEncoded,
		})
		require.NoError(gt, l2Client.SendTransaction(t.Ctx(), tx))
		sequencer.ActL2StartBlock(t)
		seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
		sequencer.ActL2EndBlock(t)
		// rcpt, err := wait.ForReceiptOK(t.Ctx(), l2Client, tx.Hash())
		// rcpt, err := l2Client.TransactionReceipt(t.Ctx(), tx.Hash())
		// require.NoError(t, err, "couldnt get receipt")
		// return rcpt
	}

	sayBlock("About to waste Gas call")
	wasteGas()
	sayBlock("did waste gas")
	etherSeconds, etherBalance, lastUpdated, mode := readGas()
	// t.Logf("What values %v %v %v %v", etherSeconds, etherBalance, lastUpdated, mode)
	etherSecondsShouldBe, etherBalanceShouldBe, modeExpected :=
		big.NewInt(85306000042653), big.NewInt(143064000071532), uint8(1)

	currentBlock, err := l2Client.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)

	lastUpdatedExpected := big.NewInt(int64(currentBlock.Time()))

	if etherSeconds.Cmp(etherSecondsShouldBe) != 0 {
		t.Fatalf("incorrect etherseconds %v %v", etherSeconds, etherSecondsShouldBe)
	}

	if etherBalance.Cmp(etherBalanceShouldBe) != 0 {
		t.Fatalf("incorrect etherBalance %v %v", etherBalance, etherBalanceShouldBe)
	}

	if lastUpdated.Cmp(lastUpdatedExpected) != 0 {
		t.Fatalf("incorrect last timestamp updated %v %v", lastUpdated, lastUpdatedExpected)
	}

	if mode != modeExpected {
		t.Fatalf("incorrect mode %v %v", mode, modeExpected)
	}
}
