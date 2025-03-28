package cmd

import (
	"context"
	"fmt"
	"html/template"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/mengbin92/ethr-did-registry/conf"
	"github.com/mengbin92/ethr-did-registry/db"
	"github.com/mengbin92/ethr-did-registry/logger"
	"github.com/mengbin92/ethr-did-registry/middleware"
	"github.com/mengbin92/ethr-did-registry/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	endpoint, keyPath, password string
	network                     *big.Int
)

func Execute() {
	// TODO: implement

	// 初始化配置
	conf.LoadConfig()

	endpoint = viper.GetString("eth.endpoint")
	keyPath = viper.GetString("keystore.path")
	password = viper.GetString("keystore.password")

	run()
}

func run() {
	if err := setEngine(loadETHClient(), loadAuth(), loadDB()).Run(fmt.Sprintf(":%d", viper.GetInt("server.port"))); err != nil {
		log.Error("Failed to run server: ", err)
		panic(err)
	}

}

func loadETHClient() *ethclient.Client {
	// 连接到私链
	client, err := ethclient.Dial(viper.GetString("eth.endpoint"))
	if err != nil {
		log.Error("Failed to connect to the Ethereum client: ", err)
		panic(err)
	}

	network, err = client.NetworkID(context.Background())
	if err != nil {
		log.Error("Failed to get network ID: ", err)
		panic(err)
	}

	return client
}

func loadAuth() *bind.TransactOpts {
	key, err := utils.LoadKey(viper.GetString("keystore.path"), viper.GetString("keystore.password"))
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, network) // 1 是主网链 ID
	if err != nil {
		log.Error("Failed to create authorized transactor: ", err)
		panic(err)
	}
	return auth
}

func loadDB() *gorm.DB {
	db.Init(viper.GetString("database.driver"), viper.GetString("database.source"))
	return db.Get()
}

func setEngine(client *ethclient.Client, auth *bind.TransactOpts, db *gorm.DB) *gin.Engine {
	gin.SetMode(viper.GetString("server.mode"))
	r := gin.New()

	// 注册自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"formatUnixTime": func(ts string) string {
			timestamp, err := strconv.ParseInt(ts, 10, 64)
			if err != nil {
				return "时间格式错误"
			}
			return formatUnixTime(timestamp)
		},
	})

	// 设置中间件
	r.Use(gin.Recovery())
	r.Use(middleware.SetLoggerMiddleware(logger.DefaultLogger(viper.GetInt("log.level"), viper.GetString("log.format"))))
	r.Use(middleware.SetEthClientMiddleware(client))
	r.Use(middleware.SetNetworkMiddleware(network))
	r.Use(middleware.SetTransactOptsMiddleware(auth))
	r.Use(middleware.SetDBMiddleware(db))
	r.Use(middleware.SetLogMiddleware(logger.DefaultLogger(viper.GetInt("log.level"), viper.GetString("log.format"))))

	//设置路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	return r
}

func formatUnixTime(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}
