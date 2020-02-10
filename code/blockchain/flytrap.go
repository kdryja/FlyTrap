// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// FlytrapABI is the input ABI used to generate the binding from.
const FlytrapABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addPub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FlytrapFuncSigs maps the 4-byte function signature to its string representation.
var FlytrapFuncSigs = map[string]string{
	"70480275": "addAdmin(address)",
	"37bb1812": "addPub(address,bytes32)",
	"c6c460f7": "addSub(address,bytes32)",
	"429b62e5": "admins(address)",
	"254be18a": "pub(address,bytes32)",
	"aa5e259e": "sub(address,bytes32)",
}

// FlytrapBin is the compiled bytecode used for deploying new contracts.
var FlytrapBin = "0x608060405234801561001057600080fd5b50336000908152602081905260409020805460ff191660011790556102ed8061003a6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063254be18a1461006757806337bb1812146100a7578063429b62e5146100d557806370480275146100fb578063aa5e259e14610121578063c6c460f71461014d575b600080fd5b6100936004803603604081101561007d57600080fd5b506001600160a01b038135169060200135610179565b604080519115158252519081900360200190f35b6100d3600480360360408110156100bd57600080fd5b506001600160a01b038135169060200135610199565b005b610093600480360360208110156100eb57600080fd5b50356001600160a01b03166101ec565b6100d36004803603602081101561011157600080fd5b50356001600160a01b0316610201565b6100936004803603604081101561013757600080fd5b506001600160a01b038135169060200135610246565b6100d36004803603604081101561016357600080fd5b506001600160a01b038135169060200135610266565b600160209081526000928352604080842090915290825290205460ff1681565b3360009081526020819052604090205460ff1615156001146101ba57600080fd5b6001600160a01b039091166000908152600160208181526040808420948452939052919020805460ff19169091179055565b60006020819052908152604090205460ff1681565b3360009081526020819052604090205460ff16151560011461022257600080fd5b6001600160a01b03166000908152602081905260409020805460ff19166001179055565b600260209081526000928352604080842090915290825290205460ff1681565b3360009081526020819052604090205460ff16151560011461028757600080fd5b6001600160a01b03909116600090815260026020908152604080832093835292905220805460ff1916600117905556fea26469706673582212203a0a4a4491739981faebe00c17b5801e32101e5fdd59ac236acc6951e9f2f45d64736f6c63430006020033"

// DeployFlytrap deploys a new Ethereum contract, binding an instance of Flytrap to it.
func DeployFlytrap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Flytrap, error) {
	parsed, err := abi.JSON(strings.NewReader(FlytrapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FlytrapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Flytrap{FlytrapCaller: FlytrapCaller{contract: contract}, FlytrapTransactor: FlytrapTransactor{contract: contract}, FlytrapFilterer: FlytrapFilterer{contract: contract}}, nil
}

// Flytrap is an auto generated Go binding around an Ethereum contract.
type Flytrap struct {
	FlytrapCaller     // Read-only binding to the contract
	FlytrapTransactor // Write-only binding to the contract
	FlytrapFilterer   // Log filterer for contract events
}

// FlytrapCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlytrapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlytrapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlytrapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlytrapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlytrapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlytrapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlytrapSession struct {
	Contract     *Flytrap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlytrapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlytrapCallerSession struct {
	Contract *FlytrapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FlytrapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlytrapTransactorSession struct {
	Contract     *FlytrapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FlytrapRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlytrapRaw struct {
	Contract *Flytrap // Generic contract binding to access the raw methods on
}

// FlytrapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlytrapCallerRaw struct {
	Contract *FlytrapCaller // Generic read-only contract binding to access the raw methods on
}

// FlytrapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlytrapTransactorRaw struct {
	Contract *FlytrapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlytrap creates a new instance of Flytrap, bound to a specific deployed contract.
func NewFlytrap(address common.Address, backend bind.ContractBackend) (*Flytrap, error) {
	contract, err := bindFlytrap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flytrap{FlytrapCaller: FlytrapCaller{contract: contract}, FlytrapTransactor: FlytrapTransactor{contract: contract}, FlytrapFilterer: FlytrapFilterer{contract: contract}}, nil
}

// NewFlytrapCaller creates a new read-only instance of Flytrap, bound to a specific deployed contract.
func NewFlytrapCaller(address common.Address, caller bind.ContractCaller) (*FlytrapCaller, error) {
	contract, err := bindFlytrap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlytrapCaller{contract: contract}, nil
}

// NewFlytrapTransactor creates a new write-only instance of Flytrap, bound to a specific deployed contract.
func NewFlytrapTransactor(address common.Address, transactor bind.ContractTransactor) (*FlytrapTransactor, error) {
	contract, err := bindFlytrap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlytrapTransactor{contract: contract}, nil
}

// NewFlytrapFilterer creates a new log filterer instance of Flytrap, bound to a specific deployed contract.
func NewFlytrapFilterer(address common.Address, filterer bind.ContractFilterer) (*FlytrapFilterer, error) {
	contract, err := bindFlytrap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlytrapFilterer{contract: contract}, nil
}

// bindFlytrap binds a generic wrapper to an already deployed contract.
func bindFlytrap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlytrapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flytrap *FlytrapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Flytrap.Contract.FlytrapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flytrap *FlytrapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flytrap.Contract.FlytrapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flytrap *FlytrapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flytrap.Contract.FlytrapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flytrap *FlytrapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Flytrap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flytrap *FlytrapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flytrap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flytrap *FlytrapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flytrap.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) constant returns(bool)
func (_Flytrap *FlytrapCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Flytrap.contract.Call(opts, out, "admins", arg0)
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) constant returns(bool)
func (_Flytrap *FlytrapSession) Admins(arg0 common.Address) (bool, error) {
	return _Flytrap.Contract.Admins(&_Flytrap.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) constant returns(bool)
func (_Flytrap *FlytrapCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _Flytrap.Contract.Admins(&_Flytrap.CallOpts, arg0)
}

// Pub is a free data retrieval call binding the contract method 0x254be18a.
//
// Solidity: function pub(address , bytes32 ) constant returns(bool)
func (_Flytrap *FlytrapCaller) Pub(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Flytrap.contract.Call(opts, out, "pub", arg0, arg1)
	return *ret0, err
}

// Pub is a free data retrieval call binding the contract method 0x254be18a.
//
// Solidity: function pub(address , bytes32 ) constant returns(bool)
func (_Flytrap *FlytrapSession) Pub(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Flytrap.Contract.Pub(&_Flytrap.CallOpts, arg0, arg1)
}

// Pub is a free data retrieval call binding the contract method 0x254be18a.
//
// Solidity: function pub(address , bytes32 ) constant returns(bool)
func (_Flytrap *FlytrapCallerSession) Pub(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Flytrap.Contract.Pub(&_Flytrap.CallOpts, arg0, arg1)
}

// Sub is a free data retrieval call binding the contract method 0xaa5e259e.
//
// Solidity: function sub(address , bytes32 ) constant returns(bool)
func (_Flytrap *FlytrapCaller) Sub(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Flytrap.contract.Call(opts, out, "sub", arg0, arg1)
	return *ret0, err
}

// Sub is a free data retrieval call binding the contract method 0xaa5e259e.
//
// Solidity: function sub(address , bytes32 ) constant returns(bool)
func (_Flytrap *FlytrapSession) Sub(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Flytrap.Contract.Sub(&_Flytrap.CallOpts, arg0, arg1)
}

// Sub is a free data retrieval call binding the contract method 0xaa5e259e.
//
// Solidity: function sub(address , bytes32 ) constant returns(bool)
func (_Flytrap *FlytrapCallerSession) Sub(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Flytrap.Contract.Sub(&_Flytrap.CallOpts, arg0, arg1)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address person) returns()
func (_Flytrap *FlytrapTransactor) AddAdmin(opts *bind.TransactOpts, person common.Address) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addAdmin", person)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address person) returns()
func (_Flytrap *FlytrapSession) AddAdmin(person common.Address) (*types.Transaction, error) {
	return _Flytrap.Contract.AddAdmin(&_Flytrap.TransactOpts, person)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address person) returns()
func (_Flytrap *FlytrapTransactorSession) AddAdmin(person common.Address) (*types.Transaction, error) {
	return _Flytrap.Contract.AddAdmin(&_Flytrap.TransactOpts, person)
}

// AddPub is a paid mutator transaction binding the contract method 0x37bb1812.
//
// Solidity: function addPub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactor) AddPub(opts *bind.TransactOpts, person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addPub", person, topic)
}

// AddPub is a paid mutator transaction binding the contract method 0x37bb1812.
//
// Solidity: function addPub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapSession) AddPub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.AddPub(&_Flytrap.TransactOpts, person, topic)
}

// AddPub is a paid mutator transaction binding the contract method 0x37bb1812.
//
// Solidity: function addPub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactorSession) AddPub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.AddPub(&_Flytrap.TransactOpts, person, topic)
}

// AddSub is a paid mutator transaction binding the contract method 0xc6c460f7.
//
// Solidity: function addSub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactor) AddSub(opts *bind.TransactOpts, person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addSub", person, topic)
}

// AddSub is a paid mutator transaction binding the contract method 0xc6c460f7.
//
// Solidity: function addSub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapSession) AddSub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.AddSub(&_Flytrap.TransactOpts, person, topic)
}

// AddSub is a paid mutator transaction binding the contract method 0xc6c460f7.
//
// Solidity: function addSub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactorSession) AddSub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.AddSub(&_Flytrap.TransactOpts, person, topic)
}
