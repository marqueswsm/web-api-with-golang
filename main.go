package main

import (
	"github.com/marqueswsm/web-api-with-golang/database"
	"github.com/marqueswsm/web-api-with-golang/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
