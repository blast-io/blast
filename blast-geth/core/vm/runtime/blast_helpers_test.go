package runtime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/params"
)

var EOA_ADDR = common.HexToAddress("0xB5Fbaf79959606E6eb265A9c466010487B30AAad")

// Assumes optimism and op-geth are in same directory
// AND optimism is not fully cleaned (i.e. .devnet folder, forge-artifacts has been generated)
var _, b, _, _ = runtime.Caller(0)
var basePath = filepath.Dir(b)
var BASE_OPTIMISM_PATH = filepath.Join(basePath, "../../../../optimism")

var BlastPath = BASE_OPTIMISM_PATH + "/packages/contracts-bedrock/forge-artifacts/Blast.sol/Blast.json"
var gasPath = BASE_OPTIMISM_PATH + "/packages/contracts-bedrock/forge-artifacts/Gas.sol/Gas.json"
var sharesPath = BASE_OPTIMISM_PATH + "/packages/contracts-bedrock/forge-artifacts/Shares.sol/Shares.json"
var devnetPath = BASE_OPTIMISM_PATH + "/packages/contracts-bedrock/deploy-config/devnetL1-template.json"

var addrToAbiPath = map[string]string{
	params.BlastSharesAddress.String():               sharesPath,
	params.BlastAccountConfigurationAddress.String(): BlastPath,
	params.BlastGasAddress.String():                  gasPath,
}

var defaultAddr = common.HexToAddress("0x1")
var regolithZeroTime = uint64(0)

var expectedInitialSharePrice = int64(1e9)
var shareCountSlot = common.BigToHash(big.NewInt(51))
var sharePriceSlot = common.BigToHash(common.Big1)

func getCountFromState(db *state.StateDB) *big.Int {
	countFromState := db.GetState(params.BlastSharesAddress, shareCountSlot).Big()
	return countFromState
}

func getPriceFromState(db *state.StateDB) *big.Int {
	price := db.GetState(params.BlastSharesAddress, sharePriceSlot).Big()
	return price
}

func l2Alias(a common.Address) common.Address {
	offset, _ := new(big.Int).SetString("0x1111000000000000000000000000000000001111", 0)
	newBigAddr := a.Big().Add(a.Big(), offset)
	return common.BigToAddress(newBigAddr)
}

func readJson(filepath string, t *testing.T) map[string]interface{} {
	testBytecodeRaw, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read file at path: %s, error: %v", filepath, err)
	}
	var testBytecodeJson map[string]interface{}
	err = json.Unmarshal(testBytecodeRaw, &testBytecodeJson)
	if err != nil {
		t.Fatal(err)
	}
	return testBytecodeJson
}

func getAbi(filepath string, t *testing.T) abi.ABI {
	testBytecodeJson := readJson(filepath, t)
	abiRaw, ok := testBytecodeJson["abi"]
	if !ok {
		t.Fatal("unable to find abi")
	}
	jsonData, err := json.Marshal(abiRaw)
	abiStr := string(jsonData)
	abi, err := abi.JSON(strings.NewReader(abiStr))

	if err != nil {
		t.Fatal(err)
	}
	return abi
}

func getDevnetVar(name string, t *testing.T) string {
	json := readJson(devnetPath, t)
	value, ok := json[name]
	if !ok {
		t.Fatal("unable to find var")
	}
	return value.(string)
}

func deployBlast(statedb *state.StateDB, t *testing.T) {
	BlastPath := BASE_OPTIMISM_PATH + "/packages/contracts-bedrock/forge-artifacts/Blast.sol/Blast.json"
	bytecode := getBytecodeFromForgeArtifacts(BlastPath, t)
	statedb.CreateAccount(params.BlastAccountConfigurationAddress)
	statedb.SetCode(params.BlastAccountConfigurationAddress, bytecode)
}

func deployGasPredeploy(statedb *state.StateDB, t *testing.T) {
	gasPath := BASE_OPTIMISM_PATH + "/packages/contracts-bedrock/forge-artifacts/Gas.sol/Gas.json"
	bytecode := getBytecodeFromForgeArtifacts(gasPath, t)
	statedb.CreateAccount(params.BlastGasAddress)
	statedb.SetCode(params.BlastGasAddress, bytecode)
}

func deploySharesPredeploy(statedb *state.StateDB, t *testing.T) {
	sharesPath := BASE_OPTIMISM_PATH + "/packages/contracts-bedrock/forge-artifacts/Shares.sol/Shares.json"
	bytecode := getBytecodeFromForgeArtifacts(sharesPath, t)
	statedb.CreateAccount(params.BlastSharesAddress)
	statedb.SetCode(params.BlastSharesAddress, bytecode)
}

func deployRawPredeploys(t *testing.T) *state.StateDB {
	db, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	deploySharesPredeploy(db, t)
	deployGasPredeploy(db, t)
	deployBlast(db, t)
	return db
}

func getBytecodeFromForStandardContract(filepath string, t *testing.T) []byte {
	testBytecodeJson := readJson(filepath, t)
	bytecodeObject, ok := testBytecodeJson["data"].(map[string]interface{})["bytecode"].(map[string]interface{})["object"]
	if !ok {
		t.Fatal("unable to parse bytecode")
	}
	if _, ok := bytecodeObject.(string); !ok {
		t.Fatal("bytecodeObject is not a string")
	}
	return common.Hex2Bytes(bytecodeObject.(string))
}

func getBytecodeFromForgeArtifacts(filepath string, t *testing.T) []byte {
	testBytecodeJson := readJson(filepath, t)
	bytecodeObject, ok := testBytecodeJson["deployedBytecode"].(map[string]interface{})["object"]
	if !ok {
		t.Fatal("unable to parse bytecode")
	}
	if _, ok := bytecodeObject.(string); !ok {
		t.Fatal("bytecodeObject is not a string")
	}
	return common.Hex2Bytes(bytecodeObject.(string))
}

func setStateGeneral(allocJson map[string]interface{}, addrRaw string, db *state.StateDB, t *testing.T) {
	addr := common.HexToAddress(addrRaw)
	db.CreateAccount(addr)
	// set code
	contractParams, ok := allocJson[addrRaw]
	if !ok {
		t.Fatal("unable to find contract")
	}

	// sets code
	code, ok := contractParams.(map[string]interface{})["code"]
	if ok {
		if _, ok := code.(string); !ok {
			t.Fatal("code is not a string")
		}
		codeStr := code.(string)
		codeBytes := common.Hex2Bytes(codeStr[2:])
		db.SetCode(addr, codeBytes)
	}
	// set storage
	storage, ok := contractParams.(map[string]interface{})["storage"]
	if ok {
		for slot, value := range storage.(map[string]interface{}) {
			if _, ok := value.(string); !ok {
				t.Fatal("value is not a string")
			}
			db.SetState(addr, common.HexToHash(slot), common.HexToHash(value.(string)))
		}
	}

	// sets flags
	flags, ok := contractParams.(map[string]interface{})["flags"]
	if ok {
		if _, ok := flags.(float64); !ok {
			t.Fatal("flags is not a float64")
		}
		flagsFloat, _ := flags.(float64)
		flagsUint8 := uint8(flagsFloat)

		db.SetFlags(addr, flagsUint8)
	}

	// sets balance
	balance, ok := contractParams.(map[string]interface{})["balance"]
	if ok {
		if _, ok := balance.(string); !ok {
			t.Fatal("balance is not a string")
		}
		balanceString, _ := balance.(string)
		_, ok := new(big.Int).SetString(balanceString, 0)
		if !ok {
			t.Fatal("balance cannot be converted to big int")
		}
		// DO NOT set balance, messes with tests dealing with share count
		// db.SetBalance(addr, balanceBigInt)
	}
}

func setState(allocJson map[string]interface{}, addr common.Address, hasStorage bool, db *state.StateDB, t *testing.T) {
	db.CreateAccount(addr)
	// set code
	aParsed := addr.String()[2:]
	contractParams, ok := allocJson[aParsed]
	if !ok {
		t.Fatal("unable to find contract")
	}
	code, ok := contractParams.(map[string]interface{})["code"]
	if !ok {
		t.Fatal("unable to get code")
	}
	if _, ok := code.(string); !ok {
		t.Fatal("code is not a string")
	}
	codeStr := code.(string)
	codeBytes := common.Hex2Bytes(codeStr[2:])
	db.SetCode(addr, codeBytes)
	// set storage
	if hasStorage {
		storage, ok := contractParams.(map[string]interface{})["storage"]
		if !ok {
			t.Fatal("unable to get storage")
		}
		for slot, value := range storage.(map[string]interface{}) {
			if _, ok := value.(string); !ok {
				t.Fatal("value is not a string")
			}
			db.SetState(addr, common.HexToHash(slot), common.HexToHash(value.(string)))
		}

	}
	flags, ok := contractParams.(map[string]interface{})["flags"]
	if !ok {
		t.Fatal("unable to get flags")
	}
	if _, ok := flags.(float64); !ok {
		t.Fatal("flags is not a float64")
	}
	flagsFloat, _ := flags.(float64)
	flagsUint8 := uint8(flagsFloat)

	db.SetFlags(addr, flagsUint8)
}

func deployGenesisWithDbOld(t *testing.T, db *state.StateDB) *state.StateDB {
	path := BASE_OPTIMISM_PATH + "/.devnet/genesis-l2.json"
	genesisJson := readJson(path, t)
	allocJson, ok := genesisJson["alloc"]
	if !ok {
		t.Fatal("unable to parse genesis json")
	}

	setState(allocJson.(map[string]interface{}), params.BlastAccountConfigurationAddress, false, db, t)
	setState(allocJson.(map[string]interface{}), params.BlastGasAddress, true, db, t)
	setState(allocJson.(map[string]interface{}), params.BlastSharesAddress, true, db, t)
	return db
}

func deployGenesisWithDb(t *testing.T, db *state.StateDB) {
	path := BASE_OPTIMISM_PATH + "/.devnet/genesis-l2.json"
	genesisJson := readJson(path, t)
	allocJson, ok := genesisJson["alloc"]
	if !ok {
		t.Fatal("unable to parse genesis json")
	}
	allocJsonMap := allocJson.(map[string]interface{})
	for key := range allocJsonMap {
		setStateGeneral(allocJson.(map[string]interface{}), key, db, t)
	}
}

func TestDeploy2(t *testing.T) {
	db, err := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}
	deployGenesisWithDb(t, db)
}

func setup(t *testing.T) (*Config, *vm.EVM) {
	cfg := new(Config)
	setDefaults(cfg, new(uint64))
	configureOptimism(cfg)

	if cfg.State == nil {
		cfg.State, _ = state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	}
	addr := common.HexToAddress("0x8")
	cfg.State.CreateAccount(addr)
	cfg.State.AddBalance(addr, big.NewInt(1000000000000))

	env := NewEnv(cfg)
	var rules = cfg.ChainConfig.Rules(env.Context.BlockNumber, env.Context.Random != nil, env.Context.Time)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, &addr, vm.ActivePrecompiles(rules), nil)

	deployGenesisWithDb(t, cfg.State)
	return cfg, env
}

func configureOptimism(cfg *Config) {
	cfg.ChainConfig.BedrockBlock = big.NewInt(0)
	cfg.ChainConfig.RegolithTime = &regolithZeroTime
	cfg.ChainConfig.CanyonTime = &regolithZeroTime
	cfg.ChainConfig.Optimism = &params.OptimismConfig{
		EIP1559Elasticity:        6,
		EIP1559Denominator:       50,
		EIP1559DenominatorCanyon: 250,
	}
}
func enableTracer(cfg *Config) {
	cfg.EVMConfig.Tracer = logger.NewStructLogger(nil)
}

func setupDb(t *testing.T) *state.StateDB {
	state, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	deployGenesisWithDb(t, state)
	return state
}

type TransactionParams struct {
	sender      common.Address
	to          *common.Address
	input       []byte
	value       *big.Int
	blockNumber *big.Int
	gasPrice    *big.Int
}

func staticCall(tp *TransactionParams, statedb *state.StateDB, t *testing.T) []byte {
	cfg := new(Config)
	if tp.gasPrice != nil {
		cfg.GasPrice = tp.gasPrice
	} else {
		cfg.GasPrice = common.Big2
	}

	cfg.BaseFee = common.Big1
	cfg.GasLimit = 100_000_000
	if statedb != nil {
		cfg.State = statedb
	}
	if tp.value != nil {
		cfg.Value = tp.value
	}

	if tp.blockNumber != nil {
		cfg.BlockNumber = tp.blockNumber
		// Convert block number to time
		cfg.Time = new(big.Int).Mul(cfg.BlockNumber, big.NewInt(1)).Uint64()
	}
	setDefaults(cfg, new(uint64))

	if cfg.State == nil {
		cfg.State, _ = state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	}
	// ensure balance is high enough to cover gas
	cfg.State.CreateAccount(tp.sender)
	cfg.State.AddBalance(tp.sender, new(big.Int).SetUint64(300_000_000))

	// add additional value if non-zero value
	if tp.value != nil {
		cfg.State.AddBalance(tp.sender, tp.value)
	}

	env := NewEnv(cfg)
	var rules = cfg.ChainConfig.Rules(env.Context.BlockNumber, env.Context.Random != nil, env.Context.Time)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, tp.to, vm.ActivePrecompiles(rules), nil)

	gasTracker := vm.NewGasTracker()
	ret, _, err := env.StaticCall(vm.AccountRef(tp.sender), *tp.to, tp.input, cfg.GasLimit, gasTracker)
	if err != nil {
		t.Fatal(err)
	}
	return ret
}
func stateTransition(tp *TransactionParams, statedb *state.StateDB, t *testing.T) *core.ExecutionResult {
	result, _ := stateTransitionLogs(tp, statedb, t)
	return result
}

func stateTransitionLogs(tp *TransactionParams, statedb *state.StateDB, t *testing.T) (*core.ExecutionResult, *Config) {
	cfg := new(Config)
	if tp.gasPrice != nil {
		cfg.GasPrice = tp.gasPrice
	} else {
		cfg.GasPrice = common.Big2
	}

	cfg.BaseFee = common.Big1
	cfg.GasLimit = 100_000_000
	if statedb != nil {
		cfg.State = statedb
	}
	if tp.value != nil {
		cfg.Value = tp.value
	}

	if tp.blockNumber != nil {
		cfg.BlockNumber = tp.blockNumber
		// Convert block number to time
		cfg.Time = new(big.Int).Mul(cfg.BlockNumber, big.NewInt(1)).Uint64()
	}
	setDefaults(cfg, new(uint64))
	configureOptimism(cfg)
	enableTracer(cfg)

	if cfg.State == nil {
		cfg.State, _ = state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	}
	// ensure balance is high enough to cover gas
	cfg.State.CreateAccount(tp.sender)
	cfg.State.AddBalance(tp.sender, new(big.Int).SetUint64(300_000_000))

	// add additional value if non-zero value
	if tp.value != nil {
		cfg.State.AddBalance(tp.sender, tp.value)
	}

	msg := &core.Message{
		From:      tp.sender,
		GasLimit:  cfg.GasLimit,
		GasPrice:  cfg.GasPrice,
		GasFeeCap: cfg.GasPrice,
		GasTipCap: new(big.Int).Sub(cfg.GasPrice, cfg.BaseFee),
		To:        tp.to,
		Data:      tp.input,
		Value:     cfg.Value,
	}

	env := NewEnv(cfg)
	// TODO(blast): refactor this for clarity
	env.Context.L1CostFunc = types.NewL1CostFunc(cfg.ChainConfig, cfg.State)
	var rules = cfg.ChainConfig.Rules(env.Context.BlockNumber, env.Context.Random != nil, env.Context.Time)
	cfg.State.Prepare(rules, cfg.Origin, cfg.Coinbase, msg.To, vm.ActivePrecompiles(rules), nil)
	gasPool := new(core.GasPool).AddGas(200_000_000)
	st := core.NewStateTransition(env, msg, gasPool)
	executionResult, err := st.TransitionDb()
	if err != nil {
		t.Fatal(err)
	}
	return executionResult, cfg
}

func getAddr(a uint64) common.Address {
	b := new(big.Int).SetUint64(a)
	return common.BigToAddress(b)
}

func getPublicVar(methodName string, predeployAddr common.Address, env *vm.EVM, t *testing.T) interface{} {
	abi := getAbi(addrToAbiPath[predeployAddr.String()], t)
	input, err := abi.Pack(methodName)
	if err != nil {
		t.Fatal(err)
	}

	gasTracker := vm.NewGasTracker()
	ret, _, err := env.StaticCall(vm.AccountRef(getAddr(1)), predeployAddr, input, uint64(10_000), gasTracker)
	if err != nil {
		t.Fatal(err)
	}
	unpack, err := abi.Unpack(methodName, ret)
	if err != nil {
		t.Fatal(err)
	}
	return unpack[0]
}

func getVar(methodName string, predeployAddr common.Address, db *state.StateDB, t *testing.T) interface{} {
	cfg := new(Config)
	setDefaults(cfg, new(uint64))
	cfg.State = db
	env := NewEnv(cfg)
	return getPublicVar(methodName, predeployAddr, env, t)
}

func compareGasConfig(varName string, env *vm.EVM, t *testing.T) *big.Int {
	devnetVarRaw := getDevnetVar(varName, t)
	contractVar := getPublicVar(varName, params.BlastGasAddress, env, t).(*big.Int)

	devnetVar, ok := new(big.Int).SetString(devnetVarRaw[2:], 16)
	if !ok {
		t.Fatal("could not convert devnet var into big int")
	}
	if devnetVar.Cmp(contractVar) != 0 {
		t.Fatal("devnet and contract not equivalent")
	}
	return devnetVar

}

func setGasMode(db *state.StateDB, addr common.Address, t *testing.T) {
	setGasModeAtBlock(db, addr, common.Big0, t)
}

func setGasModeAtBlock(db *state.StateDB, addr common.Address, blockNumber *big.Int, t *testing.T) {
	abi := getAbi(addrToAbiPath[params.BlastGasAddress.String()], t)
	input, err := abi.Pack("setGasMode", addr, uint8(1))
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender:      params.BlastAccountConfigurationAddress,
		to:          &(params.BlastGasAddress),
		input:       input,
		blockNumber: blockNumber,
	}
	stateTransition(&tp, db, t)
}

func readGasVars(db *state.StateDB, addr common.Address, t *testing.T) (*big.Int, *big.Int, *big.Int, uint8) {
	return readGasVarsAtBlock(db, addr, common.Big0, t)
}

func readGasMode(db *state.StateDB, addr common.Address, t *testing.T) uint8 {
	_, _, _, mode := readGasVars(db, addr, t)
	return mode
}

func readGasVarsAtBlock(db *state.StateDB, addr common.Address, blockNumber *big.Int, t *testing.T) (*big.Int, *big.Int, *big.Int, uint8) {
	abi := getAbi(addrToAbiPath[params.BlastGasAddress.String()], t)
	input, err := abi.Pack("readGasParams", addr)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender:      EOA_ADDR,
		to:          &(params.BlastGasAddress),
		input:       input,
		blockNumber: blockNumber,
	}
	ret := staticCall(&tp, db, t)
	decodedRes, err := abi.Unpack("readGasParams", ret)

	etherSeconds := decodedRes[0].(*big.Int)
	etherBalance := decodedRes[1].(*big.Int)
	lastUpdated := decodedRes[2].(*big.Int)
	mode := decodedRes[3].(uint8)
	return etherSeconds, etherBalance, lastUpdated, mode
}

func readYield(db *state.StateDB, addr common.Address, t *testing.T) *big.Int {
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("readClaimableYield", addr)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender: EOA_ADDR,
		to:     &(params.BlastAccountConfigurationAddress),
		input:  input,
	}
	ret := staticCall(&tp, db, t)
	decodedRes, err := abi.Unpack("readClaimableYield", ret)

	claimableYield := decodedRes[0].(*big.Int)
	return claimableYield
}

func readGovernor(db *state.StateDB, addr common.Address, t *testing.T) common.Address {
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("governorMap", addr)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender: EOA_ADDR,
		to:     &(params.BlastAccountConfigurationAddress),
		input:  input,
	}
	ret := staticCall(&tp, db, t)
	decodedRes, err := abi.Unpack("governorMap", ret)

	governor := decodedRes[0].(common.Address)
	return governor
}

func readYieldMode(db *state.StateDB, addr common.Address, t *testing.T) uint8 {
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("readYieldConfiguration", addr)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender: EOA_ADDR,
		to:     &(params.BlastAccountConfigurationAddress),
		input:  input,
	}
	ret := staticCall(&tp, db, t)
	decodedRes, err := abi.Unpack("readYieldConfiguration", ret)

	mode := decodedRes[0].(uint8)
	return mode
}

func deployTestContract(db *state.StateDB, filepath string, t *testing.T) common.Address {
	bytecode := getBytecodeFromForStandardContract(filepath, t)
	cfg := new(Config)
	setDefaults(cfg, new(uint64))
	cfg.State = db
	_, addr, _, err := Create(bytecode, cfg)
	if err != nil {
		t.Error(err)
	}
	return addr
}

func distributeYield(db *state.StateDB, amount *big.Int, t *testing.T) *big.Int {
	cfg := new(Config)
	setDefaults(cfg, new(uint64))
	cfg.State = db
	env := NewEnv(cfg)
	sharesReporter := getPublicVar("REPORTER", params.BlastSharesAddress, env, t).(common.Address)
	abi := getAbi(addrToAbiPath[params.BlastSharesAddress.String()], t)
	input, err := abi.Pack("addValue", amount)
	if err != nil {
		t.Fatal(err)
	}
	gasTracker := vm.NewGasTracker()
	_, _, err = env.Call(vm.AccountRef(l2Alias(sharesReporter)), params.BlastSharesAddress, input, uint64(100_000), big.NewInt(0), gasTracker)
	if err != nil {
		t.Fatal(err)
	}
	newPrice := getPublicVar("price", params.BlastSharesAddress, env, t).(*big.Int)
	return newPrice
}

func readClaimableYield(db *state.StateDB, addr common.Address, t *testing.T) *big.Int {
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("readClaimableYield", addr)
	if err != nil {
		t.Fatal(err)
	}

	readClaimableTp := TransactionParams{
		sender: EOA_ADDR,
		to:     &params.BlastAccountConfigurationAddress,
		input:  input,
	}
	result := staticCall(&readClaimableTp, db, t)
	claimableYield := new(big.Int).SetBytes(result)
	return claimableYield
}

func claimAllYield(db *state.StateDB, addr common.Address, t *testing.T, recipient *common.Address) *big.Int {
	gov := readGovernor(db, addr, t)
	if gov == common.HexToAddress("0x0") {
		gov = addr
	}

	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	if recipient == nil {
		recipient = &gov
	}
	input, err := abi.Pack("claimAllYield", addr, *recipient)
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

	claimedYield := new(big.Int).SetBytes(result.ReturnData)
	return claimedYield
}
func claimYield(db *state.StateDB, addr common.Address, t *testing.T, amount *big.Int, recipient *common.Address) *big.Int {
	yield, err := claimYieldWithErr(db, addr, t, amount, recipient)
	if err != nil {
		t.Fatal(err)
	}
	return yield
}

func claimYieldWithErr(db *state.StateDB, addr common.Address, t *testing.T, amount *big.Int, recipient *common.Address) (*big.Int, error) {
	gov := readGovernor(db, addr, t)
	if gov == common.HexToAddress("0x0") {
		gov = addr
	}

	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	if recipient == nil {
		recipient = &gov
	}
	input, err := abi.Pack("claimYield", addr, *recipient, amount)
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
		return nil, result.Err
	}

	claimedYield := new(big.Int).SetBytes(result.ReturnData)
	return claimedYield, nil
}

func claimGas(db *state.StateDB, addr common.Address, recipient *common.Address, gasToClaim *big.Int, gasSecondsToConsume *big.Int, blockNumber *big.Int, t *testing.T) *big.Int {
	gov := readGovernor(db, addr, t)
	if gov == common.HexToAddress("0x0") {
		gov = addr
	}
	if recipient == nil {
		recipient = &gov
	}
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("claimGas", addr, *recipient, gasToClaim, gasSecondsToConsume)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender:      gov,
		to:          &params.BlastAccountConfigurationAddress,
		input:       input,
		blockNumber: blockNumber,
	}
	result, cfg := stateTransitionLogs(&tp, db, t)
	tracer := cfg.EVMConfig.Tracer
	structLogger := tracer.(*logger.StructLogger)
	tracingResult, err := structLogger.GetResult()
	if err != nil {
		t.Fatal(err)
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, tracingResult, "", "\t")
	if error != nil {
		t.Fatal(error)
	}
	fmt.Println(string(prettyJSON.Bytes()))

	if result.Err != nil {
		t.Fatal(result.Err)
	}
	claimedGas := new(big.Int).SetBytes(result.ReturnData)
	return claimedGas
}

func claimGasAtMinClaimRate(db *state.StateDB, addr common.Address, recipient *common.Address, claimRate *big.Int, blockNumber *big.Int, t *testing.T) *big.Int {
	gov := readGovernor(db, addr, t)
	if gov == common.HexToAddress("0x0") {
		gov = addr
	}
	if recipient == nil {
		recipient = &gov
	}
	abi := getAbi(addrToAbiPath[params.BlastAccountConfigurationAddress.String()], t)
	input, err := abi.Pack("claimGasAtMinClaimRate", addr, *recipient, claimRate)
	if err != nil {
		t.Fatal(err)
	}
	tp := TransactionParams{
		sender:      gov,
		to:          &params.BlastAccountConfigurationAddress,
		input:       input,
		blockNumber: blockNumber,
	}

	result, cfg := stateTransitionLogs(&tp, db, t)
	tracer := cfg.EVMConfig.Tracer
	structLogger := tracer.(*logger.StructLogger)
	tracingResult, err := structLogger.GetResult()
	if err != nil {
		t.Fatal(err)
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, tracingResult, "", "\t")
	if error != nil {
		t.Fatal(error)
	}

	if result.Err != nil {
		t.Fatal(result.Err)
	}
	claimedGas := new(big.Int).SetBytes(result.ReturnData)
	return claimedGas
}

func readBps(db *state.StateDB, gasToClaim *big.Int, gasSecondsToConsume *big.Int, t *testing.T) (claimRateBps, gasSecondsNormalized *big.Int) {
	abi := getAbi(addrToAbiPath[params.BlastGasAddress.String()], t)
	input, err := abi.Pack("getClaimRateBps", gasSecondsToConsume, gasToClaim)
	if err != nil {
		t.Fatal(err)
	}

	tp := TransactionParams{
		to:     &params.BlastGasAddress,
		sender: EOA_ADDR,
		input:  input,
	}
	ret := staticCall(&tp, db, t)
	results, err := abi.Unpack("getClaimRateBps", ret)
	if err != nil {
		t.Fatal(err)
	}
	claimRateBps = results[0].(*big.Int)
	gasSecondsNormalized = results[1].(*big.Int)
	return claimRateBps, gasSecondsNormalized
}
