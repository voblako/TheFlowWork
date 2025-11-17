package main

import (
	"log"

	"github.com/voblako/TheFlowWork/api"
	"github.com/voblako/TheFlowWork/storage"
)

func main() {
	config := api.Config{
		ListenAddr: ":3000",
		Storage: storage.Config{
			//"postgres://username:password@localhost:5432/database_name"
			DatabaseURL: "postgres://main:qwerty@localhost:5432/flowwork",
		},
	}
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatalln(err)
	}

	server.Start()
}
