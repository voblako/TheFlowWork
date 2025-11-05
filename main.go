package main

import (
	"log"

	"github.com/voblako/TheFlowWork/api"
)

func main() {
	config := api.Config{
		ListenAddr: ":3000",
	}
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatalln(err)
	}

	server.Start()
}
