package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mengbin92/did/lib/factory"
	"github.com/pkg/errors"
)

type BasicClient struct{}

// GetBlockNumber returns the current block number of the chain.
func (BasicClient) GetBlockNumber(ctx context.Context) (uint64, error) {
	bn, err := factory.ETH(ctx).BlockNumber(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get block number failed")
	}
	return bn, nil
}

// GetBlockByNumber returns the block data of the given block number.
func (BasicClient) GetBlockByNumber(ctx context.Context, blockNumber int64) (*types.Block, error) {
	block, err := factory.ETH(ctx).BlockByNumber(ctx, big.NewInt(blockNumber))
	if err != nil {
		return nil, errors.Wrap(err, "get block by number failed")
	}
	return block, nil
}

// GetBlockByHash returns the block data of the given block hash.
func (BasicClient) GetBlockByHash(ctx context.Context, blockHash string) (*types.Block, error) {
	block, err := factory.ETH(ctx).BlockByHash(ctx, common.HexToHash(blockHash))
	if err != nil {
		return nil, errors.Wrap(err, "get block hash failed")
	}
	return block, nil
}

// GetTransactionByHash returns the transaction data of the given transaction hash.
func (BasicClient) GetTransactionByHash(ctx context.Context, txHash string) (*types.Transaction, error) {
	tx, _, err := factory.ETH(ctx).TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, errors.Wrap(err, "get transaction by hash failed")
	}
	return tx, nil
}

// GetTransactionReceipt returns the receipt data of the given transaction hash.
func (BasicClient) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	receipt, err := factory.ETH(ctx).TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, errors.Wrap(err, "get transaction receipt failed")
	}
	return receipt, nil
}

// GetBalance returns the balance of the given address.
func (BasicClient) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	balance, err := factory.ETH(ctx).BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return nil, errors.Wrap(err, "get balance failed")
	}
	return balance, nil
}