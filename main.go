package main

import "github.com/marqueswsm/web-api-with-golang/server"

func main() {
	server := server.NewServer()

	server.Run()
}
