package main

import (
	"bufio"
	"log"
	"os"
)

type Message struct {
	originalFile *os.File
	newFile      *os.File
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

func main() {

	var (
		m   = Message{}
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

	err = m.MoveToNewLocation()
	if err != nil {
		log.Fatalf("Failed to move content from one file to another:\n\t- %s", err.Error())
	}

	// Close files
	m.CloseFiles()
}
