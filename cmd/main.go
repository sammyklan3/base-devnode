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
		log.Fatalf("Failed to connect: %v", err)
	}

	blockHash := "0x08322071401b3a69b00a956560393c57beaac47c7dad5adc2feae8b8f13cc253"

	block, err := cli.GetBlockByHash(blockHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Block Number: %s\n", block.Number().String())
	fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
	fmt.Printf("Block Transactions: %d\n", len(block.Transactions()))
}
