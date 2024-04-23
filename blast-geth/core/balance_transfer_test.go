package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
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

func TestClaimableBalanceIncreaseWithRevert(t *testing.T) {
	// 1) Setup the account to have some claimable yield
	contractRawAccount := uint64(1)
	recipientRawAccount := uint64(2)

	precompile := vm.PrecompiledContractsCancun[common.BytesToAddress([]byte{1, 0})] // Blast precompile
	db, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	setSharePrice(db, 1)
	addClaimableAccount(db, contractRawAccount)
	addBalance(db, contractRawAccount, 10)
	setSharePrice(db, 2)
	addBalance(db, contractRawAccount, 10)
	if !assertBalance(db, contractRawAccount, 20) {
		t.Fail()
	}
	expectedClaimable := uint64(10)
	if !assertClaimable(db, contractRawAccount, expectedClaimable) {
		t.Fail()
	}

	// 2) Simulate a claim transaction claiming the entire contract's claimable amount to recipient
	snapshot := db.Snapshot() // Take a snapshot of the current db state

	data := []byte{0x99, 0x6c, 0xba, 0x68} // claimSelector
	contract := make([]byte, 32)
	contract[31] = byte(contractRawAccount)
	recipient := make([]byte, 32)
	recipient[31] = byte(recipientRawAccount)
	amount := make([]byte, 32)
	amount[31] = byte(expectedClaimable)
	data = append(data, contract...)
	data = append(data, recipient...)
	data = append(data, amount...)
	cpy := make([]byte, len(data))
	copy(cpy, data)
	precompile.Run(params.BlastAccountConfigurationAddress, cpy, db, false)

	// No more claimable, but balance intact
	if !assertBalance(db, contractRawAccount, 20) {
		t.Fail()
	}
	if !assertClaimable(db, contractRawAccount, 0) {
		t.Fail()
	}
	if !assertBalance(db, recipientRawAccount, expectedClaimable) {
		t.Fail()
	}

	// 2a) Simulate another operation in this transaction that will revert, making the entire transaction revert
	_, err := precompile.Run(params.BlastAccountConfigurationAddress, cpy, db, true) // this is not doing anything, but just simulating a revert
	if err != nil {
		db.RevertToSnapshot(snapshot)
	}

	// 3) State should match the snapshot state, but it doesn't, both balance and claimable doesn't match.
	// Claimable fails, since the account has been changed to Automatic instead of Claimable because of the revert
	var forceFail bool
	if !assertClaimable(db, contractRawAccount, expectedClaimable) {
		actualClaimable := db.GetClaimableAmount(getAddr(contractRawAccount))
		t.Logf("Amount claimable not matching after revert: want: %d have: %d\n", expectedClaimable, actualClaimable)
		forceFail = true
	}

	// Balance fails as the share value includes the claimable yield which is kind of hidden, but it's not either
	if !assertBalance(db, contractRawAccount, 20) {
		actualBalance := getBalance(db, contractRawAccount)
		t.Logf("Balance not matching after revert: want: %d have: %d\n", 20, actualBalance)
		forceFail = true
	}

	if !assertBalance(db, recipientRawAccount, 0) {
		actualBalance := getBalance(db, recipientRawAccount)
		t.Logf("Recipient balance not matching after revert: want: %d have: %d\n", expectedClaimable, actualBalance)
		forceFail = true
	}

	if db.GetFlags(getAddr(contractRawAccount)) != types.YieldClaimable {
		actualType := db.GetFlags(getAddr(contractRawAccount))
		t.Logf("Contract type not matching after revert: want: %d have: %d\n", types.YieldClaimable, actualType)
		forceFail = true
	}

	if forceFail {
		// Add balance to confirm it adds properly
		addBalance(db, contractRawAccount, 1)
		if !assertBalance(db, contractRawAccount, 31) {
			actualBalance := getBalance(db, contractRawAccount)
			t.Logf("Balance(after revert) not matching after revert: want: %d have: %d\n", 31, actualBalance)
		}

		t.Fail()
	}
}
