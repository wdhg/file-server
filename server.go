package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const dir = "files/"

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Create is called on POST requests to create new files
func Create(filename string, contents string) error {
	filename = dir + filename
	// check if file is valid
	filenameAbs, _ := filepath.Abs(filename)
	dirAbs, _ := filepath.Abs(dir)
	if !strings.Contains(filenameAbs, dirAbs) {
		return errors.New("File above allocated directory")
	}
	// check if file already exists
	if pathExists(filename) {
		return errors.New("File already exists")
	}
	dir := filepath.Dir(filename)
	// create directory if it doesn't exist
	if !pathExists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
	file, _ := os.Create(filename)
	fmt.Fprintf(file, contents)
	return nil
}

func main() {
	fmt.Println("hello")
}
