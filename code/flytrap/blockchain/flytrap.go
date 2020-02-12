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

// FlytrapABI is the input ABI used to generate the binding from.
const FlytrapABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumFlytrap.ActionPerformed\",\"name\":\"_action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"ACLChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addPub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"}],\"name\":\"addTopic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"revokePub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"revokeSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"topics\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"verifyPub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"verifySub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FlytrapFuncSigs maps the 4-byte function signature to its string representation.
var FlytrapFuncSigs = map[string]string{
	"37bb1812": "addPub(address,bytes32)",
	"c6c460f7": "addSub(address,bytes32)",
	"0c921708": "addTopic(bytes32,uint256,uint256)",
	"8b99d3f8": "revokePub(address,bytes32)",
	"16ce2894": "revokeSub(address,bytes32)",
	"0f2fbeec": "topics(bytes32)",
	"7d1bae0a": "verifyPub(address,bytes32)",
	"d1d4a771": "verifySub(address,bytes32)",
}

// FlytrapBin is the compiled bytecode used for deploying new contracts.
var FlytrapBin = "0x608060405234801561001057600080fd5b506040516106313803806106318339818101604052602081101561003357600080fd5b5051600080546001600160a01b031916331790556001556105d8806100596000396000f3fe60806040526004361061007b5760003560e01c80637d1bae0a1161004e5780637d1bae0a1461017d5780638b99d3f8146101ca578063c6c460f714610203578063d1d4a7711461023c5761007b565b80630c921708146100805780630f2fbeec146100ab57806316ce28941461010b57806337bb181214610144575b600080fd5b6100a96004803603606081101561009657600080fd5b5080359060208101359060400135610275565b005b3480156100b757600080fd5b506100d5600480360360208110156100ce57600080fd5b503561035c565b60408051951515865260208601949094528484019290925260608401526001600160a01b03166080830152519081900360a00190f35b34801561011757600080fd5b506100a96004803603604081101561012e57600080fd5b506001600160a01b038135169060200135610398565b34801561015057600080fd5b506100a96004803603604081101561016757600080fd5b506001600160a01b038135169060200135610432565b34801561018957600080fd5b506101b6600480360360408110156101a057600080fd5b506001600160a01b038135169060200135610490565b604080519115158252519081900360200190f35b3480156101d657600080fd5b506100a9600480360360408110156101ed57600080fd5b506001600160a01b0381351690602001356104be565b34801561020f57600080fd5b506100a96004803603604081101561022657600080fd5b506001600160a01b038135169060200135610518565b34801561024857600080fd5b506101b66004803603604081101561025f57600080fd5b506001600160a01b038135169060200135610574565b60008381526002602052604090205460ff161561029157600080fd5b6001543410156102a057600080fd5b6040805160a0810182526001808252602080830187815283850187815260608501878152336080870190815260008b81526002958690528881209751885460ff1916901515178855935195870195909555905192850192909255905160038401559051600490920180546001600160a01b0319166001600160a01b039384161790558054925192909116916108fc3480159190910292909190818181858888f19350505050158015610356573d6000803e3d6000fd5b50505050565b60026020819052600091825260409091208054600182015492820154600383015460049093015460ff909216939290916001600160a01b031685565b6000818152600260205260409020600401546001600160a01b031633146103be57600080fd5b60008181526002602090815260408083206001600160a01b03861684526006019091529020805460ff1916905560045b6040805183815290516001600160a01b0385169133917f886837a844a399b455af6a2285c7475dffd9add103ab5cf83a65461ce1dcb07a9181900360200190a45050565b6000818152600260205260409020600401546001600160a01b0316331461045857600080fd5b60008181526002602090815260408083206001600160a01b03861684526005019091529020805460ff191660019081179091556103ee565b60009081526002602090815260408083206001600160a01b0394909416835260059093019052205460ff1690565b6000818152600260205260409020600401546001600160a01b031633146104e457600080fd5b60008181526002602090815260408083206001600160a01b03861684526005019091529020805460ff1916905560036103ee565b6000818152600260205260409020600401546001600160a01b0316331461053e57600080fd5b60008181526002602081815260408084206001600160a01b0387168552600601909152909120805460ff191660011790556103ee565b60009081526002602090815260408083206001600160a01b0394909416835260069093019052205460ff169056fea26469706673582212206bbdccf395d7a77b6ad7575d6d47cd729d5db7652ea758b4776d5f5eaaa9ffae64736f6c63430006020033"

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

// FlytrapACLChangeIterator is returned from FilterACLChange and is used to iterate over the raw logs and unpacked data for ACLChange events raised by the Flytrap contract.
type FlytrapACLChangeIterator struct {
	Event *FlytrapACLChange // Event containing the contract specifics and raw log

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
func (it *FlytrapACLChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlytrapACLChange)
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
		it.Event = new(FlytrapACLChange)
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
func (it *FlytrapACLChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlytrapACLChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlytrapACLChange represents a ACLChange event raised by the Flytrap contract.
type FlytrapACLChange struct {
	Src    common.Address
	Target common.Address
	Action uint8
	Name   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterACLChange is a free log retrieval operation binding the contract event 0x886837a844a399b455af6a2285c7475dffd9add103ab5cf83a65461ce1dcb07a.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, bytes32 _name)
func (_Flytrap *FlytrapFilterer) FilterACLChange(opts *bind.FilterOpts, _src []common.Address, _target []common.Address, _action []uint8) (*FlytrapACLChangeIterator, error) {

	var _srcRule []interface{}
	for _, _srcItem := range _src {
		_srcRule = append(_srcRule, _srcItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _actionRule []interface{}
	for _, _actionItem := range _action {
		_actionRule = append(_actionRule, _actionItem)
	}

	logs, sub, err := _Flytrap.contract.FilterLogs(opts, "ACLChange", _srcRule, _targetRule, _actionRule)
	if err != nil {
		return nil, err
	}
	return &FlytrapACLChangeIterator{contract: _Flytrap.contract, event: "ACLChange", logs: logs, sub: sub}, nil
}

// WatchACLChange is a free log subscription operation binding the contract event 0x886837a844a399b455af6a2285c7475dffd9add103ab5cf83a65461ce1dcb07a.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, bytes32 _name)
func (_Flytrap *FlytrapFilterer) WatchACLChange(opts *bind.WatchOpts, sink chan<- *FlytrapACLChange, _src []common.Address, _target []common.Address, _action []uint8) (event.Subscription, error) {

	var _srcRule []interface{}
	for _, _srcItem := range _src {
		_srcRule = append(_srcRule, _srcItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _actionRule []interface{}
	for _, _actionItem := range _action {
		_actionRule = append(_actionRule, _actionItem)
	}

	logs, sub, err := _Flytrap.contract.WatchLogs(opts, "ACLChange", _srcRule, _targetRule, _actionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlytrapACLChange)
				if err := _Flytrap.contract.UnpackLog(event, "ACLChange", log); err != nil {
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

// ParseACLChange is a log parse operation binding the contract event 0x886837a844a399b455af6a2285c7475dffd9add103ab5cf83a65461ce1dcb07a.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, bytes32 _name)
func (_Flytrap *FlytrapFilterer) ParseACLChange(log types.Log) (*FlytrapACLChange, error) {
	event := new(FlytrapACLChange)
	if err := _Flytrap.contract.UnpackLog(event, "ACLChange", log); err != nil {
		return nil, err
	}
	return event, nil
}
