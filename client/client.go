package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type BaseClient struct {
	RPCURL string
	Client *ethclient.Client
}

// NewBaseClient creates a new client to Base
func NewBaseClient(rpcURL string) (*BaseClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &BaseClient{
		RPCURL: rpcURL,
		Client: client,
	}, nil
}

// GetLatestBlockNumber fetches the latest block number
func (bc *BaseClient) GetLatestBlockNumber() (*big.Int, error) {
	header, err := bc.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}
