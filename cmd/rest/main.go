package main

import (
	"context"

	"go_microservices_skeleton/config"
	"go_microservices_skeleton/server"
)

func init() {
	config.Load()
}

func main() {
	conf := config.Gateway()
	ctx := context.Background()
	server.StartRESTGateway(ctx, conf)
}
