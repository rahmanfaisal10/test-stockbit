package main

import (
	"test-stockbit/config"
	"test-stockbit/pkg/transport"
)

func main() {
	cfg := config.Get()

	recSrv := transport.NewGRPCServer()
	// fmt.Println(recSrv)
	transport.GRPCServerRun(
		cfg.PortGRPCServer,
		recSrv,
	)
}
