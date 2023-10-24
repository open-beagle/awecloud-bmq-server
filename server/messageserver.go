package server

import (
	"context"
	"log"
	"os"

	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/server"

	"github.com/open-beagle/awecloud-bmq-server/pkg/conf"
)

func NewMessageServer() {
	log.Printf("message server : uses default config file.")
	cfg := genDefaultConfig()
	svr, err := server.NewService(cfg)
	if err != nil {
		log.Fatalf("message server : %v", err)
		os.Exit(1)
	}
	svr.Run(context.Background())
	log.Printf("message server : run.")
}

func genDefaultConfig() (cfg config.ServerCommonConf) {
	cfg = config.GetDefaultServerConf()
	cfg.Token = conf.Message.Token
	cfg.BindPort = conf.Message.Port
	return cfg
}
