package message_test

import (
	"context"
	"testing"
	"time"

	"github.com/aldan-ionos/go-training/protos/message"
)

var newLineList = []byte{}

func TestMessage(t *testing.T) {
	var (
		m   = message.NewMessage()
		err error
	)

	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MutexProfile, profile.ProfilePath(".")).Stop()

	// Read original file
	err = m.OpenFile("crimeandpunishment.txt")
	if err != nil {
		t.Errorf("Failed to open originalFile:\n\t- %s", err.Error())
	}

	// Open New File
	err = m.CreateFile("newfile.txt")
	if err != nil {
		t.Errorf("Failed to craete newFile:\n\t- %s", err.Error())
	}

	m.WaitGroup.Add(1)
	go m.ReadOriginalFile()

	m.WaitGroup.Add(1)
	go func(m *message.Message) {
		defer m.WaitGroup.Done()
		method := 2
		switch method {
		case 1:
			for {
				select {
				case <-time.After(1 * time.Second):
					// m.CloseFiles()
					return
				default:
					nextLine, err := m.GetNextLine(context.Background(), &message.Void{})
					if err != nil {
						t.Errorf("Failed to get new line:\n\t- %s", err.Error())
						return
					}

					if len(nextLine.NextLine) == 0 && err == nil {
						// m.CloseFiles()
						return
					}

					err = m.SaveNewLineToFile(nextLine.NextLine)
					if err != nil {
						t.Errorf("Failed to append new line to file:\n\t- %s", err.Error())
						return
					}

				}

			}
		case 2:

			noOfLines := 0
			for {
				select {
				case <-time.After(4 * time.Second):
					// m.CloseFiles()
					return
				default:
					nextLine, err := m.GetNextLine(context.Background(), &message.Void{})
					if err != nil {
						t.Errorf("Failed to get new line:\n\t- %s", err.Error())
						return
					}

					if len(nextLine.NextLine) == 0 && err == nil {
						// m.CloseFiles()
						return
					}

					if noOfLines == 50 {
						err = m.SaveNewLineToFile(newLineList)
						if err != nil {
							t.Errorf("Failed to append new line to file:\n\t- %s", err.Error())
							return
						}
						noOfLines = 0
						newLineList = []byte{}

					} else {
						newLineList = append(newLineList, nextLine.NextLine...)
					}
					noOfLines++

				}

			}
		}

	}(m)

	m.WaitGroup.Wait()

	// Close files
	m.CloseFiles()
}

func BenchmarkMessage(b *testing.B) {
	var (
		m   *message.Message
		err error
	)

	for i := 0; i < b.N; i++ {
		m = message.NewMessage()

		// Read original file
		err = m.OpenFile("crimeandpunishment.txt")
		if err != nil {
			b.Errorf("Failed to open originalFile:\n\t- %s", err.Error())
		}

		// Open New File
		err = m.CreateFile("newfile.txt")
		if err != nil {
			b.Errorf("Failed to craete newFile:\n\t- %s", err.Error())
		}

		m.WaitGroup.Add(1)
		go m.ReadOriginalFile()

		m.WaitGroup.Add(1)
		go func(m *message.Message) {
			defer m.WaitGroup.Done()
			for {
				select {
				case <-time.After(4 * time.Second):
					// m.CloseFiles()
					return
				default:
					nextLine, err := m.GetNextLine(context.Background(), &message.Void{})
					if err != nil {
						b.Errorf("Failed to get new line:\n\t- %s", err.Error())
						return
					}

					if len(nextLine.NextLine) == 0 && err == nil {
						// m.CloseFiles()
						return
					}

					err = m.SaveNewLineToFile(nextLine.NextLine)
					if err != nil {
						b.Errorf("Failed to append new line to file:\n\t- %s", err.Error())
						return
					}

				}

			}
		}(m)

		m.WaitGroup.Wait()

		// Close files
		m.CloseFiles()
	}
}
