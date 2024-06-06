// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pythrebase

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

// PythrebaseMetaData contains all meta data concerning the Pythrebase contract.
var PythrebaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pyth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidExpo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"numerator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"denominator\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"int32\",\"name\":\"multiplier\",\"type\":\"int32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"numerator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"denominator\",\"type\":\"bytes32\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PythrebaseABI is the input ABI used to generate the binding from.
// Deprecated: Use PythrebaseMetaData.ABI instead.
var PythrebaseABI = PythrebaseMetaData.ABI

// Pythrebase is an auto generated Go binding around an Ethereum contract.
type Pythrebase struct {
	PythrebaseCaller     // Read-only binding to the contract
	PythrebaseTransactor // Write-only binding to the contract
	PythrebaseFilterer   // Log filterer for contract events
}

// PythrebaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type PythrebaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythrebaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PythrebaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythrebaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PythrebaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythrebaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PythrebaseSession struct {
	Contract     *Pythrebase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PythrebaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PythrebaseCallerSession struct {
	Contract *PythrebaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PythrebaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PythrebaseTransactorSession struct {
	Contract     *PythrebaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PythrebaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type PythrebaseRaw struct {
	Contract *Pythrebase // Generic contract binding to access the raw methods on
}

// PythrebaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PythrebaseCallerRaw struct {
	Contract *PythrebaseCaller // Generic read-only contract binding to access the raw methods on
}

// PythrebaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PythrebaseTransactorRaw struct {
	Contract *PythrebaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPythrebase creates a new instance of Pythrebase, bound to a specific deployed contract.
func NewPythrebase(address common.Address, backend bind.ContractBackend) (*Pythrebase, error) {
	contract, err := bindPythrebase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pythrebase{PythrebaseCaller: PythrebaseCaller{contract: contract}, PythrebaseTransactor: PythrebaseTransactor{contract: contract}, PythrebaseFilterer: PythrebaseFilterer{contract: contract}}, nil
}

// NewPythrebaseCaller creates a new read-only instance of Pythrebase, bound to a specific deployed contract.
func NewPythrebaseCaller(address common.Address, caller bind.ContractCaller) (*PythrebaseCaller, error) {
	contract, err := bindPythrebase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PythrebaseCaller{contract: contract}, nil
}

// NewPythrebaseTransactor creates a new write-only instance of Pythrebase, bound to a specific deployed contract.
func NewPythrebaseTransactor(address common.Address, transactor bind.ContractTransactor) (*PythrebaseTransactor, error) {
	contract, err := bindPythrebase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PythrebaseTransactor{contract: contract}, nil
}

// NewPythrebaseFilterer creates a new log filterer instance of Pythrebase, bound to a specific deployed contract.
func NewPythrebaseFilterer(address common.Address, filterer bind.ContractFilterer) (*PythrebaseFilterer, error) {
	contract, err := bindPythrebase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PythrebaseFilterer{contract: contract}, nil
}

// bindPythrebase binds a generic wrapper to an already deployed contract.
func bindPythrebase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PythrebaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pythrebase *PythrebaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pythrebase.Contract.PythrebaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pythrebase *PythrebaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pythrebase.Contract.PythrebaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pythrebase *PythrebaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pythrebase.Contract.PythrebaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pythrebase *PythrebaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pythrebase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pythrebase *PythrebaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pythrebase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pythrebase *PythrebaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pythrebase.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x0776f244.
//
// Solidity: function getPrice(bytes32 numerator, bytes32 denominator) view returns(int256)
func (_Pythrebase *PythrebaseCaller) GetPrice(opts *bind.CallOpts, numerator [32]byte, denominator [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Pythrebase.contract.Call(opts, &out, "getPrice", numerator, denominator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x0776f244.
//
// Solidity: function getPrice(bytes32 numerator, bytes32 denominator) view returns(int256)
func (_Pythrebase *PythrebaseSession) GetPrice(numerator [32]byte, denominator [32]byte) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice(&_Pythrebase.CallOpts, numerator, denominator)
}

// GetPrice is a free data retrieval call binding the contract method 0x0776f244.
//
// Solidity: function getPrice(bytes32 numerator, bytes32 denominator) view returns(int256)
func (_Pythrebase *PythrebaseCallerSession) GetPrice(numerator [32]byte, denominator [32]byte) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice(&_Pythrebase.CallOpts, numerator, denominator)
}

// GetPrice0 is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns(int256)
func (_Pythrebase *PythrebaseCaller) GetPrice0(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Pythrebase.contract.Call(opts, &out, "getPrice0", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice0 is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns(int256)
func (_Pythrebase *PythrebaseSession) GetPrice0(id [32]byte) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice0(&_Pythrebase.CallOpts, id)
}

// GetPrice0 is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns(int256)
func (_Pythrebase *PythrebaseCallerSession) GetPrice0(id [32]byte) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice0(&_Pythrebase.CallOpts, id)
}

// GetPrice1 is a free data retrieval call binding the contract method 0x566881a4.
//
// Solidity: function getPrice(bytes32 id, int32 multiplier) view returns(int256)
func (_Pythrebase *PythrebaseCaller) GetPrice1(opts *bind.CallOpts, id [32]byte, multiplier int32) (*big.Int, error) {
	var out []interface{}
	err := _Pythrebase.contract.Call(opts, &out, "getPrice1", id, multiplier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice1 is a free data retrieval call binding the contract method 0x566881a4.
//
// Solidity: function getPrice(bytes32 id, int32 multiplier) view returns(int256)
func (_Pythrebase *PythrebaseSession) GetPrice1(id [32]byte, multiplier int32) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice1(&_Pythrebase.CallOpts, id, multiplier)
}

// GetPrice1 is a free data retrieval call binding the contract method 0x566881a4.
//
// Solidity: function getPrice(bytes32 id, int32 multiplier) view returns(int256)
func (_Pythrebase *PythrebaseCallerSession) GetPrice1(id [32]byte, multiplier int32) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice1(&_Pythrebase.CallOpts, id, multiplier)
}

// GetPrice2 is a free data retrieval call binding the contract method 0x66887067.
//
// Solidity: function getPrice(bytes32 numerator, bytes32 denominator, int32 expo) view returns(int256)
func (_Pythrebase *PythrebaseCaller) GetPrice2(opts *bind.CallOpts, numerator [32]byte, denominator [32]byte, expo int32) (*big.Int, error) {
	var out []interface{}
	err := _Pythrebase.contract.Call(opts, &out, "getPrice2", numerator, denominator, expo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice2 is a free data retrieval call binding the contract method 0x66887067.
//
// Solidity: function getPrice(bytes32 numerator, bytes32 denominator, int32 expo) view returns(int256)
func (_Pythrebase *PythrebaseSession) GetPrice2(numerator [32]byte, denominator [32]byte, expo int32) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice2(&_Pythrebase.CallOpts, numerator, denominator, expo)
}

// GetPrice2 is a free data retrieval call binding the contract method 0x66887067.
//
// Solidity: function getPrice(bytes32 numerator, bytes32 denominator, int32 expo) view returns(int256)
func (_Pythrebase *PythrebaseCallerSession) GetPrice2(numerator [32]byte, denominator [32]byte, expo int32) (*big.Int, error) {
	return _Pythrebase.Contract.GetPrice2(&_Pythrebase.CallOpts, numerator, denominator, expo)
}
