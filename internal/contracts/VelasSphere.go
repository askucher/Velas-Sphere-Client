// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethdepositcontract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthdepositcontractABI is the input ABI used to generate the binding from.
const EthdepositcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"banNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"proposePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pull\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_places\",\"type\":\"uint256\"}],\"name\":\"depositWithNodes\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pull\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_places\",\"type\":\"uint256\"}],\"name\":\"changePool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height_end\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"changeNodePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405264174876e800600055603260015534801561001e57600080fd5b5060016006600001819055506001600660010181905550600160066002018190555060016006600301819055506000600660040160006101000a81548160ff021916908315150217905550603e600381905550610ea5806100806000396000f3fe60806040526004361061007b5760003560e01c80634f1c58e31161004e5780634f1c58e3146101e057806351cff8d9146102395780638e8b45c61461028a57806396ea8630146102c25761007b565b80632790fb77146100855780632eefc412146100ca5780633296972f146101365780633faba59e1461018f575b61008361034f565b005b34801561009157600080fd5b506100c8600480360360408110156100a857600080fd5b81019080803590602001909291908035906020019092919050505061045f565b005b610134600480360360a08110156100e057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506104dc565b005b34801561014257600080fd5b5061018d6004803603608081101561015957600080fd5b81019080803590602001909291908035906020019092919080359060200190929190803590602001909291905050506105d6565b005b34801561019b57600080fd5b506101de600480360360208110156101b257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610670565b005b3480156101ec57600080fd5b506102376004803603608081101561020357600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610714565b005b34801561024557600080fd5b506102886004803603602081101561025c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107b0565b005b6102c0600480360360408110156102a057600080fd5b81019080803590602001909291908035906020019092919050505061089b565b005b3480156102ce57600080fd5b5061034d600480360360e08110156102e557600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506108cc565b005b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000341161039f57600080fd5b348160080160008282540192505081905550600015158160090160009054906101000a900460ff161515141561043f57600681600001600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a900460ff168160040160006101000a81548160ff02191690831515021790555090505060016004600082825401925050819055505b60018160090160006101000a81548160ff02191690831515021790555050565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905081816006016001018190555082816006016000018190555060018160050160006101000a81548160ff021916908315150217905550505050565b6000600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600015158160020160009054906101000a900460ff1615151461054157600080fd5b60018160020160006101000a81548160ff021916908315150217905550600054341461056c57600080fd5b610574610c46565b816008016001018190555060055481600801600001819055508481600301600001819055508381600301600101819055508281600301600201819055508181600301600301819055506001600260008282540192505081905550505050505050565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905084816000016000018190555083816000016001018190555082816000016002018190555081816000016003018190555060018160000160040160006101000a81548160ff0219169083151502179055505050505050565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600a0160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160016101000a81548160ff0219169083151502179055505050565b6000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600115158160020160009054906101000a900460ff1615151461077957600080fd5b8481600301600001819055508381600301600101819055508281600301600201819055508181600301600301819055505050505050565b6000600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600101541161080457600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549081150290604051600060405180830381858888f1935050505015801561088c573d6000803e3d6000fd5b50600081600101819055505050565b6108a361034f565b6000821480156108b35750600081145b156108bd576108c8565b6108c7828261045f565b5b5050565b6000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020160009054906101000a900460ff1661092a57600080fd5b6000600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506001151581600a0160008460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160019054906101000a900460ff16151514156109f457600080fd5b6000600e60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600d01541415610a8857878160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b868160080160000160008282540192505081905550858160080160010160008282540192505081905550848160080160020160008282540192505081905550838160080160030160008282540192505081905550600181600d01600082825401925050819055506000610b018360000189898989610ca6565b90508082600c01600082825401925050819055506001548a014310610c385760035482600d015410610b4457610b3b8983600c0154610ce2565b50505050610c3d565b600e60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382016000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff02191690555050600882016000808201600090556001820160009055600282016000905560038201600090555050600c820160009055600d820160009055505050505050610c3d565b505050505b50505050505050565b600080600b600060055481526020019081526020016000209050605e816001015410610c7e5760016005600082825401925050819055505b600081600101546001901b905060018260010160008282540192505081905550809250505090565b60008085876000015402810190508487600101540281019050828760030154028101905083876002015402810190508091505095945050505050565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080154811115610d3157600080fd5b80600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080160008282540392505081905550600e60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382016000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff02191690555050600882016000808201600090556001820160009055600282016000905560038201600090555050600c820160009055600d8201600090555050505056fea265627a7a72315820c45d8e79752fd2de80a2767668e2af20cfec5e8a0bb8996df08316ef394510a364736f6c63430005100032"

// DeployEthdepositcontract deploys a new Ethereum contract, binding an instance of Ethdepositcontract to it.
func DeployEthdepositcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethdepositcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthdepositcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthdepositcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethdepositcontract{EthdepositcontractCaller: EthdepositcontractCaller{contract: contract}, EthdepositcontractTransactor: EthdepositcontractTransactor{contract: contract}, EthdepositcontractFilterer: EthdepositcontractFilterer{contract: contract}}, nil
}

// Ethdepositcontract is an auto generated Go binding around an Ethereum contract.
type Ethdepositcontract struct {
	EthdepositcontractCaller     // Read-only binding to the contract
	EthdepositcontractTransactor // Write-only binding to the contract
	EthdepositcontractFilterer   // Log filterer for contract events
}

// EthdepositcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthdepositcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthdepositcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthdepositcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthdepositcontractSession struct {
	Contract     *Ethdepositcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthdepositcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthdepositcontractCallerSession struct {
	Contract *EthdepositcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// EthdepositcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthdepositcontractTransactorSession struct {
	Contract     *EthdepositcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// EthdepositcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthdepositcontractRaw struct {
	Contract *Ethdepositcontract // Generic contract binding to access the raw methods on
}

// EthdepositcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthdepositcontractCallerRaw struct {
	Contract *EthdepositcontractCaller // Generic read-only contract binding to access the raw methods on
}

// EthdepositcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthdepositcontractTransactorRaw struct {
	Contract *EthdepositcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthdepositcontract creates a new instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontract(address common.Address, backend bind.ContractBackend) (*Ethdepositcontract, error) {
	contract, err := bindEthdepositcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethdepositcontract{EthdepositcontractCaller: EthdepositcontractCaller{contract: contract}, EthdepositcontractTransactor: EthdepositcontractTransactor{contract: contract}, EthdepositcontractFilterer: EthdepositcontractFilterer{contract: contract}}, nil
}

// NewEthdepositcontractCaller creates a new read-only instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractCaller(address common.Address, caller bind.ContractCaller) (*EthdepositcontractCaller, error) {
	contract, err := bindEthdepositcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractCaller{contract: contract}, nil
}

// NewEthdepositcontractTransactor creates a new write-only instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*EthdepositcontractTransactor, error) {
	contract, err := bindEthdepositcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractTransactor{contract: contract}, nil
}

// NewEthdepositcontractFilterer creates a new log filterer instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*EthdepositcontractFilterer, error) {
	contract, err := bindEthdepositcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractFilterer{contract: contract}, nil
}

// bindEthdepositcontract binds a generic wrapper to an already deployed contract.
func bindEthdepositcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthdepositcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethdepositcontract *EthdepositcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethdepositcontract.Contract.EthdepositcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethdepositcontract *EthdepositcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.EthdepositcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethdepositcontract *EthdepositcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.EthdepositcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethdepositcontract *EthdepositcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethdepositcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethdepositcontract *EthdepositcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethdepositcontract *EthdepositcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.contract.Transact(opts, method, params...)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) BanNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "banNode", _node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractSession) BanNode(_node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.BanNode(&_Ethdepositcontract.TransactOpts, _node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) BanNode(_node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.BanNode(&_Ethdepositcontract.TransactOpts, _node)
}

// ChangeNodePricing is a paid mutator transaction binding the contract method 0x4f1c58e3.
//
// Solidity: function changeNodePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ChangeNodePricing(opts *bind.TransactOpts, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "changeNodePricing", _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ChangeNodePricing is a paid mutator transaction binding the contract method 0x4f1c58e3.
//
// Solidity: function changeNodePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ChangeNodePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeNodePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ChangeNodePricing is a paid mutator transaction binding the contract method 0x4f1c58e3.
//
// Solidity: function changeNodePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ChangeNodePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeNodePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ChangePool is a paid mutator transaction binding the contract method 0x2790fb77.
//
// Solidity: function changePool(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ChangePool(opts *bind.TransactOpts, _pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "changePool", _pull, _places)
}

// ChangePool is a paid mutator transaction binding the contract method 0x2790fb77.
//
// Solidity: function changePool(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ChangePool(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangePool(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// ChangePool is a paid mutator transaction binding the contract method 0x2790fb77.
//
// Solidity: function changePool(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ChangePool(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangePool(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) CreateInvoice(opts *bind.TransactOpts, height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "createInvoice", height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) CreateInvoice(height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) CreateInvoice(height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) DepositWithNodes(opts *bind.TransactOpts, _pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "depositWithNodes", _pull, _places)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractSession) DepositWithNodes(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) DepositWithNodes(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ProposePricing(opts *bind.TransactOpts, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "proposePricing", _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ProposePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ProposePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ProposePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ProposePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x2eefc412.
//
// Solidity: function registerNode(address addr, uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) RegisterNode(opts *bind.TransactOpts, addr common.Address, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "registerNode", addr, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x2eefc412.
//
// Solidity: function registerNode(address addr, uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) RegisterNode(addr common.Address, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, addr, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x2eefc412.
//
// Solidity: function registerNode(address addr, uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) RegisterNode(addr common.Address, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, addr, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) Withdraw(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "withdraw", addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractSession) Withdraw(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) Withdraw(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, addr)
}
