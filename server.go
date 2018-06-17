package main

import (
	"errors"
	"fmt"
	"os"
)

const dir = "files/"

// Create is called on POST requests to create new files
func Create(file string) error {
	// append "files/" to start
	file = dir + file
	if _, err := os.Stat(file); err == nil {
		return errors.New("File already exists")
	}
	os.Create(file)
	return nil
}

func main() {
	fmt.Println("hello")
}
