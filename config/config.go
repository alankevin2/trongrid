package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	ApiKey  string
	BaseURL string
	Tokens  map[string]string
}

func GetConfig(network string) Config {
	var fileName string
	switch network {
	case "Mainnet":
		fileName = "provider-mainnet.yml"
	case "Shasta":
		fileName = "provider-testnet-shasta.yml"
	}

	currentPath, _ := os.Getwd()
	fullpath := path.Join(currentPath, "config", fileName)
	_, err := os.Stat(fullpath)
	// 如果找不到，代表當前執行環境不是以此pkg為主，而是被別人vendor引用
	if err != nil {
		pkgPath := reflect.TypeOf(Config{}).PkgPath() + "../../../config/"
		fullpath = path.Join(currentPath, "vendor", pkgPath, fileName)
	}
	viper.SetConfigFile(fullpath)
	viper.SetConfigType("yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}

	return Config{
		ApiKey:  viper.GetString("root.api-key"),
		BaseURL: viper.GetString("root.url"),
		Tokens:  viper.GetStringMapString("root.tokens"),
	}
}
