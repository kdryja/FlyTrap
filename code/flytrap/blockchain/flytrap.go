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
const FlytrapABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumFlytrap.ActionPerformed\",\"name\":\"_action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"ACLChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addPub\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"addSub\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"bytes2\",\"name\":\"country\",\"type\":\"bytes2\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"}],\"name\":\"addTopic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"enumFlytrap.ActionPerformed\",\"name\":\"alert\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"desc\",\"type\":\"bytes32\"}],\"name\":\"logAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"revokePub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"revokeSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"topics\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"country\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"verifyPub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"verifySub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FlytrapFuncSigs maps the 4-byte function signature to its string representation.
var FlytrapFuncSigs = map[string]string{
	"37bb1812": "addPub(address,bytes32)",
	"c6c460f7": "addSub(address,bytes32)",
	"31a0a410": "addTopic(bytes32,bytes2,uint256,uint256)",
	"7bd7b2ac": "logAlert(address,uint8,bytes32)",
	"8b99d3f8": "revokePub(address,bytes32)",
	"16ce2894": "revokeSub(address,bytes32)",
	"0f2fbeec": "topics(bytes32)",
	"7d1bae0a": "verifyPub(address,bytes32)",
	"d1d4a771": "verifySub(address,bytes32)",
}

// FlytrapBin is the compiled bytecode used for deploying new contracts.
var FlytrapBin = "0x608060405234801561001057600080fd5b506040516108da3803806108da8339818101604052602081101561003357600080fd5b5051600080546001600160a01b03191633179055600155610881806100596000396000f3fe6080604052600436106100865760003560e01c80637bd7b2ac116100595780637bd7b2ac1461019d5780637d1bae0a146101df5780638b99d3f81461023c578063c6c460f714610275578063d1d4a771146102a157610086565b80630f2fbeec1461008b57806316ce2894146100fd57806331a0a4101461013857806337bb181214610171575b600080fd5b34801561009757600080fd5b506100b5600480360360208110156100ae57600080fd5b50356102da565b60408051961515875260208701959095528585019390935260608501919091526001600160a01b031660808401526001600160f01b03191660a0830152519081900360c00190f35b34801561010957600080fd5b506101366004803603604081101561012057600080fd5b506001600160a01b038135169060200135610322565b005b6101366004803603608081101561014e57600080fd5b508035906001600160f01b031960208201351690604081013590606001356103c3565b6101366004803603604081101561018757600080fd5b506001600160a01b038135169060200135610529565b3480156101a957600080fd5b50610136600480360360608110156101c057600080fd5b506001600160a01b038135169060ff602082013516906040013561061a565b3480156101eb57600080fd5b506102186004803603604081101561020257600080fd5b506001600160a01b038135169060200135610672565b6040805192151583526001600160f01b031990911660208301528051918290030190f35b34801561024857600080fd5b506101366004803603604081101561025f57600080fd5b506001600160a01b0381351690602001356106ba565b6101366004803603604081101561028b57600080fd5b506001600160a01b038135169060200135610714565b3480156102ad57600080fd5b50610218600480360360408110156102c457600080fd5b506001600160a01b038135169060200135610803565b60026020819052600091825260409091208054600182015492820154600383015460049093015460ff909216939290916001600160a01b03811690600160a01b900460f01b86565b6000818152600260205260409020600401546001600160a01b0316331461034857600080fd5b60008181526002602090815260408083206001600160a01b03861684526006019091529020805460ff1916905560045b6040805183815242602082015281516001600160a01b0386169233927f8866e5b4c1d0cc6912c273f44188b035e935ede829489860f0c8bae1d1f8c570929081900390910190a45050565b60008481526002602052604090205460ff16156103df57600080fd5b6001543410156103ee57600080fd5b6040805160c081018252600180825260208083018881528385018781526060850187815233608087019081526001600160f01b03198b1660a0880190815260008d8152600296879052989098209651875490151560ff19909116178755925186860155905192850192909255905160038401555160049092018054935160f01c600160a01b0261ffff60a01b196001600160a01b03949094166001600160a01b0319909516949094179290921692909217905554156104e357600080546040516001600160a01b03909116913480156108fc02929091818181858888f193505050501580156104e1573d6000803e3d6000fd5b505b6000604080518681524260208201528151339283927f8866e5b4c1d0cc6912c273f44188b035e935ede829489860f0c8bae1d1f8c570929081900390910190a450505050565b6000818152600260205260409020600401546001600160a01b031633148061056257506000818152600260208190526040909120015415155b61056b57600080fd5b6000818152600260208190526040909120015434101561058a57600080fd5b60008181526002602081815260408084206001600160a01b0387168552600581018352908420805460ff1916600117905592849052819052015415610613576000818152600260205260408082206004015490516001600160a01b03909116913480156108fc02929091818181858888f19350505050158015610611573d6000803e3d6000fd5b505b6001610378565b81600681111561062657fe5b6040805183815242602082015281516001600160a01b0387169233927f8866e5b4c1d0cc6912c273f44188b035e935ede829489860f0c8bae1d1f8c570929081900390910190a4505050565b60008181526002602081815260408084206001600160a01b0396909616845260058601825283205493909252905260049091015460ff90911691600160a01b90910460f01b90565b6000818152600260205260409020600401546001600160a01b031633146106e057600080fd5b60008181526002602090815260408083206001600160a01b03861684526005019091529020805460ff191690556003610378565b6000818152600260205260409020600401546001600160a01b031633148061074c575060008181526002602052604090206003015415155b61075557600080fd5b60008181526002602052604090206003015434101561077357600080fd5b60008181526002602081815260408084206001600160a01b0387168552600681018352908420805460ff19166001179055928490525260030154156107fc576000818152600260205260408082206004015490516001600160a01b03909116913480156108fc02929091818181858888f193505050501580156107fa573d6000803e3d6000fd5b505b6002610378565b60008181526002602081815260408084206001600160a01b0396909616845260068601825283205493909252905260049091015460ff90911691600160a01b90910460f01b9056fea26469706673582212203da5e6a40f53fd880444dafd2b4274fd57f0d232188673cad67bf4f75a85830e64736f6c63430006020033"

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
// Solidity: function topics(bytes32 ) constant returns(bool isValue, bytes32 name, uint256 addPubCost, uint256 addSubCost, address owner, bytes2 country)
func (_Flytrap *FlytrapCaller) Topics(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsValue    bool
	Name       [32]byte
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
	Country    [2]byte
}, error) {
	ret := new(struct {
		IsValue    bool
		Name       [32]byte
		AddPubCost *big.Int
		AddSubCost *big.Int
		Owner      common.Address
		Country    [2]byte
	})
	out := ret
	err := _Flytrap.contract.Call(opts, out, "topics", arg0)
	return *ret, err
}

// Topics is a free data retrieval call binding the contract method 0x0f2fbeec.
//
// Solidity: function topics(bytes32 ) constant returns(bool isValue, bytes32 name, uint256 addPubCost, uint256 addSubCost, address owner, bytes2 country)
func (_Flytrap *FlytrapSession) Topics(arg0 [32]byte) (struct {
	IsValue    bool
	Name       [32]byte
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
	Country    [2]byte
}, error) {
	return _Flytrap.Contract.Topics(&_Flytrap.CallOpts, arg0)
}

// Topics is a free data retrieval call binding the contract method 0x0f2fbeec.
//
// Solidity: function topics(bytes32 ) constant returns(bool isValue, bytes32 name, uint256 addPubCost, uint256 addSubCost, address owner, bytes2 country)
func (_Flytrap *FlytrapCallerSession) Topics(arg0 [32]byte) (struct {
	IsValue    bool
	Name       [32]byte
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
	Country    [2]byte
}, error) {
	return _Flytrap.Contract.Topics(&_Flytrap.CallOpts, arg0)
}

// VerifyPub is a free data retrieval call binding the contract method 0x7d1bae0a.
//
// Solidity: function verifyPub(address person, bytes32 topic) constant returns(bool, bytes2)
func (_Flytrap *FlytrapCaller) VerifyPub(opts *bind.CallOpts, person common.Address, topic [32]byte) (bool, [2]byte, error) {
	var (
		ret0 = new(bool)
		ret1 = new([2]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Flytrap.contract.Call(opts, out, "verifyPub", person, topic)
	return *ret0, *ret1, err
}

// VerifyPub is a free data retrieval call binding the contract method 0x7d1bae0a.
//
// Solidity: function verifyPub(address person, bytes32 topic) constant returns(bool, bytes2)
func (_Flytrap *FlytrapSession) VerifyPub(person common.Address, topic [32]byte) (bool, [2]byte, error) {
	return _Flytrap.Contract.VerifyPub(&_Flytrap.CallOpts, person, topic)
}

// VerifyPub is a free data retrieval call binding the contract method 0x7d1bae0a.
//
// Solidity: function verifyPub(address person, bytes32 topic) constant returns(bool, bytes2)
func (_Flytrap *FlytrapCallerSession) VerifyPub(person common.Address, topic [32]byte) (bool, [2]byte, error) {
	return _Flytrap.Contract.VerifyPub(&_Flytrap.CallOpts, person, topic)
}

// VerifySub is a free data retrieval call binding the contract method 0xd1d4a771.
//
// Solidity: function verifySub(address person, bytes32 topic) constant returns(bool, bytes2)
func (_Flytrap *FlytrapCaller) VerifySub(opts *bind.CallOpts, person common.Address, topic [32]byte) (bool, [2]byte, error) {
	var (
		ret0 = new(bool)
		ret1 = new([2]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Flytrap.contract.Call(opts, out, "verifySub", person, topic)
	return *ret0, *ret1, err
}

// VerifySub is a free data retrieval call binding the contract method 0xd1d4a771.
//
// Solidity: function verifySub(address person, bytes32 topic) constant returns(bool, bytes2)
func (_Flytrap *FlytrapSession) VerifySub(person common.Address, topic [32]byte) (bool, [2]byte, error) {
	return _Flytrap.Contract.VerifySub(&_Flytrap.CallOpts, person, topic)
}

// VerifySub is a free data retrieval call binding the contract method 0xd1d4a771.
//
// Solidity: function verifySub(address person, bytes32 topic) constant returns(bool, bytes2)
func (_Flytrap *FlytrapCallerSession) VerifySub(person common.Address, topic [32]byte) (bool, [2]byte, error) {
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

// AddTopic is a paid mutator transaction binding the contract method 0x31a0a410.
//
// Solidity: function addTopic(bytes32 topic, bytes2 country, uint256 addPubCost, uint256 addSubCost) returns()
func (_Flytrap *FlytrapTransactor) AddTopic(opts *bind.TransactOpts, topic [32]byte, country [2]byte, addPubCost *big.Int, addSubCost *big.Int) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addTopic", topic, country, addPubCost, addSubCost)
}

// AddTopic is a paid mutator transaction binding the contract method 0x31a0a410.
//
// Solidity: function addTopic(bytes32 topic, bytes2 country, uint256 addPubCost, uint256 addSubCost) returns()
func (_Flytrap *FlytrapSession) AddTopic(topic [32]byte, country [2]byte, addPubCost *big.Int, addSubCost *big.Int) (*types.Transaction, error) {
	return _Flytrap.Contract.AddTopic(&_Flytrap.TransactOpts, topic, country, addPubCost, addSubCost)
}

// AddTopic is a paid mutator transaction binding the contract method 0x31a0a410.
//
// Solidity: function addTopic(bytes32 topic, bytes2 country, uint256 addPubCost, uint256 addSubCost) returns()
func (_Flytrap *FlytrapTransactorSession) AddTopic(topic [32]byte, country [2]byte, addPubCost *big.Int, addSubCost *big.Int) (*types.Transaction, error) {
	return _Flytrap.Contract.AddTopic(&_Flytrap.TransactOpts, topic, country, addPubCost, addSubCost)
}

// LogAlert is a paid mutator transaction binding the contract method 0x7bd7b2ac.
//
// Solidity: function logAlert(address person, uint8 alert, bytes32 desc) returns()
func (_Flytrap *FlytrapTransactor) LogAlert(opts *bind.TransactOpts, person common.Address, alert uint8, desc [32]byte) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "logAlert", person, alert, desc)
}

// LogAlert is a paid mutator transaction binding the contract method 0x7bd7b2ac.
//
// Solidity: function logAlert(address person, uint8 alert, bytes32 desc) returns()
func (_Flytrap *FlytrapSession) LogAlert(person common.Address, alert uint8, desc [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.LogAlert(&_Flytrap.TransactOpts, person, alert, desc)
}

// LogAlert is a paid mutator transaction binding the contract method 0x7bd7b2ac.
//
// Solidity: function logAlert(address person, uint8 alert, bytes32 desc) returns()
func (_Flytrap *FlytrapTransactorSession) LogAlert(person common.Address, alert uint8, desc [32]byte) (*types.Transaction, error) {
	return _Flytrap.Contract.LogAlert(&_Flytrap.TransactOpts, person, alert, desc)
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
	Src       common.Address
	Target    common.Address
	Action    uint8
	Name      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterACLChange is a free log retrieval operation binding the contract event 0x8866e5b4c1d0cc6912c273f44188b035e935ede829489860f0c8bae1d1f8c570.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, bytes32 _name, uint256 _timestamp)
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

// WatchACLChange is a free log subscription operation binding the contract event 0x8866e5b4c1d0cc6912c273f44188b035e935ede829489860f0c8bae1d1f8c570.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, bytes32 _name, uint256 _timestamp)
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

// ParseACLChange is a log parse operation binding the contract event 0x8866e5b4c1d0cc6912c273f44188b035e935ede829489860f0c8bae1d1f8c570.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, bytes32 _name, uint256 _timestamp)
func (_Flytrap *FlytrapFilterer) ParseACLChange(log types.Log) (*FlytrapACLChange, error) {
	event := new(FlytrapACLChange)
	if err := _Flytrap.contract.UnpackLog(event, "ACLChange", log); err != nil {
		return nil, err
	}
	return event, nil
}
