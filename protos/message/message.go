package main

import (
	"bufio"
	"log"
	"os"
	sync "sync"
)

type Message struct {
	originalFile *os.File
	newFile      *os.File

	messageChan chan []byte
	waitGroup   sync.WaitGroup
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
	defer m.waitGroup.Done()

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

func (m *Message) SaveNewLineToFile() {
	defer m.waitGroup.Done()

	for {
		select {
		case newLine, ok := <-m.messageChan:
			if ok {
				_, err := m.newFile.Write(newLine)
				if err != nil {
					log.Fatalf("Failed to write line to new file:\n\t- %s", err.Error())
					return
				}
			} else {
				return
			}
		}
	}

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

// func (m *Message) GetNextLine(ctx context.Context, line *Line) (*Line, error) {
// 	return line, nil
// }

func main() {

	var (
		m   = NewMessage()
		err error
	)

	// Read original file
	err = m.OpenFile("crimeandpunishment.txt")
	if err != nil {
		log.Fatalf("Failed to open originalFile:\n\t- %s", err.Error())
	}

	// Open New File
	err = m.CreateFile("newfile.txt")
	if err != nil {
		log.Fatalf("Failed to craete newFile:\n\t- %s", err.Error())
	}

	// err = m.MoveToNewLocation()
	// if err != nil {
	// 	log.Fatalf("Failed to move content from one file to another:\n\t- %s", err.Error())
	// }

	m.waitGroup.Add(1)
	go m.ReadOriginalFile()

	m.waitGroup.Add(1)
	m.SaveNewLineToFile()

	m.waitGroup.Wait()

	// Close files
	m.CloseFiles()
}
