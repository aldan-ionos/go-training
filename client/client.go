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

type Client struct{}

func (c *Client) SaveNewLineToFile(messageClient message.MessageServiceClient, filename string) error {
	m := message.NewMessage()
	err := m.CreateFile(filename)
	if err != nil {
		return err
	}

	for {
		select {
		case <-time.After(4 * time.Second):
			m.CloseFiles()
			return nil
		default:
			nextLine, err := messageClient.GetNextLine(context.Background(), nil)
			if err != nil {
				return err
			}

			err = m.SaveNewLineToFile(nextLine.NextLine)
			if err != nil {
				return err
			}

		}

	}

}

func (c *Client) StartClient(port int) {
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
	if err = c.SaveNewLineToFile(messageClient, "crimeandpunishment.txt"); err != nil {
		log.Fatalf("Failed to save next line to file.\n\t- %s", err.Error())
		os.Exit(1)
	}

}
