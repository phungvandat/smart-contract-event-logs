package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	PrintJSON(logs)
}

func PrintJSON(val interface{}) {
	b, _ := json.MarshalIndent(val, "", "\t")
	fmt.Println(string(b))
}
