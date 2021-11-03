// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abimarket2

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
)

// NFTMarket2MarketItem is an auto generated low-level Go binding around an user-defined struct.
type NFTMarket2MarketItem struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
}

// Abimarket2MetaData contains all meta data concerning the Abimarket2 contract.
var Abimarket2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"name\":\"MarketItemCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createMarketItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"createMarketSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMarketItems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket2.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMyNFTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket2.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchItemsCreated\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket2.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"inStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"appendUintToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// Abimarket2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Abimarket2MetaData.ABI instead.
var Abimarket2ABI = Abimarket2MetaData.ABI

// Abimarket2 is an auto generated Go binding around an Ethereum contract.
type Abimarket2 struct {
	Abimarket2Caller     // Read-only binding to the contract
	Abimarket2Transactor // Write-only binding to the contract
	Abimarket2Filterer   // Log filterer for contract events
}

// Abimarket2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Abimarket2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abimarket2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Abimarket2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abimarket2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Abimarket2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abimarket2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Abimarket2Session struct {
	Contract     *Abimarket2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Abimarket2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Abimarket2CallerSession struct {
	Contract *Abimarket2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Abimarket2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Abimarket2TransactorSession struct {
	Contract     *Abimarket2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Abimarket2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Abimarket2Raw struct {
	Contract *Abimarket2 // Generic contract binding to access the raw methods on
}

// Abimarket2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Abimarket2CallerRaw struct {
	Contract *Abimarket2Caller // Generic read-only contract binding to access the raw methods on
}

// Abimarket2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Abimarket2TransactorRaw struct {
	Contract *Abimarket2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAbimarket2 creates a new instance of Abimarket2, bound to a specific deployed contract.
func NewAbimarket2(address common.Address, backend bind.ContractBackend) (*Abimarket2, error) {
	contract, err := bindAbimarket2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abimarket2{Abimarket2Caller: Abimarket2Caller{contract: contract}, Abimarket2Transactor: Abimarket2Transactor{contract: contract}, Abimarket2Filterer: Abimarket2Filterer{contract: contract}}, nil
}

// NewAbimarket2Caller creates a new read-only instance of Abimarket2, bound to a specific deployed contract.
func NewAbimarket2Caller(address common.Address, caller bind.ContractCaller) (*Abimarket2Caller, error) {
	contract, err := bindAbimarket2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Abimarket2Caller{contract: contract}, nil
}

// NewAbimarket2Transactor creates a new write-only instance of Abimarket2, bound to a specific deployed contract.
func NewAbimarket2Transactor(address common.Address, transactor bind.ContractTransactor) (*Abimarket2Transactor, error) {
	contract, err := bindAbimarket2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Abimarket2Transactor{contract: contract}, nil
}

// NewAbimarket2Filterer creates a new log filterer instance of Abimarket2, bound to a specific deployed contract.
func NewAbimarket2Filterer(address common.Address, filterer bind.ContractFilterer) (*Abimarket2Filterer, error) {
	contract, err := bindAbimarket2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Abimarket2Filterer{contract: contract}, nil
}

// bindAbimarket2 binds a generic wrapper to an already deployed contract.
func bindAbimarket2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Abimarket2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abimarket2 *Abimarket2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abimarket2.Contract.Abimarket2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abimarket2 *Abimarket2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket2.Contract.Abimarket2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abimarket2 *Abimarket2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abimarket2.Contract.Abimarket2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abimarket2 *Abimarket2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abimarket2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abimarket2 *Abimarket2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abimarket2 *Abimarket2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abimarket2.Contract.contract.Transact(opts, method, params...)
}

// AppendUintToString is a free data retrieval call binding the contract method 0xc80667e3.
//
// Solidity: function appendUintToString(string inStr, uint256 val) pure returns(string str)
func (_Abimarket2 *Abimarket2Caller) AppendUintToString(opts *bind.CallOpts, inStr string, val *big.Int) (string, error) {
	var out []interface{}
	err := _Abimarket2.contract.Call(opts, &out, "appendUintToString", inStr, val)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AppendUintToString is a free data retrieval call binding the contract method 0xc80667e3.
//
// Solidity: function appendUintToString(string inStr, uint256 val) pure returns(string str)
func (_Abimarket2 *Abimarket2Session) AppendUintToString(inStr string, val *big.Int) (string, error) {
	return _Abimarket2.Contract.AppendUintToString(&_Abimarket2.CallOpts, inStr, val)
}

// AppendUintToString is a free data retrieval call binding the contract method 0xc80667e3.
//
// Solidity: function appendUintToString(string inStr, uint256 val) pure returns(string str)
func (_Abimarket2 *Abimarket2CallerSession) AppendUintToString(inStr string, val *big.Int) (string, error) {
	return _Abimarket2.Contract.AppendUintToString(&_Abimarket2.CallOpts, inStr, val)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2Caller) FetchItemsCreated(opts *bind.CallOpts) ([]NFTMarket2MarketItem, error) {
	var out []interface{}
	err := _Abimarket2.contract.Call(opts, &out, "fetchItemsCreated")

	if err != nil {
		return *new([]NFTMarket2MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarket2MarketItem)).(*[]NFTMarket2MarketItem)

	return out0, err

}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2Session) FetchItemsCreated() ([]NFTMarket2MarketItem, error) {
	return _Abimarket2.Contract.FetchItemsCreated(&_Abimarket2.CallOpts)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2CallerSession) FetchItemsCreated() ([]NFTMarket2MarketItem, error) {
	return _Abimarket2.Contract.FetchItemsCreated(&_Abimarket2.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2Caller) FetchMarketItems(opts *bind.CallOpts) ([]NFTMarket2MarketItem, error) {
	var out []interface{}
	err := _Abimarket2.contract.Call(opts, &out, "fetchMarketItems")

	if err != nil {
		return *new([]NFTMarket2MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarket2MarketItem)).(*[]NFTMarket2MarketItem)

	return out0, err

}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2Session) FetchMarketItems() ([]NFTMarket2MarketItem, error) {
	return _Abimarket2.Contract.FetchMarketItems(&_Abimarket2.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2CallerSession) FetchMarketItems() ([]NFTMarket2MarketItem, error) {
	return _Abimarket2.Contract.FetchMarketItems(&_Abimarket2.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2Caller) FetchMyNFTs(opts *bind.CallOpts) ([]NFTMarket2MarketItem, error) {
	var out []interface{}
	err := _Abimarket2.contract.Call(opts, &out, "fetchMyNFTs")

	if err != nil {
		return *new([]NFTMarket2MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarket2MarketItem)).(*[]NFTMarket2MarketItem)

	return out0, err

}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2Session) FetchMyNFTs() ([]NFTMarket2MarketItem, error) {
	return _Abimarket2.Contract.FetchMyNFTs(&_Abimarket2.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket2 *Abimarket2CallerSession) FetchMyNFTs() ([]NFTMarket2MarketItem, error) {
	return _Abimarket2.Contract.FetchMyNFTs(&_Abimarket2.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket2 *Abimarket2Caller) GetListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abimarket2.contract.Call(opts, &out, "getListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket2 *Abimarket2Session) GetListingPrice() (*big.Int, error) {
	return _Abimarket2.Contract.GetListingPrice(&_Abimarket2.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket2 *Abimarket2CallerSession) GetListingPrice() (*big.Int, error) {
	return _Abimarket2.Contract.GetListingPrice(&_Abimarket2.CallOpts)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket2 *Abimarket2Transactor) CreateMarketItem(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket2.contract.Transact(opts, "createMarketItem", nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket2 *Abimarket2Session) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket2.Contract.CreateMarketItem(&_Abimarket2.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket2 *Abimarket2TransactorSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket2.Contract.CreateMarketItem(&_Abimarket2.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket2 *Abimarket2Transactor) CreateMarketSale(opts *bind.TransactOpts, nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket2.contract.Transact(opts, "createMarketSale", nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket2 *Abimarket2Session) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket2.Contract.CreateMarketSale(&_Abimarket2.TransactOpts, nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket2 *Abimarket2TransactorSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket2.Contract.CreateMarketSale(&_Abimarket2.TransactOpts, nftContract, itemId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Abimarket2 *Abimarket2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Abimarket2 *Abimarket2Session) Initialize() (*types.Transaction, error) {
	return _Abimarket2.Contract.Initialize(&_Abimarket2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Abimarket2 *Abimarket2TransactorSession) Initialize() (*types.Transaction, error) {
	return _Abimarket2.Contract.Initialize(&_Abimarket2.TransactOpts)
}

// Abimarket2MarketItemCreatedIterator is returned from FilterMarketItemCreated and is used to iterate over the raw logs and unpacked data for MarketItemCreated events raised by the Abimarket2 contract.
type Abimarket2MarketItemCreatedIterator struct {
	Event *Abimarket2MarketItemCreated // Event containing the contract specifics and raw log

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
func (it *Abimarket2MarketItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abimarket2MarketItemCreated)
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
		it.Event = new(Abimarket2MarketItemCreated)
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
func (it *Abimarket2MarketItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abimarket2MarketItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abimarket2MarketItemCreated represents a MarketItemCreated event raised by the Abimarket2 contract.
type Abimarket2MarketItemCreated struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketItemCreated is a free log retrieval operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_Abimarket2 *Abimarket2Filterer) FilterMarketItemCreated(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*Abimarket2MarketItemCreatedIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Abimarket2.contract.FilterLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Abimarket2MarketItemCreatedIterator{contract: _Abimarket2.contract, event: "MarketItemCreated", logs: logs, sub: sub}, nil
}

// WatchMarketItemCreated is a free log subscription operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_Abimarket2 *Abimarket2Filterer) WatchMarketItemCreated(opts *bind.WatchOpts, sink chan<- *Abimarket2MarketItemCreated, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Abimarket2.contract.WatchLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abimarket2MarketItemCreated)
				if err := _Abimarket2.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
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

// ParseMarketItemCreated is a log parse operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_Abimarket2 *Abimarket2Filterer) ParseMarketItemCreated(log types.Log) (*Abimarket2MarketItemCreated, error) {
	event := new(Abimarket2MarketItemCreated)
	if err := _Abimarket2.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
