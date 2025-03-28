package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile("./conf/conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("load config error: %s", err.Error()))
	}
}
