// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deploy

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
	_ = abi.ConvertType
)

// ENGCStakingMetaData contains all meta data concerning the ENGCStaking contract.
var ENGCStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_engcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"engcToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610d47380380610d4783398181016040528101906100319190610120565b815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600181905550505061015e565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100ab82610082565b9050919050565b5f6100bc826100a1565b9050919050565b6100cc816100b2565b81146100d6575f5ffd5b50565b5f815190506100e7816100c3565b92915050565b5f819050919050565b6100ff816100ed565b8114610109575f5ffd5b50565b5f8151905061011a816100f6565b92915050565b5f5f604083850312156101365761013561007e565b5b5f610143858286016100d9565b92505060206101548582860161010c565b9150509250929050565b610bdc8061016b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610086575f3560e01c8063817b1cd211610059578063817b1cd214610114578063b6b55f2514610132578063c7b8981c1461014e578063f40f0f521461015857610086565b806316934fc41461008a5780631811585b146100bc5780632e1a7d4d146100da5780637b0a47ee146100f6575b5f5ffd5b6100a4600480360381019061009f91906106e2565b610188565b6040516100b393929190610725565b60405180910390f35b6100c46101ae565b6040516100d191906107b5565b60405180910390f35b6100f460048036038101906100ef91906107f8565b6101d2565b005b6100fe610341565b60405161010b9190610823565b60405180910390f35b61011c610347565b6040516101299190610823565b60405180910390f35b61014c600480360381019061014791906107f8565b61034d565b005b6101566104ab565b005b610172600480360381019061016d91906106e2565b6105e2565b60405161017f9190610823565b60405180910390f35b6002602052805f5260405f205f91509050805f0154908060010154908060020154905083565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905081815f01541015610258576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024f90610896565b60405180910390fd5b5f610262336105f3565b905082825f015f82825461027691906108e1565b925050819055508260035f82825461028e91906108e1565b925050819055505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb3383866102de9190610914565b6040518363ffffffff1660e01b81526004016102fb929190610956565b6020604051808303815f875af1158015610317573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061033b91906109b2565b50505050565b60015481565b60035481565b5f811161038f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038690610a27565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016103ec93929190610a45565b6020604051808303815f875af1158015610408573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061042c91906109b2565b505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905081815f015f82825461047f9190610914565b925050819055504281600201819055508160035f8282546104a09190610914565b925050819055505050565b5f6104b5336105f3565b90505f81116104f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f090610ac4565b60405180910390fd5b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090504281600201819055505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b815260040161059d929190610956565b6020604051808303815f875af11580156105b9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105dd91906109b2565b505050565b5f6105ec826105f3565b9050919050565b5f5f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f81600201544261064591906108e1565b90505f670de0b6b3a764000082600154855f01546106639190610ae2565b61066d9190610ae2565b6106779190610b50565b9050809350505050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106b182610688565b9050919050565b6106c1816106a7565b81146106cb575f5ffd5b50565b5f813590506106dc816106b8565b92915050565b5f602082840312156106f7576106f6610684565b5b5f610704848285016106ce565b91505092915050565b5f819050919050565b61071f8161070d565b82525050565b5f6060820190506107385f830186610716565b6107456020830185610716565b6107526040830184610716565b949350505050565b5f819050919050565b5f61077d61077861077384610688565b61075a565b610688565b9050919050565b5f61078e82610763565b9050919050565b5f61079f82610784565b9050919050565b6107af81610795565b82525050565b5f6020820190506107c85f8301846107a6565b92915050565b6107d78161070d565b81146107e1575f5ffd5b50565b5f813590506107f2816107ce565b92915050565b5f6020828403121561080d5761080c610684565b5b5f61081a848285016107e4565b91505092915050565b5f6020820190506108365f830184610716565b92915050565b5f82825260208201905092915050565b7f496e73756666696369656e74207374616b65642062616c616e636500000000005f82015250565b5f610880601b8361083c565b915061088b8261084c565b602082019050919050565b5f6020820190508181035f8301526108ad81610874565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6108eb8261070d565b91506108f68361070d565b925082820390508181111561090e5761090d6108b4565b5b92915050565b5f61091e8261070d565b91506109298361070d565b9250828201905080821115610941576109406108b4565b5b92915050565b610950816106a7565b82525050565b5f6040820190506109695f830185610947565b6109766020830184610716565b9392505050565b5f8115159050919050565b6109918161097d565b811461099b575f5ffd5b50565b5f815190506109ac81610988565b92915050565b5f602082840312156109c7576109c6610684565b5b5f6109d48482850161099e565b91505092915050565b7f43616e6e6f74207374616b65203020746f6b656e7300000000000000000000005f82015250565b5f610a1160158361083c565b9150610a1c826109dd565b602082019050919050565b5f6020820190508181035f830152610a3e81610a05565b9050919050565b5f606082019050610a585f830186610947565b610a656020830185610947565b610a726040830184610716565b949350505050565b7f4e6f207265776172647320746f207769746864726177000000000000000000005f82015250565b5f610aae60168361083c565b9150610ab982610a7a565b602082019050919050565b5f6020820190508181035f830152610adb81610aa2565b9050919050565b5f610aec8261070d565b9150610af78361070d565b9250828202610b058161070d565b91508282048414831517610b1c57610b1b6108b4565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610b5a8261070d565b9150610b658361070d565b925082610b7557610b74610b23565b5b82820490509291505056fea2646970667358221220c4eeab0d999d7c21438a1c8129797ff350de688ceb7cf61b59e0e7137211cd4e64736f6c637828302e382e32372d646576656c6f702e323032342e382e33312b636f6d6d69742e34306530303562650059",
}

// ENGCStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use ENGCStakingMetaData.ABI instead.
var ENGCStakingABI = ENGCStakingMetaData.ABI

// ENGCStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ENGCStakingMetaData.Bin instead.
var ENGCStakingBin = ENGCStakingMetaData.Bin

// DeployENGCStaking deploys a new Ethereum contract, binding an instance of ENGCStaking to it.
func DeployENGCStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _engcToken common.Address, _rewardRate *big.Int) (common.Address, *types.Transaction, *ENGCStaking, error) {
	parsed, err := ENGCStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ENGCStakingBin), backend, _engcToken, _rewardRate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENGCStaking{ENGCStakingCaller: ENGCStakingCaller{contract: contract}, ENGCStakingTransactor: ENGCStakingTransactor{contract: contract}, ENGCStakingFilterer: ENGCStakingFilterer{contract: contract}}, nil
}

// ENGCStaking is an auto generated Go binding around an Ethereum contract.
type ENGCStaking struct {
	ENGCStakingCaller     // Read-only binding to the contract
	ENGCStakingTransactor // Write-only binding to the contract
	ENGCStakingFilterer   // Log filterer for contract events
}

// ENGCStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENGCStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENGCStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENGCStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENGCStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENGCStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENGCStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENGCStakingSession struct {
	Contract     *ENGCStaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENGCStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENGCStakingCallerSession struct {
	Contract *ENGCStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ENGCStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENGCStakingTransactorSession struct {
	Contract     *ENGCStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ENGCStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENGCStakingRaw struct {
	Contract *ENGCStaking // Generic contract binding to access the raw methods on
}

// ENGCStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENGCStakingCallerRaw struct {
	Contract *ENGCStakingCaller // Generic read-only contract binding to access the raw methods on
}

// ENGCStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENGCStakingTransactorRaw struct {
	Contract *ENGCStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENGCStaking creates a new instance of ENGCStaking, bound to a specific deployed contract.
func NewENGCStaking(address common.Address, backend bind.ContractBackend) (*ENGCStaking, error) {
	contract, err := bindENGCStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENGCStaking{ENGCStakingCaller: ENGCStakingCaller{contract: contract}, ENGCStakingTransactor: ENGCStakingTransactor{contract: contract}, ENGCStakingFilterer: ENGCStakingFilterer{contract: contract}}, nil
}

// NewENGCStakingCaller creates a new read-only instance of ENGCStaking, bound to a specific deployed contract.
func NewENGCStakingCaller(address common.Address, caller bind.ContractCaller) (*ENGCStakingCaller, error) {
	contract, err := bindENGCStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENGCStakingCaller{contract: contract}, nil
}

// NewENGCStakingTransactor creates a new write-only instance of ENGCStaking, bound to a specific deployed contract.
func NewENGCStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*ENGCStakingTransactor, error) {
	contract, err := bindENGCStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENGCStakingTransactor{contract: contract}, nil
}

// NewENGCStakingFilterer creates a new log filterer instance of ENGCStaking, bound to a specific deployed contract.
func NewENGCStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*ENGCStakingFilterer, error) {
	contract, err := bindENGCStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENGCStakingFilterer{contract: contract}, nil
}

// bindENGCStaking binds a generic wrapper to an already deployed contract.
func bindENGCStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ENGCStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENGCStaking *ENGCStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENGCStaking.Contract.ENGCStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENGCStaking *ENGCStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENGCStaking.Contract.ENGCStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENGCStaking *ENGCStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENGCStaking.Contract.ENGCStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENGCStaking *ENGCStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENGCStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENGCStaking *ENGCStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENGCStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENGCStaking *ENGCStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENGCStaking.Contract.contract.Transact(opts, method, params...)
}

// EngcToken is a free data retrieval call binding the contract method 0x1811585b.
//
// Solidity: function engcToken() view returns(address)
func (_ENGCStaking *ENGCStakingCaller) EngcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ENGCStaking.contract.Call(opts, &out, "engcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EngcToken is a free data retrieval call binding the contract method 0x1811585b.
//
// Solidity: function engcToken() view returns(address)
func (_ENGCStaking *ENGCStakingSession) EngcToken() (common.Address, error) {
	return _ENGCStaking.Contract.EngcToken(&_ENGCStaking.CallOpts)
}

// EngcToken is a free data retrieval call binding the contract method 0x1811585b.
//
// Solidity: function engcToken() view returns(address)
func (_ENGCStaking *ENGCStakingCallerSession) EngcToken() (common.Address, error) {
	return _ENGCStaking.Contract.EngcToken(&_ENGCStaking.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _staker) view returns(uint256)
func (_ENGCStaking *ENGCStakingCaller) PendingReward(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ENGCStaking.contract.Call(opts, &out, "pendingReward", _staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _staker) view returns(uint256)
func (_ENGCStaking *ENGCStakingSession) PendingReward(_staker common.Address) (*big.Int, error) {
	return _ENGCStaking.Contract.PendingReward(&_ENGCStaking.CallOpts, _staker)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _staker) view returns(uint256)
func (_ENGCStaking *ENGCStakingCallerSession) PendingReward(_staker common.Address) (*big.Int, error) {
	return _ENGCStaking.Contract.PendingReward(&_ENGCStaking.CallOpts, _staker)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ENGCStaking *ENGCStakingCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ENGCStaking.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ENGCStaking *ENGCStakingSession) RewardRate() (*big.Int, error) {
	return _ENGCStaking.Contract.RewardRate(&_ENGCStaking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ENGCStaking *ENGCStakingCallerSession) RewardRate() (*big.Int, error) {
	return _ENGCStaking.Contract.RewardRate(&_ENGCStaking.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 rewardDebt, uint256 depositTime)
func (_ENGCStaking *ENGCStakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount      *big.Int
	RewardDebt  *big.Int
	DepositTime *big.Int
}, error) {
	var out []interface{}
	err := _ENGCStaking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		RewardDebt  *big.Int
		DepositTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DepositTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 rewardDebt, uint256 depositTime)
func (_ENGCStaking *ENGCStakingSession) Stakes(arg0 common.Address) (struct {
	Amount      *big.Int
	RewardDebt  *big.Int
	DepositTime *big.Int
}, error) {
	return _ENGCStaking.Contract.Stakes(&_ENGCStaking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 rewardDebt, uint256 depositTime)
func (_ENGCStaking *ENGCStakingCallerSession) Stakes(arg0 common.Address) (struct {
	Amount      *big.Int
	RewardDebt  *big.Int
	DepositTime *big.Int
}, error) {
	return _ENGCStaking.Contract.Stakes(&_ENGCStaking.CallOpts, arg0)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_ENGCStaking *ENGCStakingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ENGCStaking.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_ENGCStaking *ENGCStakingSession) TotalStaked() (*big.Int, error) {
	return _ENGCStaking.Contract.TotalStaked(&_ENGCStaking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_ENGCStaking *ENGCStakingCallerSession) TotalStaked() (*big.Int, error) {
	return _ENGCStaking.Contract.TotalStaked(&_ENGCStaking.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_ENGCStaking *ENGCStakingTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ENGCStaking.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_ENGCStaking *ENGCStakingSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _ENGCStaking.Contract.Deposit(&_ENGCStaking.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_ENGCStaking *ENGCStakingTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _ENGCStaking.Contract.Deposit(&_ENGCStaking.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ENGCStaking *ENGCStakingTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ENGCStaking.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ENGCStaking *ENGCStakingSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _ENGCStaking.Contract.Withdraw(&_ENGCStaking.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ENGCStaking *ENGCStakingTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _ENGCStaking.Contract.Withdraw(&_ENGCStaking.TransactOpts, _amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_ENGCStaking *ENGCStakingTransactor) WithdrawRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENGCStaking.contract.Transact(opts, "withdrawRewards")
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_ENGCStaking *ENGCStakingSession) WithdrawRewards() (*types.Transaction, error) {
	return _ENGCStaking.Contract.WithdrawRewards(&_ENGCStaking.TransactOpts)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_ENGCStaking *ENGCStakingTransactorSession) WithdrawRewards() (*types.Transaction, error) {
	return _ENGCStaking.Contract.WithdrawRewards(&_ENGCStaking.TransactOpts)
}
