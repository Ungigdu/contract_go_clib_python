// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SimpleDemoABI is the input ABI used to generate the binding from.
const SimpleDemoABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"map\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"updateMap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"Update\",\"type\":\"event\"}]"

// SimpleDemo is an auto generated Go binding around an Ethereum contract.
type SimpleDemo struct {
	SimpleDemoCaller     // Read-only binding to the contract
	SimpleDemoTransactor // Write-only binding to the contract
	SimpleDemoFilterer   // Log filterer for contract events
}

// SimpleDemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleDemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleDemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleDemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleDemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleDemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleDemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleDemoSession struct {
	Contract     *SimpleDemo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleDemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleDemoCallerSession struct {
	Contract *SimpleDemoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimpleDemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleDemoTransactorSession struct {
	Contract     *SimpleDemoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleDemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleDemoRaw struct {
	Contract *SimpleDemo // Generic contract binding to access the raw methods on
}

// SimpleDemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleDemoCallerRaw struct {
	Contract *SimpleDemoCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleDemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleDemoTransactorRaw struct {
	Contract *SimpleDemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleDemo creates a new instance of SimpleDemo, bound to a specific deployed contract.
func NewSimpleDemo(address common.Address, backend bind.ContractBackend) (*SimpleDemo, error) {
	contract, err := bindSimpleDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleDemo{SimpleDemoCaller: SimpleDemoCaller{contract: contract}, SimpleDemoTransactor: SimpleDemoTransactor{contract: contract}, SimpleDemoFilterer: SimpleDemoFilterer{contract: contract}}, nil
}

// NewSimpleDemoCaller creates a new read-only instance of SimpleDemo, bound to a specific deployed contract.
func NewSimpleDemoCaller(address common.Address, caller bind.ContractCaller) (*SimpleDemoCaller, error) {
	contract, err := bindSimpleDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleDemoCaller{contract: contract}, nil
}

// NewSimpleDemoTransactor creates a new write-only instance of SimpleDemo, bound to a specific deployed contract.
func NewSimpleDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleDemoTransactor, error) {
	contract, err := bindSimpleDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleDemoTransactor{contract: contract}, nil
}

// NewSimpleDemoFilterer creates a new log filterer instance of SimpleDemo, bound to a specific deployed contract.
func NewSimpleDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleDemoFilterer, error) {
	contract, err := bindSimpleDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleDemoFilterer{contract: contract}, nil
}

// bindSimpleDemo binds a generic wrapper to an already deployed contract.
func bindSimpleDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleDemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleDemo *SimpleDemoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleDemo.Contract.SimpleDemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleDemo *SimpleDemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleDemo.Contract.SimpleDemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleDemo *SimpleDemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleDemo.Contract.SimpleDemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleDemo *SimpleDemoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleDemo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleDemo *SimpleDemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleDemo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleDemo *SimpleDemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleDemo.Contract.contract.Transact(opts, method, params...)
}

// Map is a free data retrieval call binding the contract method 0x23814fc5.
//
// Solidity: function map(string ) constant returns(string)
func (_SimpleDemo *SimpleDemoCaller) Map(opts *bind.CallOpts, arg0 string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SimpleDemo.contract.Call(opts, out, "map", arg0)
	return *ret0, err
}

// Map is a free data retrieval call binding the contract method 0x23814fc5.
//
// Solidity: function map(string ) constant returns(string)
func (_SimpleDemo *SimpleDemoSession) Map(arg0 string) (string, error) {
	return _SimpleDemo.Contract.Map(&_SimpleDemo.CallOpts, arg0)
}

// Map is a free data retrieval call binding the contract method 0x23814fc5.
//
// Solidity: function map(string ) constant returns(string)
func (_SimpleDemo *SimpleDemoCallerSession) Map(arg0 string) (string, error) {
	return _SimpleDemo.Contract.Map(&_SimpleDemo.CallOpts, arg0)
}

// UpdateMap is a paid mutator transaction binding the contract method 0x915ab755.
//
// Solidity: function updateMap(string key, string value) returns()
func (_SimpleDemo *SimpleDemoTransactor) UpdateMap(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _SimpleDemo.contract.Transact(opts, "updateMap", key, value)
}

// UpdateMap is a paid mutator transaction binding the contract method 0x915ab755.
//
// Solidity: function updateMap(string key, string value) returns()
func (_SimpleDemo *SimpleDemoSession) UpdateMap(key string, value string) (*types.Transaction, error) {
	return _SimpleDemo.Contract.UpdateMap(&_SimpleDemo.TransactOpts, key, value)
}

// UpdateMap is a paid mutator transaction binding the contract method 0x915ab755.
//
// Solidity: function updateMap(string key, string value) returns()
func (_SimpleDemo *SimpleDemoTransactorSession) UpdateMap(key string, value string) (*types.Transaction, error) {
	return _SimpleDemo.Contract.UpdateMap(&_SimpleDemo.TransactOpts, key, value)
}

// SimpleDemoUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the SimpleDemo contract.
type SimpleDemoUpdateIterator struct {
	Event *SimpleDemoUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleDemoUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleDemoUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SimpleDemoUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SimpleDemoUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleDemoUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleDemoUpdate represents a Update event raised by the SimpleDemo contract.
type SimpleDemoUpdate struct {
	Key      string
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x659b5aacd6c43d9bafa4b755ebfae73517d92eef9628aec8a9b7077d90b67d57.
//
// Solidity: event Update(string key, address operator)
func (_SimpleDemo *SimpleDemoFilterer) FilterUpdate(opts *bind.FilterOpts) (*SimpleDemoUpdateIterator, error) {

	logs, sub, err := _SimpleDemo.contract.FilterLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return &SimpleDemoUpdateIterator{contract: _SimpleDemo.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x659b5aacd6c43d9bafa4b755ebfae73517d92eef9628aec8a9b7077d90b67d57.
//
// Solidity: event Update(string key, address operator)
func (_SimpleDemo *SimpleDemoFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *SimpleDemoUpdate) (event.Subscription, error) {

	logs, sub, err := _SimpleDemo.contract.WatchLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleDemoUpdate)
				if err := _SimpleDemo.contract.UnpackLog(event, "Update", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdate is a log parse operation binding the contract event 0x659b5aacd6c43d9bafa4b755ebfae73517d92eef9628aec8a9b7077d90b67d57.
//
// Solidity: event Update(string key, address operator)
func (_SimpleDemo *SimpleDemoFilterer) ParseUpdate(log types.Log) (*SimpleDemoUpdate, error) {
	event := new(SimpleDemoUpdate)
	if err := _SimpleDemo.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	return event, nil
}
