package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type File struct {
	path     string
	contents string
	valid    bool
}

var getFiles = []File{
	{"test.txt", "test file\n", true},
	{"../test.txt", "", false},
}
var createFiles = []File{
	{"test.txt", "test file\n", true},
	{"test/test.txt", "test file\n", true},
	{"../test.txt", "test file\n", false},
}
var updateFiles = []File{
	{"test.txt", "test file\n", true},
	{"../test.txt", "test file\n", false},
}
var deleteFiles = []File{
	{"test.txt", "", true},
	{"../test.txt", "", false},
}

func TestGet(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)
	defer os.RemoveAll(dir)

	for _, file := range getFiles {
		if !file.valid {
			if _, err := Get(file.path); err == nil {
				t.Errorf("Not returning error when trying to access unaccessible file %s", file.path)
			}
			continue
		}

		// test reading nonexistent file
		_, err := Get(file.path)
		if err == nil {
			t.Errorf("Not returning error when trying to read nonexistent file %s", file.path)
		}
		// make the file
		os.MkdirAll(dir+filepath.Dir(file.path), os.ModePerm)
		fileWriter, _ := os.Create(dir + file.path)
		fmt.Fprintf(fileWriter, file.contents)
		// test reading the file
		contents, err := Get(file.path)
		if err != nil {
			t.Errorf("Returning error when trying to read valid file %s", file.path)
		}
		if contents != file.contents {
			t.Errorf("Not writing to file %s properly", file.path)
		}
	}
}

func TestCreate(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)
	defer os.RemoveAll(dir)

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
}

func TestUpdate(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)
	defer os.RemoveAll(dir)

	for _, file := range updateFiles {
		if !file.valid {
			err := Update(file.path, file.contents)
			if err == nil {
				t.Errorf("Not returning error when trying to access unaccessible file %s", file.path)
			}
			continue
		}

		// test updating nonexistent file
		err := Update(file.path, file.contents)
		if err == nil {
			t.Errorf("Not returning error when updating noexistent file %s", file.path)
		}
		os.MkdirAll(dir+filepath.Dir(file.path), os.ModePerm)
		os.Create(dir + file.path)
		// test updating a file
		err = Update(file.path, file.contents)
		if err != nil {
			t.Errorf("Returning error when updating file %s", file.path)
		}
		// test file contents
		if dat, _ := ioutil.ReadFile(dir + file.path); string(dat) != file.contents {
			t.Errorf("%s contents not being updated", file.path)
		}
	}
}

func TestDelete(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)
	defer os.RemoveAll(dir)

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
		// make the file
		os.MkdirAll(filepath.Dir(dir+file.path), os.ModePerm)
		os.Create(dir + file.path)
		// test deleting file
		err = Delete(file.path)
		if err != nil {
			t.Errorf("Error returned when deleting %s", file.path)
		}
		if _, err := os.Stat(dir + file.path); !os.IsNotExist(err) {
			t.Errorf("%s not deleted", file.path)
		}
	}
}
