package storage

import "os"

const EmailFilePath = "emails.txt"

type EmailWriter interface {
	Write(email string) error
}

type FileEmailWriter struct {
	filepath string
}

func NewFileEmailWriter() *FileEmailWriter {
	return &FileEmailWriter{filepath: EmailFilePath}
}

func (w *FileEmailWriter) Write(email string) error {
	f, err := os.OpenFile(w.filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(email + "\n")
	return err
}
