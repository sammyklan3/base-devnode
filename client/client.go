package client

import (
	"context"
	"fmt"
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
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}
	return &BaseClient{
		RPCURL: rpcURL,
		Client: client,
	}, nil
}

func (bc *BaseClient) Close() {
	bc.Client.Close()
}

// GetLatestBlockNumber fetches the latest block number
func (bc *BaseClient) GetLatestBlockNumber(ctx context.Context) (*big.Int, error) {
	header, err := bc.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block number: %w", err)
	}
	return header.Number, nil
}

// GetTransactionByHash fetches a transaction and its receipt by hash
func (bc *BaseClient) GetTransactionByHash(ctx context.Context, txHash string) (*types.Transaction, bool, error) {
	hash := common.HexToHash(txHash)
	tx, isPending, err := bc.Client.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get transaction: %w", err)
	}
	return tx, isPending, nil
}

// GetBalance returns the ETH balance of the given address
func (bc *BaseClient) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid address: %s", address)
	}
	addr := common.HexToAddress(address)
	balance, err := bc.Client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	return balance, nil
}

// GetChainID returns the chain ID of the connected Ethereum network
func (bc *BaseClient) GetChainID(ctx context.Context) (*big.Int, error) {
	chainID, err := bc.Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	return chainID, nil
}

// GetBlockByNumber fetches the block for the given number
func (bc *BaseClient) GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*types.Block, error) {
	block, err := bc.Client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get block by number: %w", err)
	}
	return block, nil
}

// GetBlockByHash fetches the block for the given hash
func (bc *BaseClient) GetBlockByHash(ctx context.Context, blockHash string) (*types.Block, error) {
	hash := common.HexToHash(blockHash)
	block, err := bc.Client.BlockByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get block by hash: %w", err)
	}
	return block, nil
}

// SendRawTransaction broadcasts a pre-signed raw transaction to the network
func (bc *BaseClient) SendRawTransaction(ctx context.Context, rawTxHex string) (string, error) {
	var txHash common.Hash
	err := bc.Client.Client().CallContext(
		ctx,
		&txHash,
		"eth_sendRawTransaction",
		rawTxHex,
	)
	if err != nil {
		return "", fmt.Errorf("failed to send raw transaction: %w", err)
	}
	return txHash.Hex(), nil
}

// GetTransactionReceipt fetches the receipt of a transaction by its hash
func (bc *BaseClient) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	hash := common.HexToHash(txHash)
	receipt, err := bc.Client.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction receipt: %w", err)
	}
	return receipt, nil
}

// GetNonce returns the nonce for the given address
func (bc *BaseClient) GetNonce(ctx context.Context, address string) (uint64, error) {
	addr := common.HexToAddress(address)
	nonce, err := bc.Client.NonceAt(ctx, addr, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get nonce: %w", err)
	}
	return nonce, nil
}

// EstimateGas estimates the gas required for a transaction. Call this before sending txs to get a reliable gas estimate.
func (bc *BaseClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	gas, err := bc.Client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, fmt.Errorf("failed to estimate gas: %w", err)
	}
	return gas, nil
}
