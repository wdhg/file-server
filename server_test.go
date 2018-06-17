package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	Create("test.txt")
	if _, err := os.Stat("files/test.txt"); os.IsNotExist(err) {
		t.Error("File not created")
	}
}
