package main

import (
	"fmt"
	"log"

	"base-devnode/client"
)

func main() {
	rpc := "https://mainnet.base.org"
	txHash := "0x8164abb7e4207cf26038299ecdb74c4a485615907368040aaf866a15c0c4c8d0"

	cli, err := client.NewBaseClient(rpc)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	tx, isPending, err := cli.GetTransactionByHash(txHash)
	if err != nil {
		log.Fatalf("Failed to fetch transaction: %v", err)
	}

	fmt.Printf("🔎 Transaction Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("📤 From: (unavailable directly, needs tx sender recovery)\n")
	fmt.Printf("📥 To: %s\n", tx.To().Hex())
	fmt.Printf("⛽️ Gas: %d\n", tx.Gas())
	fmt.Printf("💰 Value: %s\n", tx.Value().String())
	fmt.Printf("⌛ Pending: %v\n", isPending)
}
