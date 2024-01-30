package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var sharePriceSlot = common.BigToHash(common.Big1)

func setSharePrice(state *state.StateDB, price uint64) {
	priceInt := new(big.Int).SetUint64(price)
	priceHash := common.BigToHash(priceInt)
	state.SetState(params.BlastSharesAddress, sharePriceSlot, priceHash)
}

func defaults(state *state.StateDB) {
	setSharePrice(state, 1)
}

func addAccount(state *state.StateDB, a uint64, yield uint8) {
	addr := getAddr(a)
	state.CreateAccount(addr)
	state.SetFlags(addr, yield)
	state.SetBalance(addr, big.NewInt(0))
}

func setFlag(state *state.StateDB, addrRaw uint64, yield uint8) {
	addr := getAddr(addrRaw)
	state.SetFlags(addr, yield)
}

func addVoidAccount(state *state.StateDB, a uint64) {
	addAccount(state, a, types.YieldDisabled)
}

func addClaimableAccount(state *state.StateDB, a uint64) {
	addAccount(state, a, types.YieldClaimable)
}

func addAutomaticAccount(state *state.StateDB, a uint64) {
	addAccount(state, a, types.YieldAutomatic)
}

func getAddr(a uint64) common.Address {
	return common.BigToAddress(new(big.Int).SetUint64(a))
}

func addBalance(state *state.StateDB, addrRaw uint64, amount uint64) {
	addr := getAddr(addrRaw)
	state.AddBalance(addr, new(big.Int).SetUint64(amount))
}

func subBalance(state *state.StateDB, addrRaw uint64, amount uint64) {
	addr := getAddr(addrRaw)
	state.SubBalance(addr, new(big.Int).SetUint64(amount))
}

func getBalance(state *state.StateDB, addrRaw uint64) uint64 {
	addr := getAddr(addrRaw)
	bal := state.GetBalance(addr)
	return bal.Uint64()
}

func assertBalance(state *state.StateDB, addrRaw uint64, expectedBalance uint64) bool {
	bal := getBalance(state, addrRaw)
	return bal == expectedBalance
}

func assertClaimable(state *state.StateDB, addrRaw uint64, expectedClaimableAmount uint64) bool {
	addr := getAddr(addrRaw)
	am := state.GetClaimableAmount(addr)
	return am.Uint64() == expectedClaimableAmount
}

func TestBalanceTransferOnVoid(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addVoidAccount(state, 1)
	addBalance(state, 1, 10)
	if !assertBalance(state, 1, 10) {
		t.Fail()
	}
}

func TestBalanceTransferOnVoidTwo(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addVoidAccount(state, 1)
	addBalance(state, 1, 10)

	setSharePrice(state, 2)

	addBalance(state, 1, 10)
	subBalance(state, 1, 15)
	if !assertBalance(state, 1, 5) {
		t.Fail()
	}
}
func TestBalanceTransferOnAutomatic(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addAutomaticAccount(state, 1)
	addBalance(state, 1, 10)
	if !assertBalance(state, 1, 10) {
		t.Fail()
	}
}

func TestBalanceTransferOnAutomaticWithSub(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addAutomaticAccount(state, 1)
	addBalance(state, 1, 10)
	addBalance(state, 1, 10)
	subBalance(state, 1, 15)
	if !assertBalance(state, 1, 5) {
		t.Fail()
	}
}

func TestBalanceTransferOnClaimable(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addClaimableAccount(state, 1)
	addBalance(state, 1, 10)
	if !assertBalance(state, 1, 10) {
		t.Fail()
	}
}

func TestBalanceTransferOnClaimableWithSub(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addClaimableAccount(state, 1)
	addBalance(state, 1, 10)
	addBalance(state, 1, 10)
	subBalance(state, 1, 15)
	if !assertBalance(state, 1, 5) {
		t.Fail()
	}
}

func TestBalanceTransferOnSharePriceIncreaseForAutomaticCase(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	defaults(state)

	addAutomaticAccount(state, 1)
	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	if !assertBalance(state, 1, 20) {
		t.Fail()
	}
}

func TestAutomaticBalanceIncrease(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)

	addAutomaticAccount(state, 1)

	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	addBalance(state, 1, 10)

	if !assertBalance(state, 1, 30) {
		t.Fail()
	}
}

func TestClaimableBalanceIncrease(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)

	addClaimableAccount(state, 1)

	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	addBalance(state, 1, 10)

	if !assertBalance(state, 1, 20) {
		t.Fail()
	}
	if !assertClaimable(state, 1, 10) {
		t.Fail()
	}
}

func TestFlagsSwitch(t *testing.T) {
	// Test cases go here
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)

	addClaimableAccount(state, 1)

	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	addBalance(state, 1, 10)
	if !assertBalance(state, 1, 20) {
		t.Fail()
	}
	if !assertClaimable(state, 1, 10) {
		t.Fail()
	}

	setFlag(state, 1, types.YieldAutomatic)
	if !assertBalance(state, 1, 30) {
		t.Fail()
	}
	if !assertClaimable(state, 1, 0) {
		t.Fail()
	}
}
func TestFlagSwitchAutomaticToVoid(t *testing.T) {
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)
	addAccount(state, 1, types.YieldAutomatic)
	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	setFlag(state, 1, types.YieldDisabled)
	if !assertBalance(state, 1, 20) {
		t.Fail()
	}
}

func TestFlagSwitchAutomaticToClaimable(t *testing.T) {
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)
	addAccount(state, 1, types.YieldAutomatic)
	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	setFlag(state, 1, types.YieldClaimable)
	if !assertBalance(state, 1, 20) {
		t.Fail()
	}
}

func TestFlagSwitchVoidToAutomatic(t *testing.T) {
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)
	addAccount(state, 1, types.YieldDisabled)
	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	setFlag(state, 1, types.YieldAutomatic)
	if !assertBalance(state, 1, 10) {
		t.Fail()
	}
}

func TestFlagSwitchVoidToClaimable(t *testing.T) {
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)
	addAccount(state, 1, types.YieldDisabled)
	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	setFlag(state, 1, types.YieldClaimable)
	if !assertBalance(state, 1, 10) {
		t.Fail()
	}
	if !assertClaimable(state, 1, 0) {
		t.Fail()
	}
}

func TestFlagSwitchClaimableToAutomatic(t *testing.T) {
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)
	addAccount(state, 1, types.YieldClaimable)
	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	setFlag(state, 1, types.YieldAutomatic)
	if !assertBalance(state, 1, 20) {
		t.Fail()
	}
	if !assertClaimable(state, 1, 0) {
		t.Fail()
	}
}

func TestFlagSwitchClaimableToVoid(t *testing.T) {
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(state, 1)
	addAccount(state, 1, types.YieldClaimable)
	addBalance(state, 1, 10)
	setSharePrice(state, 2)
	setFlag(state, 1, types.YieldDisabled)
	if !assertBalance(state, 1, 20) {
		t.Fail()
	}
	if !assertClaimable(state, 1, 0) {
		t.Fail()
	}
}
