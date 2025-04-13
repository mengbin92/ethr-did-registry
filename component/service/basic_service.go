package service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mengbin92/did/component/chain/client"
)

type BlockchainService struct{}

func (BlockchainService) GetBlockNumber(ctx context.Context) (uint64, error) {
	bn, err := client.BasicClient{}.GetBlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return bn, nil
}

func (BlockchainService) GetBlockByNumber(ctx context.Context, blockNumber int64) (*types.Block, error) {
	block, err := client.BasicClient{}.GetBlockByNumber(ctx, blockNumber)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (BlockchainService) GetBlockByHash(ctx context.Context, blockHash string) (*types.Block, error) {
	block, err := client.BasicClient{}.GetBlockByHash(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (BlockchainService) GetTransactionByHash(ctx context.Context, txHash string) (*types.Transaction, error) {
	tx, err := client.BasicClient{}.GetTransactionByHash(ctx, txHash)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (BlockchainService) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	receipt, err := client.BasicClient{}.GetTransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (BlockchainService) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	balance, err := client.BasicClient{}.GetBalance(ctx, address)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
