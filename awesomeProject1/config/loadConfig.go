package config

import (
	"awesomeProject1/config/envConfig"
	"fmt"
	"github.com/imdario/mergo"
	"os"
)

var Config = envConfig.Default

func init() {
	//mergo.Merge(&Config, envConfig.TQHNET, mergo.WithOverride)
	// 根据环境变量合并不同的配置
	env := os.Getenv("GO_ENV")

	switch env {
	case "production":
		mergo.Merge(&Config, envConfig.Production, mergo.WithOverride)
		break
	case "dev":
		mergo.Merge(&Config, envConfig.Dev, mergo.WithOverride)
		break
	case "orlando":
		mergo.Merge(&Config, envConfig.Orlando, mergo.WithOverride)
		break
	case "TQHNET":
		mergo.Merge(&Config, envConfig.TQHNET, mergo.WithOverride)
		break
	case "lichaolong":
		mergo.Merge(&Config, envConfig.Lichaolong, mergo.WithOverride)
		break
	}
	fmt.Printf("\n环境变量初始化：GO_ENV=" + env + "\n")
}
