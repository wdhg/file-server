package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var createFiles = []struct {
	name     string
	contents string
	valid    bool
}{
	{"test.txt", "test file\n", true},
	{"test/test.txt", "test file\n", true},
	{"../test.txt", "test file\n", false},
}
var deleteFiles = []struct {
	name string
	path string
}{
	{"test.txt", ""},
}

func TestCreate(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)

	for _, testFile := range createFiles {
		if testFile.valid {
			testCreateValid(t, testFile.name, testFile.contents)
		} else {
			testCreateInvalid(t, testFile.name, testFile.contents)
		}
	}

	os.RemoveAll(dir)
}

func testCreateValid(t *testing.T, file, contents string) {
	// test file creating
	err := Create(file, contents)
	if err != nil {
		t.Errorf("Throwing error on making %s for first time", file)
	}
	if _, err = os.Stat(dir + file); os.IsNotExist(err) {
		t.Errorf("%s not created", file)
	}
	// test file contents
	if dat, _ := ioutil.ReadFile(dir + file); string(dat) != contents {
		t.Errorf("%s contents not being writen", file)
	}
	// test attempting to recreate files
	err = Create(file, contents)
	if err == nil {
		t.Errorf("Recreating %s not throwing error", file)
	}
}

func testCreateInvalid(t *testing.T, file, contents string) {
	err := Create(file, contents)
	if err == nil {
		t.Errorf("Can create file %s", file)
	}
}

func TestDelete(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)

	for _, testFile := range deleteFiles {
		file := dir + testFile.path + testFile.name
		os.MkdirAll(dir+testFile.path, os.ModePerm)
		os.Create(file)
		// test deleting file
		err := Delete(file)
		if err != nil {
			t.Errorf("Throwing error on deleting %s", file)
		}
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			t.Errorf("%s not deleted", file)
		}
	}

	os.RemoveAll(dir)
}
