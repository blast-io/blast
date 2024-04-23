package runtime

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

func TestStateTransition(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db

	address := common.HexToAddress("0xfa")
	state.SetCode(address, []byte{
		byte(vm.PUSH1), 10,
		byte(vm.PUSH1), 0,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 32,
		byte(vm.PUSH1), 0,
		byte(vm.RETURN),
	})
	tp := TransactionParams{
		sender: common.HexToAddress("0x0b"),
		to:     &address,
		input:  []byte{},
	}
	result := stateTransition(&tp, state, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	num := new(big.Int).SetBytes(result.ReturnData)
	if num.Cmp(big.NewInt(10)) != 0 {
		t.Error("Expected 10, got", num)
	}
}

func TestGenesisSharePrice(t *testing.T) {
	cfg, env := setup(t)

	addr := common.HexToAddress("0x8")
	tracker := vm.NewGasTracker()

	abi := getAbi(sharesPath, t)
	input, err := abi.Pack("price")
	if err != nil {
		t.Fatal(err)
	}

	ret, _, err := env.Call(vm.AccountRef(addr), params.BlastSharesAddress, input, uint64(10_000), big.NewInt(0), tracker)
	if err != nil {
		t.Fatal(err)
	}
	unpack, err := abi.Unpack("price", ret)
	if err != nil {
		t.Fatal(err)
	}
	contractPrice := unpack[0].(*big.Int)
	if contractPrice.Cmp(big.NewInt(expectedInitialSharePrice)) != 0 {
		t.Fatalf("unexpected share price")
	}

	statePrice := getPriceFromState(cfg.State)
	if statePrice.Cmp(contractPrice) != 0 {
		t.Fatalf("unexpected share price. State price: %d, Contract price: %d", statePrice, contractPrice)
	}
}

func TestGenesisShareCount(t *testing.T) {
	_, env := setup(t)
	tracker := vm.NewGasTracker()

	abi := getAbi(sharesPath, t)
	input, err := abi.Pack("count")
	if err != nil {
		t.Fatal(err)
	}

	ret, _, err := env.Call(vm.AccountRef(getAddr(1)), params.BlastSharesAddress, input, uint64(10_000), big.NewInt(0), tracker)
	if err != nil {
		t.Fatal(err)
	}
	unpack, err := abi.Unpack("count", ret)
	if err != nil {
		t.Fatal(err)
	}
	if unpack[0].(*big.Int).Cmp(big.NewInt(0x0)) != 0 {
		t.Fatalf("unexpected share price")
	}
}

func TestShareCountIncreasing(t *testing.T) {
	cfg, env := setup(t)
	initialCount := getPublicVar("count", params.BlastSharesAddress, env, t).(*big.Int)
	if initialCount.Cmp(common.Big0) != 0 {
		t.Fatal("initial share count doesnt match")
	}
	sharesReporter := getPublicVar("REPORTER", params.BlastSharesAddress, env, t).(common.Address)
	t.Log(sharesReporter, l2Alias(sharesReporter))

	cfg.State.AddBalance(getAddr(0x132), big.NewInt(1e18))
	newCount := getPublicVar("count", params.BlastSharesAddress, env, t).(*big.Int)
	expectedNewCount := big.NewInt(1e18 / expectedInitialSharePrice)
	stateCount := getCountFromState(cfg.State)

	if newCount.Cmp(expectedNewCount) != 0 {
		t.Log(cfg.State.GetBalanceValues(getAddr(0x132)))
		t.Fatalf("new count not correct: got %d, expected %d, state %d\n", newCount, expectedNewCount, stateCount)
	}
}

func TestDistributeYield(t *testing.T) {
	cfg, env := setup(t)
	sharesReporter := getPublicVar("REPORTER", params.BlastSharesAddress, env, t).(common.Address)
	cfg.State.AddBalance(getAddr(0x321), big.NewInt(1e18))
	newCount := getPublicVar("count", params.BlastSharesAddress, env, t).(*big.Int)
	oldPrice := getPublicVar("price", params.BlastSharesAddress, env, t).(*big.Int)

	abi := getAbi(addrToAbiPath[params.BlastSharesAddress.String()], t)
	input, err := abi.Pack("addValue", new(big.Int).Mul(newCount, big.NewInt(10000)))
	if err != nil {
		t.Fatal(err)
	}

	oldPending := getPublicVar("pending", params.BlastSharesAddress, env, t).(*big.Int)
	cfg.State.AddBalance(sharesReporter, big.NewInt(100_000))
	gasTracker := vm.NewGasTracker()
	_, _, err = env.Call(vm.AccountRef(l2Alias(sharesReporter)), params.BlastSharesAddress, input, uint64(100_000), big.NewInt(0), gasTracker)
	if err != nil {
		t.Fatal(err)
	}
	newPrice := getPublicVar("price", params.BlastSharesAddress, env, t).(*big.Int)
	if newPrice.Cmp(oldPrice) <= 0 {
		t.Fatalf("price must increase after yield distribution")
	}
	pending := getPublicVar("pending", params.BlastSharesAddress, env, t).(*big.Int)
	t.Log(oldPrice, newPrice)
	t.Log(newCount)
	t.Log(oldPending, pending)
}

func TestCountPredeploySlot(t *testing.T) {
	cfg, env := setup(t)
	cfg.State.AddBalance(getAddr(0x322), big.NewInt(2e18))
	newCount := getPublicVar("count", params.BlastSharesAddress, env, t).(*big.Int)
	countFromState := cfg.State.GetState(params.BlastSharesAddress, shareCountSlot).Big()
	t.Log(newCount, countFromState)
	if newCount.Cmp(countFromState) != 0 {
		t.Fatalf("count slot not set correctly: got %d, expected %d\n", newCount.Uint64(), countFromState.Uint64())
	}
	balValues := cfg.State.GetBalanceValues(getAddr(0x322))
	if balValues.Shares.Cmp(newCount) != 0 {
		t.Fatalf("count slot not set correctly: got %d, expected %d\n", newCount.Uint64(), balValues.Shares.Uint64())
	}

	if newCount.Cmp(common.Big0) == 0 {
		t.Fatalf("newCount must be > 0")
	}
}
