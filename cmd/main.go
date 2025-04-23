package main

import (
	"log"

	"github.com/Komilov31/TaskManagerApi/cmd/api"
)

func main() {
	server := api.NewServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
