package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mengbin92/did/component/service"
)

type BasicAPI struct {
}

func (BasicAPI) GetBlockNumber(ctx *gin.Context) {
	bn, err := service.BlockchainService{}.GetBlockNumber(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"block_number": bn,
	})
}

func (BasicAPI) GetBlockByNumber(ctx *gin.Context) {
	blockNumber := ctx.Query("block_number")
	bn, err := strconv.ParseInt(blockNumber, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid block number",
		})
		return
	}
	block, err := service.BlockchainService{}.GetBlockByNumber(ctx.Request.Context(), bn)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"block": block,
	})
}

func (BasicAPI) GetBlockByHash(ctx *gin.Context) {
	blockHash := ctx.Query("block_hash")
	block, err := service.BlockchainService{}.GetBlockByHash(ctx.Request.Context(), blockHash)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"block": block,
	})
}

func (BasicAPI) GetTransactionByHash(ctx *gin.Context) {
	txHash := ctx.Query("tx_hash")
	tx, err := service.BlockchainService{}.GetTransactionByHash(ctx.Request.Context(), txHash)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"transaction": tx,
	})
}

func (BasicAPI) GetTransactionReceipt(ctx *gin.Context) {
	txHash := ctx.Query("tx_hash")
	receipt, err := service.BlockchainService{}.GetTransactionReceipt(ctx.Request.Context(), txHash)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"receipt": receipt,
	})
}

func (BasicAPI) GetBalance(ctx *gin.Context) {
	address := ctx.Query("address")
	balance, err := service.BlockchainService{}.GetBalance(ctx.Request.Context(), address)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}
