package config

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

type Values struct {
	PrivateKey   string
	EthClientUrl string
	Port         int
	// TODO update to exchanges struct
	BinanceUrl string

	// TODO update to Tokens struct
	WBTCAddress string
	WETHAddress string
	USDCAddress string

	GinMode string
}

func envArrayToAddressSlice(s string) []common.Address {
	env := strings.Split(s, ",")
	slc := []common.Address{}
	for _, ep := range env {
		slc = append(slc, common.HexToAddress(strings.TrimSpace(ep)))
	}

	return slc
}

func variableNotSetOrIsNil(env string) bool {
	return !viper.IsSet(env) || viper.GetString(env) == ""
}

func GetValues() *Values {
	// Default variables
	viper.SetDefault("port", 8080)
	viper.SetDefault("debug_mode", false)
	viper.SetDefault("gin_mode", gin.ReleaseMode)

	// Read in from .env file if available
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			// Can ignore
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Read in from environment variables
	_ = viper.BindEnv("port")
	_ = viper.BindEnv("debug_mode")
	_ = viper.BindEnv("gin_mode")

	// Validate required variables
	if variableNotSetOrIsNil("eth_client_url") {
		panic("Fatal config error: eth_client_url not set")
	}

	if variableNotSetOrIsNil("private_key") {
		panic("Fatal config error: private_key not set")
	}

	// Return Values
	privateKey := viper.GetString("private_key")
	ethClientUrl := viper.GetString("eth_client_url")
	port := viper.GetInt("port")

	ginMode := viper.GetString("gin_mode")
	return &Values{
		PrivateKey:   privateKey,
		EthClientUrl: ethClientUrl,
		Port:         port,
		GinMode:      ginMode,
	}
}
