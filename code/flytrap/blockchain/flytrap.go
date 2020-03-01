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
const FlytrapABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumFlytrap.ActionPerformed\",\"name\":\"_action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"ACLChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"addPub\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"addSub\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"},{\"internalType\":\"bytes2\",\"name\":\"country\",\"type\":\"bytes2\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"sensitive\",\"type\":\"bool\"}],\"name\":\"addTopic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"enumFlytrap.ActionPerformed\",\"name\":\"alert\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"logAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"revokePub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"revokeSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"topics\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"sensitive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"addPubCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addSubCost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"country\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"verifyPub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"verifySub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FlytrapFuncSigs maps the 4-byte function signature to its string representation.
var FlytrapFuncSigs = map[string]string{
	"eb7e13b4": "addPub(address,string,string)",
	"4f7c039e": "addSub(address,string,string)",
	"f029d981": "addTopic(string,bytes2,uint256,uint256,string,bool)",
	"435e9dd7": "logAlert(address,uint8,string,string)",
	"820b479a": "revokePub(address,string,string)",
	"8b2bb7a5": "revokeSub(address,string,string)",
	"f75c6c72": "topics(string)",
	"b1486230": "verifyPub(address,string)",
	"51f70b1c": "verifySub(address,string)",
}

// FlytrapBin is the compiled bytecode used for deploying new contracts.
var FlytrapBin = "0x608060405234801561001057600080fd5b50604051611d31380380611d318339818101604052602081101561003357600080fd5b5051600080546001600160a01b03191633179055600155611cd8806100596000396000f3fe6080604052600436106100865760003560e01c80638b2bb7a5116100595780638b2bb7a514610546578063b14862301461068c578063eb7e13b41461074d578063f029d98114610886578063f75c6c72146109cd57610086565b8063435e9dd71461008b5780634f7c039e146101db57806351f70b1c14610314578063820b479a14610400575b600080fd5b34801561009757600080fd5b506101d9600480360360808110156100ae57600080fd5b6001600160a01b038235169160ff60208201351691810190606081016040820135600160201b8111156100e057600080fd5b8201836020820111156100f257600080fd5b803590602001918460018302840111600160201b8311171561011357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561016557600080fd5b82018360208201111561017757600080fd5b803590602001918460018302840111600160201b8311171561019857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b3e945050505050565b005b6101d9600480360360608110156101f157600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561021b57600080fd5b82018360208201111561022d57600080fd5b803590602001918460018302840111600160201b8311171561024e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102a057600080fd5b8201836020820111156102b257600080fd5b803590602001918460018302840111600160201b831117156102d357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c6c945050505050565b34801561032057600080fd5b506103d56004803603604081101561033757600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561036157600080fd5b82018360208201111561037357600080fd5b803590602001918460018302840111600160201b8311171561039457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611089945050505050565b6040805193151584526001600160f01b03199092166020840152151582820152519081900360600190f35b34801561040c57600080fd5b506101d96004803603606081101561042357600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561044d57600080fd5b82018360208201111561045f57600080fd5b803590602001918460018302840111600160201b8311171561048057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156104d257600080fd5b8201836020820111156104e457600080fd5b803590602001918460018302840111600160201b8311171561050557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506111fc945050505050565b34801561055257600080fd5b506101d96004803603606081101561056957600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561059357600080fd5b8201836020820111156105a557600080fd5b803590602001918460018302840111600160201b831117156105c657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561061857600080fd5b82018360208201111561062a57600080fd5b803590602001918460018302840111600160201b8311171561064b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611317945050505050565b34801561069857600080fd5b506103d5600480360360408110156106af57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156106d957600080fd5b8201836020820111156106eb57600080fd5b803590602001918460018302840111600160201b8311171561070c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611432945050505050565b6101d96004803603606081101561076357600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561078d57600080fd5b82018360208201111561079f57600080fd5b803590602001918460018302840111600160201b831117156107c057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561081257600080fd5b82018360208201111561082457600080fd5b803590602001918460018302840111600160201b8311171561084557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506114e7945050505050565b6101d9600480360360c081101561089c57600080fd5b810190602081018135600160201b8111156108b657600080fd5b8201836020820111156108c857600080fd5b803590602001918460018302840111600160201b831117156108e957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092956001600160f01b031985351695602086013595604081013595509193509150608081019060600135600160201b81111561095757600080fd5b82018360208201111561096957600080fd5b803590602001918460018302840111600160201b8311171561098a57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050505035151590506117e7565b3480156109d957600080fd5b50610a7e600480360360208110156109f057600080fd5b810190602081018135600160201b811115610a0a57600080fd5b820183602082011115610a1c57600080fd5b803590602001918460018302840111600160201b83111715610a3d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611b1f945050505050565b60408051881515815287151560208083019190915260608201879052608082018690526001600160a01b03851660a08301526001600160f01b0319841660c083015260e0928201838152885193830193909352875191929161010084019189019080838360005b83811015610afd578181015183820152602001610ae5565b50505050905090810190601f168015610b2a5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b826007811115610b4a57fe5b846001600160a01b0316336001600160a01b03167f671de2fbee27783d42cd5da9bd3d0025c7f8a762d877c4beb088fb4e9d7a541d858542604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610bc9578181015183820152602001610bb1565b50505050905090810190601f168015610bf65780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610c29578181015183820152602001610c11565b50505050905090810190601f168015610c565780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a450505050565b336001600160a01b03166002836040518082805190602001908083835b60208310610ca85780518252601f199092019160209182019101610c89565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600401546001600160a01b0316929092149150819050610d5b575060006002836040518082805190602001908083835b60208310610d225780518252601f199092019160209182019101610d03565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060030154115b610d6457600080fd5b6002826040518082805190602001908083835b60208310610d965780518252601f199092019160209182019101610d77565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600301543410159150610dd5905057600080fd5b60016002836040518082805190602001908083835b60208310610e095780518252601f199092019160209182019101610dea565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160a01b038a16600090815260069091018352908120805460ff1916961515969096179095558651600294889450925082918401908083835b60208310610e905780518252601f199092019160209182019101610e71565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600301541115610f65576002826040518082805190602001908083835b60208310610eff5780518252601f199092019160209182019101610ee0565b51815160001960209485036101000a0190811690199190911617905292019485525060405193849003018320600401546001600160a01b0316923480156108fc02935091506000818181858888f19350505050158015610f63573d6000803e3d6000fd5b505b60025b836001600160a01b0316336001600160a01b03167f671de2fbee27783d42cd5da9bd3d0025c7f8a762d877c4beb088fb4e9d7a541d858542604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610fe7578181015183820152602001610fcf565b50505050905090810190601f1680156110145780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561104757818101518382015260200161102f565b50505050905090810190601f1680156110745780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a4505050565b60008060006002846040518082805190602001908083835b602083106110c05780518252601f1990920191602091820191016110a1565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160a01b038b166000908152600690910183522054885160ff9091169460029450899350918291908401908083835b6020831061113d5780518252601f19909201916020918201910161111e565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420600401548951600160a01b90910460f01b94600294508a9350918291908401908083835b602083106111ab5780518252601f19909201916020918201910161118c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160019054906101000a900460ff169250925092509250925092565b336001600160a01b03166002836040518082805190602001908083835b602083106112385780518252601f199092019160209182019101611219565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600401546001600160a01b0316929092149150611281905057600080fd5b60006002836040518082805190602001908083835b602083106112b55780518252601f199092019160209182019101611296565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160a01b0389166000908152600590910190915292909220805460ff1916931515939093179092555060039050610f68565b336001600160a01b03166002836040518082805190602001908083835b602083106113535780518252601f199092019160209182019101611334565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600401546001600160a01b031692909214915061139c905057600080fd5b60006002836040518082805190602001908083835b602083106113d05780518252601f1990920191602091820191016113b1565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160a01b0389166000908152600690910190915292909220805460ff1916931515939093179092555060049050610f68565b60008060006002846040518082805190602001908083835b602083106114695780518252601f19909201916020918201910161144a565b51815160001960209485036101000a019081169019919091161790529201948552506040805194859003820185206001600160a01b038b166000908152600590910183522054885160ff9091169460029450899350918291908401908083836020831061113d5780518252601f19909201916020918201910161111e565b336001600160a01b03166002836040518082805190602001908083835b602083106115235780518252601f199092019160209182019101611504565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600401546001600160a01b03169290921491508190506115d6575060006002836040518082805190602001908083835b6020831061159d5780518252601f19909201916020918201910161157e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154115b6115df57600080fd5b6002826040518082805190602001908083835b602083106116115780518252601f1990920191602091820191016115f2565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600201543410159150611650905057600080fd5b60016002836040518082805190602001908083835b602083106116845780518252601f199092019160209182019101611665565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160a01b038a16600090815260059091018352908120805460ff1916961515969096179095558651600294889450925082918401908083835b6020831061170b5780518252601f1990920191602091820191016116ec565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206002015411156117e0576002826040518082805190602001908083835b6020831061177a5780518252601f19909201916020918201910161175b565b51815160001960209485036101000a0190811690199190911617905292019485525060405193849003018320600401546001600160a01b0316923480156108fc02935091506000818181858888f193505050501580156117de573d6000803e3d6000fd5b505b6001610f68565b6002866040518082805190602001908083835b602083106118195780518252601f1990920191602091820191016117fa565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff16159150611856905057600080fd5b60015434101561186557600080fd5b6040518060e001604052806001151581526020018215158152602001878152602001858152602001848152602001336001600160a01b03168152602001866001600160f01b0319168152506002876040518082805190602001908083835b602083106118e25780518252601f1990920191602091820191016118c3565b518151600019602094850361010090810a919091019182169119929092161790915293909101958652604080519687900382019096208751815489840151151590950261ff001991151560ff1990961695909517169390931783559486015180519295611959955060018701945001919050611c07565b50606082015160028201556080820151600382015560a08201516004909101805460c09093015160f01c600160a01b0261ffff60a01b196001600160a01b039093166001600160a01b03199094169390931791909116919091179055600154156119f957600080546040516001600160a01b03909116913480156108fc02929091818181858888f193505050501580156119f7573d6000803e3d6000fd5b505b6000336001600160a01b0316336001600160a01b03167f671de2fbee27783d42cd5da9bd3d0025c7f8a762d877c4beb088fb4e9d7a541d898642604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015611a7a578181015183820152602001611a62565b50505050905090810190601f168015611aa75780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015611ada578181015183820152602001611ac2565b50505050905090810190601f168015611b075780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a4505050505050565b8051602081830181018051600280835293830194830194909420939052825460018085018054604080516101009483161585026000190190921696909604601f810186900486028201860190965285815260ff8085169793909404909316949193909190830182828015611bd45780601f10611ba957610100808354040283529160200191611bd4565b820191906000526020600020905b815481529060010190602001808311611bb757829003601f168201915b50505060028401546003850154600490950154939490939092506001600160a01b0381169150600160a01b900460f01b87565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611c4857805160ff1916838001178555611c75565b82800160010185558215611c75579182015b82811115611c75578251825591602001919060010190611c5a565b50611c81929150611c85565b5090565b611c9f91905b80821115611c815760008155600101611c8b565b9056fea2646970667358221220c026d5173d603f5715c49d2e0090a1c1d9b00342ad50bdd4f4ec7acb5c8a501864736f6c63430006030033"

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

// Topics is a free data retrieval call binding the contract method 0xf75c6c72.
//
// Solidity: function topics(string ) constant returns(bool isValue, bool sensitive, string name, uint256 addPubCost, uint256 addSubCost, address owner, bytes2 country)
func (_Flytrap *FlytrapCaller) Topics(opts *bind.CallOpts, arg0 string) (struct {
	IsValue    bool
	Sensitive  bool
	Name       string
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
	Country    [2]byte
}, error) {
	ret := new(struct {
		IsValue    bool
		Sensitive  bool
		Name       string
		AddPubCost *big.Int
		AddSubCost *big.Int
		Owner      common.Address
		Country    [2]byte
	})
	out := ret
	err := _Flytrap.contract.Call(opts, out, "topics", arg0)
	return *ret, err
}

// Topics is a free data retrieval call binding the contract method 0xf75c6c72.
//
// Solidity: function topics(string ) constant returns(bool isValue, bool sensitive, string name, uint256 addPubCost, uint256 addSubCost, address owner, bytes2 country)
func (_Flytrap *FlytrapSession) Topics(arg0 string) (struct {
	IsValue    bool
	Sensitive  bool
	Name       string
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
	Country    [2]byte
}, error) {
	return _Flytrap.Contract.Topics(&_Flytrap.CallOpts, arg0)
}

// Topics is a free data retrieval call binding the contract method 0xf75c6c72.
//
// Solidity: function topics(string ) constant returns(bool isValue, bool sensitive, string name, uint256 addPubCost, uint256 addSubCost, address owner, bytes2 country)
func (_Flytrap *FlytrapCallerSession) Topics(arg0 string) (struct {
	IsValue    bool
	Sensitive  bool
	Name       string
	AddPubCost *big.Int
	AddSubCost *big.Int
	Owner      common.Address
	Country    [2]byte
}, error) {
	return _Flytrap.Contract.Topics(&_Flytrap.CallOpts, arg0)
}

// VerifyPub is a free data retrieval call binding the contract method 0xb1486230.
//
// Solidity: function verifyPub(address person, string topic) constant returns(bool, bytes2, bool)
func (_Flytrap *FlytrapCaller) VerifyPub(opts *bind.CallOpts, person common.Address, topic string) (bool, [2]byte, bool, error) {
	var (
		ret0 = new(bool)
		ret1 = new([2]byte)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Flytrap.contract.Call(opts, out, "verifyPub", person, topic)
	return *ret0, *ret1, *ret2, err
}

// VerifyPub is a free data retrieval call binding the contract method 0xb1486230.
//
// Solidity: function verifyPub(address person, string topic) constant returns(bool, bytes2, bool)
func (_Flytrap *FlytrapSession) VerifyPub(person common.Address, topic string) (bool, [2]byte, bool, error) {
	return _Flytrap.Contract.VerifyPub(&_Flytrap.CallOpts, person, topic)
}

// VerifyPub is a free data retrieval call binding the contract method 0xb1486230.
//
// Solidity: function verifyPub(address person, string topic) constant returns(bool, bytes2, bool)
func (_Flytrap *FlytrapCallerSession) VerifyPub(person common.Address, topic string) (bool, [2]byte, bool, error) {
	return _Flytrap.Contract.VerifyPub(&_Flytrap.CallOpts, person, topic)
}

// VerifySub is a free data retrieval call binding the contract method 0x51f70b1c.
//
// Solidity: function verifySub(address person, string topic) constant returns(bool, bytes2, bool)
func (_Flytrap *FlytrapCaller) VerifySub(opts *bind.CallOpts, person common.Address, topic string) (bool, [2]byte, bool, error) {
	var (
		ret0 = new(bool)
		ret1 = new([2]byte)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Flytrap.contract.Call(opts, out, "verifySub", person, topic)
	return *ret0, *ret1, *ret2, err
}

// VerifySub is a free data retrieval call binding the contract method 0x51f70b1c.
//
// Solidity: function verifySub(address person, string topic) constant returns(bool, bytes2, bool)
func (_Flytrap *FlytrapSession) VerifySub(person common.Address, topic string) (bool, [2]byte, bool, error) {
	return _Flytrap.Contract.VerifySub(&_Flytrap.CallOpts, person, topic)
}

// VerifySub is a free data retrieval call binding the contract method 0x51f70b1c.
//
// Solidity: function verifySub(address person, string topic) constant returns(bool, bytes2, bool)
func (_Flytrap *FlytrapCallerSession) VerifySub(person common.Address, topic string) (bool, [2]byte, bool, error) {
	return _Flytrap.Contract.VerifySub(&_Flytrap.CallOpts, person, topic)
}

// AddPub is a paid mutator transaction binding the contract method 0xeb7e13b4.
//
// Solidity: function addPub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactor) AddPub(opts *bind.TransactOpts, person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addPub", person, topic, reason)
}

// AddPub is a paid mutator transaction binding the contract method 0xeb7e13b4.
//
// Solidity: function addPub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapSession) AddPub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.AddPub(&_Flytrap.TransactOpts, person, topic, reason)
}

// AddPub is a paid mutator transaction binding the contract method 0xeb7e13b4.
//
// Solidity: function addPub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactorSession) AddPub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.AddPub(&_Flytrap.TransactOpts, person, topic, reason)
}

// AddSub is a paid mutator transaction binding the contract method 0x4f7c039e.
//
// Solidity: function addSub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactor) AddSub(opts *bind.TransactOpts, person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addSub", person, topic, reason)
}

// AddSub is a paid mutator transaction binding the contract method 0x4f7c039e.
//
// Solidity: function addSub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapSession) AddSub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.AddSub(&_Flytrap.TransactOpts, person, topic, reason)
}

// AddSub is a paid mutator transaction binding the contract method 0x4f7c039e.
//
// Solidity: function addSub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactorSession) AddSub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.AddSub(&_Flytrap.TransactOpts, person, topic, reason)
}

// AddTopic is a paid mutator transaction binding the contract method 0xf029d981.
//
// Solidity: function addTopic(string topic, bytes2 country, uint256 addPubCost, uint256 addSubCost, string reason, bool sensitive) returns()
func (_Flytrap *FlytrapTransactor) AddTopic(opts *bind.TransactOpts, topic string, country [2]byte, addPubCost *big.Int, addSubCost *big.Int, reason string, sensitive bool) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "addTopic", topic, country, addPubCost, addSubCost, reason, sensitive)
}

// AddTopic is a paid mutator transaction binding the contract method 0xf029d981.
//
// Solidity: function addTopic(string topic, bytes2 country, uint256 addPubCost, uint256 addSubCost, string reason, bool sensitive) returns()
func (_Flytrap *FlytrapSession) AddTopic(topic string, country [2]byte, addPubCost *big.Int, addSubCost *big.Int, reason string, sensitive bool) (*types.Transaction, error) {
	return _Flytrap.Contract.AddTopic(&_Flytrap.TransactOpts, topic, country, addPubCost, addSubCost, reason, sensitive)
}

// AddTopic is a paid mutator transaction binding the contract method 0xf029d981.
//
// Solidity: function addTopic(string topic, bytes2 country, uint256 addPubCost, uint256 addSubCost, string reason, bool sensitive) returns()
func (_Flytrap *FlytrapTransactorSession) AddTopic(topic string, country [2]byte, addPubCost *big.Int, addSubCost *big.Int, reason string, sensitive bool) (*types.Transaction, error) {
	return _Flytrap.Contract.AddTopic(&_Flytrap.TransactOpts, topic, country, addPubCost, addSubCost, reason, sensitive)
}

// LogAlert is a paid mutator transaction binding the contract method 0x435e9dd7.
//
// Solidity: function logAlert(address person, uint8 alert, string desc, string reason) returns()
func (_Flytrap *FlytrapTransactor) LogAlert(opts *bind.TransactOpts, person common.Address, alert uint8, desc string, reason string) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "logAlert", person, alert, desc, reason)
}

// LogAlert is a paid mutator transaction binding the contract method 0x435e9dd7.
//
// Solidity: function logAlert(address person, uint8 alert, string desc, string reason) returns()
func (_Flytrap *FlytrapSession) LogAlert(person common.Address, alert uint8, desc string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.LogAlert(&_Flytrap.TransactOpts, person, alert, desc, reason)
}

// LogAlert is a paid mutator transaction binding the contract method 0x435e9dd7.
//
// Solidity: function logAlert(address person, uint8 alert, string desc, string reason) returns()
func (_Flytrap *FlytrapTransactorSession) LogAlert(person common.Address, alert uint8, desc string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.LogAlert(&_Flytrap.TransactOpts, person, alert, desc, reason)
}

// RevokePub is a paid mutator transaction binding the contract method 0x820b479a.
//
// Solidity: function revokePub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactor) RevokePub(opts *bind.TransactOpts, person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "revokePub", person, topic, reason)
}

// RevokePub is a paid mutator transaction binding the contract method 0x820b479a.
//
// Solidity: function revokePub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapSession) RevokePub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokePub(&_Flytrap.TransactOpts, person, topic, reason)
}

// RevokePub is a paid mutator transaction binding the contract method 0x820b479a.
//
// Solidity: function revokePub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactorSession) RevokePub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokePub(&_Flytrap.TransactOpts, person, topic, reason)
}

// RevokeSub is a paid mutator transaction binding the contract method 0x8b2bb7a5.
//
// Solidity: function revokeSub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactor) RevokeSub(opts *bind.TransactOpts, person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.contract.Transact(opts, "revokeSub", person, topic, reason)
}

// RevokeSub is a paid mutator transaction binding the contract method 0x8b2bb7a5.
//
// Solidity: function revokeSub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapSession) RevokeSub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokeSub(&_Flytrap.TransactOpts, person, topic, reason)
}

// RevokeSub is a paid mutator transaction binding the contract method 0x8b2bb7a5.
//
// Solidity: function revokeSub(address person, string topic, string reason) returns()
func (_Flytrap *FlytrapTransactorSession) RevokeSub(person common.Address, topic string, reason string) (*types.Transaction, error) {
	return _Flytrap.Contract.RevokeSub(&_Flytrap.TransactOpts, person, topic, reason)
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
	Name      string
	Reason    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterACLChange is a free log retrieval operation binding the contract event 0x671de2fbee27783d42cd5da9bd3d0025c7f8a762d877c4beb088fb4e9d7a541d.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, string _name, string reason, uint256 _timestamp)
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

// WatchACLChange is a free log subscription operation binding the contract event 0x671de2fbee27783d42cd5da9bd3d0025c7f8a762d877c4beb088fb4e9d7a541d.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, string _name, string reason, uint256 _timestamp)
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

// ParseACLChange is a log parse operation binding the contract event 0x671de2fbee27783d42cd5da9bd3d0025c7f8a762d877c4beb088fb4e9d7a541d.
//
// Solidity: event ACLChange(address indexed _src, address indexed _target, uint8 indexed _action, string _name, string reason, uint256 _timestamp)
func (_Flytrap *FlytrapFilterer) ParseACLChange(log types.Log) (*FlytrapACLChange, error) {
	event := new(FlytrapACLChange)
	if err := _Flytrap.contract.UnpackLog(event, "ACLChange", log); err != nil {
		return nil, err
	}
	return event, nil
}
