package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	// test file creating
	err := Create("test.txt")
	if err != nil {
		t.Error("Throwing error on making file for first time")
	}
	if _, err = os.Stat("files/test.txt"); os.IsNotExist(err) {
		t.Error("File not created")
	}
	// test attempting to recreate files
	err = Create("test.txt")
	if err == nil {
		t.Error("Recreating file not throwing error")
	}
}
