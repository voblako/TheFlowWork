package main

import (
	"log"

	config2 "github.com/voblako/TheFlowWork/internal/config"
	"github.com/voblako/TheFlowWork/internal/api"
)


func main() {

	conf := config2.MustLoadConf()

	server, err := api.NewServer(conf)
	if err != nil {
		log.Fatalln(err)
	}

	server.Start()
}
