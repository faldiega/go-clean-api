package main

import (
	"go-clean-api/internal/initialize"
	"go-clean-api/internal/server"
)

func main() {

	container := initialize.NewContainer()
	server.Start(container)
}
