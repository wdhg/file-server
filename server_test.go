package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	// setup code. clear files
	os.RemoveAll("files/")
	os.Mkdir("files/", os.ModePerm)

	// test file creating
	err := Create("test.txt", "test file\n")
	if err != nil {
		t.Error("Throwing error on making file for first time")
	}
	if _, err = os.Stat("files/test.txt"); os.IsNotExist(err) {
		t.Error("File not created")
	}
	// test file contents
	if dat, _ := ioutil.ReadFile("files/test.txt"); string(dat) != "test file\n" {
		t.Error("File content not being writen")
	}
	// test attempting to recreate files
	err = Create("test.txt", "")
	if err == nil {
		t.Error("Recreating file not throwing error")
	}
}
