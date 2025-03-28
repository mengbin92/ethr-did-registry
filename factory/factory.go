package factory

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mengbin92/ethr-did-registry/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(utils.ContextKey("DB"))
	if v == nil {
		panic("db is not exist")
	}
	if db, ok := v.(*gorm.DB); ok {
		return db
	}
	panic("db is not exist")
}

func Redis(ctx context.Context) *redis.Client {
	v := ctx.Value(utils.ContextKey("Redis"))
	if v == nil {
		panic("redis is not exist")
	}
	if redis, ok := v.(*redis.Client); ok {
		return redis
	}
	panic("redis is not exist")
}

func ETH(ctx context.Context) *ethclient.Client {
	v := ctx.Value(utils.ContextKey("ETH"))
	if v == nil {
		panic("ethclient is not exist")
	}
	if eth, ok := v.(*ethclient.Client); ok {
		return eth
	}
	panic("ethclient is not exist")
}

func Logger(ctx context.Context) *zap.Logger {
	v := ctx.Value(utils.ContextKey("LOGGER"))
	if v == nil {
		panic("zap.Logger is not exist")
	}
	if log, ok := v.(*zap.Logger); ok {
		return log
	}
	panic("zap.Logger is not exist")
}

func TransactOpts(ctx context.Context) *bind.TransactOpts {
	v := ctx.Value(utils.ContextKey("AUTH"))
	if v == nil {
		panic("bind.TransactOpts is not exist")
	}
	if opts, ok := v.(*bind.TransactOpts); ok {
		return opts
	}
	panic("bind.TransactOpts is not exist")
}

func Network(ctx context.Context) *big.Int {
	v := ctx.Value(utils.ContextKey("NETWORK"))
	if v == nil {
		panic("network id is not exist")
	}
	if network, ok := v.(*big.Int); ok {
		return network
	}
	panic("network id is not exist")
}
