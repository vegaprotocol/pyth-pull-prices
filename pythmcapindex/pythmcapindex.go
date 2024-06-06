// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pythmcapindex

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

// IndexAsset is an auto generated low-level Go binding around an user-defined struct.
type IndexAsset struct {
	PythId      [32]byte
	TotalSupply *big.Int
}

// PythmcapindexMetaData contains all meta data concerning the Pythmcapindex contract.
var PythmcapindexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pyth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidExpo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTopN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pythId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"totalSupply\",\"type\":\"uint96\"}],\"internalType\":\"structIndexAsset[]\",\"name\":\"indexAssets\",\"type\":\"tuple[]\"}],\"name\":\"getIndexPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pythId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"totalSupply\",\"type\":\"uint96\"}],\"internalType\":\"structIndexAsset[]\",\"name\":\"indexAssets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"topN\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"}],\"name\":\"getIndexPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pythId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"totalSupply\",\"type\":\"uint96\"}],\"internalType\":\"structIndexAsset[]\",\"name\":\"indexAssets\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"baseAssetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"topN\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"}],\"name\":\"getIndexPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PythmcapindexABI is the input ABI used to generate the binding from.
// Deprecated: Use PythmcapindexMetaData.ABI instead.
var PythmcapindexABI = PythmcapindexMetaData.ABI

// Pythmcapindex is an auto generated Go binding around an Ethereum contract.
type Pythmcapindex struct {
	PythmcapindexCaller     // Read-only binding to the contract
	PythmcapindexTransactor // Write-only binding to the contract
	PythmcapindexFilterer   // Log filterer for contract events
}

// PythmcapindexCaller is an auto generated read-only Go binding around an Ethereum contract.
type PythmcapindexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythmcapindexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PythmcapindexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythmcapindexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PythmcapindexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythmcapindexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PythmcapindexSession struct {
	Contract     *Pythmcapindex    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PythmcapindexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PythmcapindexCallerSession struct {
	Contract *PythmcapindexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PythmcapindexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PythmcapindexTransactorSession struct {
	Contract     *PythmcapindexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PythmcapindexRaw is an auto generated low-level Go binding around an Ethereum contract.
type PythmcapindexRaw struct {
	Contract *Pythmcapindex // Generic contract binding to access the raw methods on
}

// PythmcapindexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PythmcapindexCallerRaw struct {
	Contract *PythmcapindexCaller // Generic read-only contract binding to access the raw methods on
}

// PythmcapindexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PythmcapindexTransactorRaw struct {
	Contract *PythmcapindexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPythmcapindex creates a new instance of Pythmcapindex, bound to a specific deployed contract.
func NewPythmcapindex(address common.Address, backend bind.ContractBackend) (*Pythmcapindex, error) {
	contract, err := bindPythmcapindex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pythmcapindex{PythmcapindexCaller: PythmcapindexCaller{contract: contract}, PythmcapindexTransactor: PythmcapindexTransactor{contract: contract}, PythmcapindexFilterer: PythmcapindexFilterer{contract: contract}}, nil
}

// NewPythmcapindexCaller creates a new read-only instance of Pythmcapindex, bound to a specific deployed contract.
func NewPythmcapindexCaller(address common.Address, caller bind.ContractCaller) (*PythmcapindexCaller, error) {
	contract, err := bindPythmcapindex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PythmcapindexCaller{contract: contract}, nil
}

// NewPythmcapindexTransactor creates a new write-only instance of Pythmcapindex, bound to a specific deployed contract.
func NewPythmcapindexTransactor(address common.Address, transactor bind.ContractTransactor) (*PythmcapindexTransactor, error) {
	contract, err := bindPythmcapindex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PythmcapindexTransactor{contract: contract}, nil
}

// NewPythmcapindexFilterer creates a new log filterer instance of Pythmcapindex, bound to a specific deployed contract.
func NewPythmcapindexFilterer(address common.Address, filterer bind.ContractFilterer) (*PythmcapindexFilterer, error) {
	contract, err := bindPythmcapindex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PythmcapindexFilterer{contract: contract}, nil
}

// bindPythmcapindex binds a generic wrapper to an already deployed contract.
func bindPythmcapindex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PythmcapindexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pythmcapindex *PythmcapindexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pythmcapindex.Contract.PythmcapindexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pythmcapindex *PythmcapindexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pythmcapindex.Contract.PythmcapindexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pythmcapindex *PythmcapindexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pythmcapindex.Contract.PythmcapindexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pythmcapindex *PythmcapindexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pythmcapindex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pythmcapindex *PythmcapindexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pythmcapindex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pythmcapindex *PythmcapindexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pythmcapindex.Contract.contract.Transact(opts, method, params...)
}

// GetIndexPrice is a free data retrieval call binding the contract method 0x61f608df.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets) view returns(int256)
func (_Pythmcapindex *PythmcapindexCaller) GetIndexPrice(opts *bind.CallOpts, indexAssets []IndexAsset) (*big.Int, error) {
	var out []interface{}
	err := _Pythmcapindex.contract.Call(opts, &out, "getIndexPrice", indexAssets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexPrice is a free data retrieval call binding the contract method 0x61f608df.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets) view returns(int256)
func (_Pythmcapindex *PythmcapindexSession) GetIndexPrice(indexAssets []IndexAsset) (*big.Int, error) {
	return _Pythmcapindex.Contract.GetIndexPrice(&_Pythmcapindex.CallOpts, indexAssets)
}

// GetIndexPrice is a free data retrieval call binding the contract method 0x61f608df.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets) view returns(int256)
func (_Pythmcapindex *PythmcapindexCallerSession) GetIndexPrice(indexAssets []IndexAsset) (*big.Int, error) {
	return _Pythmcapindex.Contract.GetIndexPrice(&_Pythmcapindex.CallOpts, indexAssets)
}

// GetIndexPrice0 is a free data retrieval call binding the contract method 0x6ca50725.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets, uint8 topN, int32 expo) view returns(int256)
func (_Pythmcapindex *PythmcapindexCaller) GetIndexPrice0(opts *bind.CallOpts, indexAssets []IndexAsset, topN uint8, expo int32) (*big.Int, error) {
	var out []interface{}
	err := _Pythmcapindex.contract.Call(opts, &out, "getIndexPrice0", indexAssets, topN, expo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexPrice0 is a free data retrieval call binding the contract method 0x6ca50725.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets, uint8 topN, int32 expo) view returns(int256)
func (_Pythmcapindex *PythmcapindexSession) GetIndexPrice0(indexAssets []IndexAsset, topN uint8, expo int32) (*big.Int, error) {
	return _Pythmcapindex.Contract.GetIndexPrice0(&_Pythmcapindex.CallOpts, indexAssets, topN, expo)
}

// GetIndexPrice0 is a free data retrieval call binding the contract method 0x6ca50725.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets, uint8 topN, int32 expo) view returns(int256)
func (_Pythmcapindex *PythmcapindexCallerSession) GetIndexPrice0(indexAssets []IndexAsset, topN uint8, expo int32) (*big.Int, error) {
	return _Pythmcapindex.Contract.GetIndexPrice0(&_Pythmcapindex.CallOpts, indexAssets, topN, expo)
}

// GetIndexPrice1 is a free data retrieval call binding the contract method 0xa587fa91.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets, bytes32 baseAssetId, uint8 topN, int32 expo) view returns(int256)
func (_Pythmcapindex *PythmcapindexCaller) GetIndexPrice1(opts *bind.CallOpts, indexAssets []IndexAsset, baseAssetId [32]byte, topN uint8, expo int32) (*big.Int, error) {
	var out []interface{}
	err := _Pythmcapindex.contract.Call(opts, &out, "getIndexPrice1", indexAssets, baseAssetId, topN, expo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexPrice1 is a free data retrieval call binding the contract method 0xa587fa91.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets, bytes32 baseAssetId, uint8 topN, int32 expo) view returns(int256)
func (_Pythmcapindex *PythmcapindexSession) GetIndexPrice1(indexAssets []IndexAsset, baseAssetId [32]byte, topN uint8, expo int32) (*big.Int, error) {
	return _Pythmcapindex.Contract.GetIndexPrice1(&_Pythmcapindex.CallOpts, indexAssets, baseAssetId, topN, expo)
}

// GetIndexPrice1 is a free data retrieval call binding the contract method 0xa587fa91.
//
// Solidity: function getIndexPrice((bytes32,uint96)[] indexAssets, bytes32 baseAssetId, uint8 topN, int32 expo) view returns(int256)
func (_Pythmcapindex *PythmcapindexCallerSession) GetIndexPrice1(indexAssets []IndexAsset, baseAssetId [32]byte, topN uint8, expo int32) (*big.Int, error) {
	return _Pythmcapindex.Contract.GetIndexPrice1(&_Pythmcapindex.CallOpts, indexAssets, baseAssetId, topN, expo)
}
