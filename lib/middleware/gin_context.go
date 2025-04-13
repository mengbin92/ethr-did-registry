package middleware

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/mengbin92/did/lib/utils"
	"go.uber.org/zap"
)

func SetLogMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), utils.ContextKey("LOGGER"), logger))
		ctx.Next()
	}
}

func SetETHMiddleware(client *ethclient.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), utils.ContextKey("ETH_CLIENT"), client))
		ctx.Next()
	}
}
