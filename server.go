package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const dir = "files/"

func exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func accessible(file string) bool {
	x, _ := filepath.Abs(file)
	y, _ := filepath.Abs(dir)
	return strings.Index(x, y) == 0
}

// Create is called on POST requests to create new files
func Create(file string, contents string) error {
	file = dir + file
	if !accessible(file) {
		return errors.New("File above allocated directory")
	}
	// check if file already exists
	if exists(file) {
		return errors.New("File already exists")
	}
	os.MkdirAll(filepath.Dir(file), os.ModePerm)
	fileWriter, _ := os.Create(file)
	fmt.Fprintf(fileWriter, contents)
	return nil
}

func Delete(file string) error {
	os.Remove(file)
	return nil
}

func main() {
	fmt.Println("hello")
}
