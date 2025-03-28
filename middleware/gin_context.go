package middleware

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/mengbin92/ethr-did-registry/utils"
	"go.uber.org/zap"
)

func SetEthClientMiddleware(client *ethclient.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), utils.ContextKey("ETH"), client))
		ctx.Next()
	}
}

func SetNetworkMiddleware(network *big.Int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), utils.ContextKey("NETWORK"), network))
		ctx.Next()
	}
}

func SetTransactOptsMiddleware(bindOpts *bind.TransactOpts) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), utils.ContextKey("AUTH"), bindOpts))
		ctx.Next()
	}
}

func SetLogMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), utils.ContextKey("LOGGER"), logger))
		ctx.Next()
	}
}
