// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abimarket

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

// NFTMarketMarketItem is an auto generated low-level Go binding around an user-defined struct.
type NFTMarketMarketItem struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
}

// AbimarketMetaData contains all meta data concerning the Abimarket contract.
var AbimarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"name\":\"MarketItemCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getListingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createMarketItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"createMarketSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMarketItems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMyNFTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchItemsCreated\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarket.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbimarketABI is the input ABI used to generate the binding from.
// Deprecated: Use AbimarketMetaData.ABI instead.
var AbimarketABI = AbimarketMetaData.ABI

// Abimarket is an auto generated Go binding around an Ethereum contract.
type Abimarket struct {
	AbimarketCaller     // Read-only binding to the contract
	AbimarketTransactor // Write-only binding to the contract
	AbimarketFilterer   // Log filterer for contract events
}

// AbimarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbimarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbimarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbimarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbimarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbimarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbimarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbimarketSession struct {
	Contract     *Abimarket        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbimarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbimarketCallerSession struct {
	Contract *AbimarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AbimarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbimarketTransactorSession struct {
	Contract     *AbimarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbimarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbimarketRaw struct {
	Contract *Abimarket // Generic contract binding to access the raw methods on
}

// AbimarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbimarketCallerRaw struct {
	Contract *AbimarketCaller // Generic read-only contract binding to access the raw methods on
}

// AbimarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbimarketTransactorRaw struct {
	Contract *AbimarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbimarket creates a new instance of Abimarket, bound to a specific deployed contract.
func NewAbimarket(address common.Address, backend bind.ContractBackend) (*Abimarket, error) {
	contract, err := bindAbimarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abimarket{AbimarketCaller: AbimarketCaller{contract: contract}, AbimarketTransactor: AbimarketTransactor{contract: contract}, AbimarketFilterer: AbimarketFilterer{contract: contract}}, nil
}

// NewAbimarketCaller creates a new read-only instance of Abimarket, bound to a specific deployed contract.
func NewAbimarketCaller(address common.Address, caller bind.ContractCaller) (*AbimarketCaller, error) {
	contract, err := bindAbimarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbimarketCaller{contract: contract}, nil
}

// NewAbimarketTransactor creates a new write-only instance of Abimarket, bound to a specific deployed contract.
func NewAbimarketTransactor(address common.Address, transactor bind.ContractTransactor) (*AbimarketTransactor, error) {
	contract, err := bindAbimarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbimarketTransactor{contract: contract}, nil
}

// NewAbimarketFilterer creates a new log filterer instance of Abimarket, bound to a specific deployed contract.
func NewAbimarketFilterer(address common.Address, filterer bind.ContractFilterer) (*AbimarketFilterer, error) {
	contract, err := bindAbimarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbimarketFilterer{contract: contract}, nil
}

// bindAbimarket binds a generic wrapper to an already deployed contract.
func bindAbimarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbimarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abimarket *AbimarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abimarket.Contract.AbimarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abimarket *AbimarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket.Contract.AbimarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abimarket *AbimarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abimarket.Contract.AbimarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abimarket *AbimarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abimarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abimarket *AbimarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abimarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abimarket *AbimarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abimarket.Contract.contract.Transact(opts, method, params...)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketCaller) FetchItemsCreated(opts *bind.CallOpts) ([]NFTMarketMarketItem, error) {
	var out []interface{}
	err := _Abimarket.contract.Call(opts, &out, "fetchItemsCreated")

	if err != nil {
		return *new([]NFTMarketMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarketMarketItem)).(*[]NFTMarketMarketItem)

	return out0, err

}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketSession) FetchItemsCreated() ([]NFTMarketMarketItem, error) {
	return _Abimarket.Contract.FetchItemsCreated(&_Abimarket.CallOpts)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketCallerSession) FetchItemsCreated() ([]NFTMarketMarketItem, error) {
	return _Abimarket.Contract.FetchItemsCreated(&_Abimarket.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketCaller) FetchMarketItems(opts *bind.CallOpts) ([]NFTMarketMarketItem, error) {
	var out []interface{}
	err := _Abimarket.contract.Call(opts, &out, "fetchMarketItems")

	if err != nil {
		return *new([]NFTMarketMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarketMarketItem)).(*[]NFTMarketMarketItem)

	return out0, err

}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketSession) FetchMarketItems() ([]NFTMarketMarketItem, error) {
	return _Abimarket.Contract.FetchMarketItems(&_Abimarket.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketCallerSession) FetchMarketItems() ([]NFTMarketMarketItem, error) {
	return _Abimarket.Contract.FetchMarketItems(&_Abimarket.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketCaller) FetchMyNFTs(opts *bind.CallOpts) ([]NFTMarketMarketItem, error) {
	var out []interface{}
	err := _Abimarket.contract.Call(opts, &out, "fetchMyNFTs")

	if err != nil {
		return *new([]NFTMarketMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarketMarketItem)).(*[]NFTMarketMarketItem)

	return out0, err

}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketSession) FetchMyNFTs() ([]NFTMarketMarketItem, error) {
	return _Abimarket.Contract.FetchMyNFTs(&_Abimarket.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_Abimarket *AbimarketCallerSession) FetchMyNFTs() ([]NFTMarketMarketItem, error) {
	return _Abimarket.Contract.FetchMyNFTs(&_Abimarket.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket *AbimarketCaller) GetListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abimarket.contract.Call(opts, &out, "getListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket *AbimarketSession) GetListingPrice() (*big.Int, error) {
	return _Abimarket.Contract.GetListingPrice(&_Abimarket.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_Abimarket *AbimarketCallerSession) GetListingPrice() (*big.Int, error) {
	return _Abimarket.Contract.GetListingPrice(&_Abimarket.CallOpts)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket *AbimarketTransactor) CreateMarketItem(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket.contract.Transact(opts, "createMarketItem", nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket *AbimarketSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket.Contract.CreateMarketItem(&_Abimarket.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_Abimarket *AbimarketTransactorSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Abimarket.Contract.CreateMarketItem(&_Abimarket.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket *AbimarketTransactor) CreateMarketSale(opts *bind.TransactOpts, nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket.contract.Transact(opts, "createMarketSale", nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket *AbimarketSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket.Contract.CreateMarketSale(&_Abimarket.TransactOpts, nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Abimarket *AbimarketTransactorSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Abimarket.Contract.CreateMarketSale(&_Abimarket.TransactOpts, nftContract, itemId)
}

// AbimarketMarketItemCreatedIterator is returned from FilterMarketItemCreated and is used to iterate over the raw logs and unpacked data for MarketItemCreated events raised by the Abimarket contract.
type AbimarketMarketItemCreatedIterator struct {
	Event *AbimarketMarketItemCreated // Event containing the contract specifics and raw log

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
func (it *AbimarketMarketItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbimarketMarketItemCreated)
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
		it.Event = new(AbimarketMarketItemCreated)
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
func (it *AbimarketMarketItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbimarketMarketItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbimarketMarketItemCreated represents a MarketItemCreated event raised by the Abimarket contract.
type AbimarketMarketItemCreated struct {
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
func (_Abimarket *AbimarketFilterer) FilterMarketItemCreated(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*AbimarketMarketItemCreatedIterator, error) {

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

	logs, sub, err := _Abimarket.contract.FilterLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AbimarketMarketItemCreatedIterator{contract: _Abimarket.contract, event: "MarketItemCreated", logs: logs, sub: sub}, nil
}

// WatchMarketItemCreated is a free log subscription operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_Abimarket *AbimarketFilterer) WatchMarketItemCreated(opts *bind.WatchOpts, sink chan<- *AbimarketMarketItemCreated, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Abimarket.contract.WatchLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbimarketMarketItemCreated)
				if err := _Abimarket.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
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
func (_Abimarket *AbimarketFilterer) ParseMarketItemCreated(log types.Log) (*AbimarketMarketItemCreated, error) {
	event := new(AbimarketMarketItemCreated)
	if err := _Abimarket.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
