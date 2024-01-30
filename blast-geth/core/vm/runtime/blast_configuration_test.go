package runtime

import (
	"math/big"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var selfConfigureContractPath = filepath.Join(basePath, "./blast_contract_tests/self_configure_contract/SelfConfigureContract.json")
var interleaveContractPath = filepath.Join(basePath, "./blast_contract_tests/interleave_contract/InterleaveContract.json")
var simulateContractPath = filepath.Join(basePath, "./blast_contract_tests/simulate_contract_test/SimulateContract.json")
var recursiveContractPath = filepath.Join(basePath, "./blast_contract_tests/recursive_contract_tests/D1.json")

func TestDeployNewContract(t *testing.T) {
	state := setupDb(t) // -> puts env vars in db
	deployTestContract(state, selfConfigureContractPath, t)
}

func TestParams(t *testing.T) {
	db := setupDb(t) // -> puts env vars in db
	addr := deployTestContract(db, selfConfigureContractPath, t)
	_, _, _, mode := readGasVars(db, addr, t)
	if mode != 1 {
		t.Error("mode not set")
	}
	if db.GetFlags(addr) != types.YieldClaimable {
		t.Error("yield not set correctly")
	}
}

func TestInvalidContractCall(t *testing.T) {
	db := setupDb(t) // -> puts env vars in db
	addr := deployTestContract(db, selfConfigureContractPath, t)
	// gov := common.HexToAddress("0x0000000000000000000000000000000000000002")
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("configureClaimableYield")
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender: addr,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&tp, db, t)
	if result.Err == nil {
		t.Error("result should have reverted")
	}
}

func TestValidContractCall(t *testing.T) {
	db := setupDb(t) // -> puts env vars in db
	addr := deployTestContract(db, selfConfigureContractPath, t)
	gov := common.HexToAddress("0x0000000000000000000000000000000000000002")
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("configureAutomaticYieldOnBehalf", addr)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender: gov,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&tp, db, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	mode := readYieldMode(db, addr, t)
	if mode != types.YieldAutomatic {
		t.Fatal("result must be of type automatic")
	}
}
func TestConfigureContract(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	gov := common.HexToAddress("0x0000000000000000000000000000000000000002")
	newGov := common.HexToAddress("0x0000000000000000000000000000000000000003")
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("configureContract", addr, uint8(types.YieldAutomatic), uint8(0), newGov)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender: gov,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&tp, db, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	mode := readYieldMode(db, addr, t)
	if mode != types.YieldAutomatic {
		t.Fatal("result must be of type automatic")
	}
	gasMode := readGasMode(db, addr, t)
	if gasMode != 0 {
		t.Fatal("result must be of type void")
	}
	governor := readGovernor(db, addr, t)
	if governor != newGov {
		t.Fatal("result must be the same as the governor address")
	}
}

func TestConfigureOnNilCase(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	gov := readGovernor(db, addr, t)
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("configureContract", addr, uint8(types.YieldAutomatic), uint8(0), addr)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender: gov,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&tp, db, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	input, err = abi.Pack("configureVoidYield")
	if err != nil {
		t.Fatal(err)
	}

	tp = TransactionParams{
		sender: addr,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}

	result = stateTransition(&tp, db, t)

	mode := readYieldMode(db, addr, t)
	if mode != types.YieldDisabled {
		t.Fatal("result must be of type disabled")
	}
}

func TestReadClaimableYield(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("readClaimableYield", addr)
	if err != nil {
		t.Fatal(err)
	}
	gov := readGovernor(db, addr, t)
	tp := TransactionParams{
		sender: gov,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&tp, db, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	claimableYield := new(big.Int).SetBytes(result.ReturnData)
	if claimableYield.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("Claimable yield must be zero initially")
	}
}

func TestGetYieldConfiguration(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("readYieldConfiguration", addr)
	if err != nil {
		t.Fatal(err)
	}
	gov := readGovernor(db, addr, t)
	tp := TransactionParams{
		sender: gov,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&tp, db, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	configuration := new(big.Int).SetBytes(result.ReturnData)
	if configuration.Cmp(big.NewInt(2)) != 0 {
		t.Fatal("configuration must be automatic")
	}
}

func TestReadClaimableYieldWithYield(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	t.Log("addr", addr)
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("readClaimableYield", addr)
	if err != nil {
		t.Fatal(err)
	}
	gov := readGovernor(db, addr, t)
	readClaimableTp := TransactionParams{
		sender: gov,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&readClaimableTp, db, t)
	if result.Err != nil {
		t.Fatal(result.Err)
	}
	claimableYield := new(big.Int).SetBytes(result.ReturnData)
	if claimableYield.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("Claimable yield must be zero initially")
	}
	t.Log(claimableYield)

	price := getVar("price", params.BlastSharesAddress, db, t).(*big.Int)
	db.AddBalance(addr, new(big.Int).Mul(price, big.NewInt(10))) // 10 shares
	if db.GetBalanceValues(addr).Shares.Cmp(big.NewInt(10)) != 0 {
		t.Log("shares", db.GetBalanceValues(addr).Shares)
		t.Log(db.GetBalanceValues(addr).Remainder)
		t.Log(db.GetBalanceValues(addr).Fixed)
		t.Log(db.GetBalanceValues(addr).Flags)
		t.Error("expected 10 shares")
	}
	count := getVar("count", params.BlastSharesAddress, db, t).(*big.Int)
	distributeYield(db, count, t)
	newClaimableYield := readClaimableYield(db, addr, t)
	if newClaimableYield.Cmp(big.NewInt(10)) != 0 {
		t.Error("claimable yield did not update")
	}
}

func TestClaimYield(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	price := getVar("price", params.BlastSharesAddress, db, t).(*big.Int)
	db.AddBalance(addr, new(big.Int).Mul(price, big.NewInt(10))) // 10 shares
	if db.GetBalanceValues(addr).Shares.Cmp(big.NewInt(10)) != 0 {
		t.Error("expected 10 shares")
	}

	count := getVar("count", params.BlastSharesAddress, db, t).(*big.Int)
	distributeYield(db, count, t)
	newClaimableYield := readClaimableYield(db, addr, t)
	if newClaimableYield.Cmp(big.NewInt(10)) != 0 {
		t.Error("claimable yield did not update")
	}

	rec := getAddr(5)
	claimedYield := claimYield(db, addr, t, big.NewInt(5), &rec)
	if claimedYield.Cmp(big.NewInt(5)) != 0 {
		t.Error("expected yield to be 5")
	}
	if db.GetBalance(rec).Cmp(big.NewInt(5)) != 0 {
		t.Error("yield did not transfer to correct account")
	}
	if db.GetClaimableAmount(addr).Cmp(big.NewInt(5)) != 0 {
		t.Error("all yield must be claimed")
	}

}

// Tests that claiming all yield at max does not revert
func TestClaimYieldMax(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	price := getVar("price", params.BlastSharesAddress, db, t).(*big.Int)
	db.AddBalance(addr, new(big.Int).Mul(price, big.NewInt(10))) // 10 shares
	if db.GetBalanceValues(addr).Shares.Cmp(big.NewInt(10)) != 0 {
		t.Error("expected 10 shares")
	}

	count := getVar("count", params.BlastSharesAddress, db, t).(*big.Int)
	distributeYield(db, count, t)
	newClaimableYield := readClaimableYield(db, addr, t)
	if newClaimableYield.Cmp(big.NewInt(10)) != 0 {
		t.Error("claimable yield did not update")
	}

	rec := getAddr(5)
	claimedYield := claimYield(db, addr, t, big.NewInt(10), &rec)
	if claimedYield.Cmp(big.NewInt(10)) != 0 {
		t.Error("expected yield to be 10")
	}
	if db.GetBalance(rec).Cmp(big.NewInt(10)) != 0 {
		t.Error("yield did not transfer to correct account")
	}
	if db.GetClaimableAmount(addr).Cmp(big.NewInt(0)) != 0 {
		t.Error("all yield must be claimed")
	}
}

func TestClaimYieldRevert(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	price := getVar("price", params.BlastSharesAddress, db, t).(*big.Int)
	db.AddBalance(addr, new(big.Int).Mul(price, big.NewInt(10))) // 10 shares
	if db.GetBalanceValues(addr).Shares.Cmp(big.NewInt(10)) != 0 {
		t.Error("expected 10 shares")
	}

	count := getVar("count", params.BlastSharesAddress, db, t).(*big.Int)
	distributeYield(db, count, t)
	newClaimableYield := readClaimableYield(db, addr, t)
	if newClaimableYield.Cmp(big.NewInt(10)) != 0 {
		t.Error("claimable yield did not update")
	}

	rec := getAddr(5)
	_, err := claimYieldWithErr(db, addr, t, big.NewInt(10000), &rec)
	if err == nil {
		t.Fatal("expected execution revert")
	}

}

func TestClaimYieldWrongUser(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	price := getVar("price", params.BlastSharesAddress, db, t).(*big.Int)
	db.AddBalance(addr, new(big.Int).Mul(price, big.NewInt(10))) // 10 shares
	count := getVar("count", params.BlastSharesAddress, db, t).(*big.Int)
	distributeYield(db, count, t)
	newClaimableYield := readClaimableYield(db, addr, t)
	if newClaimableYield.Cmp(big.NewInt(10)) != 0 {
		t.Error("claimable yield did not update")
	}

	rec := getAddr(5)
	claimedYield := claimYield(db, addr, t, big.NewInt(5), &rec)
	if claimedYield.Cmp(big.NewInt(5)) != 0 {
		t.Error("expected yield to be 5")
	}
	if db.GetBalance(rec).Cmp(big.NewInt(5)) != 0 {
		t.Error("yield did not transfer to correct account")
	}

	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("claimAllYield", addr, rec)
	if err != nil {
		t.Fatal(err)
	}

	tp := TransactionParams{
		sender: addr,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := stateTransition(&tp, db, t)
	if result.Err == nil {
		t.Fatal("transaction should revert")
	}

}

func TestClaimAllYield(t *testing.T) {
	db := setupDb(t)
	addr := deployTestContract(db, selfConfigureContractPath, t)
	price := getVar("price", params.BlastSharesAddress, db, t).(*big.Int)
	db.AddBalance(addr, new(big.Int).Mul(price, big.NewInt(10))) // 10 shares
	if db.GetBalanceValues(addr).Shares.Cmp(big.NewInt(10)) != 0 {
		t.Error("expected 10 shares")
	}

	count := getVar("count", params.BlastSharesAddress, db, t).(*big.Int)
	distributeYield(db, count, t)
	newClaimableYield := readClaimableYield(db, addr, t)
	if newClaimableYield.Cmp(big.NewInt(10)) != 0 {
		t.Error("claimable yield did not update")
	}

	rec := getAddr(5)
	claimedYield := claimAllYield(db, addr, t, &rec)
	expYield := big.NewInt(10)
	if claimedYield.Cmp(expYield) != 0 {
		t.Error("expected yield to be 10")
	}
	if db.GetBalance(rec).Cmp(expYield) != 0 {
		t.Error("yield did not transfer to correct account")
	}
	if db.GetClaimableAmount(addr).Cmp(common.Big0) != 0 {
		t.Error("all yield must be claimed")
	}
}
