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

func moveToNewLocation(originalFile, newFile *os.File) error {
	scanner := bufio.NewScanner(originalFile)

	for scanner.Scan() {
		newFile.Write(append(scanner.Bytes(), []byte("\n")...))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func openFile(filename string, option int) (*os.File, error) {
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
	originalFile, err := openFile("crimeandpunishment.txt", READ)
	if err != nil {
		log.Fatalf("Failed to open originalFile:\n\t- %s", err.Error())
	}

	// Open New File
	newFile, err := openFile("newfile.txt", WRITE)
	if err != nil {
		log.Fatalf("Failed to craete newFile:\n\t- %s", err.Error())
	}

	err = moveToNewLocation(originalFile, newFile)
	if err != nil {
		log.Fatalf("Failed to move content from one file to another:\n\t- %s", err.Error())
	}

	// Close files
	defer originalFile.Close()
	defer newFile.Close()
}
