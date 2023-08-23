package main

import (
	"os"
)

func CreateFile(document *Document) (error) {
	f, err := os.Create(document.Name)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(document.Content)

	if err != nil {
		return err
	}

	return nil
}
