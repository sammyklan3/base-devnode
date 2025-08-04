package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

// GetTransactionByHash fetches a transaction and its receipt by hash
func (bc *BaseClient) GetTransactionByHash(txHash string) (*types.Transaction, bool, error) {
	hash := common.HexToHash(txHash)
	tx, isPending, err := bc.Client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, false, err
	}
	return tx, isPending, nil
}

// GetBalance returns the ETH balance of the given address
func (bc *BaseClient) GetBalance(address string) (*big.Int, error) {
	addr := common.HexToAddress(address)
	balance, err := bc.Client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetChainID returns the chain ID of the connected Ethereum network
func (bc *BaseClient) GetChainID() (*big.Int, error) {
	chainID, err := bc.Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	return chainID, nil
}

// GetBlockByNumber fetches the block for the given number
func (bc *BaseClient) GetBlockByNumber(blockNumber *big.Int) (*types.Block, error) {
	block, err := bc.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetBlockByHash fetches the block for the given hash
func (bc *BaseClient) GetBlockByHash(blockHash string) (*types.Block, error) {
	hash := common.HexToHash(blockHash)
	block, err := bc.Client.BlockByHash(context.Background(), hash)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// SendRawTransaction broadcasts a pre-signed raw transaction to the network
func (bc *BaseClient) SendRawTransaction(rawTxHex string) (string, error) {
	var txHash common.Hash
	err := bc.Client.Client().CallContext(
		context.Background(),
		&txHash,
		"eth_sendRawTransaction",
		rawTxHex,
	)
	if err != nil {
		return "", err
	}
	return txHash.Hex(), nil
}

// GetTransactionReceipt fetches the receipt of a transaction by its hash
func (bc *BaseClient) GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	hash := common.HexToHash(txHash)
	receipt, err := bc.Client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

// GetNonce returns the nonce for the given address
func (bc *BaseClient) GetNonce(address string) (uint64, error) {
	addr := common.HexToAddress(address)
	nonce, err := bc.Client.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// EstimateGas estimates the gas required for a transaction. Call this before sending txs to get a reliable gas estimate.
func (bc *BaseClient) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	gas, err := bc.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, err
	}
	return gas, nil
}