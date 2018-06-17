package main

import (
	"errors"
	"fmt"
	"os"
)

const dir = "files/"

// Create is called on POST requests to create new files
func Create(filename string, contents string) error {
	// append "files/" to start
	filename = dir + filename
	if _, err := os.Stat(filename); err == nil {
		return errors.New("File already exists")
	}
	file, _ := os.Create(filename)
	fmt.Fprintf(file, contents)
	return nil
}

func main() {
	fmt.Println("hello")
}
