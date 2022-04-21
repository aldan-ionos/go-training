package server

import (
	"fmt"
	"log"
	"net"

	"github.com/aldan-ionos/go-training/protos/message"
	"google.golang.org/grpc"
)

type Server struct{}

func (s Server) StartServer(protocol string, port int) {
	listener, err := net.Listen(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d.\n\t- %s", port, err.Error())
	}

	grpcServer := grpc.NewServer()
	messageServer := message.NewMessage()
	message.RegisterMessageServiceServer(grpcServer, messageServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server gRPC server over port %d.\n\t- %s", port, err.Error())
	}
}
