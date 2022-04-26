package message

import (
	"bufio"
	context "context"
	"log"
	"os"
	sync "sync"
)

type Message struct {
	originalFile *os.File
	newFile      *os.File

	messageChan chan []byte
	WaitGroup   sync.WaitGroup
}

func NewMessage() *Message {
	m := &Message{}
	m.messageChan = make(chan []byte)

	return m
}

func (m *Message) MoveToNewLocation() error {
	scanner := bufio.NewScanner(m.originalFile)

	for scanner.Scan() {
		m.newFile.Write(append(scanner.Bytes(), []byte("\n")...))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (m *Message) ReadOriginalFile() {
	defer m.WaitGroup.Done()

	scanner := bufio.NewScanner(m.originalFile)

	for scanner.Scan() {
		// m.newFile.Write(append(scanner.Bytes(), []byte("\n")...))
		m.messageChan <- append(scanner.Bytes(), []byte("\n")...)
		//fmt.Printf(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan original file:\n\t- %s", err.Error())
		return
	}

	close(m.messageChan)

}

func (m *Message) SaveNewLineToFile(newLine []byte) error {
	_, err := m.newFile.Write(newLine)
	return err
}

func (m *Message) OpenFile(filename string) error {
	var err error

	m.originalFile, err = os.Open(filename)
	return err
}

func (m *Message) CreateFile(filename string) error {
	var err error
	m.newFile, err = os.Create(filename)
	return err
}

func (m *Message) GetFiles() (*os.File, *os.File) {
	return m.originalFile, m.newFile
}

func (m *Message) CloseFiles() {
	m.originalFile.Close()
	m.newFile.Close()
}

func (m *Message) GetNextLine(ctx context.Context, _ *Void) (*Line, error) {
	newLine, ok := <-m.messageChan
	if ok {
		return &Line{NextLine: newLine}, nil
	}

	return &Line{}, nil

}
