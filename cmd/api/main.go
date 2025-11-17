package main

import (
	"log"

	config2 "github.com/voblako/TheFlowWork/internal/config"
	"github.com/voblako/TheFlowWork/internal/http-server"
)


func main() {

	conf := config2.MustLoadConf()

	server, err := http_server.NewServer(conf)
	if err != nil {
		log.Fatalln(err)
	}

	server.Start()
}
