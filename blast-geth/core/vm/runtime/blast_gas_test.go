package runtime

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

func TestGasConstants(t *testing.T) {
	_, env := setup(t)
	compareGasConfig("zeroClaimRate", env, t)
	compareGasConfig("baseClaimRate", env, t)
	compareGasConfig("ceilClaimRate", env, t)
	compareGasConfig("baseGasSeconds", env, t)
	compareGasConfig("ceilGasSeconds", env, t)
}

func TestGasUpdate(t *testing.T) {
	cfg, env := setup(t)
	abi := getAbi(addrToAbiPath[params.BlastGasAddress.String()], t)
	input, err := abi.Pack("updateAdminParameters", big.NewInt(1), big.NewInt(2), big.NewInt(2), big.NewInt(3), big.NewInt(3))
	if err != nil {
		t.Fatal(err)
	}
	admin := getPublicVar("admin", params.BlastGasAddress, env, t).(common.Address)
	cfg.State.AddBalance(admin, new(big.Int).SetUint64(100_000_000))

	gasTracker := vm.NewGasTracker()
	env = NewEnv(cfg)

	var rules = cfg.ChainConfig.Rules(env.Context.BlockNumber, env.Context.Random != nil, env.Context.Time)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, &params.BlastGasAddress, vm.ActivePrecompiles(rules), nil)
	_, _, err = env.Call(vm.AccountRef(admin), params.BlastGasAddress, input, uint64(900_000), big.NewInt(0), gasTracker)
	if err != nil {
		t.Fatal(err)
	}
	// Check if the updated variables actually got updated
	updatedTaxRateNum := getPublicVar("zeroClaimRate", params.BlastGasAddress, env, t).(*big.Int)
	if updatedTaxRateNum.Cmp(big.NewInt(1)) != 0 {
		t.Fatal("taxRateNum did not update correctly")
	}
	ceilClaimRate := getPublicVar("ceilClaimRate", params.BlastGasAddress, env, t).(*big.Int)
	if ceilClaimRate.Cmp(big.NewInt(3)) != 0 {
		t.Fatal("ceil claim rate did not update correctly")
	}
	baseGasSeconds := getPublicVar("baseGasSeconds", params.BlastGasAddress, env, t).(*big.Int)
	if baseGasSeconds.Cmp(big.NewInt(2)) != 0 {
		t.Fatal("baseGasSecodns did not update correctly")
	}
}

func TestReadGasVars(t *testing.T) {
	cfg, env := setup(t)
	abi := getAbi(addrToAbiPath[params.BlastGasAddress.String()], t)

	// Call the readGasVars function
	userAddr := getAddr(0xA1)
	input, err := abi.Pack("readGasParams", userAddr)
	if err != nil {
		t.Fatal(err)
	}

	gasTracker := vm.NewGasTracker()
	env = NewEnv(cfg)

	var rules = cfg.ChainConfig.Rules(env.Context.BlockNumber, env.Context.Random != nil, env.Context.Time)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, &params.BlastGasAddress, vm.ActivePrecompiles(rules), nil)
	res, _, err := env.Call(vm.AccountRef(userAddr), params.BlastGasAddress, input, uint64(900_000), big.NewInt(0), gasTracker)
	if err != nil {
		t.Fatal(err)
	}
	decodedRes, err := abi.Unpack("readGasParams", res)

	etherSeconds := decodedRes[0].(*big.Int)
	etherBalance := decodedRes[1].(*big.Int)
	lastUpdated := decodedRes[2].(*big.Int)
	mode := decodedRes[3].(uint8)
	if etherSeconds.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("etherSeconds is not zero")
	}
	if etherBalance.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("etherBalance is not zero")
	}
	if lastUpdated.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("lastUpdated is not zero")
	}
	if mode != 0 {
		t.Fatal("mode is not zero")
	}
}

func TestGasAccumulationNilCase(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db

	to := common.HexToAddress("0xfd")
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender: common.HexToAddress("0x0e"),
		to:     &to,
		input:  nil,
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	etherSeconds, etherBalance, lastUpdated, mode := readGasVars(state, to, t)
	if etherSeconds.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", lastUpdated)
	}
	if mode != 0 {
		t.Error("Expected 0, got", mode)
	}
}

func TestSetGasMode(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	addr := common.HexToAddress("0x0b")
	setGasMode(state, addr, t)

	etherSeconds, etherBalance, lastUpdated, mode := readGasVars(state, addr, t)
	if etherSeconds.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", lastUpdated)
	}
	if mode != 1 {
		t.Error("Expected 1, got", mode)
	}
}

func TestGasAccumulation(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xfd")
	setGasMode(state, to, t)
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender: common.HexToAddress("0x0b"),
		to:     &to,
		input:  nil,
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	etherSeconds, etherBalance, lastUpdated, mode := readGasVarsAtBlock(state, to, common.Big1, t)
	if etherSeconds.Cmp(big.NewInt(0)) <= 0 {
		t.Error("Expected etherSeconds to be greater than 0, got", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(0)) <= 0 {
		t.Error("Expected 0, got", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", lastUpdated)
	}
	if mode != 1 {
		t.Error("Expected 1, got", mode)
	}
}

func TestGasAccumulationEtherSeconds(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xfd")
	setGasMode(state, to, t)
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender:      common.HexToAddress("0x0b"),
		to:          &to,
		input:       nil,
		blockNumber: big.NewInt(0),
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}

	etherSeconds, etherBalance, lastUpdated, _ := readGasVarsAtBlock(state, to, big.NewInt(7*24*60*60), t)
	t.Log(result.UsedGas)
	t.Log(etherBalance, etherSeconds, lastUpdated)
	if etherSeconds.Cmp(new(big.Int).Mul(big.NewInt(36), (big.NewInt(7*24*60*60)))) != 0 {
		t.Error("Expected etherSeconds to be greater than 0, got", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Error("Expected 36, got", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", lastUpdated)
	}
}

func TestGasAccOverTime(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xff")
	setGasMode(state, to, t)
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})

	tp := TransactionParams{
		sender:      common.HexToAddress("0x0b"),
		to:          &to,
		input:       nil,
		blockNumber: big.NewInt(0),
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}

	etherSeconds, etherBalance, lastUpdated, _ := readGasVarsAtBlock(state, to, big.NewInt(7*24*60*60), t)
	if etherSeconds.Cmp(new(big.Int).Mul(big.NewInt(36), (big.NewInt(7*24*60*60)))) != 0 {
		t.Error("Expected etherSeconds to be greater than 0, got", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Error("Expected 36, got", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(0)) != 0 {
		t.Error("Expected 0, got", lastUpdated)
	}
}

func TestGasClaim(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, simulateContractPath, t)
	abi := getAbi(simulateContractPath, t)
	input, err := abi.Pack("readConfigurationSimulator", addr)
	if err != nil {
		t.Error(err)
	}
	tp := TransactionParams{
		sender: EOA_ADDR,
		to:     &addr,
		input:  input,
	}
	stateTransition(&tp, db, t)

	etherSeconds, etherBalance, _, _ := readGasVarsAtBlock(db, addr, big.NewInt(100), t)
	t.Log(etherSeconds, etherBalance)
	if etherSeconds.Cmp(common.Big0) == 0 {
		t.Fail()
	}

	if etherBalance.Cmp(common.Big0) == 0 {
		t.Fail()
	}

	bps, secs := readBps(db, etherBalance, etherSeconds, t)
	expGasToConsume := new(big.Int).Div(new(big.Int).Mul(bps, etherBalance), big.NewInt(10_000))

	rec := getAddr(0x32)
	gasClaimed := claimGas(db, addr, &rec, etherBalance, etherSeconds, big.NewInt(100), t)
	if gasClaimed.Cmp(expGasToConsume) != 0 {
		t.Log(gasClaimed, etherBalance, expGasToConsume)
		t.Fatalf("not all gas claimed")
	}
	newSec, newBal, _, _ := readGasVarsAtBlock(db, addr, big.NewInt(100), t)
	if secs.Cmp(new(big.Int).Sub(etherSeconds, newSec)) != 0 {
		t.Fatalf("ether seconds mismatch")
	}
	if newBal.Cmp(common.Big0) != 0 {
		t.Fatalf("all sec should be consumed")
	}
	if db.GetBalance(rec).Cmp(gasClaimed) != 0 {
		t.Fatalf("gas did not transfer to rec")
	}
}

func TestGasPackingMode(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xfd")
	setGasMode(state, to, t)
	etherSeconds, etherBalance, lastUpdated, mode := readGasVars(state, to, t)
	if etherSeconds.Cmp(common.Big0) != 0 {
		t.Fatalf("Expected etherSeconds to be 0, got %s", etherSeconds)
	}
	if etherBalance.Cmp(common.Big0) != 0 {
		t.Fatalf("Expected etherBalance to be 0, got %s", etherBalance)
	}
	if lastUpdated.Cmp(common.Big0) != 0 {
		t.Fatalf("Expected lastUpdated to be 0, got %s", lastUpdated)
	}
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}

}

func TestGasPackingBalance(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xfd")
	setGasMode(state, to, t)
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender:      common.HexToAddress("0x0b"),
		to:          &to,
		input:       nil,
		blockNumber: big.NewInt(0),
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	etherSeconds, etherBalance, lastUpdated, mode := readGasVars(state, to, t)
	if etherSeconds.Cmp(common.Big0) != 0 {
		t.Fatalf("Expected etherSeconds to be 0, got %s", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Fatalf("Expected etherBalance to be 36, got %s", etherBalance)
	}
	if lastUpdated.Cmp(common.Big0) != 0 {
		t.Fatalf("Expected lastUpdated to be 0, got %s", lastUpdated)
	}
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}

}

func TestGasPackingBalanceEtherSeconds(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xfd")
	setGasMode(state, to, t)
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender:      common.HexToAddress("0x0b"),
		to:          &to,
		input:       nil,
		blockNumber: big.NewInt(0),
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	etherSeconds, etherBalance, lastUpdated, mode := readGasVarsAtBlock(state, to, big.NewInt(100), t)
	if etherSeconds.Cmp(big.NewInt(3600)) != 0 {
		t.Fatalf("Expected etherSeconds to be 0, got %s", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Fatalf("Expected etherBalance to be 36, got %s", etherBalance)
	}
	if lastUpdated.Cmp(common.Big0) != 0 {
		t.Fatalf("Expected lastUpdated to be 0, got %s", lastUpdated)
	}
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}

}

func TestGasPackingBalanceWithUpdate(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xfd")
	setGasMode(state, to, t)
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender:      common.HexToAddress("0x0b"),
		to:          &to,
		input:       nil,
		blockNumber: big.NewInt(50),
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	etherSeconds, etherBalance, lastUpdated, mode := readGasVarsAtBlock(state, to, big.NewInt(100), t)
	if etherSeconds.Cmp(big.NewInt(1800)) != 0 {
		t.Fatalf("Expected etherSeconds to be 0, got %s", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Fatalf("Expected etherBalance to be 36, got %s", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(50)) != 0 {
		t.Fatalf("Expected lastUpdated to be 50, got %s", lastUpdated)
	}
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}

	setGasModeAtBlock(state, to, big.NewInt(75), t)
	etherSeconds, etherBalance, lastUpdated, mode = readGasVarsAtBlock(state, to, big.NewInt(100), t)
	if etherSeconds.Cmp(big.NewInt(1800)) != 0 {
		t.Fatalf("Expected etherSeconds to be 3600, got %s", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Fatalf("Expected etherBalance to be 36, got %s", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(75)) != 0 {
		t.Fatalf("Expected lastUpdated to be 50, got %s", lastUpdated)
	}
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}

}

func TestGasPackingBalanceWith100YearsBuffer(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	to := common.HexToAddress("0xfd")
	setGasMode(state, to, t)
	state.SetCode(to, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender:      common.HexToAddress("0x0b"),
		to:          &to,
		input:       nil,
		blockNumber: big.NewInt(0),
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	etherSeconds, etherBalance, lastUpdated, mode := readGasVarsAtBlock(state, to, big.NewInt(100), t)
	if etherSeconds.Cmp(big.NewInt(3600)) != 0 {
		t.Fatalf("Expected etherSeconds to be 0, got %s", etherSeconds)
	}
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Fatalf("Expected etherBalance to be 36, got %s", etherBalance)
	}
	if lastUpdated.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected lastUpdated to be 50, got %s", lastUpdated)
	}
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}

	var years int64 = 100 * 365 * 24 * 60 * 60
	yearsInt := big.NewInt(years)
	t.Log("#bytes", len(yearsInt.Bytes()))
	setGasModeAtBlock(state, to, yearsInt, t)
	etherSeconds, etherBalance, lastUpdated, mode = readGasVarsAtBlock(state, to, yearsInt, t)
	if etherBalance.Cmp(big.NewInt(36)) != 0 {
		t.Fatalf("Expected etherBalance to be 36, got %s", etherBalance)
	}
	if lastUpdated.Cmp(yearsInt) != 0 {
		t.Fatalf("Expected lastUpdated to be %s, got %s", yearsInt, lastUpdated)
	}
	if etherSeconds.Cmp(new(big.Int).Mul(yearsInt, big.NewInt(36))) != 0 {
		exp := new(big.Int).Mul(yearsInt, big.NewInt(36))
		t.Fatalf("Expected etherSeconds to be %s, got %s", exp, etherSeconds)
	}
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}

}

func TestGasCode(t *testing.T) {
	cfg, _ := setup(t) // -> puts env vars in db
	state := cfg.State
	codeHash := state.GetCodeHash(params.BlastGasAddress)
	code := state.GetCode(params.BlastGasAddress)
	misc.EnsureUpdateGas(cfg.ChainConfig, params.BlastTestnetPredeployForkTime, state)
	afterCodeHash := state.GetCodeHash(params.BlastGasAddress)
	afterCode := state.GetCode(params.BlastGasAddress)
	t.Log(codeHash, code, afterCodeHash, afterCode)

}
