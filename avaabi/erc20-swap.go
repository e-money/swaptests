// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avaabi

import (
	"math/big"
	"strings"

	"github.com/ava-labs/coreth/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20AtomicSwapperABI is the input ABI used to generate the binding from.
const ERC20AtomicSwapperABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumber\",\"type\":\"bytes32\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_bep2Addr\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expireSeconds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bep2Amount\",\"type\":\"uint256\"}],\"name\":\"HTLT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20ContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_swapSender\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_bep2SenderAddr\",\"type\":\"bytes20\"}],\"name\":\"calSwapID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_randomNumber\",\"type\":\"bytes32\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_secondsSpan\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_bep2SenderAddr\",\"type\":\"bytes20\"},{\"internalType\":\"bytes20\",\"name\":\"_bep2RecipientAddr\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bep2Amount\",\"type\":\"uint256\"}],\"name\":\"htlt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"isSwapExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"queryOpenSwap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_expireSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"refundable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC20AtomicSwapperFuncSigs maps the 4-byte function signature to its string representation.
var ERC20AtomicSwapperFuncSigs = map[string]string{
	"49404437": "ERC20ContractAddr()",
	"7ef3e92e": "calSwapID(bytes32,address,bytes20)",
	"84cc9dfb": "claim(bytes32,bytes32)",
	"9b58e0a1": "claimable(bytes32)",
	"91fda287": "htlt(bytes32,uint64,uint256,address,bytes20,bytes20,uint256,uint256)",
	"50f7a03b": "isSwapExist(bytes32)",
	"b48017b1": "queryOpenSwap(bytes32)",
	"7249fbb6": "refund(bytes32)",
	"9fb31475": "refundable(bytes32)",
}

// ERC20AtomicSwapperBin is the compiled bytecode used for deploying new contracts.
var ERC20AtomicSwapperBin = "0x608060405234801561001057600080fd5b5060405161111b38038061111b8339818101604052602081101561003357600080fd5b5051600280546001600160a01b0319166001600160a01b039092169190911790556110b8806100636000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806384cc9dfb1161006657806384cc9dfb1461015857806391fda2871461017b5780639b58e0a1146101e45780639fb3147514610201578063b48017b11461021e57610093565b8063494044371461009857806350f7a03b146100bc5780637249fbb6146100ed5780637ef3e92e1461010a575b600080fd5b6100a0610283565b604080516001600160a01b039092168252519081900360200190f35b6100d9600480360360208110156100d257600080fd5b5035610292565b604080519115158252519081900360200190f35b6100d96004803603602081101561010357600080fd5b50356102ba565b6101466004803603606081101561012057600080fd5b5080359060208101356001600160a01b031690604001356001600160601b0319166104f4565b60408051918252519081900360200190f35b6100d96004803603604081101561016e57600080fd5b5080359060200135610690565b6100d9600480360361010081101561019257600080fd5b5080359067ffffffffffffffff602082013516906040810135906001600160a01b03606082013516906001600160601b0319608082013581169160a08101359091169060c08101359060e001356109fc565b6100d9600480360360208110156101fa57600080fd5b5035610e50565b6100d96004803603602081101561021757600080fd5b5035610e93565b61023b6004803603602081101561023457600080fd5b5035610eb6565b6040805196875267ffffffffffffffff90951660208701528585019390935260608501919091526001600160a01b0390811660808501521660a0830152519081900360c00190f35b6002546001600160a01b031681565b60008060008381526001602052604090205460ff1660038111156102b257fe5b141592915050565b600081600160008281526001602052604090205460ff1660038111156102dc57fe5b14610323576040805162461bcd60e51b81526020600482015260126024820152711cddd85c081a5cc81b9bdd081bdc195b995960721b604482015290519081900360640190fd5b6000838152602081905260409020600101548390421015610381576040805162461bcd60e51b81526020600482015260136024820152721cddd85c081a5cc81b9bdd08195e1c1a5c9959606a1b604482015290519081900360640190fd5b60008481526001602081815260408084208054600360ff19909116811790915584835281852090810180548254600280850180548a87559886018a90558990556001600160e01b03198316909355600493840180546001600160a01b03191690559154845163a9059cbb60e01b81526001600160a01b03600160401b9093048316948101859052602481018490529451939792969591169363a9059cbb9360448083019491928390030190829087803b15801561043d57600080fd5b505af1158015610451573d6000803e3d6000fd5b505050506040513d602081101561046757600080fd5b50516104a45760405162461bcd60e51b8152600401808060200182810382526034815260200180610fb66034913960400191505060405180910390fd5b60408051828152905188916001600160a01b0386169133917f04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9919081900360200190a45060019695505050505050565b60006001600160601b031982166105c7576002848460405160200180838152602001826001600160a01b031660601b8152601401925050506040516020818303038152906040526040518082805190602001908083835b6020831061056a5780518252601f19909201916020918201910161054b565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156105a9573d6000803e3d6000fd5b5050506040513d60208110156105be57600080fd5b50519050610689565b6040805160208082018790526001600160601b0319606087901b81168385015285166054830152825160488184030181526068909201928390528151600293918291908401908083835b602083106106305780518252601f199092019160209182019101610611565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561066f573d6000803e3d6000fd5b5050506040513d602081101561068457600080fd5b505190505b9392505050565b600082600160008281526001602052604090205460ff1660038111156106b257fe5b146106f9576040805162461bcd60e51b81526020600482015260126024820152711cddd85c081a5cc81b9bdd081bdc195b995960721b604482015290519081900360640190fd5b60008481526020819052604090206001015484904210610760576040805162461bcd60e51b815260206004820152601760248201527f7377617020697320616c72656164792065787069726564000000000000000000604482015290519081900360640190fd5b6000858152602081815260409182902060030154825180830188905260c09190911b6001600160c01b03191681840152825180820360280181526048909101928390528051889388936002939282918401908083835b602083106107d55780518252601f1990920191602091820191016107b6565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610814573d6000803e3d6000fd5b5050506040513d602081101561082957600080fd5b505160008381526020819052604090206002015414610886576040805162461bcd60e51b815260206004820152601460248201527334b73b30b634b2103930b73237b6a73ab6b132b960611b604482015290519081900360640190fd5b6000878152600160208181526040808420805460ff1916600290811790915584835281852060048082018054835484860180548b87559986018b90558a9055600390940180546001600160e01b03191690556001600160a01b031981169091559254845163a9059cbb60e01b81526001600160a01b03948516928101839052602481018490529451919792969593169363a9059cbb936044808301949193928390030190829087803b15801561093b57600080fd5b505af115801561094f573d6000803e3d6000fd5b505050506040513d602081101561096557600080fd5b50516109a25760405162461bcd60e51b815260040180806020018281038252602c815260200180611010602c913960400191505060405180910390fd5b60408051828152602081018b905281518c926001600160a01b0387169233927f9f46b1606087bdf4183ec7dfdbe68e4ab9129a6a37901c16a7b320ae11a96018929181900390910190a45060019998505050505050505050565b600080610a0a8a33886104f4565b90506000808281526001602052604090205460ff166003811115610a2a57fe5b14610a7c576040805162461bcd60e51b815260206004820152601960248201527f73776170206973206f70656e65642070726576696f75736c7900000000000000604482015290519081900360640190fd5b603c8810158015610a90575062093a808811155b610acb5760405162461bcd60e51b8152600401808060200182810382526026815260200180610fea6026913960400191505060405180910390fd5b6001600160a01b038716610b105760405162461bcd60e51b815260040180806020018281038252602181526020018061103c6021913960400191505060405180910390fd5b60008411610b65576040805162461bcd60e51b815260206004820152601e60248201527f5f6f7574416d6f756e74206d757374206265206d6f7265207468616e20300000604482015290519081900360640190fd5b42888a67ffffffffffffffff160111610baf5760405162461bcd60e51b815260040180806020018281038252602681526020018061105d6026913960400191505060405180910390fd5b610bb7610f48565b6040518060c001604052808681526020018b67ffffffffffffffff168b0181526020018c81526020018b67ffffffffffffffff168152602001336001600160a01b03168152602001896001600160a01b031681525090508060008084815260200190815260200160002060008201518160000155602082015181600101556040820151816002015560608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550905050600180600084815260200190815260200160002060006101000a81548160ff02191690836003811115610cfa57fe5b0217905550600254604080516323b872dd60e01b81523360048201523060248201526044810188905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b158015610d5957600080fd5b505af1158015610d6d573d6000803e3d6000fd5b505050506040513d6020811015610d8357600080fd5b5051610dc05760405162461bcd60e51b8152600401808060200182810382526038815260200180610f7e6038913960400191505060405180910390fd5b602080820151604080518e815267ffffffffffffffff8e16938101939093526001600160601b031989168382015260608301919091526080820187905260a082018690525183916001600160a01b038b169133917fb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1919081900360c00190a45060019a9950505050505050505050565b60008181526020819052604081206001015442108015610e8d575060015b60008381526001602052604090205460ff166003811115610e8b57fe5b145b92915050565b6000818152602081905260408120600101544210801590610e8d57506001610e6e565b600080600080600080610ec7610f48565b50505060009485525050506020828152604092839020835160c081018552815480825260018301549382018490526002830154958201869052600383015467ffffffffffffffff811660608401819052600160401b9091046001600160a01b03908116608085018190526004909501541660a0909301839052959693945092565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091529056fe6661696c656420746f207472616e7366657220636c69656e7420617373657420746f207377617020636f6e747261637420616464726573734661696c656420746f207472616e73666572206c6f636b6564206173736574206261636b20746f20737761702063726561746f725f7365636f6e64735370616e2073686f756c6420626520696e205b36302c203630343830305d4661696c656420746f207472616e73666572206c6f636b656420617373657420746f20726563697069656e745f726563697069656e74416464722073686f756c64206e6f74206265207a65726f4e6577207377617073206e65656420746f2065787069726520696e2074686520667574757265a264697066735822122037e92017c82d7ca57215baa669cccb0f0f2e2aa8bc7e8b3cea912fc5f3c44a6f64736f6c634300060c0033"

// DeployERC20AtomicSwapper deploys a new Ethereum contract, binding an instance of ERC20AtomicSwapper to it.
func DeployERC20AtomicSwapper(auth *TransactOpts, backend ContractBackend, _erc20Contract common.Address) (common.Address, *types.Transaction, *ERC20AtomicSwapper, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AtomicSwapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := DeployContract(auth, parsed, common.FromHex(ERC20AtomicSwapperBin), backend, _erc20Contract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20AtomicSwapper{ERC20AtomicSwapperCaller: ERC20AtomicSwapperCaller{contract: contract}, ERC20AtomicSwapperTransactor: ERC20AtomicSwapperTransactor{contract: contract}, ERC20AtomicSwapperFilterer: ERC20AtomicSwapperFilterer{contract: contract}}, nil
}

// ERC20AtomicSwapper is an auto generated Go binding around an Ethereum contract.
type ERC20AtomicSwapper struct {
	ERC20AtomicSwapperCaller     // Read-only binding to the contract
	ERC20AtomicSwapperTransactor // Write-only binding to the contract
	ERC20AtomicSwapperFilterer   // Log filterer for contract events
}

// ERC20AtomicSwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20AtomicSwapperCaller struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AtomicSwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20AtomicSwapperTransactor struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AtomicSwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20AtomicSwapperFilterer struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AtomicSwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20AtomicSwapperSession struct {
	Contract     *ERC20AtomicSwapper // Generic contract binding to set the session for
	CallOpts     CallOpts            // Call options to use throughout this session
	TransactOpts TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20AtomicSwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20AtomicSwapperCallerSession struct {
	Contract *ERC20AtomicSwapperCaller // Generic contract caller binding to set the session for
	CallOpts CallOpts                  // Call options to use throughout this session
}

// ERC20AtomicSwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20AtomicSwapperTransactorSession struct {
	Contract     *ERC20AtomicSwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC20AtomicSwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20AtomicSwapperRaw struct {
	Contract *ERC20AtomicSwapper // Generic contract binding to access the raw methods on
}

// ERC20AtomicSwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20AtomicSwapperCallerRaw struct {
	Contract *ERC20AtomicSwapperCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20AtomicSwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20AtomicSwapperTransactorRaw struct {
	Contract *ERC20AtomicSwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20AtomicSwapper creates a new instance of ERC20AtomicSwapper, bound to a specific deployed contract.
func NewERC20AtomicSwapper(address common.Address, backend ContractBackend) (*ERC20AtomicSwapper, error) {
	contract, err := bindERC20AtomicSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20AtomicSwapper{ERC20AtomicSwapperCaller: ERC20AtomicSwapperCaller{contract: contract}, ERC20AtomicSwapperTransactor: ERC20AtomicSwapperTransactor{contract: contract}, ERC20AtomicSwapperFilterer: ERC20AtomicSwapperFilterer{contract: contract}}, nil
}

// NewERC20AtomicSwapperCaller creates a new read-only instance of ERC20AtomicSwapper, bound to a specific deployed contract.
func NewERC20AtomicSwapperCaller(address common.Address, caller ContractCaller) (*ERC20AtomicSwapperCaller, error) {
	contract, err := bindERC20AtomicSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AtomicSwapperCaller{contract: contract}, nil
}

// NewERC20AtomicSwapperTransactor creates a new write-only instance of ERC20AtomicSwapper, bound to a specific deployed contract.
func NewERC20AtomicSwapperTransactor(address common.Address, transactor ContractTransactor) (*ERC20AtomicSwapperTransactor, error) {
	contract, err := bindERC20AtomicSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AtomicSwapperTransactor{contract: contract}, nil
}

// NewERC20AtomicSwapperFilterer creates a new log filterer instance of ERC20AtomicSwapper, bound to a specific deployed contract.
func NewERC20AtomicSwapperFilterer(address common.Address, filterer ContractFilterer) (*ERC20AtomicSwapperFilterer, error) {
	contract, err := bindERC20AtomicSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20AtomicSwapperFilterer{contract: contract}, nil
}

// bindERC20AtomicSwapper binds a generic wrapper to an already deployed contract.
func bindERC20AtomicSwapper(address common.Address, caller ContractCaller, transactor ContractTransactor, filterer ContractFilterer) (*BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AtomicSwapperABI))
	if err != nil {
		return nil, err
	}
	return NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AtomicSwapper *ERC20AtomicSwapperRaw) Call(opts *CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AtomicSwapper.Contract.ERC20AtomicSwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AtomicSwapper *ERC20AtomicSwapperRaw) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.ERC20AtomicSwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AtomicSwapper *ERC20AtomicSwapperRaw) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.ERC20AtomicSwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCallerRaw) Call(opts *CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AtomicSwapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactorRaw) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactorRaw) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.contract.Transact(opts, method, params...)
}

// ERC20ContractAddr is a free data retrieval call binding the contract method 0x49404437.
//
// Solidity: function ERC20ContractAddr() view returns(address)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCaller) ERC20ContractAddr(opts *CallOpts) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _ERC20AtomicSwapper.contract.Call(opts, out, "ERC20ContractAddr")
	return *ret0, err
}

// ERC20ContractAddr is a free data retrieval call binding the contract method 0x49404437.
//
// Solidity: function ERC20ContractAddr() view returns(address)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) ERC20ContractAddr() (common.Address, error) {
	return _ERC20AtomicSwapper.Contract.ERC20ContractAddr(&_ERC20AtomicSwapper.CallOpts)
}

// ERC20ContractAddr is a free data retrieval call binding the contract method 0x49404437.
//
// Solidity: function ERC20ContractAddr() view returns(address)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCallerSession) ERC20ContractAddr() (common.Address, error) {
	return _ERC20AtomicSwapper.Contract.ERC20ContractAddr(&_ERC20AtomicSwapper.CallOpts)
}

// CalSwapID is a free data retrieval call binding the contract method 0x7ef3e92e.
//
// Solidity: function calSwapID(bytes32 _randomNumberHash, address _swapSender, bytes20 _bep2SenderAddr) pure returns(bytes32)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCaller) CalSwapID(opts *CallOpts, _randomNumberHash [32]byte, _swapSender common.Address, _bep2SenderAddr [20]byte) ([32]byte, error) {
	ret0 := new([32]byte)
	out := ret0
	err := _ERC20AtomicSwapper.contract.Call(opts, out, "calSwapID", _randomNumberHash, _swapSender, _bep2SenderAddr)
	return *ret0, err
}

// CalSwapID is a free data retrieval call binding the contract method 0x7ef3e92e.
//
// Solidity: function calSwapID(bytes32 _randomNumberHash, address _swapSender, bytes20 _bep2SenderAddr) pure returns(bytes32)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) CalSwapID(_randomNumberHash [32]byte, _swapSender common.Address, _bep2SenderAddr [20]byte) ([32]byte, error) {
	return _ERC20AtomicSwapper.Contract.CalSwapID(&_ERC20AtomicSwapper.CallOpts, _randomNumberHash, _swapSender, _bep2SenderAddr)
}

// CalSwapID is a free data retrieval call binding the contract method 0x7ef3e92e.
//
// Solidity: function calSwapID(bytes32 _randomNumberHash, address _swapSender, bytes20 _bep2SenderAddr) pure returns(bytes32)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCallerSession) CalSwapID(_randomNumberHash [32]byte, _swapSender common.Address, _bep2SenderAddr [20]byte) ([32]byte, error) {
	return _ERC20AtomicSwapper.Contract.CalSwapID(&_ERC20AtomicSwapper.CallOpts, _randomNumberHash, _swapSender, _bep2SenderAddr)
}

// Claimable is a free data retrieval call binding the contract method 0x9b58e0a1.
//
// Solidity: function claimable(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCaller) Claimable(opts *CallOpts, _swapID [32]byte) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _ERC20AtomicSwapper.contract.Call(opts, out, "claimable", _swapID)
	return *ret0, err
}

// Claimable is a free data retrieval call binding the contract method 0x9b58e0a1.
//
// Solidity: function claimable(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) Claimable(_swapID [32]byte) (bool, error) {
	return _ERC20AtomicSwapper.Contract.Claimable(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// Claimable is a free data retrieval call binding the contract method 0x9b58e0a1.
//
// Solidity: function claimable(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCallerSession) Claimable(_swapID [32]byte) (bool, error) {
	return _ERC20AtomicSwapper.Contract.Claimable(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// IsSwapExist is a free data retrieval call binding the contract method 0x50f7a03b.
//
// Solidity: function isSwapExist(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCaller) IsSwapExist(opts *CallOpts, _swapID [32]byte) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _ERC20AtomicSwapper.contract.Call(opts, out, "isSwapExist", _swapID)
	return *ret0, err
}

// IsSwapExist is a free data retrieval call binding the contract method 0x50f7a03b.
//
// Solidity: function isSwapExist(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) IsSwapExist(_swapID [32]byte) (bool, error) {
	return _ERC20AtomicSwapper.Contract.IsSwapExist(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// IsSwapExist is a free data retrieval call binding the contract method 0x50f7a03b.
//
// Solidity: function isSwapExist(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCallerSession) IsSwapExist(_swapID [32]byte) (bool, error) {
	return _ERC20AtomicSwapper.Contract.IsSwapExist(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// QueryOpenSwap is a free data retrieval call binding the contract method 0xb48017b1.
//
// Solidity: function queryOpenSwap(bytes32 _swapID) view returns(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _expireSeconds, uint256 _outAmount, address _sender, address _recipient)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCaller) QueryOpenSwap(opts *CallOpts, _swapID [32]byte) (struct {
	RandomNumberHash [32]byte
	Timestamp        uint64
	ExpireSeconds    *big.Int
	OutAmount        *big.Int
	Sender           common.Address
	Recipient        common.Address
}, error) {
	ret := new(struct {
		RandomNumberHash [32]byte
		Timestamp        uint64
		ExpireSeconds    *big.Int
		OutAmount        *big.Int
		Sender           common.Address
		Recipient        common.Address
	})
	out := ret
	err := _ERC20AtomicSwapper.contract.Call(opts, out, "queryOpenSwap", _swapID)
	return *ret, err
}

// QueryOpenSwap is a free data retrieval call binding the contract method 0xb48017b1.
//
// Solidity: function queryOpenSwap(bytes32 _swapID) view returns(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _expireSeconds, uint256 _outAmount, address _sender, address _recipient)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) QueryOpenSwap(_swapID [32]byte) (struct {
	RandomNumberHash [32]byte
	Timestamp        uint64
	ExpireSeconds    *big.Int
	OutAmount        *big.Int
	Sender           common.Address
	Recipient        common.Address
}, error) {
	return _ERC20AtomicSwapper.Contract.QueryOpenSwap(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// QueryOpenSwap is a free data retrieval call binding the contract method 0xb48017b1.
//
// Solidity: function queryOpenSwap(bytes32 _swapID) view returns(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _expireSeconds, uint256 _outAmount, address _sender, address _recipient)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCallerSession) QueryOpenSwap(_swapID [32]byte) (struct {
	RandomNumberHash [32]byte
	Timestamp        uint64
	ExpireSeconds    *big.Int
	OutAmount        *big.Int
	Sender           common.Address
	Recipient        common.Address
}, error) {
	return _ERC20AtomicSwapper.Contract.QueryOpenSwap(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// Refundable is a free data retrieval call binding the contract method 0x9fb31475.
//
// Solidity: function refundable(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCaller) Refundable(opts *CallOpts, _swapID [32]byte) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _ERC20AtomicSwapper.contract.Call(opts, out, "refundable", _swapID)
	return *ret0, err
}

// Refundable is a free data retrieval call binding the contract method 0x9fb31475.
//
// Solidity: function refundable(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) Refundable(_swapID [32]byte) (bool, error) {
	return _ERC20AtomicSwapper.Contract.Refundable(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// Refundable is a free data retrieval call binding the contract method 0x9fb31475.
//
// Solidity: function refundable(bytes32 _swapID) view returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperCallerSession) Refundable(_swapID [32]byte) (bool, error) {
	return _ERC20AtomicSwapper.Contract.Refundable(&_ERC20AtomicSwapper.CallOpts, _swapID)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 _swapID, bytes32 _randomNumber) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactor) Claim(opts *TransactOpts, _swapID [32]byte, _randomNumber [32]byte) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.contract.Transact(opts, "claim", _swapID, _randomNumber)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 _swapID, bytes32 _randomNumber) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) Claim(_swapID [32]byte, _randomNumber [32]byte) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.Claim(&_ERC20AtomicSwapper.TransactOpts, _swapID, _randomNumber)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 _swapID, bytes32 _randomNumber) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactorSession) Claim(_swapID [32]byte, _randomNumber [32]byte) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.Claim(&_ERC20AtomicSwapper.TransactOpts, _swapID, _randomNumber)
}

// Htlt is a paid mutator transaction binding the contract method 0x91fda287.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _secondsSpan, address _recipientAddr, bytes20 _bep2SenderAddr, bytes20 _bep2RecipientAddr, uint256 _outAmount, uint256 _bep2Amount) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactor) Htlt(opts *TransactOpts, _randomNumberHash [32]byte, _timestamp uint64, _secondsSpan *big.Int, _recipientAddr common.Address, _bep2SenderAddr [20]byte, _bep2RecipientAddr [20]byte, _outAmount *big.Int, _bep2Amount *big.Int) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.contract.Transact(opts, "htlt", _randomNumberHash, _timestamp, _secondsSpan, _recipientAddr, _bep2SenderAddr, _bep2RecipientAddr, _outAmount, _bep2Amount)
}

// Htlt is a paid mutator transaction binding the contract method 0x91fda287.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _secondsSpan, address _recipientAddr, bytes20 _bep2SenderAddr, bytes20 _bep2RecipientAddr, uint256 _outAmount, uint256 _bep2Amount) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) Htlt(_randomNumberHash [32]byte, _timestamp uint64, _secondsSpan *big.Int, _recipientAddr common.Address, _bep2SenderAddr [20]byte, _bep2RecipientAddr [20]byte, _outAmount *big.Int, _bep2Amount *big.Int) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.Htlt(&_ERC20AtomicSwapper.TransactOpts, _randomNumberHash, _timestamp, _secondsSpan, _recipientAddr, _bep2SenderAddr, _bep2RecipientAddr, _outAmount, _bep2Amount)
}

// Htlt is a paid mutator transaction binding the contract method 0x91fda287.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _secondsSpan, address _recipientAddr, bytes20 _bep2SenderAddr, bytes20 _bep2RecipientAddr, uint256 _outAmount, uint256 _bep2Amount) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactorSession) Htlt(_randomNumberHash [32]byte, _timestamp uint64, _secondsSpan *big.Int, _recipientAddr common.Address, _bep2SenderAddr [20]byte, _bep2RecipientAddr [20]byte, _outAmount *big.Int, _bep2Amount *big.Int) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.Htlt(&_ERC20AtomicSwapper.TransactOpts, _randomNumberHash, _timestamp, _secondsSpan, _recipientAddr, _bep2SenderAddr, _bep2RecipientAddr, _outAmount, _bep2Amount)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _swapID) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactor) Refund(opts *TransactOpts, _swapID [32]byte) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.contract.Transact(opts, "refund", _swapID)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _swapID) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperSession) Refund(_swapID [32]byte) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.Refund(&_ERC20AtomicSwapper.TransactOpts, _swapID)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _swapID) returns(bool)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperTransactorSession) Refund(_swapID [32]byte) (*types.Transaction, error) {
	return _ERC20AtomicSwapper.Contract.Refund(&_ERC20AtomicSwapper.TransactOpts, _swapID)
}

// ERC20AtomicSwapperClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the ERC20AtomicSwapper contract.
type ERC20AtomicSwapperClaimedIterator struct {
	Event *ERC20AtomicSwapperClaimed // Event containing the contract specifics and raw log

	contract *BoundContract // Generic contract to use for unpacking event data
	event    string         // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20AtomicSwapperClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20AtomicSwapperClaimed)
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
		it.Event = new(ERC20AtomicSwapperClaimed)
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
func (it *ERC20AtomicSwapperClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20AtomicSwapperClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20AtomicSwapperClaimed represents a Claimed event raised by the ERC20AtomicSwapper contract.
type ERC20AtomicSwapperClaimed struct {
	MsgSender        common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	RandomNumber     [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9f46b1606087bdf4183ec7dfdbe68e4ab9129a6a37901c16a7b320ae11a96018.
//
// Solidity: event Claimed(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, bytes32 _randomNumber)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) FilterClaimed(opts *FilterOpts, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ERC20AtomicSwapperClaimedIterator, error) {
	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20AtomicSwapper.contract.FilterLogs(opts, "Claimed", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20AtomicSwapperClaimedIterator{contract: _ERC20AtomicSwapper.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9f46b1606087bdf4183ec7dfdbe68e4ab9129a6a37901c16a7b320ae11a96018.
//
// Solidity: event Claimed(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, bytes32 _randomNumber)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) WatchClaimed(opts *WatchOpts, sink chan<- *ERC20AtomicSwapperClaimed, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {
	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20AtomicSwapper.contract.WatchLogs(opts, "Claimed", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20AtomicSwapperClaimed)
				if err := _ERC20AtomicSwapper.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9f46b1606087bdf4183ec7dfdbe68e4ab9129a6a37901c16a7b320ae11a96018.
//
// Solidity: event Claimed(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, bytes32 _randomNumber)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) ParseClaimed(log types.Log) (*ERC20AtomicSwapperClaimed, error) {
	event := new(ERC20AtomicSwapperClaimed)
	if err := _ERC20AtomicSwapper.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20AtomicSwapperHTLTIterator is returned from FilterHTLT and is used to iterate over the raw logs and unpacked data for HTLT events raised by the ERC20AtomicSwapper contract.
type ERC20AtomicSwapperHTLTIterator struct {
	Event *ERC20AtomicSwapperHTLT // Event containing the contract specifics and raw log

	contract *BoundContract // Generic contract to use for unpacking event data
	event    string         // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20AtomicSwapperHTLTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20AtomicSwapperHTLT)
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
		it.Event = new(ERC20AtomicSwapperHTLT)
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
func (it *ERC20AtomicSwapperHTLTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20AtomicSwapperHTLTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20AtomicSwapperHTLT represents a HTLT event raised by the ERC20AtomicSwapper contract.
type ERC20AtomicSwapperHTLT struct {
	MsgSender        common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	Timestamp        uint64
	Bep2Addr         [20]byte
	ExpireSeconds    *big.Int
	OutAmount        *big.Int
	Bep2Amount       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHTLT is a free log retrieval operation binding the contract event 0xb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1.
//
// Solidity: event HTLT(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _bep2Addr, uint256 _expireSeconds, uint256 _outAmount, uint256 _bep2Amount)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) FilterHTLT(opts *FilterOpts, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ERC20AtomicSwapperHTLTIterator, error) {
	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20AtomicSwapper.contract.FilterLogs(opts, "HTLT", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20AtomicSwapperHTLTIterator{contract: _ERC20AtomicSwapper.contract, event: "HTLT", logs: logs, sub: sub}, nil
}

// WatchHTLT is a free log subscription operation binding the contract event 0xb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1.
//
// Solidity: event HTLT(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _bep2Addr, uint256 _expireSeconds, uint256 _outAmount, uint256 _bep2Amount)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) WatchHTLT(opts *WatchOpts, sink chan<- *ERC20AtomicSwapperHTLT, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {
	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20AtomicSwapper.contract.WatchLogs(opts, "HTLT", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20AtomicSwapperHTLT)
				if err := _ERC20AtomicSwapper.contract.UnpackLog(event, "HTLT", log); err != nil {
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

// ParseHTLT is a log parse operation binding the contract event 0xb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1.
//
// Solidity: event HTLT(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _bep2Addr, uint256 _expireSeconds, uint256 _outAmount, uint256 _bep2Amount)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) ParseHTLT(log types.Log) (*ERC20AtomicSwapperHTLT, error) {
	event := new(ERC20AtomicSwapperHTLT)
	if err := _ERC20AtomicSwapper.contract.UnpackLog(event, "HTLT", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20AtomicSwapperRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the ERC20AtomicSwapper contract.
type ERC20AtomicSwapperRefundedIterator struct {
	Event *ERC20AtomicSwapperRefunded // Event containing the contract specifics and raw log

	contract *BoundContract // Generic contract to use for unpacking event data
	event    string         // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20AtomicSwapperRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20AtomicSwapperRefunded)
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
		it.Event = new(ERC20AtomicSwapperRefunded)
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
func (it *ERC20AtomicSwapperRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20AtomicSwapperRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20AtomicSwapperRefunded represents a Refunded event raised by the ERC20AtomicSwapper contract.
type ERC20AtomicSwapperRefunded struct {
	MsgSender        common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9.
//
// Solidity: event Refunded(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) FilterRefunded(opts *FilterOpts, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ERC20AtomicSwapperRefundedIterator, error) {
	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20AtomicSwapper.contract.FilterLogs(opts, "Refunded", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20AtomicSwapperRefundedIterator{contract: _ERC20AtomicSwapper.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9.
//
// Solidity: event Refunded(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) WatchRefunded(opts *WatchOpts, sink chan<- *ERC20AtomicSwapperRefunded, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {
	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20AtomicSwapper.contract.WatchLogs(opts, "Refunded", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20AtomicSwapperRefunded)
				if err := _ERC20AtomicSwapper.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9.
//
// Solidity: event Refunded(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash)
func (_ERC20AtomicSwapper *ERC20AtomicSwapperFilterer) ParseRefunded(log types.Log) (*ERC20AtomicSwapperRefunded, error) {
	event := new(ERC20AtomicSwapperRefunded)
	if err := _ERC20AtomicSwapper.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20      // Generic contract binding to set the session for
	CallOpts     CallOpts     // Call options to use throughout this session
	TransactOpts TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts CallOpts      // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts TransactOpts      // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller ContractCaller, transactor ContractTransactor, filterer ContractFilterer) (*BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *CallOpts, account common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *BoundContract // Generic contract to use for unpacking event data
	event    string         // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *BoundContract // Generic contract to use for unpacking event data
	event    string         // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
