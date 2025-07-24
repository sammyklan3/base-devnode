package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

// GetTransactionByHash fetches a transaction and its receipt by hash
func (bc *BaseClient) GetTransactionByHash(txHash string) (*types.Transaction, bool, error) {
	hash := common.HexToHash(txHash)
	tx, isPending, err := bc.Client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, false, err
	}
	return tx, isPending, nil
}
