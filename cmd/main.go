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

	signedTxHex := "0xf86b808504a817c80082520894f1d...etc"

	txHash, err := cli.SendRawTransaction(signedTxHex)
	if err != nil {
		log.Fatal("Broadcast failed:", err)
	}

	fmt.Println("Transaction broadcasted! Hash:", txHash)
}
