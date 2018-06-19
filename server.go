package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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

func Get(file string) (string, error) {
	file = dir + file
	if !accessible(file) {
		return "", errors.New("File above allocated directory")
	}
	if !exists(file) {
		return "", errors.New("File does not exist")
	}
	data, _ := ioutil.ReadFile(file)
	return string(data), nil
}

func Create(file string, contents string) error {
	file = dir + file
	if !accessible(file) {
		return errors.New("File above allocated directory")
	}
	if exists(file) {
		return errors.New("File already exists")
	}
	os.MkdirAll(filepath.Dir(file), os.ModePerm)
	fileWriter, _ := os.Create(file)
	fmt.Fprintf(fileWriter, contents)
	return nil
}

func Update(file, contents string) error {
	file = dir + file
	if !accessible(file) {
		return errors.New("File above allocated directory")
	}
	if !exists(file) {
		return errors.New("File does not exist")
	}
	ioutil.WriteFile(dir+file, []byte(contents), os.ModePerm)
	return nil
}

func Delete(file string) error {
	file = dir + file
	if !accessible(file) {
		return errors.New("File above allocated directory")
	}
	os.Remove(file)
	return nil
}

func main() {
	fmt.Println("hello")
}
