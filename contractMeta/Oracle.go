// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMeta

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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"query\",\"type\":\"string\"}],\"name\":\"QueryEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cbAddress\",\"type\":\"address\"}],\"name\":\"setCbAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506105c0806100606000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637c26192914610046578063b62a717c14610076578063c281d19e14610092575b600080fd5b610060600480360381019061005b919061033b565b6100b0565b60405161006d919061039d565b60405180910390f35b610090600480360381019061008b9190610416565b61011f565b005b61009a6101bb565b6040516100a79190610452565b60405180910390f35b6000434283336040516020016100c9949392919061050e565b6040516020818303038152906040528051906020012090507f3040d97d9df6583359820eecf54620d1e08c89c5d0b4ad40b23f48cd8a7408b4818360405161011292919061055a565b60405180910390a1919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461017757600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610248826101ff565b810181811067ffffffffffffffff8211171561026757610266610210565b5b80604052505050565b600061027a6101e1565b9050610286828261023f565b919050565b600067ffffffffffffffff8211156102a6576102a5610210565b5b6102af826101ff565b9050602081019050919050565b82818337600083830152505050565b60006102de6102d98461028b565b610270565b9050828152602081018484840111156102fa576102f96101fa565b5b6103058482856102bc565b509392505050565b600082601f830112610322576103216101f5565b5b81356103328482602086016102cb565b91505092915050565b600060208284031215610351576103506101eb565b5b600082013567ffffffffffffffff81111561036f5761036e6101f0565b5b61037b8482850161030d565b91505092915050565b6000819050919050565b61039781610384565b82525050565b60006020820190506103b2600083018461038e565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103e3826103b8565b9050919050565b6103f3816103d8565b81146103fe57600080fd5b50565b600081359050610410816103ea565b92915050565b60006020828403121561042c5761042b6101eb565b5b600061043a84828501610401565b91505092915050565b61044c816103d8565b82525050565b60006020820190506104676000830184610443565b92915050565b6000819050919050565b6104808161046d565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104c05780820151818401526020810190506104a5565b838111156104cf576000848401525b50505050565b60006104e082610486565b6104ea8185610491565b93506104fa8185602086016104a2565b610503816101ff565b840191505092915050565b60006080820190506105236000830187610477565b6105306020830186610477565b818103604083015261054281856104d5565b90506105516060830184610443565b95945050505050565b600060408201905061056f600083018561038e565b818103602083015261058181846104d5565b9050939250505056fea2646970667358221220e9b5a48d97106dc99af500b3b24b303e62f5c2831ab39c67264530eb2d255dd764736f6c634300080d0033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() view returns(address)
func (_Oracle *OracleCaller) CbAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "cbAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() view returns(address)
func (_Oracle *OracleSession) CbAddress() (common.Address, error) {
	return _Oracle.Contract.CbAddress(&_Oracle.CallOpts)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() view returns(address)
func (_Oracle *OracleCallerSession) CbAddress() (common.Address, error) {
	return _Oracle.Contract.CbAddress(&_Oracle.CallOpts)
}

// Query is a paid mutator transaction binding the contract method 0x7c261929.
//
// Solidity: function query(string _query) returns(bytes32 id)
func (_Oracle *OracleTransactor) Query(opts *bind.TransactOpts, _query string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "query", _query)
}

// Query is a paid mutator transaction binding the contract method 0x7c261929.
//
// Solidity: function query(string _query) returns(bytes32 id)
func (_Oracle *OracleSession) Query(_query string) (*types.Transaction, error) {
	return _Oracle.Contract.Query(&_Oracle.TransactOpts, _query)
}

// Query is a paid mutator transaction binding the contract method 0x7c261929.
//
// Solidity: function query(string _query) returns(bytes32 id)
func (_Oracle *OracleTransactorSession) Query(_query string) (*types.Transaction, error) {
	return _Oracle.Contract.Query(&_Oracle.TransactOpts, _query)
}

// SetCbAddress is a paid mutator transaction binding the contract method 0xb62a717c.
//
// Solidity: function setCbAddress(address _cbAddress) returns()
func (_Oracle *OracleTransactor) SetCbAddress(opts *bind.TransactOpts, _cbAddress common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setCbAddress", _cbAddress)
}

// SetCbAddress is a paid mutator transaction binding the contract method 0xb62a717c.
//
// Solidity: function setCbAddress(address _cbAddress) returns()
func (_Oracle *OracleSession) SetCbAddress(_cbAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetCbAddress(&_Oracle.TransactOpts, _cbAddress)
}

// SetCbAddress is a paid mutator transaction binding the contract method 0xb62a717c.
//
// Solidity: function setCbAddress(address _cbAddress) returns()
func (_Oracle *OracleTransactorSession) SetCbAddress(_cbAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetCbAddress(&_Oracle.TransactOpts, _cbAddress)
}

// OracleQueryEventIterator is returned from FilterQueryEvent and is used to iterate over the raw logs and unpacked data for QueryEvent events raised by the Oracle contract.
type OracleQueryEventIterator struct {
	Event *OracleQueryEvent // Event containing the contract specifics and raw log

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
func (it *OracleQueryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleQueryEvent)
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
		it.Event = new(OracleQueryEvent)
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
func (it *OracleQueryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleQueryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleQueryEvent represents a QueryEvent event raised by the Oracle contract.
type OracleQueryEvent struct {
	Id    [32]byte
	Query string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterQueryEvent is a free log retrieval operation binding the contract event 0x3040d97d9df6583359820eecf54620d1e08c89c5d0b4ad40b23f48cd8a7408b4.
//
// Solidity: event QueryEvent(bytes32 id, string query)
func (_Oracle *OracleFilterer) FilterQueryEvent(opts *bind.FilterOpts) (*OracleQueryEventIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "QueryEvent")
	if err != nil {
		return nil, err
	}
	return &OracleQueryEventIterator{contract: _Oracle.contract, event: "QueryEvent", logs: logs, sub: sub}, nil
}

// WatchQueryEvent is a free log subscription operation binding the contract event 0x3040d97d9df6583359820eecf54620d1e08c89c5d0b4ad40b23f48cd8a7408b4.
//
// Solidity: event QueryEvent(bytes32 id, string query)
func (_Oracle *OracleFilterer) WatchQueryEvent(opts *bind.WatchOpts, sink chan<- *OracleQueryEvent) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "QueryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleQueryEvent)
				if err := _Oracle.contract.UnpackLog(event, "QueryEvent", log); err != nil {
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

// ParseQueryEvent is a log parse operation binding the contract event 0x3040d97d9df6583359820eecf54620d1e08c89c5d0b4ad40b23f48cd8a7408b4.
//
// Solidity: event QueryEvent(bytes32 id, string query)
func (_Oracle *OracleFilterer) ParseQueryEvent(log types.Log) (*OracleQueryEvent, error) {
	event := new(OracleQueryEvent)
	if err := _Oracle.contract.UnpackLog(event, "QueryEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
