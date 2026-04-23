package main

import (
	"go-simple-api/internal/initialize"
	"go-simple-api/internal/server"
)

func main() {

	container := initialize.NewContainer()
	server.Start(container)
}
