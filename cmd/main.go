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

	balance, err := cli.GetBalance("0x89aeE487218f9e0f986c30445A9B7ebE135c1029")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("💰 Balance in wei: %s\n", balance.String())
}
