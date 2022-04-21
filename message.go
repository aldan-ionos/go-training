package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	READ = iota
	WRITE
)

type Message struct{}

func (m Message) MoveToNewLocation(originalFile, newFile *os.File) error {
	scanner := bufio.NewScanner(originalFile)

	for scanner.Scan() {
		newFile.Write(append(scanner.Bytes(), []byte("\n")...))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (m Message) OpenFile(filename string, option int) (*os.File, error) {
	var (
		file *os.File
		err  error
	)

	switch option {
	case READ:
		file, err = os.Open(filename)
	case WRITE:
		file, err = os.Create(filename)
	default:
		err = fmt.Errorf("unknown option %d. Expected either %d or %d", option, READ, WRITE)
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return file, nil
}

func main() {
	// Read original file
	m := Message{}
	originalFile, err := m.OpenFile("crimeandpunishment.txt", READ)
	if err != nil {
		log.Fatalf("Failed to open originalFile:\n\t- %s", err.Error())
	}

	// Open New File
	newFile, err := m.OpenFile("newfile.txt", WRITE)
	if err != nil {
		log.Fatalf("Failed to craete newFile:\n\t- %s", err.Error())
	}

	err = m.MoveToNewLocation(originalFile, newFile)
	if err != nil {
		log.Fatalf("Failed to move content from one file to another:\n\t- %s", err.Error())
	}

	// Close files
	defer originalFile.Close()
	defer newFile.Close()
}
