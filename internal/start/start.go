package start

import (
	"context"
	"fmt"
	"log"

	"github.com/armatrix/priceFeed/internal/config"
	"github.com/armatrix/priceFeed/internal/handlers"
	"github.com/armatrix/priceFeed/internal/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func QAMod() {
	conf := config.GetValues()

	logr := logger.NewZeroLogr().
		WithName("backendstackup_bundler").
		WithValues("Mod", "QAMod")

	rpc, err := rpc.Dial(conf.EthClientUrl)
	if err != nil {
		log.Fatal(err)
	}

	eth := ethclient.NewClient(rpc)

	chainID, err := eth.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Print("current chainId: ", chainID)

	// Init HTTP server
	gin.SetMode(conf.GinMode)
	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}
	r.Use(
		cors.Default(),
		logger.WithLogr(logr),
		gin.Recovery(),
	)

	r.GET("/ping", handlers.Ping())

	if err := r.Run(fmt.Sprintf("run gin service on: %d", conf.Port)); err != nil {
		log.Fatal(err)
	}
}
