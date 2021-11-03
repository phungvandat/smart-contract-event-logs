package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/phungvandat/abigen/contracts/abimarket3"
)

func main() {
	godotenv.Load()

	client, err := ethclient.Dial(os.Getenv("RAW_URL"))
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("SMART_CONTRACT_ADDRESS"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
		Topics: [][]common.Hash{
			{crypto.Keccak256Hash([]byte("MarketItemSold(uint256,address,uint256,address,address,uint256)"))},
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	nftmarketAbi, err := abi.JSON(strings.NewReader(abimarket3.Abimarket3ABI))
	if err != nil {
		panic(err)
	}

	type marketItemSold struct {
		Seller common.Address
		Owner  common.Address
		Price  *big.Int
	}

	var sumAmount = new(big.Int)
	for idx := range logs {
		var item marketItemSold
		err := nftmarketAbi.UnpackIntoInterface(&item, "MarketItemSold", logs[idx].Data)
		if err != nil {
			panic(err)
		}
		sumAmount = sumAmount.Add(sumAmount, item.Price)
		PrintJSON(item)
	}
	fmt.Println("Sum Amount: ", sumAmount.String())
}

func PrintJSON(val interface{}) {
	b, _ := json.MarshalIndent(val, "", "\t")
	fmt.Println(string(b))
}
