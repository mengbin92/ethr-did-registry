package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mengbin92/did/component/apis"
)

func SetBasicRouters(r *gin.Engine) {
	v1 := r.Group("/api/blockchain/v1/basic")
	v1.GET("/block_number", apis.BasicAPI{}.GetBlockNumber)
	v1.GET("/block_by_hash", apis.BasicAPI{}.GetBlockByHash)
	v1.GET("/transaction_by_hash", apis.BasicAPI{}.GetTransactionByHash)
	v1.GET("/transaction_receipt_by_hash", apis.BasicAPI{}.GetTransactionReceipt)
	v1.GET("/get_balance", apis.BasicAPI{}.GetBalance)
}
