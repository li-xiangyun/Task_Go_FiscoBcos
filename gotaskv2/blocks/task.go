// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocks

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_owner\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllBonus\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllComments\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllDesc\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllIssuers\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllStatus\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllWorkers\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"qryOneTask\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"take\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Task is an auto generated Go binding around a Solidity contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around a Solidity contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around a Solidity contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around a Solidity contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) constant returns(uint256 balance)
func (_Task *TaskCaller) BalanceOf(opts *bind.CallOpts, _owner string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) constant returns(uint256 balance)
func (_Task *TaskSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) constant returns(uint256 balance)
func (_Task *TaskCallerSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, _owner)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) constant returns(bool)
func (_Task *TaskCaller) Login(opts *bind.CallOpts, username string, passwd string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "login", username, passwd)
	return *ret0, err
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) constant returns(bool)
func (_Task *TaskSession) Login(username string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, username, passwd)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) constant returns(bool)
func (_Task *TaskCallerSession) Login(username string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, username, passwd)
}

// QryAllBonus is a free data retrieval call binding the contract method 0xb59c5cc8.
//
// Solidity: function qryAllBonus() constant returns(uint256[])
func (_Task *TaskCaller) QryAllBonus(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "qryAllBonus")
	return *ret0, err
}

// QryAllBonus is a free data retrieval call binding the contract method 0xb59c5cc8.
//
// Solidity: function qryAllBonus() constant returns(uint256[])
func (_Task *TaskSession) QryAllBonus() ([]*big.Int, error) {
	return _Task.Contract.QryAllBonus(&_Task.CallOpts)
}

// QryAllBonus is a free data retrieval call binding the contract method 0xb59c5cc8.
//
// Solidity: function qryAllBonus() constant returns(uint256[])
func (_Task *TaskCallerSession) QryAllBonus() ([]*big.Int, error) {
	return _Task.Contract.QryAllBonus(&_Task.CallOpts)
}

// QryAllComments is a free data retrieval call binding the contract method 0x10d49e51.
//
// Solidity: function qryAllComments() constant returns(string[])
func (_Task *TaskCaller) QryAllComments(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "qryAllComments")
	return *ret0, err
}

// QryAllComments is a free data retrieval call binding the contract method 0x10d49e51.
//
// Solidity: function qryAllComments() constant returns(string[])
func (_Task *TaskSession) QryAllComments() ([]string, error) {
	return _Task.Contract.QryAllComments(&_Task.CallOpts)
}

// QryAllComments is a free data retrieval call binding the contract method 0x10d49e51.
//
// Solidity: function qryAllComments() constant returns(string[])
func (_Task *TaskCallerSession) QryAllComments() ([]string, error) {
	return _Task.Contract.QryAllComments(&_Task.CallOpts)
}

// QryAllDesc is a free data retrieval call binding the contract method 0x8fb7fccd.
//
// Solidity: function qryAllDesc() constant returns(string[])
func (_Task *TaskCaller) QryAllDesc(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "qryAllDesc")
	return *ret0, err
}

// QryAllDesc is a free data retrieval call binding the contract method 0x8fb7fccd.
//
// Solidity: function qryAllDesc() constant returns(string[])
func (_Task *TaskSession) QryAllDesc() ([]string, error) {
	return _Task.Contract.QryAllDesc(&_Task.CallOpts)
}

// QryAllDesc is a free data retrieval call binding the contract method 0x8fb7fccd.
//
// Solidity: function qryAllDesc() constant returns(string[])
func (_Task *TaskCallerSession) QryAllDesc() ([]string, error) {
	return _Task.Contract.QryAllDesc(&_Task.CallOpts)
}

// QryAllIssuers is a free data retrieval call binding the contract method 0x01d7d462.
//
// Solidity: function qryAllIssuers() constant returns(string[])
func (_Task *TaskCaller) QryAllIssuers(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "qryAllIssuers")
	return *ret0, err
}

// QryAllIssuers is a free data retrieval call binding the contract method 0x01d7d462.
//
// Solidity: function qryAllIssuers() constant returns(string[])
func (_Task *TaskSession) QryAllIssuers() ([]string, error) {
	return _Task.Contract.QryAllIssuers(&_Task.CallOpts)
}

// QryAllIssuers is a free data retrieval call binding the contract method 0x01d7d462.
//
// Solidity: function qryAllIssuers() constant returns(string[])
func (_Task *TaskCallerSession) QryAllIssuers() ([]string, error) {
	return _Task.Contract.QryAllIssuers(&_Task.CallOpts)
}

// QryAllStatus is a free data retrieval call binding the contract method 0x41f12eb7.
//
// Solidity: function qryAllStatus() constant returns(uint8[])
func (_Task *TaskCaller) QryAllStatus(opts *bind.CallOpts) ([]uint8, error) {
	var (
		ret0 = new([]uint8)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "qryAllStatus")
	return *ret0, err
}

// QryAllStatus is a free data retrieval call binding the contract method 0x41f12eb7.
//
// Solidity: function qryAllStatus() constant returns(uint8[])
func (_Task *TaskSession) QryAllStatus() ([]uint8, error) {
	return _Task.Contract.QryAllStatus(&_Task.CallOpts)
}

// QryAllStatus is a free data retrieval call binding the contract method 0x41f12eb7.
//
// Solidity: function qryAllStatus() constant returns(uint8[])
func (_Task *TaskCallerSession) QryAllStatus() ([]uint8, error) {
	return _Task.Contract.QryAllStatus(&_Task.CallOpts)
}

// QryAllWorkers is a free data retrieval call binding the contract method 0xc03f62db.
//
// Solidity: function qryAllWorkers() constant returns(string[])
func (_Task *TaskCaller) QryAllWorkers(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "qryAllWorkers")
	return *ret0, err
}

// QryAllWorkers is a free data retrieval call binding the contract method 0xc03f62db.
//
// Solidity: function qryAllWorkers() constant returns(string[])
func (_Task *TaskSession) QryAllWorkers() ([]string, error) {
	return _Task.Contract.QryAllWorkers(&_Task.CallOpts)
}

// QryAllWorkers is a free data retrieval call binding the contract method 0xc03f62db.
//
// Solidity: function qryAllWorkers() constant returns(string[])
func (_Task *TaskCallerSession) QryAllWorkers() ([]string, error) {
	return _Task.Contract.QryAllWorkers(&_Task.CallOpts)
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) constant returns(string, string, uint256, string, uint8)
func (_Task *TaskCaller) QryOneTask(opts *bind.CallOpts, taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Task.contract.Call(opts, out, "qryOneTask", taskID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) constant returns(string, string, uint256, string, uint8)
func (_Task *TaskSession) QryOneTask(taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	return _Task.Contract.QryOneTask(&_Task.CallOpts, taskID)
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) constant returns(string, string, uint256, string, uint8)
func (_Task *TaskCallerSession) QryOneTask(taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	return _Task.Contract.QryOneTask(&_Task.CallOpts, taskID)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Task *TaskCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Task *TaskSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Task *TaskCallerSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactor) Commit(opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "commit", worker, passwd, taskID)
}

func (_Task *TaskTransactor) AsyncCommit(handler func(*types.Receipt, error), opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "commit", worker, passwd, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskSession) Commit(worker string, passwd string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, worker, passwd, taskID)
}

func (_Task *TaskSession) AsyncCommit(handler func(*types.Receipt, error), worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncCommit(handler, &_Task.TransactOpts, worker, passwd, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Commit(worker string, passwd string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, worker, passwd, taskID)
}

func (_Task *TaskTransactorSession) AsyncCommit(handler func(*types.Receipt, error), worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncCommit(handler, &_Task.TransactOpts, worker, passwd, taskID)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskTransactor) Confirm(opts *bind.TransactOpts, issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "confirm", issuer, passwd, taskID, comment, status)
}

func (_Task *TaskTransactor) AsyncConfirm(handler func(*types.Receipt, error), opts *bind.TransactOpts, issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "confirm", issuer, passwd, taskID, comment, status)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskSession) Confirm(issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

func (_Task *TaskSession) AsyncConfirm(handler func(*types.Receipt, error), issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.Contract.AsyncConfirm(handler, &_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskTransactorSession) Confirm(issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

func (_Task *TaskTransactorSession) AsyncConfirm(handler func(*types.Receipt, error), issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.Contract.AsyncConfirm(handler, &_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskTransactor) Issue(opts *bind.TransactOpts, issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "issue", issuer, passwd, desc, bonus)
}

func (_Task *TaskTransactor) AsyncIssue(handler func(*types.Receipt, error), opts *bind.TransactOpts, issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "issue", issuer, passwd, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskSession) Issue(issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, passwd, desc, bonus)
}

func (_Task *TaskSession) AsyncIssue(handler func(*types.Receipt, error), issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncIssue(handler, &_Task.TransactOpts, issuer, passwd, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskTransactorSession) Issue(issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, passwd, desc, bonus)
}

func (_Task *TaskTransactorSession) AsyncIssue(handler func(*types.Receipt, error), issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncIssue(handler, &_Task.TransactOpts, issuer, passwd, desc, bonus)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskTransactor) Mint(opts *bind.TransactOpts, _to string, _value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "mint", _to, _value)
}

func (_Task *TaskTransactor) AsyncMint(handler func(*types.Receipt, error), opts *bind.TransactOpts, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskSession) Mint(_to string, _value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, _to, _value)
}

func (_Task *TaskSession) AsyncMint(handler func(*types.Receipt, error), _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncMint(handler, &_Task.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskTransactorSession) Mint(_to string, _value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, _to, _value)
}

func (_Task *TaskTransactorSession) AsyncMint(handler func(*types.Receipt, error), _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncMint(handler, &_Task.TransactOpts, _to, _value)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskTransactor) Register(opts *bind.TransactOpts, username string, passwd string) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "register", username, passwd)
}

func (_Task *TaskTransactor) AsyncRegister(handler func(*types.Receipt, error), opts *bind.TransactOpts, username string, passwd string) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "register", username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskSession) Register(username string, passwd string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, username, passwd)
}

func (_Task *TaskSession) AsyncRegister(handler func(*types.Receipt, error), username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.AsyncRegister(handler, &_Task.TransactOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskTransactorSession) Register(username string, passwd string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, username, passwd)
}

func (_Task *TaskTransactorSession) AsyncRegister(handler func(*types.Receipt, error), username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.AsyncRegister(handler, &_Task.TransactOpts, username, passwd)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactor) Take(opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "take", worker, passwd, taskID)
}

func (_Task *TaskTransactor) AsyncTake(handler func(*types.Receipt, error), opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "take", worker, passwd, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskSession) Take(worker string, passwd string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, worker, passwd, taskID)
}

func (_Task *TaskSession) AsyncTake(handler func(*types.Receipt, error), worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTake(handler, &_Task.TransactOpts, worker, passwd, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Take(worker string, passwd string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, worker, passwd, taskID)
}

func (_Task *TaskTransactorSession) AsyncTake(handler func(*types.Receipt, error), worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTake(handler, &_Task.TransactOpts, worker, passwd, taskID)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskTransactor) Transfer(opts *bind.TransactOpts, _from string, _to string, _value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "transfer", _from, _to, _value)
}

func (_Task *TaskTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, _from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "transfer", _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, _from, _to, _value)
}

func (_Task *TaskSession) AsyncTransfer(handler func(*types.Receipt, error), _from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTransfer(handler, &_Task.TransactOpts, _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskTransactorSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, _from, _to, _value)
}

func (_Task *TaskTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), _from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTransfer(handler, &_Task.TransactOpts, _from, _to, _value)
}
