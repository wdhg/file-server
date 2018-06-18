package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var testFiles = []struct {
	name     string
	contents string
	valid    bool
}{
	{"test.txt", "test file\n", true},
	{"test/test.txt", "test file\n", true},
	{"../test.txt", "test file\n", false},
}

func TestCreate(t *testing.T) {
	// setup code
	os.Mkdir(dir, os.ModePerm)

	for _, testFile := range testFiles {
		file, contents := testFile.name, testFile.contents
		if testFile.valid {
			doTestCreateValid(t, file, contents)
		} else {
			doTestCreateInvalid(t, file, contents)
		}
	}

	// teardown code
	os.RemoveAll(dir)
}

func doTestCreateValid(t *testing.T, file, contents string) {
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

func doTestCreateInvalid(t *testing.T, file, contents string) {
	err := Create(file, contents)
	if err == nil {
		t.Errorf("Can create file %s", file)
	}
}
