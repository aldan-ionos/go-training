package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aldan-ionos/go-training/protos/message"
	"google.golang.org/grpc"
)

func SaveNewLineToFile(messageClient message.MessageServiceClient, filePath string) error {
	m := message.NewMessage()
	err := m.CreateFile(filePath)
	if err != nil {
		return err
	}

	for {
		select {
		case <-time.After(4 * time.Second):
			m.CloseFiles()
			return nil
		default:
			nextLine, err := messageClient.GetNextLine(context.Background(), &message.Void{})
			if err != nil {
				return err
			}

			if len(nextLine.NextLine) == 0 && err == nil {
				m.CloseFiles()
				return nil
			}

			err = m.SaveNewLineToFile(nextLine.NextLine)
			if err != nil {
				return err
			}

		}

	}

}

func StartClient(port int, filePath string) {
	var (
		conn *grpc.ClientConn
		err  error
	)

	defer conn.Close()

	conn, err = grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to port %d.\n\t- %s", port, err.Error())
		os.Exit(1)
	}

	messageClient := message.NewMessageServiceClient(conn)
	if err = SaveNewLineToFile(messageClient, filePath); err != nil {
		log.Fatalf("Failed to save next line to file.\n\t- %s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)

}
