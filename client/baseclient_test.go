package client

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var rpcURL = "https://mainnet.base.org"

func init() {
	if val := os.Getenv("RPC_URL"); val != "" {
		rpcURL = val
	}
}

func TestNewBaseClient(t *testing.T) {
	cli, err := NewBaseClient(rpcURL)
	assert.NoError(t, err)
	assert.NotNil(t, cli)
}

func TestGetLatestBlockNumber(t *testing.T) {
	cli, err := NewBaseClient(rpcURL)
	assert.NoError(t, err)

	number, err := cli.GetLatestBlockNumber()
	assert.NoError(t, err)
	assert.True(t, number.Cmp(big.NewInt(0)) > 0)

	t.Log("Latest block number:", number.String())
}

func TestGetBalance(t *testing.T) {
	cli, err := NewBaseClient(rpcURL)
	assert.NoError(t, err)

	addr := "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
	balance, err := cli.GetBalance(addr)
	assert.NoError(t, err)
	assert.NotNil(t, balance)

	t.Log("Balance:", balance.String())
}

func TestGetChainID(t *testing.T) {
	cli, err := NewBaseClient(rpcURL)
	assert.NoError(t, err)

	chainID, err := cli.GetChainID()
	assert.NoError(t, err)
	assert.NotNil(t, chainID)

	t.Log("Chain ID:", chainID.String())
}

func TestEstimateGas(t *testing.T) {
	cli, err := NewBaseClient(rpcURL)
	assert.NoError(t, err)

	to := common.HexToAddress("0x0000000000000000000000000000000000000000")
	msg := ethereum.CallMsg{
		To:   &to,
		Gas:  21000,
		Data: nil,
	}

	gas, err := cli.EstimateGas(msg)
	assert.NoError(t, err)
	assert.True(t, gas > 0)

	t.Log("Estimated gas:", gas)
}
