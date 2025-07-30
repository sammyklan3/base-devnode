package main

import (
	"fmt"
	"log"
	"math/big"

	"base-devnode/client"
)

func main() {
	rpc := "https://base.llamarpc.com"

	cli, err := client.NewBaseClient(rpc)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
}
