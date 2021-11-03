// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abimarket3

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

// NFTMarket3MarketItem is an auto generated low-level Go binding around an user-defined struct.
type NFTMarket3MarketItem struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
}

// Abimarket3MetaData contains all meta data concerning the Abimarket3 contract.
var Abimarket3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"name\":\"MarketItemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MarketItemSold\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createMarketItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"createMarketSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMarketItems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket3.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMyNFTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket3.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchItemsCreated\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket3.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"inStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"appendUintToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// Abimarket3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Abimarket3MetaData.ABI instead.
var Abimarket3ABI = Abimarket3MetaData.ABI

// Abimarket3 is an auto generated Go binding around an Ethereum contract.
type Abimarket3 struct {
	Abimarket3Caller     // Read-only binding to the contract
	Abimarket3Transactor // Write-only binding to the contract
	Abimarket3Filterer   // Log filterer for contract events
}

// Abimarket3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Abimarket3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abimarket3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Abimarket3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abimarket3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Abimarket3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abimarket3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Abimarket3Session struct {
	Contract     *Abimarket3       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Abimarket3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Abimarket3CallerSession struct {
	Contract *Abimarket3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Abimarket3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Abimarket3TransactorSession struct {
	Contract     *Abimarket3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Abimarket3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Abimarket3Raw struct {
	Contract *Abimarket3 // Generic contract binding to access the raw methods on
}

// Abimarket3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Abimarket3CallerRaw struct {
	Contract *Abimarket3Caller // Generic read-only contract binding to access the raw methods on
}

// Abimarket3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Abimarket3TransactorRaw struct {
	Contract *Abimarket3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAbimarket3 creates a new instance of Abimarket3, bound to a specific deployed contract.
func NewAbimarket3(address common.Address, backend bind.ContractBackend) (*Abimarket3, error) {
	contract, err := bindAbimarket3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abimarket3{Abimarket3Caller: Abimarket3Caller{contract: contract}, Abimarket3Transactor: Abimarket3Transactor{contract: contract}, Abimarket3Filterer: Abimarket3Filterer{contract: contract}}, nil
}

// NewAbimarket3Caller creates a new read-only instance of Abimarket3, bound to a specific deployed contract.
func NewAbimarket3Caller(address common.Address, caller bind.ContractCaller) (*Abimarket3Caller, error) {
	contract, err := bindAbimarket3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Abimarket3Caller{contract: contract}, nil
}

// NewAbimarket3Transactor creates a new write-only instance of Abimarket3, bound to a specific deployed contract.
func NewAbimarket3Transactor(address common.Address, transactor bind.ContractTransactor) (*Abimarket3Transactor, error) {
	contract, err := bindAbimarket3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Abimarket3Transactor{contract: contract}, nil
}

// NewAbimarket3Filterer creates a new log filterer instance of Abimarket3, bound to a specific deployed contract.
func NewAbimarket3Filterer(address common.Address, filterer bind.ContractFilterer) (*Abimarket3Filterer, error) {
	contract, err := bindAbimarket3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Abimarket3Filterer{contract: contract}, nil
}

// bindAbimarket3 binds a generic wrapper to an already deployed contract.
func bindAbimarket3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Abimarket3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abimarket3 *Abimarket3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abimarket3.Contract.Abimarket3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abimarket3 *Abimarket3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket3.Contract.Abimarket3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abimarket3 *Abimarket3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abimarket3.Contract.Abimarket3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abimarket3 *Abimarket3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abimarket3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abimarket3 *Abimarket3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abimarket3 *Abimarket3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abimarket3.Contract.contract.Transact(opts, method, params...)
}

// AppendUintToString is a free data retrieval call binding the contract method 0xc80667e3.
//
// Solidity: function appendUintToString(string inStr, uint256 val) pure returns(string str)
func (_Abimarket3 *Abimarket3Caller) AppendUintToString(opts *bind.CallOpts, inStr string, val *big.Int) (string, error) {
	var out []interface{}
	err := _Abimarket3.contract.Call(opts, &out, "appendUintToString", inStr, val)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AppendUintToString is a free data retrieval call binding the contract method 0xc80667e3.
//
// Solidity: function appendUintToString(string inStr, uint256 val) pure returns(string str)
func (_Abimarket3 *Abimarket3Session) AppendUintToString(inStr string, val *big.Int) (string, error) {
	return _Abimarket3.Contract.AppendUintToString(&_Abimarket3.CallOpts, inStr, val)
}

// AppendUintToString is a free data retrieval call binding the contract method 0xc80667e3.
//
// Solidity: function appendUintToString(string inStr, uint256 val) pure returns(string str)
func (_Abimarket3 *Abimarket3CallerSession) AppendUintToString(inStr string, val *big.Int) (string, error) {
	return _Abimarket3.Contract.AppendUintToString(&_Abimarket3.CallOpts, inStr, val)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3Caller) FetchItemsCreated(opts *bind.CallOpts) ([]NFTMarket3MarketItem, error) {
	var out []interface{}
	err := _Abimarket3.contract.Call(opts, &out, "fetchItemsCreated")

	if err != nil {
		return *new([]NFTMarket3MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarket3MarketItem)).(*[]NFTMarket3MarketItem)

	return out0, err

}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3Session) FetchItemsCreated() ([]NFTMarket3MarketItem, error) {
	return _Abimarket3.Contract.FetchItemsCreated(&_Abimarket3.CallOpts)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3CallerSession) FetchItemsCreated() ([]NFTMarket3MarketItem, error) {
	return _Abimarket3.Contract.FetchItemsCreated(&_Abimarket3.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3Caller) FetchMarketItems(opts *bind.CallOpts) ([]NFTMarket3MarketItem, error) {
	var out []interface{}
	err := _Abimarket3.contract.Call(opts, &out, "fetchMarketItems")

	if err != nil {
		return *new([]NFTMarket3MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarket3MarketItem)).(*[]NFTMarket3MarketItem)

	return out0, err

}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3Session) FetchMarketItems() ([]NFTMarket3MarketItem, error) {
	return _Abimarket3.Contract.FetchMarketItems(&_Abimarket3.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3CallerSession) FetchMarketItems() ([]NFTMarket3MarketItem, error) {
	return _Abimarket3.Contract.FetchMarketItems(&_Abimarket3.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3Caller) FetchMyNFTs(opts *bind.CallOpts) ([]NFTMarket3MarketItem, error) {
	var out []interface{}
	err := _Abimarket3.contract.Call(opts, &out, "fetchMyNFTs")

	if err != nil {
		return *new([]NFTMarket3MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarket3MarketItem)).(*[]NFTMarket3MarketItem)

	return out0, err

}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3Session) FetchMyNFTs() ([]NFTMarket3MarketItem, error) {
	return _Abimarket3.Contract.FetchMyNFTs(&_Abimarket3.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket3 *Abimarket3CallerSession) FetchMyNFTs() ([]NFTMarket3MarketItem, error) {
	return _Abimarket3.Contract.FetchMyNFTs(&_Abimarket3.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket3 *Abimarket3Caller) GetListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abimarket3.contract.Call(opts, &out, "getListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket3 *Abimarket3Session) GetListingPrice() (*big.Int, error) {
	return _Abimarket3.Contract.GetListingPrice(&_Abimarket3.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket3 *Abimarket3CallerSession) GetListingPrice() (*big.Int, error) {
	return _Abimarket3.Contract.GetListingPrice(&_Abimarket3.CallOpts)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket3 *Abimarket3Transactor) CreateMarketItem(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket3.contract.Transact(opts, "createMarketItem", nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket3 *Abimarket3Session) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket3.Contract.CreateMarketItem(&_Abimarket3.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket3 *Abimarket3TransactorSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket3.Contract.CreateMarketItem(&_Abimarket3.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket3 *Abimarket3Transactor) CreateMarketSale(opts *bind.TransactOpts, nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket3.contract.Transact(opts, "createMarketSale", nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket3 *Abimarket3Session) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket3.Contract.CreateMarketSale(&_Abimarket3.TransactOpts, nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket3 *Abimarket3TransactorSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket3.Contract.CreateMarketSale(&_Abimarket3.TransactOpts, nftContract, itemId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Abimarket3 *Abimarket3Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket3.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Abimarket3 *Abimarket3Session) Initialize() (*types.Transaction, error) {
	return _Abimarket3.Contract.Initialize(&_Abimarket3.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Abimarket3 *Abimarket3TransactorSession) Initialize() (*types.Transaction, error) {
	return _Abimarket3.Contract.Initialize(&_Abimarket3.TransactOpts)
}

// Abimarket3MarketItemCreatedIterator is returned from FilterMarketItemCreated and is used to iterate over the raw logs and unpacked data for MarketItemCreated events raised by the Abimarket3 contract.
type Abimarket3MarketItemCreatedIterator struct {
	Event *Abimarket3MarketItemCreated // Event containing the contract specifics and raw log

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
func (it *Abimarket3MarketItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abimarket3MarketItemCreated)
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
		it.Event = new(Abimarket3MarketItemCreated)
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
func (it *Abimarket3MarketItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abimarket3MarketItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abimarket3MarketItemCreated represents a MarketItemCreated event raised by the Abimarket3 contract.
type Abimarket3MarketItemCreated struct {
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
func (_Abimarket3 *Abimarket3Filterer) FilterMarketItemCreated(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*Abimarket3MarketItemCreatedIterator, error) {

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

	logs, sub, err := _Abimarket3.contract.FilterLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Abimarket3MarketItemCreatedIterator{contract: _Abimarket3.contract, event: "MarketItemCreated", logs: logs, sub: sub}, nil
}

// WatchMarketItemCreated is a free log subscription operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_Abimarket3 *Abimarket3Filterer) WatchMarketItemCreated(opts *bind.WatchOpts, sink chan<- *Abimarket3MarketItemCreated, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Abimarket3.contract.WatchLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abimarket3MarketItemCreated)
				if err := _Abimarket3.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
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
func (_Abimarket3 *Abimarket3Filterer) ParseMarketItemCreated(log types.Log) (*Abimarket3MarketItemCreated, error) {
	event := new(Abimarket3MarketItemCreated)
	if err := _Abimarket3.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abimarket3MarketItemSoldIterator is returned from FilterMarketItemSold and is used to iterate over the raw logs and unpacked data for MarketItemSold events raised by the Abimarket3 contract.
type Abimarket3MarketItemSoldIterator struct {
	Event *Abimarket3MarketItemSold // Event containing the contract specifics and raw log

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
func (it *Abimarket3MarketItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abimarket3MarketItemSold)
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
		it.Event = new(Abimarket3MarketItemSold)
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
func (it *Abimarket3MarketItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abimarket3MarketItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abimarket3MarketItemSold represents a MarketItemSold event raised by the Abimarket3 contract.
type Abimarket3MarketItemSold struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketItemSold is a free log retrieval operation binding the contract event 0x6a5f9d2b6ff1f375ebe095556dc91e68fce80a6187d7527cce81c3ab02726c9c.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price)
func (_Abimarket3 *Abimarket3Filterer) FilterMarketItemSold(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*Abimarket3MarketItemSoldIterator, error) {

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

	logs, sub, err := _Abimarket3.contract.FilterLogs(opts, "MarketItemSold", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Abimarket3MarketItemSoldIterator{contract: _Abimarket3.contract, event: "MarketItemSold", logs: logs, sub: sub}, nil
}

// WatchMarketItemSold is a free log subscription operation binding the contract event 0x6a5f9d2b6ff1f375ebe095556dc91e68fce80a6187d7527cce81c3ab02726c9c.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price)
func (_Abimarket3 *Abimarket3Filterer) WatchMarketItemSold(opts *bind.WatchOpts, sink chan<- *Abimarket3MarketItemSold, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Abimarket3.contract.WatchLogs(opts, "MarketItemSold", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abimarket3MarketItemSold)
				if err := _Abimarket3.contract.UnpackLog(event, "MarketItemSold", log); err != nil {
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

// ParseMarketItemSold is a log parse operation binding the contract event 0x6a5f9d2b6ff1f375ebe095556dc91e68fce80a6187d7527cce81c3ab02726c9c.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price)
func (_Abimarket3 *Abimarket3Filterer) ParseMarketItemSold(log types.Log) (*Abimarket3MarketItemSold, error) {
	event := new(Abimarket3MarketItemSold)
	if err := _Abimarket3.contract.UnpackLog(event, "MarketItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
