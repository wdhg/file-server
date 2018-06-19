package main

import (
	"io/ioutil"
	"os"
	"testing"
)

type File struct {
	path     string
	contents string
	valid    bool
}

var createFiles = []File{
	{"test.txt", "test file\n", true},
	{"test/test.txt", "test file\n", true},
	{"../test.txt", "test file\n", false},
}
var deleteFiles = []File{
	{"test.txt", "", true},
	{"../test.txt", "", false},
}

func TestCreate(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)

	for _, file := range createFiles {
		if !file.valid {
			err := Create(file.path, file.contents)
			if err == nil {
				t.Errorf("Can create file %s", file.path)
			}
			continue
		}
		// test file creating
		err := Create(file.path, file.contents)
		if err != nil {
			t.Errorf("Error returned when making %s for first time", file.path)
		}
		if _, err = os.Stat(dir + file.path); os.IsNotExist(err) {
			t.Errorf("%s not created", file.path)
		}
		// test file contents
		if dat, _ := ioutil.ReadFile(dir + file.path); string(dat) != file.contents {
			t.Errorf("%s contents not being writen", file.path)
		}
		// test attempting to recreate files
		err = Create(file.path, file.contents)
		if err == nil {
			t.Errorf("Recreating %s not returning error", file.path)
		}
	}

	os.RemoveAll(dir)
}

func TestDelete(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)

	for _, file := range deleteFiles {
		if !file.valid {
			// test deleting files outside allocated directory
			err := Delete(file.path)
			if err == nil {
				t.Errorf("Error not returned when trying to delete unaccessible file %s", file.path)
			}
			continue
		}
		// test deleting file that doesn't exist
		err := Delete(file.path)
		if err != nil {
			t.Errorf("Error not returned when trying to delete deleted file %s", file.path)
		}
		os.MkdirAll(dir+file.path, os.ModePerm)
		os.Create(dir + file.path)
		// test deleting file
		err = Delete(file.path)
		if err != nil {
			t.Errorf("Error returned when deleting %s", file.path)
		}
		if _, err := os.Stat(dir + file.path); !os.IsNotExist(err) {
			t.Errorf("%s not deleted", file.path)
		}
		// make sure file is deleted
		os.Remove(dir + file.path)
	}

	os.RemoveAll(dir)
}
