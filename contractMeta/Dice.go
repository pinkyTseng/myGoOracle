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

// DiceMetaData contains all meta data concerning the Dice contract.
var DiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"diceResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"newMyOracleQuery\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"_callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkResult\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showOracleResolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c22380380610c22833981810160405281019061003291906101ed565b80806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b18b78dd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061010291906101ed565b600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061021a565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101ba8261018f565b9050919050565b6101ca816101af565b81146101d557600080fd5b50565b6000815190506101e7816101c1565b92915050565b6000602082840312156102035761020261018a565b5b6000610211848285016101d8565b91505092915050565b6109f9806102296000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063569f1bfc1461005c578063833e5a5c14610066578063a21d942f14610084578063cee77be9146100a2578063fb9a9f6a146100be575b600080fd5b6100646100dc565b005b61006e61019a565b60405161007b9190610585565b60405180910390f35b61008c6101c4565b6040516100999190610639565b60405180910390f35b6100bc60048036038101906100b791906107da565b6102a6565b005b6100c66103d1565b6040516100d39190610585565b60405180910390f35b7f7320fefbde31987714fd878b64cf8254ae37e7fee0ac90b214b00498cb9cae1f604051610109906108a8565b60405180910390a160006101516040518060400160405280600681526020017f302d3130303000000000000000000000000000000000000000000000000000008152506103fa565b905080600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060046000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481526020019081526020016000208054610223906108f7565b80601f016020809104026020016040519081016040528092919081815260200182805461024f906108f7565b801561029c5780601f106102715761010080835404028352916020019161029c565b820191906000526020600020905b81548152906001019060200180831161027f57829003601f168201915b5050505050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c281d19e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610313573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103379190610954565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461036e57600080fd5b806004600084815260200190815260200160002090805190602001906103959291906104a1565b507f7d78eceb19a25d6e8704a729339c24743f404e277ec977da6474308cf5790a80816040516103c59190610639565b60405180910390a15050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637c261929836040518263ffffffff1660e01b81526004016104579190610639565b6020604051808303816000875af1158015610476573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049a9190610996565b9050919050565b8280546104ad906108f7565b90600052602060002090601f0160209004810192826104cf5760008555610516565b82601f106104e857805160ff1916838001178555610516565b82800160010185558215610516579182015b828111156105155782518255916020019190600101906104fa565b5b5090506105239190610527565b5090565b5b80821115610540576000816000905550600101610528565b5090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061056f82610544565b9050919050565b61057f81610564565b82525050565b600060208201905061059a6000830184610576565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156105da5780820151818401526020810190506105bf565b838111156105e9576000848401525b50505050565b6000601f19601f8301169050919050565b600061060b826105a0565b61061581856105ab565b93506106258185602086016105bc565b61062e816105ef565b840191505092915050565b600060208201905081810360008301526106538184610600565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6106828161066f565b811461068d57600080fd5b50565b60008135905061069f81610679565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106e7826105ef565b810181811067ffffffffffffffff82111715610706576107056106af565b5b80604052505050565b600061071961065b565b905061072582826106de565b919050565b600067ffffffffffffffff821115610745576107446106af565b5b61074e826105ef565b9050602081019050919050565b82818337600083830152505050565b600061077d6107788461072a565b61070f565b905082815260208101848484011115610799576107986106aa565b5b6107a484828561075b565b509392505050565b600082601f8301126107c1576107c06106a5565b5b81356107d184826020860161076a565b91505092915050565b600080604083850312156107f1576107f0610665565b5b60006107ff85828601610690565b925050602083013567ffffffffffffffff8111156108205761081f61066a565b5b61082c858286016107ac565b9150509250929050565b7f4d794f7261636c65207175657279207761732073656e742c207374616e64696e60008201527f6720627920666f722074686520616e737765722e2e0000000000000000000000602082015250565b60006108926035836105ab565b915061089d82610836565b604082019050919050565b600060208201905081810360008301526108c181610885565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061090f57607f821691505b602082108103610922576109216108c8565b5b50919050565b61093181610564565b811461093c57600080fd5b50565b60008151905061094e81610928565b92915050565b60006020828403121561096a57610969610665565b5b60006109788482850161093f565b91505092915050565b60008151905061099081610679565b92915050565b6000602082840312156109ac576109ab610665565b5b60006109ba84828501610981565b9150509291505056fea2646970667358221220fe9284c7c061a1ddb8c2c12ab573ae0cac0483f286aa8efeee6081f1bec66cbc64736f6c634300080d0033",
}

// DiceABI is the input ABI used to generate the binding from.
// Deprecated: Use DiceMetaData.ABI instead.
var DiceABI = DiceMetaData.ABI

// DiceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiceMetaData.Bin instead.
var DiceBin = DiceMetaData.Bin

// DeployDice deploys a new Ethereum contract, binding an instance of Dice to it.
func DeployDice(auth *bind.TransactOpts, backend bind.ContractBackend, _address common.Address) (common.Address, *types.Transaction, *Dice, error) {
	parsed, err := DiceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiceBin), backend, _address)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dice{DiceCaller: DiceCaller{contract: contract}, DiceTransactor: DiceTransactor{contract: contract}, DiceFilterer: DiceFilterer{contract: contract}}, nil
}

// Dice is an auto generated Go binding around an Ethereum contract.
type Dice struct {
	DiceCaller     // Read-only binding to the contract
	DiceTransactor // Write-only binding to the contract
	DiceFilterer   // Log filterer for contract events
}

// DiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiceSession struct {
	Contract     *Dice             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiceCallerSession struct {
	Contract *DiceCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiceTransactorSession struct {
	Contract     *DiceTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiceRaw struct {
	Contract *Dice // Generic contract binding to access the raw methods on
}

// DiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiceCallerRaw struct {
	Contract *DiceCaller // Generic read-only contract binding to access the raw methods on
}

// DiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiceTransactorRaw struct {
	Contract *DiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDice creates a new instance of Dice, bound to a specific deployed contract.
func NewDice(address common.Address, backend bind.ContractBackend) (*Dice, error) {
	contract, err := bindDice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dice{DiceCaller: DiceCaller{contract: contract}, DiceTransactor: DiceTransactor{contract: contract}, DiceFilterer: DiceFilterer{contract: contract}}, nil
}

// NewDiceCaller creates a new read-only instance of Dice, bound to a specific deployed contract.
func NewDiceCaller(address common.Address, caller bind.ContractCaller) (*DiceCaller, error) {
	contract, err := bindDice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiceCaller{contract: contract}, nil
}

// NewDiceTransactor creates a new write-only instance of Dice, bound to a specific deployed contract.
func NewDiceTransactor(address common.Address, transactor bind.ContractTransactor) (*DiceTransactor, error) {
	contract, err := bindDice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiceTransactor{contract: contract}, nil
}

// NewDiceFilterer creates a new log filterer instance of Dice, bound to a specific deployed contract.
func NewDiceFilterer(address common.Address, filterer bind.ContractFilterer) (*DiceFilterer, error) {
	contract, err := bindDice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiceFilterer{contract: contract}, nil
}

// bindDice binds a generic wrapper to an already deployed contract.
func bindDice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dice *DiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dice.Contract.DiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dice *DiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dice.Contract.DiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dice *DiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dice.Contract.DiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dice *DiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dice *DiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dice *DiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dice.Contract.contract.Transact(opts, method, params...)
}

// CheckResult is a free data retrieval call binding the contract method 0xa21d942f.
//
// Solidity: function checkResult() view returns(string)
func (_Dice *DiceCaller) CheckResult(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dice.contract.Call(opts, &out, "checkResult")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CheckResult is a free data retrieval call binding the contract method 0xa21d942f.
//
// Solidity: function checkResult() view returns(string)
func (_Dice *DiceSession) CheckResult() (string, error) {
	return _Dice.Contract.CheckResult(&_Dice.CallOpts)
}

// CheckResult is a free data retrieval call binding the contract method 0xa21d942f.
//
// Solidity: function checkResult() view returns(string)
func (_Dice *DiceCallerSession) CheckResult() (string, error) {
	return _Dice.Contract.CheckResult(&_Dice.CallOpts)
}

// ShowOracle is a free data retrieval call binding the contract method 0x833e5a5c.
//
// Solidity: function showOracle() view returns(address)
func (_Dice *DiceCaller) ShowOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dice.contract.Call(opts, &out, "showOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShowOracle is a free data retrieval call binding the contract method 0x833e5a5c.
//
// Solidity: function showOracle() view returns(address)
func (_Dice *DiceSession) ShowOracle() (common.Address, error) {
	return _Dice.Contract.ShowOracle(&_Dice.CallOpts)
}

// ShowOracle is a free data retrieval call binding the contract method 0x833e5a5c.
//
// Solidity: function showOracle() view returns(address)
func (_Dice *DiceCallerSession) ShowOracle() (common.Address, error) {
	return _Dice.Contract.ShowOracle(&_Dice.CallOpts)
}

// ShowOracleResolver is a free data retrieval call binding the contract method 0xfb9a9f6a.
//
// Solidity: function showOracleResolver() view returns(address)
func (_Dice *DiceCaller) ShowOracleResolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dice.contract.Call(opts, &out, "showOracleResolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShowOracleResolver is a free data retrieval call binding the contract method 0xfb9a9f6a.
//
// Solidity: function showOracleResolver() view returns(address)
func (_Dice *DiceSession) ShowOracleResolver() (common.Address, error) {
	return _Dice.Contract.ShowOracleResolver(&_Dice.CallOpts)
}

// ShowOracleResolver is a free data retrieval call binding the contract method 0xfb9a9f6a.
//
// Solidity: function showOracleResolver() view returns(address)
func (_Dice *DiceCallerSession) ShowOracleResolver() (common.Address, error) {
	return _Dice.Contract.ShowOracleResolver(&_Dice.CallOpts)
}

// Callback is a paid mutator transaction binding the contract method 0xcee77be9.
//
// Solidity: function _callback(bytes32 _id, string result) returns()
func (_Dice *DiceTransactor) Callback(opts *bind.TransactOpts, _id [32]byte, result string) (*types.Transaction, error) {
	return _Dice.contract.Transact(opts, "_callback", _id, result)
}

// Callback is a paid mutator transaction binding the contract method 0xcee77be9.
//
// Solidity: function _callback(bytes32 _id, string result) returns()
func (_Dice *DiceSession) Callback(_id [32]byte, result string) (*types.Transaction, error) {
	return _Dice.Contract.Callback(&_Dice.TransactOpts, _id, result)
}

// Callback is a paid mutator transaction binding the contract method 0xcee77be9.
//
// Solidity: function _callback(bytes32 _id, string result) returns()
func (_Dice *DiceTransactorSession) Callback(_id [32]byte, result string) (*types.Transaction, error) {
	return _Dice.Contract.Callback(&_Dice.TransactOpts, _id, result)
}

// Dice is a paid mutator transaction binding the contract method 0x569f1bfc.
//
// Solidity: function dice() returns()
func (_Dice *DiceTransactor) Dice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dice.contract.Transact(opts, "dice")
}

// Dice is a paid mutator transaction binding the contract method 0x569f1bfc.
//
// Solidity: function dice() returns()
func (_Dice *DiceSession) Dice() (*types.Transaction, error) {
	return _Dice.Contract.Dice(&_Dice.TransactOpts)
}

// Dice is a paid mutator transaction binding the contract method 0x569f1bfc.
//
// Solidity: function dice() returns()
func (_Dice *DiceTransactorSession) Dice() (*types.Transaction, error) {
	return _Dice.Contract.Dice(&_Dice.TransactOpts)
}

// DiceDiceResultIterator is returned from FilterDiceResult and is used to iterate over the raw logs and unpacked data for DiceResult events raised by the Dice contract.
type DiceDiceResultIterator struct {
	Event *DiceDiceResult // Event containing the contract specifics and raw log

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
func (it *DiceDiceResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiceDiceResult)
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
		it.Event = new(DiceDiceResult)
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
func (it *DiceDiceResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiceDiceResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiceDiceResult represents a DiceResult event raised by the Dice contract.
type DiceDiceResult struct {
	Result string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDiceResult is a free log retrieval operation binding the contract event 0x7d78eceb19a25d6e8704a729339c24743f404e277ec977da6474308cf5790a80.
//
// Solidity: event diceResult(string result)
func (_Dice *DiceFilterer) FilterDiceResult(opts *bind.FilterOpts) (*DiceDiceResultIterator, error) {

	logs, sub, err := _Dice.contract.FilterLogs(opts, "diceResult")
	if err != nil {
		return nil, err
	}
	return &DiceDiceResultIterator{contract: _Dice.contract, event: "diceResult", logs: logs, sub: sub}, nil
}

// WatchDiceResult is a free log subscription operation binding the contract event 0x7d78eceb19a25d6e8704a729339c24743f404e277ec977da6474308cf5790a80.
//
// Solidity: event diceResult(string result)
func (_Dice *DiceFilterer) WatchDiceResult(opts *bind.WatchOpts, sink chan<- *DiceDiceResult) (event.Subscription, error) {

	logs, sub, err := _Dice.contract.WatchLogs(opts, "diceResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiceDiceResult)
				if err := _Dice.contract.UnpackLog(event, "diceResult", log); err != nil {
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

// ParseDiceResult is a log parse operation binding the contract event 0x7d78eceb19a25d6e8704a729339c24743f404e277ec977da6474308cf5790a80.
//
// Solidity: event diceResult(string result)
func (_Dice *DiceFilterer) ParseDiceResult(log types.Log) (*DiceDiceResult, error) {
	event := new(DiceDiceResult)
	if err := _Dice.contract.UnpackLog(event, "diceResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiceNewMyOracleQueryIterator is returned from FilterNewMyOracleQuery and is used to iterate over the raw logs and unpacked data for NewMyOracleQuery events raised by the Dice contract.
type DiceNewMyOracleQueryIterator struct {
	Event *DiceNewMyOracleQuery // Event containing the contract specifics and raw log

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
func (it *DiceNewMyOracleQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiceNewMyOracleQuery)
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
		it.Event = new(DiceNewMyOracleQuery)
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
func (it *DiceNewMyOracleQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiceNewMyOracleQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiceNewMyOracleQuery represents a NewMyOracleQuery event raised by the Dice contract.
type DiceNewMyOracleQuery struct {
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewMyOracleQuery is a free log retrieval operation binding the contract event 0x7320fefbde31987714fd878b64cf8254ae37e7fee0ac90b214b00498cb9cae1f.
//
// Solidity: event newMyOracleQuery(string description)
func (_Dice *DiceFilterer) FilterNewMyOracleQuery(opts *bind.FilterOpts) (*DiceNewMyOracleQueryIterator, error) {

	logs, sub, err := _Dice.contract.FilterLogs(opts, "newMyOracleQuery")
	if err != nil {
		return nil, err
	}
	return &DiceNewMyOracleQueryIterator{contract: _Dice.contract, event: "newMyOracleQuery", logs: logs, sub: sub}, nil
}

// WatchNewMyOracleQuery is a free log subscription operation binding the contract event 0x7320fefbde31987714fd878b64cf8254ae37e7fee0ac90b214b00498cb9cae1f.
//
// Solidity: event newMyOracleQuery(string description)
func (_Dice *DiceFilterer) WatchNewMyOracleQuery(opts *bind.WatchOpts, sink chan<- *DiceNewMyOracleQuery) (event.Subscription, error) {

	logs, sub, err := _Dice.contract.WatchLogs(opts, "newMyOracleQuery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiceNewMyOracleQuery)
				if err := _Dice.contract.UnpackLog(event, "newMyOracleQuery", log); err != nil {
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

// ParseNewMyOracleQuery is a log parse operation binding the contract event 0x7320fefbde31987714fd878b64cf8254ae37e7fee0ac90b214b00498cb9cae1f.
//
// Solidity: event newMyOracleQuery(string description)
func (_Dice *DiceFilterer) ParseNewMyOracleQuery(log types.Log) (*DiceNewMyOracleQuery, error) {
	event := new(DiceNewMyOracleQuery)
	if err := _Dice.contract.UnpackLog(event, "newMyOracleQuery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
