package main

import (
	"flag"
	"os"

	"github.com/aldan-ionos/go-training/client"
	"github.com/aldan-ionos/go-training/server"
	"github.com/prometheus/common/log"
)

func main() {
	var (
		startServer    bool
		startClient    bool
		serverProtocol string
		port           int
	)
	flag.BoolVar(&startServer, "start-server", false, "Start the gRPC server.")
	flag.BoolVar(&startClient, "start-client", false, "Start the gRPC client.")
	flag.StringVar(&serverProtocol, "protocol", "tcp", "The internet protocol suite used by the gRPC server.")
	flag.IntVar(&port, "port", 9000, "The port on which the client-server communication will take place.")

	flag.Parse()

	switch {
	case startServer:
		log.Info("Starting gRPC server.")
		server.StartServer(serverProtocol, port)
	case startClient:
		log.Info("Starting gRPC client.")
		client.StartClient(port)
	default:
		log.Error("Neither the \"-start-server\" nor the \"-start-client\" flags were set.")
		os.Exit(1)
	}

}
