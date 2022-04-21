package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aldan-ionos/go-training/protos/message"
	"google.golang.org/grpc"
)

func StartServer(protocol string, port int) {
	listener, err := net.Listen(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d.\n\t- %s", port, err.Error())
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	messageServer := message.NewMessage()

	// Read original file
	err = messageServer.OpenFile("crimeandpunishment.txt")
	if err != nil {
		log.Fatalf("Failed to open originalFile:\n\t- %s", err.Error())
		os.Exit(1)
	}

	message.RegisterMessageServiceServer(grpcServer, messageServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server gRPC server over port %d.\n\t- %s", port, err.Error())
		os.Exit(1)
	}

	messageServer.WaitGroup.Add(1)
	go messageServer.ReadOriginalFile()

	messageServer.WaitGroup.Wait()

	// messageServer.CloseFiles()
}
