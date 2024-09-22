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

// EngageCoinMetaData contains all meta data concerning the EngageCoin contract.
var EngageCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originalOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611c74380380611c7483398181016040528101906100319190610544565b806040518060400160405280600a81526020017f456e67616765436f696e000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f454e47430000000000000000000000000000000000000000000000000000000081525081600390816100ad91906107ac565b5080600490816100bd91906107ac565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610130575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610127919061088a565b60405180910390fd5b61013f8161017d60201b60201c565b506101773361015261024060201b60201c565b600a61015e9190610a0b565b620f424061016c9190610a55565b61024860201b60201c565b50610b26565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f6012905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036102b8575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016102af919061088a565b60405180910390fd5b6102c95f83836102cd60201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361031d578060025f8282546103119190610a96565b925050819055506103eb565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050818110156103a6578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161039d93929190610ad8565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610432578060025f828254039250508190555061047c565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516104d99190610b0d565b60405180910390a3505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610513826104ea565b9050919050565b61052381610509565b811461052d575f5ffd5b50565b5f8151905061053e8161051a565b92915050565b5f60208284031215610559576105586104e6565b5b5f61056684828501610530565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806105ea57607f821691505b6020821081036105fd576105fc6105a6565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261065f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610624565b6106698683610624565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6106ad6106a86106a384610681565b61068a565b610681565b9050919050565b5f819050919050565b6106c683610693565b6106da6106d2826106b4565b848454610630565b825550505050565b5f5f905090565b6106f16106e2565b6106fc8184846106bd565b505050565b5b8181101561071f576107145f826106e9565b600181019050610702565b5050565b601f8211156107645761073581610603565b61073e84610615565b8101602085101561074d578190505b61076161075985610615565b830182610701565b50505b505050565b5f82821c905092915050565b5f6107845f1984600802610769565b1980831691505092915050565b5f61079c8383610775565b9150826002028217905092915050565b6107b58261056f565b67ffffffffffffffff8111156107ce576107cd610579565b5b6107d882546105d3565b6107e3828285610723565b5f60209050601f831160018114610814575f8415610802578287015190505b61080c8582610791565b865550610873565b601f19841661082286610603565b5f5b8281101561084957848901518255600182019150602085019450602081019050610824565b868310156108665784890151610862601f891682610775565b8355505b6001600288020188555050505b505050505050565b61088481610509565b82525050565b5f60208201905061089d5f83018461087b565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b600185111561092557808604811115610901576109006108a3565b5b60018516156109105780820291505b808102905061091e856108d0565b94506108e5565b94509492505050565b5f8261093d57600190506109f8565b8161094a575f90506109f8565b8160018114610960576002811461096a57610999565b60019150506109f8565b60ff84111561097c5761097b6108a3565b5b8360020a915084821115610993576109926108a3565b5b506109f8565b5060208310610133831016604e8410600b84101617156109ce5782820a9050838111156109c9576109c86108a3565b5b6109f8565b6109db84848460016108dc565b925090508184048111156109f2576109f16108a3565b5b81810290505b9392505050565b5f60ff82169050919050565b5f610a1582610681565b9150610a20836109ff565b9250610a4d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461092e565b905092915050565b5f610a5f82610681565b9150610a6a83610681565b9250828202610a7881610681565b91508282048414831517610a8f57610a8e6108a3565b5b5092915050565b5f610aa082610681565b9150610aab83610681565b9250828201905080821115610ac357610ac26108a3565b5b92915050565b610ad281610681565b82525050565b5f606082019050610aeb5f83018661087b565b610af86020830185610ac9565b610b056040830184610ac9565b949350505050565b5f602082019050610b205f830184610ac9565b92915050565b61114180610b335f395ff3fe608060405234801561000f575f5ffd5b50600436106100cd575f3560e01c8063715018a61161008a5780639a49090e116100645780639a49090e14610201578063a9059cbb1461021d578063dd62ed3e1461024d578063f2fde38b1461027d576100cd565b8063715018a6146101bb5780638da5cb5b146101c557806395d89b41146101e3576100cd565b806306fdde03146100d1578063095ea7b3146100ef57806318160ddd1461011f57806323b872dd1461013d578063313ce5671461016d57806370a082311461018b575b5f5ffd5b6100d9610299565b6040516100e69190610d94565b60405180910390f35b61010960048036038101906101049190610e45565b610329565b6040516101169190610e9d565b60405180910390f35b61012761034b565b6040516101349190610ec5565b60405180910390f35b61015760048036038101906101529190610ede565b610354565b6040516101649190610e9d565b60405180910390f35b610175610382565b6040516101829190610f49565b60405180910390f35b6101a560048036038101906101a09190610f62565b61038a565b6040516101b29190610ec5565b60405180910390f35b6101c36103cf565b005b6101cd6103e2565b6040516101da9190610f9c565b60405180910390f35b6101eb61040a565b6040516101f89190610d94565b60405180910390f35b61021b60048036038101906102169190610e45565b61049a565b005b61023760048036038101906102329190610e45565b6104b0565b6040516102449190610e9d565b60405180910390f35b61026760048036038101906102629190610fb5565b6104d2565b6040516102749190610ec5565b60405180910390f35b61029760048036038101906102929190610f62565b610554565b005b6060600380546102a890611020565b80601f01602080910402602001604051908101604052809291908181526020018280546102d490611020565b801561031f5780601f106102f65761010080835404028352916020019161031f565b820191905f5260205f20905b81548152906001019060200180831161030257829003601f168201915b5050505050905090565b5f5f6103336105d8565b90506103408185856105df565b600191505092915050565b5f600254905090565b5f5f61035e6105d8565b905061036b8582856105f1565b610376858585610683565b60019150509392505050565b5f6012905090565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6103d7610773565b6103e05f6107fa565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461041990611020565b80601f016020809104026020016040519081016040528092919081815260200182805461044590611020565b80156104905780601f1061046757610100808354040283529160200191610490565b820191905f5260205f20905b81548152906001019060200180831161047357829003601f168201915b5050505050905090565b6104a2610773565b6104ac82826108bd565b5050565b5f5f6104ba6105d8565b90506104c7818585610683565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61055c610773565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105cc575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016105c39190610f9c565b60405180910390fd5b6105d5816107fa565b50565b5f33905090565b6105ec838383600161093c565b505050565b5f6105fc84846104d2565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461067d578181101561066e578281836040517ffb8f41b200000000000000000000000000000000000000000000000000000000815260040161066593929190611050565b60405180910390fd5b61067c84848484035f61093c565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036106f3575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016106ea9190610f9c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610763575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161075a9190610f9c565b60405180910390fd5b61076e838383610b0b565b505050565b61077b6105d8565b73ffffffffffffffffffffffffffffffffffffffff166107996103e2565b73ffffffffffffffffffffffffffffffffffffffff16146107f8576107bc6105d8565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016107ef9190610f9c565b60405180910390fd5b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361092d575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016109249190610f9c565b60405180910390fd5b6109385f8383610b0b565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036109ac575f6040517fe602df050000000000000000000000000000000000000000000000000000000081526004016109a39190610f9c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a1c575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610a139190610f9c565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610b05578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610afc9190610ec5565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b5b578060025f828254610b4f91906110b2565b92505081905550610c29565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610be4578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610bdb93929190611050565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c70578060025f8282540392505081905550610cba565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d179190610ec5565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610d6682610d24565b610d708185610d2e565b9350610d80818560208601610d3e565b610d8981610d4c565b840191505092915050565b5f6020820190508181035f830152610dac8184610d5c565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610de182610db8565b9050919050565b610df181610dd7565b8114610dfb575f5ffd5b50565b5f81359050610e0c81610de8565b92915050565b5f819050919050565b610e2481610e12565b8114610e2e575f5ffd5b50565b5f81359050610e3f81610e1b565b92915050565b5f5f60408385031215610e5b57610e5a610db4565b5b5f610e6885828601610dfe565b9250506020610e7985828601610e31565b9150509250929050565b5f8115159050919050565b610e9781610e83565b82525050565b5f602082019050610eb05f830184610e8e565b92915050565b610ebf81610e12565b82525050565b5f602082019050610ed85f830184610eb6565b92915050565b5f5f5f60608486031215610ef557610ef4610db4565b5b5f610f0286828701610dfe565b9350506020610f1386828701610dfe565b9250506040610f2486828701610e31565b9150509250925092565b5f60ff82169050919050565b610f4381610f2e565b82525050565b5f602082019050610f5c5f830184610f3a565b92915050565b5f60208284031215610f7757610f76610db4565b5b5f610f8484828501610dfe565b91505092915050565b610f9681610dd7565b82525050565b5f602082019050610faf5f830184610f8d565b92915050565b5f5f60408385031215610fcb57610fca610db4565b5b5f610fd885828601610dfe565b9250506020610fe985828601610dfe565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061103757607f821691505b60208210810361104a57611049610ff3565b5b50919050565b5f6060820190506110635f830186610f8d565b6110706020830185610eb6565b61107d6040830184610eb6565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110bc82610e12565b91506110c783610e12565b92508282019050808211156110df576110de611085565b5b9291505056fea26469706673582212205a621eb41b9ed2e06bcdf5460936fb992ad2adadbff509b9cb16049df315189064736f6c637828302e382e32372d646576656c6f702e323032342e382e33312b636f6d6d69742e34306530303562650059",
}

// EngageCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use EngageCoinMetaData.ABI instead.
var EngageCoinABI = EngageCoinMetaData.ABI

// EngageCoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EngageCoinMetaData.Bin instead.
var EngageCoinBin = EngageCoinMetaData.Bin

// DeployEngageCoin deploys a new Ethereum contract, binding an instance of EngageCoin to it.
func DeployEngageCoin(auth *bind.TransactOpts, backend bind.ContractBackend, originalOwner common.Address) (common.Address, *types.Transaction, *EngageCoin, error) {
	parsed, err := EngageCoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EngageCoinBin), backend, originalOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EngageCoin{EngageCoinCaller: EngageCoinCaller{contract: contract}, EngageCoinTransactor: EngageCoinTransactor{contract: contract}, EngageCoinFilterer: EngageCoinFilterer{contract: contract}}, nil
}

// EngageCoin is an auto generated Go binding around an Ethereum contract.
type EngageCoin struct {
	EngageCoinCaller     // Read-only binding to the contract
	EngageCoinTransactor // Write-only binding to the contract
	EngageCoinFilterer   // Log filterer for contract events
}

// EngageCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type EngageCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EngageCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EngageCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EngageCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EngageCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EngageCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EngageCoinSession struct {
	Contract     *EngageCoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EngageCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EngageCoinCallerSession struct {
	Contract *EngageCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EngageCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EngageCoinTransactorSession struct {
	Contract     *EngageCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EngageCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type EngageCoinRaw struct {
	Contract *EngageCoin // Generic contract binding to access the raw methods on
}

// EngageCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EngageCoinCallerRaw struct {
	Contract *EngageCoinCaller // Generic read-only contract binding to access the raw methods on
}

// EngageCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EngageCoinTransactorRaw struct {
	Contract *EngageCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEngageCoin creates a new instance of EngageCoin, bound to a specific deployed contract.
func NewEngageCoin(address common.Address, backend bind.ContractBackend) (*EngageCoin, error) {
	contract, err := bindEngageCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EngageCoin{EngageCoinCaller: EngageCoinCaller{contract: contract}, EngageCoinTransactor: EngageCoinTransactor{contract: contract}, EngageCoinFilterer: EngageCoinFilterer{contract: contract}}, nil
}

// NewEngageCoinCaller creates a new read-only instance of EngageCoin, bound to a specific deployed contract.
func NewEngageCoinCaller(address common.Address, caller bind.ContractCaller) (*EngageCoinCaller, error) {
	contract, err := bindEngageCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EngageCoinCaller{contract: contract}, nil
}

// NewEngageCoinTransactor creates a new write-only instance of EngageCoin, bound to a specific deployed contract.
func NewEngageCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*EngageCoinTransactor, error) {
	contract, err := bindEngageCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EngageCoinTransactor{contract: contract}, nil
}

// NewEngageCoinFilterer creates a new log filterer instance of EngageCoin, bound to a specific deployed contract.
func NewEngageCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*EngageCoinFilterer, error) {
	contract, err := bindEngageCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EngageCoinFilterer{contract: contract}, nil
}

// bindEngageCoin binds a generic wrapper to an already deployed contract.
func bindEngageCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EngageCoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EngageCoin *EngageCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EngageCoin.Contract.EngageCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EngageCoin *EngageCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EngageCoin.Contract.EngageCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EngageCoin *EngageCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EngageCoin.Contract.EngageCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EngageCoin *EngageCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EngageCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EngageCoin *EngageCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EngageCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EngageCoin *EngageCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EngageCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EngageCoin *EngageCoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EngageCoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EngageCoin *EngageCoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EngageCoin.Contract.Allowance(&_EngageCoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EngageCoin *EngageCoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EngageCoin.Contract.Allowance(&_EngageCoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EngageCoin *EngageCoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EngageCoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EngageCoin *EngageCoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EngageCoin.Contract.BalanceOf(&_EngageCoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EngageCoin *EngageCoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EngageCoin.Contract.BalanceOf(&_EngageCoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EngageCoin *EngageCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EngageCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EngageCoin *EngageCoinSession) Decimals() (uint8, error) {
	return _EngageCoin.Contract.Decimals(&_EngageCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EngageCoin *EngageCoinCallerSession) Decimals() (uint8, error) {
	return _EngageCoin.Contract.Decimals(&_EngageCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EngageCoin *EngageCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EngageCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EngageCoin *EngageCoinSession) Name() (string, error) {
	return _EngageCoin.Contract.Name(&_EngageCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EngageCoin *EngageCoinCallerSession) Name() (string, error) {
	return _EngageCoin.Contract.Name(&_EngageCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EngageCoin *EngageCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EngageCoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EngageCoin *EngageCoinSession) Owner() (common.Address, error) {
	return _EngageCoin.Contract.Owner(&_EngageCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EngageCoin *EngageCoinCallerSession) Owner() (common.Address, error) {
	return _EngageCoin.Contract.Owner(&_EngageCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EngageCoin *EngageCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EngageCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EngageCoin *EngageCoinSession) Symbol() (string, error) {
	return _EngageCoin.Contract.Symbol(&_EngageCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EngageCoin *EngageCoinCallerSession) Symbol() (string, error) {
	return _EngageCoin.Contract.Symbol(&_EngageCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EngageCoin *EngageCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EngageCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EngageCoin *EngageCoinSession) TotalSupply() (*big.Int, error) {
	return _EngageCoin.Contract.TotalSupply(&_EngageCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EngageCoin *EngageCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _EngageCoin.Contract.TotalSupply(&_EngageCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.Approve(&_EngageCoin.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.Approve(&_EngageCoin.TransactOpts, spender, value)
}

// MintReward is a paid mutator transaction binding the contract method 0x9a49090e.
//
// Solidity: function mintReward(address to, uint256 amount) returns()
func (_EngageCoin *EngageCoinTransactor) MintReward(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EngageCoin.contract.Transact(opts, "mintReward", to, amount)
}

// MintReward is a paid mutator transaction binding the contract method 0x9a49090e.
//
// Solidity: function mintReward(address to, uint256 amount) returns()
func (_EngageCoin *EngageCoinSession) MintReward(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.MintReward(&_EngageCoin.TransactOpts, to, amount)
}

// MintReward is a paid mutator transaction binding the contract method 0x9a49090e.
//
// Solidity: function mintReward(address to, uint256 amount) returns()
func (_EngageCoin *EngageCoinTransactorSession) MintReward(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.MintReward(&_EngageCoin.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EngageCoin *EngageCoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EngageCoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EngageCoin *EngageCoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _EngageCoin.Contract.RenounceOwnership(&_EngageCoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EngageCoin *EngageCoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EngageCoin.Contract.RenounceOwnership(&_EngageCoin.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.Transfer(&_EngageCoin.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.Transfer(&_EngageCoin.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.TransferFrom(&_EngageCoin.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_EngageCoin *EngageCoinTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EngageCoin.Contract.TransferFrom(&_EngageCoin.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EngageCoin *EngageCoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EngageCoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EngageCoin *EngageCoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EngageCoin.Contract.TransferOwnership(&_EngageCoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EngageCoin *EngageCoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EngageCoin.Contract.TransferOwnership(&_EngageCoin.TransactOpts, newOwner)
}

// EngageCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EngageCoin contract.
type EngageCoinApprovalIterator struct {
	Event *EngageCoinApproval // Event containing the contract specifics and raw log

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
func (it *EngageCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EngageCoinApproval)
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
		it.Event = new(EngageCoinApproval)
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
func (it *EngageCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EngageCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EngageCoinApproval represents a Approval event raised by the EngageCoin contract.
type EngageCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EngageCoin *EngageCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EngageCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EngageCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EngageCoinApprovalIterator{contract: _EngageCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EngageCoin *EngageCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EngageCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EngageCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EngageCoinApproval)
				if err := _EngageCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EngageCoin *EngageCoinFilterer) ParseApproval(log types.Log) (*EngageCoinApproval, error) {
	event := new(EngageCoinApproval)
	if err := _EngageCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EngageCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EngageCoin contract.
type EngageCoinOwnershipTransferredIterator struct {
	Event *EngageCoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EngageCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EngageCoinOwnershipTransferred)
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
		it.Event = new(EngageCoinOwnershipTransferred)
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
func (it *EngageCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EngageCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EngageCoinOwnershipTransferred represents a OwnershipTransferred event raised by the EngageCoin contract.
type EngageCoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EngageCoin *EngageCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EngageCoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EngageCoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EngageCoinOwnershipTransferredIterator{contract: _EngageCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EngageCoin *EngageCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EngageCoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EngageCoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EngageCoinOwnershipTransferred)
				if err := _EngageCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EngageCoin *EngageCoinFilterer) ParseOwnershipTransferred(log types.Log) (*EngageCoinOwnershipTransferred, error) {
	event := new(EngageCoinOwnershipTransferred)
	if err := _EngageCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EngageCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EngageCoin contract.
type EngageCoinTransferIterator struct {
	Event *EngageCoinTransfer // Event containing the contract specifics and raw log

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
func (it *EngageCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EngageCoinTransfer)
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
		it.Event = new(EngageCoinTransfer)
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
func (it *EngageCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EngageCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EngageCoinTransfer represents a Transfer event raised by the EngageCoin contract.
type EngageCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EngageCoin *EngageCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EngageCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EngageCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EngageCoinTransferIterator{contract: _EngageCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EngageCoin *EngageCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EngageCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EngageCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EngageCoinTransfer)
				if err := _EngageCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EngageCoin *EngageCoinFilterer) ParseTransfer(log types.Log) (*EngageCoinTransfer, error) {
	event := new(EngageCoinTransfer)
	if err := _EngageCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
