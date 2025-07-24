package main

import (
	"fmt"
	"log"

	"base-devnode/client"
)

func main() {
	rpc := "https://mainnet.base.org" 

	cli, err := client.NewBaseClient(rpc)
	if err != nil {
		log.Fatalf("Failed to connect to Base: %v", err)
	}

	blockNumber, err := cli.GetLatestBlockNumber()
	if err != nil {
		log.Fatalf("Error fetching block number: %v", err)
	}

	fmt.Printf("🟦 Latest Base Block: %v\n", blockNumber)
}