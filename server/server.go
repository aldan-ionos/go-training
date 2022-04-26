package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aldan-ionos/go-training/protos/message"
	"google.golang.org/grpc"
)

func StartServer(protocol string, port int, filePath string) {
	listener, err := net.Listen(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d.\n\t- %s", port, err.Error())
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	messageServer := message.NewMessage()

	// Read original file
	err = messageServer.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open originalFile:\n\t- %s", err.Error())
		os.Exit(1)
	}

	message.RegisterMessageServiceServer(grpcServer, messageServer)

	messageServer.WaitGroup.Add(1)
	go messageServer.ReadOriginalFile()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server gRPC server over port %d.\n\t- %s", port, err.Error())
		os.Exit(1)
	}

	messageServer.WaitGroup.Wait()

	// messageServer.CloseFiles()
}
