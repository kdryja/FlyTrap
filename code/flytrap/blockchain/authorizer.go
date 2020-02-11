// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// AuthorizerABI is the input ABI used to generate the binding from.
const AuthorizerABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"}],\"name\":\"retrieveKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AuthorizerFuncSigs maps the 4-byte function signature to its string representation.
var AuthorizerFuncSigs = map[string]string{
	"82d38954": "addresses(address)",
	"2ded0fb5": "registerToken(bytes32,address)",
	"52437c20": "retrieveKey(bytes32)",
	"904194a3": "tokens(bytes32)",
}

// AuthorizerBin is the compiled bytecode used for deploying new contracts.
var AuthorizerBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610211806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632ded0fb51461005157806352437c201461007f57806382d38954146100b8578063904194a3146100f2575b600080fd5b61007d6004803603604081101561006757600080fd5b50803590602001356001600160a01b031661010f565b005b61009c6004803603602081101561009557600080fd5b5035610190565b604080516001600160a01b039092168252519081900360200190f35b6100de600480360360208110156100ce57600080fd5b50356001600160a01b03166101ab565b604080519115158252519081900360200190f35b61009c6004803603602081101561010857600080fd5b50356101c0565b6000546001600160a01b0316331461012657600080fd5b6001600160a01b03811660009081526002602052604090205460ff161561014c57600080fd5b600091825260016020818152604080852080546001600160a01b039095166001600160a01b031990951685179055928452600290529120805460ff19169091179055565b6000908152600160205260409020546001600160a01b031690565b60026020526000908152604090205460ff1681565b6001602052600090815260409020546001600160a01b03168156fea26469706673582212205fc8235fb666d0b2a649058f68529c710e2ba05a299599f87711a6d392eeebdb64736f6c63430006020033"

// DeployAuthorizer deploys a new Ethereum contract, binding an instance of Authorizer to it.
func DeployAuthorizer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Authorizer, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthorizerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Authorizer{AuthorizerCaller: AuthorizerCaller{contract: contract}, AuthorizerTransactor: AuthorizerTransactor{contract: contract}, AuthorizerFilterer: AuthorizerFilterer{contract: contract}}, nil
}

// Authorizer is an auto generated Go binding around an Ethereum contract.
type Authorizer struct {
	AuthorizerCaller     // Read-only binding to the contract
	AuthorizerTransactor // Write-only binding to the contract
	AuthorizerFilterer   // Log filterer for contract events
}

// AuthorizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthorizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthorizerSession struct {
	Contract     *Authorizer       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorizerCallerSession struct {
	Contract *AuthorizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AuthorizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorizerTransactorSession struct {
	Contract     *AuthorizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AuthorizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorizerRaw struct {
	Contract *Authorizer // Generic contract binding to access the raw methods on
}

// AuthorizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorizerCallerRaw struct {
	Contract *AuthorizerCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorizerTransactorRaw struct {
	Contract *AuthorizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthorizer creates a new instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizer(address common.Address, backend bind.ContractBackend) (*Authorizer, error) {
	contract, err := bindAuthorizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authorizer{AuthorizerCaller: AuthorizerCaller{contract: contract}, AuthorizerTransactor: AuthorizerTransactor{contract: contract}, AuthorizerFilterer: AuthorizerFilterer{contract: contract}}, nil
}

// NewAuthorizerCaller creates a new read-only instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizerCaller(address common.Address, caller bind.ContractCaller) (*AuthorizerCaller, error) {
	contract, err := bindAuthorizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorizerCaller{contract: contract}, nil
}

// NewAuthorizerTransactor creates a new write-only instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizerTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorizerTransactor, error) {
	contract, err := bindAuthorizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorizerTransactor{contract: contract}, nil
}

// NewAuthorizerFilterer creates a new log filterer instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizerFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthorizerFilterer, error) {
	contract, err := bindAuthorizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthorizerFilterer{contract: contract}, nil
}

// bindAuthorizer binds a generic wrapper to an already deployed contract.
func bindAuthorizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authorizer *AuthorizerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authorizer.Contract.AuthorizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authorizer *AuthorizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorizer.Contract.AuthorizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authorizer *AuthorizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authorizer.Contract.AuthorizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authorizer *AuthorizerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authorizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authorizer *AuthorizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authorizer *AuthorizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authorizer.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0x82d38954.
//
// Solidity: function addresses(address ) constant returns(bool)
func (_Authorizer *AuthorizerCaller) Addresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authorizer.contract.Call(opts, out, "addresses", arg0)
	return *ret0, err
}

// Addresses is a free data retrieval call binding the contract method 0x82d38954.
//
// Solidity: function addresses(address ) constant returns(bool)
func (_Authorizer *AuthorizerSession) Addresses(arg0 common.Address) (bool, error) {
	return _Authorizer.Contract.Addresses(&_Authorizer.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x82d38954.
//
// Solidity: function addresses(address ) constant returns(bool)
func (_Authorizer *AuthorizerCallerSession) Addresses(arg0 common.Address) (bool, error) {
	return _Authorizer.Contract.Addresses(&_Authorizer.CallOpts, arg0)
}

// RetrieveKey is a free data retrieval call binding the contract method 0x52437c20.
//
// Solidity: function retrieveKey(bytes32 token) constant returns(address)
func (_Authorizer *AuthorizerCaller) RetrieveKey(opts *bind.CallOpts, token [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Authorizer.contract.Call(opts, out, "retrieveKey", token)
	return *ret0, err
}

// RetrieveKey is a free data retrieval call binding the contract method 0x52437c20.
//
// Solidity: function retrieveKey(bytes32 token) constant returns(address)
func (_Authorizer *AuthorizerSession) RetrieveKey(token [32]byte) (common.Address, error) {
	return _Authorizer.Contract.RetrieveKey(&_Authorizer.CallOpts, token)
}

// RetrieveKey is a free data retrieval call binding the contract method 0x52437c20.
//
// Solidity: function retrieveKey(bytes32 token) constant returns(address)
func (_Authorizer *AuthorizerCallerSession) RetrieveKey(token [32]byte) (common.Address, error) {
	return _Authorizer.Contract.RetrieveKey(&_Authorizer.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x904194a3.
//
// Solidity: function tokens(bytes32 ) constant returns(address)
func (_Authorizer *AuthorizerCaller) Tokens(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Authorizer.contract.Call(opts, out, "tokens", arg0)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x904194a3.
//
// Solidity: function tokens(bytes32 ) constant returns(address)
func (_Authorizer *AuthorizerSession) Tokens(arg0 [32]byte) (common.Address, error) {
	return _Authorizer.Contract.Tokens(&_Authorizer.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x904194a3.
//
// Solidity: function tokens(bytes32 ) constant returns(address)
func (_Authorizer *AuthorizerCallerSession) Tokens(arg0 [32]byte) (common.Address, error) {
	return _Authorizer.Contract.Tokens(&_Authorizer.CallOpts, arg0)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x2ded0fb5.
//
// Solidity: function registerToken(bytes32 token, address person) returns()
func (_Authorizer *AuthorizerTransactor) RegisterToken(opts *bind.TransactOpts, token [32]byte, person common.Address) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "registerToken", token, person)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x2ded0fb5.
//
// Solidity: function registerToken(bytes32 token, address person) returns()
func (_Authorizer *AuthorizerSession) RegisterToken(token [32]byte, person common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.RegisterToken(&_Authorizer.TransactOpts, token, person)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x2ded0fb5.
//
// Solidity: function registerToken(bytes32 token, address person) returns()
func (_Authorizer *AuthorizerTransactorSession) RegisterToken(token [32]byte, person common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.RegisterToken(&_Authorizer.TransactOpts, token, person)
}
