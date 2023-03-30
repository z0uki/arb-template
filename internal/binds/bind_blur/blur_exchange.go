// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind_blur

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BlurExchangeMetaData contains all meta data concerning the BlurExchange contract.
var BlurExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BlurExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use BlurExchangeMetaData.ABI instead.
var BlurExchangeABI = BlurExchangeMetaData.ABI

// BlurExchange is an auto generated Go binding around an Ethereum contract.
type BlurExchange struct {
	BlurExchangeCaller     // Read-only binding to the contract
	BlurExchangeTransactor // Write-only binding to the contract
	BlurExchangeFilterer   // Log filterer for contract events
}

// BlurExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlurExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlurExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlurExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlurExchangeSession struct {
	Contract     *BlurExchange     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlurExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlurExchangeCallerSession struct {
	Contract *BlurExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BlurExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlurExchangeTransactorSession struct {
	Contract     *BlurExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BlurExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlurExchangeRaw struct {
	Contract *BlurExchange // Generic contract binding to access the raw methods on
}

// BlurExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlurExchangeCallerRaw struct {
	Contract *BlurExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// BlurExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlurExchangeTransactorRaw struct {
	Contract *BlurExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlurExchange creates a new instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchange(address common.Address, backend bind.ContractBackend) (*BlurExchange, error) {
	contract, err := bindBlurExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlurExchange{BlurExchangeCaller: BlurExchangeCaller{contract: contract}, BlurExchangeTransactor: BlurExchangeTransactor{contract: contract}, BlurExchangeFilterer: BlurExchangeFilterer{contract: contract}}, nil
}

// NewBlurExchangeCaller creates a new read-only instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchangeCaller(address common.Address, caller bind.ContractCaller) (*BlurExchangeCaller, error) {
	contract, err := bindBlurExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeCaller{contract: contract}, nil
}

// NewBlurExchangeTransactor creates a new write-only instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*BlurExchangeTransactor, error) {
	contract, err := bindBlurExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeTransactor{contract: contract}, nil
}

// NewBlurExchangeFilterer creates a new log filterer instance of BlurExchange, bound to a specific deployed contract.
func NewBlurExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*BlurExchangeFilterer, error) {
	contract, err := bindBlurExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeFilterer{contract: contract}, nil
}

// bindBlurExchange binds a generic wrapper to an already deployed contract.
func bindBlurExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlurExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurExchange *BlurExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurExchange.Contract.BlurExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurExchange *BlurExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.Contract.BlurExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurExchange *BlurExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurExchange.Contract.BlurExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurExchange *BlurExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurExchange *BlurExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurExchange *BlurExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurExchange.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BlurExchange *BlurExchangeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BlurExchange.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BlurExchange *BlurExchangeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BlurExchange.Contract.Fallback(&_BlurExchange.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BlurExchange *BlurExchangeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BlurExchange.Contract.Fallback(&_BlurExchange.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BlurExchange *BlurExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BlurExchange *BlurExchangeSession) Receive() (*types.Transaction, error) {
	return _BlurExchange.Contract.Receive(&_BlurExchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BlurExchange *BlurExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _BlurExchange.Contract.Receive(&_BlurExchange.TransactOpts)
}

// BlurExchangeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BlurExchange contract.
type BlurExchangeAdminChangedIterator struct {
	Event *BlurExchangeAdminChanged // Event containing the contract specifics and raw log

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
func (it *BlurExchangeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeAdminChanged)
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
		it.Event = new(BlurExchangeAdminChanged)
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
func (it *BlurExchangeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeAdminChanged represents a AdminChanged event raised by the BlurExchange contract.
type BlurExchangeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchange *BlurExchangeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BlurExchangeAdminChangedIterator, error) {

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeAdminChangedIterator{contract: _BlurExchange.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchange *BlurExchangeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BlurExchangeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeAdminChanged)
				if err := _BlurExchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchange *BlurExchangeFilterer) ParseAdminChanged(log types.Log) (*BlurExchangeAdminChanged, error) {
	event := new(BlurExchangeAdminChanged)
	if err := _BlurExchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BlurExchange contract.
type BlurExchangeBeaconUpgradedIterator struct {
	Event *BlurExchangeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurExchangeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeBeaconUpgraded)
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
		it.Event = new(BlurExchangeBeaconUpgraded)
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
func (it *BlurExchangeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeBeaconUpgraded represents a BeaconUpgraded event raised by the BlurExchange contract.
type BlurExchangeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchange *BlurExchangeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BlurExchangeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeBeaconUpgradedIterator{contract: _BlurExchange.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchange *BlurExchangeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BlurExchangeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeBeaconUpgraded)
				if err := _BlurExchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchange *BlurExchangeFilterer) ParseBeaconUpgraded(log types.Log) (*BlurExchangeBeaconUpgraded, error) {
	event := new(BlurExchangeBeaconUpgraded)
	if err := _BlurExchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BlurExchange contract.
type BlurExchangeUpgradedIterator struct {
	Event *BlurExchangeUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurExchangeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeUpgraded)
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
		it.Event = new(BlurExchangeUpgraded)
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
func (it *BlurExchangeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeUpgraded represents a Upgraded event raised by the BlurExchange contract.
type BlurExchangeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchange *BlurExchangeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BlurExchangeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlurExchange.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeUpgradedIterator{contract: _BlurExchange.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchange *BlurExchangeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BlurExchangeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlurExchange.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeUpgraded)
				if err := _BlurExchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchange *BlurExchangeFilterer) ParseUpgraded(log types.Log) (*BlurExchangeUpgraded, error) {
	event := new(BlurExchangeUpgraded)
	if err := _BlurExchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
