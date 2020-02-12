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
const AuthorizerABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"}],\"name\":\"retrieveKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AuthorizerFuncSigs maps the 4-byte function signature to its string representation.
var AuthorizerFuncSigs = map[string]string{
	"82d38954": "addresses(address)",
	"2ded0fb5": "registerToken(bytes32,address)",
	"52437c20": "retrieveKey(bytes32)",
	"904194a3": "tokens(bytes32)",
}

// AuthorizerBin is the compiled bytecode used for deploying new contracts.
var AuthorizerBin = "0x608060405234801561001057600080fd5b506040516102763803806102768339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610211806100656000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632ded0fb51461005157806352437c201461007f57806382d38954146100b8578063904194a3146100f2575b600080fd5b61007d6004803603604081101561006757600080fd5b50803590602001356001600160a01b031661010f565b005b61009c6004803603602081101561009557600080fd5b5035610190565b604080516001600160a01b039092168252519081900360200190f35b6100de600480360360208110156100ce57600080fd5b50356001600160a01b03166101ab565b604080519115158252519081900360200190f35b61009c6004803603602081101561010857600080fd5b50356101c0565b6000546001600160a01b0316331461012657600080fd5b6001600160a01b03811660009081526002602052604090205460ff161561014c57600080fd5b600091825260016020818152604080852080546001600160a01b039095166001600160a01b031990951685179055928452600290529120805460ff19169091179055565b6000908152600160205260409020546001600160a01b031690565b60026020526000908152604090205460ff1681565b6001602052600090815260409020546001600160a01b03168156fea2646970667358221220e449703bdef53dc4eb777e5439641fbcaaf0ccda4001c793ad5eaf94acb29b1d64736f6c63430006020033"

// DeployAuthorizer deploys a new Ethereum contract, binding an instance of Authorizer to it.
func DeployAuthorizer(auth *bind.TransactOpts, backend bind.ContractBackend, sender common.Address) (common.Address, *types.Transaction, *Authorizer, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthorizerBin), backend, sender)
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

// FlytrapABI is the input ABI used to generate the binding from.
const FlytrapABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addPub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"}],\"name\":\"addTopic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizer\",\"outputs\":[{\"internalType\":\"contractAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"revokePub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"revokeSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"topics\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"verifyPub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"verifySub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FlytrapFuncSigs maps the 4-byte function signature to its string representation.
var FlytrapFuncSigs = map[string]string{
	"37bb1812": "addPub(address,bytes32)",
	"c6c460f7": "addSub(address,bytes32)",
	"0c921708": "addTopic(bytes32,uint256,uint256)",
	"d09edf31": "authorizer()",
	"8b99d3f8": "revokePub(address,bytes32)",
	"16ce2894": "revokeSub(address,bytes32)",
	"0f2fbeec": "topics(bytes32)",
	"7d1bae0a": "verifyPub(address,bytes32)",
	"d1d4a771": "verifySub(address,bytes32)",
}

// FlytrapBin is the compiled bytecode used for deploying new contracts.
var FlytrapBin = "0x608060405234801561001057600080fd5b506040516109143803806109148339818101604052602081101561003357600080fd5b5051600080546001600160a01b03191633908117909155600282905560405161005b906100af565b6001600160a01b03909116815260405190819003602001906000f080158015610088573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b0392909216919091179055506100bc565b6102768061069e83390190565b6105d3806100cb6000396000f3fe6080604052600436106100865760003560e01c80637d1bae0a116100595780637d1bae0a146101885780638b99d3f8146101d5578063c6c460f71461020e578063d09edf3114610247578063d1d4a7711461027857610086565b80630c9217081461008b5780630f2fbeec146100b657806316ce28941461011657806337bb18121461014f575b600080fd5b6100b4600480360360608110156100a157600080fd5b50803590602081013590604001356102b1565b005b3480156100c257600080fd5b506100e0600480360360208110156100d957600080fd5b5035610397565b60408051951515865260208601949094528484019290925260608401526001600160a01b03166080830152519081900360a00190f35b34801561012257600080fd5b506100b46004803603604081101561013957600080fd5b506001600160a01b0381351690602001356103d4565b34801561015b57600080fd5b506100b46004803603604081101561017257600080fd5b506001600160a01b03813516906020013561042a565b34801561019457600080fd5b506101c1600480360360408110156101ab57600080fd5b506001600160a01b038135169060200135610483565b604080519115158252519081900360200190f35b3480156101e157600080fd5b506100b4600480360360408110156101f857600080fd5b506001600160a01b0381351690602001356104b1565b34801561021a57600080fd5b506100b46004803603604081101561023157600080fd5b506001600160a01b038135169060200135610507565b34801561025357600080fd5b5061025c610560565b604080516001600160a01b039092168252519081900360200190f35b34801561028457600080fd5b506101c16004803603604081101561029b57600080fd5b506001600160a01b03813516906020013561056f565b60008381526003602052604090205460ff16156102cd57600080fd5b6002543410156102dc57600080fd5b6040805160a0810182526001808252602080830187815283850187815260608501878152336080870190815260008b81526003958690528881209751885460ff19169015151788559351958701959095559051600286015551918401919091559051600490920180546001600160a01b0319166001600160a01b039384161790558054925192909116916108fc3480159190910292909190818181858888f19350505050158015610391573d6000803e3d6000fd5b50505050565b60036020819052600091825260409091208054600182015460028301549383015460049093015460ff90921693909290916001600160a01b031685565b6000818152600360205260409020600401546001600160a01b031633146103fa57600080fd5b60009081526003602090815260408083206001600160a01b0390941683526006909301905220805460ff19169055565b6000818152600360205260409020600401546001600160a01b0316331461045057600080fd5b60009081526003602090815260408083206001600160a01b0390941683526005909301905220805460ff19166001179055565b60009081526003602090815260408083206001600160a01b0394909416835260059093019052205460ff1690565b6000818152600360205260409020600401546001600160a01b031633146104d757600080fd5b60009081526003602090815260408083206001600160a01b0390941683526005909301905220805460ff19169055565b6000818152600360205260409020600401546001600160a01b0316331461052d57600080fd5b60009081526003602090815260408083206001600160a01b0390941683526006909301905220805460ff19166001179055565b6001546001600160a01b031681565b60009081526003602090815260408083206001600160a01b0394909416835260069093019052205460ff169056fea2646970667358221220feb94f44334b144936f56a1c23f4cf6d9ff800093f66845c3850d09379a73a9264736f6c63430006020033608060405234801561001057600080fd5b506040516102763803806102768339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610211806100656000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632ded0fb51461005157806352437c201461007f57806382d38954146100b8578063904194a3146100f2575b600080fd5b61007d6004803603604081101561006757600080fd5b50803590602001356001600160a01b031661010f565b005b61009c6004803603602081101561009557600080fd5b5035610190565b604080516001600160a01b039092168252519081900360200190f35b6100de600480360360208110156100ce57600080fd5b50356001600160a01b03166101ab565b604080519115158252519081900360200190f35b61009c6004803603602081101561010857600080fd5b50356101c0565b6000546001600160a01b0316331461012657600080fd5b6001600160a01b03811660009081526002602052604090205460ff161561014c57600080fd5b600091825260016020818152604080852080546001600160a01b039095166001600160a01b031990951685179055928452600290529120805460ff19169091179055565b6000908152600160205260409020546001600160a01b031690565b60026020526000908152604090205460ff1681565b6001602052600090815260409020546001600160a01b03168156fea2646970667358221220e449703bdef53dc4eb777e5439641fbcaaf0ccda4001c793ad5eaf94acb29b1d64736f6c63430006020033"

// DeployFlytrap deploys a new Ethereum contract, binding an instance of Flytrap to it.
func DeployFlytrap(auth *bind.TransactOpts, backend bind.ContractBackend, cost *big.Int) (common.Address, *types.Transaction, *Flytrap, error) {
	parsed, err := abi.JSON(strings.NewReader(FlytrapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FlytrapBin), backend, cost)
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

// Authorizer is a free data retrieval call binding the contract method 0xd09edf31.
//
// Solidity: function authorizer() constant returns(address)
func (_Flytrap *FlytrapCaller) Authorizer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Flytrap.contract.Call(opts, out, "authorizer")
	return *ret0, err
}

// Authorizer is a free data retrieval call binding the contract method 0xd09edf31.
//
// Solidity: function authorizer() constant returns(address)
func (_Flytrap *FlytrapSession) Authorizer() (common.Address, error) {
	return _Flytrap.Contract.Authorizer(&_Flytrap.CallOpts)
}

// Authorizer is a free data retrieval call binding the contract method 0xd09edf31.
//
// Solidity: function authorizer() constant returns(address)
func (_Flytrap *FlytrapCallerSession) Authorizer() (common.Address, error) {
	return _Flytrap.Contract.Authorizer(&_Flytrap.CallOpts)
}

// Topics is a free data retrieval call binding the contract method 0x0f2fbeec.
//
// Solidity: function topics(bytes32 ) constant returns(bool isValue, bytes32 name, uint256 addPubCost, uint256 addSubCost, address owner)
func (_Flytrap *FlytrapCaller) Topics(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsValue    bool
	Name       [32]byte
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
}, error) {
	ret := new(struct {
		IsValue    bool
		Name       [32]byte
		AddPubCost *big.Int
		AddSubCost *big.Int
		Owner      common.Address
	})
	out := ret
	err := _Flytrap.contract.Call(opts, out, "topics", arg0)
	return *ret, err
}

// Topics is a free data retrieval call binding the contract method 0x0f2fbeec.
//
// Solidity: function topics(bytes32 ) constant returns(bool isValue, bytes32 name, uint256 addPubCost, uint256 addSubCost, address owner)
func (_Flytrap *FlytrapSession) Topics(arg0 [32]byte) (struct {
	IsValue    bool
	Name       [32]byte
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
}, error) {
	return _Flytrap.Contract.Topics(&_Flytrap.CallOpts, arg0)
}

// Topics is a free data retrieval call binding the contract method 0x0f2fbeec.
//
// Solidity: function topics(bytes32 ) constant returns(bool isValue, bytes32 name, uint256 addPubCost, uint256 addSubCost, address owner)
func (_Flytrap *FlytrapCallerSession) Topics(arg0 [32]byte) (struct {
	IsValue    bool
	Name       [32]byte
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
}, error) {
	return _Flytrap.Contract.Topics(&_Flytrap.CallOpts, arg0)
}

// VerifyPub is a free data retrieval call binding the contract method 0x7d1bae0a.
//
// Solidity: function verifyPub(address person, bytes32 topic) constant returns(bool)
func (_Flytrap *FlytrapCaller) VerifyPub(opts *bind.CallOpts, person common.Address, topic [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Flytrap.contract.Call(opts, out, "verifyPub", person, topic)
	return *ret0, err
}

// VerifyPub is a free data retrieval call binding the contract method 0x7d1bae0a.
//
// Solidity: function verifyPub(address person, bytes32 topic) constant returns(bool)
func (_Flytrap *FlytrapSession) VerifyPub(person common.Address, topic [32]byte) (bool, error) {
	return _Flytrap.Contract.VerifyPub(&_Flytrap.CallOpts, person, topic)
}

// VerifyPub is a free data retrieval call binding the contract method 0x7d1bae0a.
//
// Solidity: function verifyPub(address person, bytes32 topic) constant returns(bool)
func (_Flytrap *FlytrapCallerSession) VerifyPub(person common.Address, topic [32]byte) (bool, error) {
	return _Flytrap.Contract.VerifyPub(&_Flytrap.CallOpts, person, topic)
}

// VerifySub is a free data retrieval call binding the contract method 0xd1d4a771.
//
// Solidity: function verifySub(address person, bytes32 topic) constant returns(bool)
func (_Flytrap *FlytrapCaller) VerifySub(opts *bind.CallOpts, person common.Address, topic [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Flytrap.contract.Call(opts, out, "verifySub", person, topic)
	return *ret0, err
}

// VerifySub is a free data retrieval call binding the contract method 0xd1d4a771.
//
// Solidity: function verifySub(address person, bytes32 topic) constant returns(bool)
func (_Flytrap *FlytrapSession) VerifySub(person common.Address, topic [32]byte) (bool, error) {
	return _Flytrap.Contract.VerifySub(&_Flytrap.CallOpts, person, topic)
}

// VerifySub is a free data retrieval call binding the contract method 0xd1d4a771.
//
// Solidity: function verifySub(address person, bytes32 topic) constant returns(bool)
func (_Flytrap *FlytrapCallerSession) VerifySub(person common.Address, topic [32]byte) (bool, error) {
	return _Flytrap.Contract.VerifySub(&_Flytrap.CallOpts, person, topic)
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

// AddTopic is a paid mutator transaction binding the contract method 0x0c921708.
//
// Solidity: function addTopic(bytes32 topic, uint256 addPubCost, uint256 addSubCost) returns()
func (_Flytrap *FlytrapTransactor) AddTopic(opts *bind.TransactOpts, topic [32]byte, addPubCost *big.Int, addSubCost *big.Int) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addTopic", topic, addPubCost, addSubCost)
}

// AddTopic is a paid mutator transaction binding the contract method 0x0c921708.
//
// Solidity: function addTopic(bytes32 topic, uint256 addPubCost, uint256 addSubCost) returns()
func (_Flytrap *FlytrapSession) AddTopic(topic [32]byte, addPubCost *big.Int, addSubCost *big.Int) (*types.Transaction, error) {
	return _Flytrap.Contract.AddTopic(&_Flytrap.TransactOpts, topic, addPubCost, addSubCost)
}

// AddTopic is a paid mutator transaction binding the contract method 0x0c921708.
//
// Solidity: function addTopic(bytes32 topic, uint256 addPubCost, uint256 addSubCost) returns()
func (_Flytrap *FlytrapTransactorSession) AddTopic(topic [32]byte, addPubCost *big.Int, addSubCost *big.Int) (*types.Transaction, error) {
	return _Flytrap.Contract.AddTopic(&_Flytrap.TransactOpts, topic, addPubCost, addSubCost)
}

// RevokePub is a paid mutator transaction binding the contract method 0x8b99d3f8.
//
// Solidity: function revokePub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactor) RevokePub(opts *bind.TransactOpts, person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "revokePub", person, topic)
}

// RevokePub is a paid mutator transaction binding the contract method 0x8b99d3f8.
//
// Solidity: function revokePub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapSession) RevokePub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokePub(&_Flytrap.TransactOpts, person, topic)
}

// RevokePub is a paid mutator transaction binding the contract method 0x8b99d3f8.
//
// Solidity: function revokePub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactorSession) RevokePub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokePub(&_Flytrap.TransactOpts, person, topic)
}

// RevokeSub is a paid mutator transaction binding the contract method 0x16ce2894.
//
// Solidity: function revokeSub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactor) RevokeSub(opts *bind.TransactOpts, person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "revokeSub", person, topic)
}

// RevokeSub is a paid mutator transaction binding the contract method 0x16ce2894.
//
// Solidity: function revokeSub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapSession) RevokeSub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokeSub(&_Flytrap.TransactOpts, person, topic)
}

// RevokeSub is a paid mutator transaction binding the contract method 0x16ce2894.
//
// Solidity: function revokeSub(address person, bytes32 topic) returns()
func (_Flytrap *FlytrapTransactorSession) RevokeSub(person common.Address, topic [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokeSub(&_Flytrap.TransactOpts, person, topic)
}
