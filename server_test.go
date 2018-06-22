package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestServer(t *testing.T) {
	os.Mkdir(dir, os.ModePerm)
	defer os.RemoveAll(dir)
	gin.SetMode(gin.TestMode)
	router := CreateRouter()
	writer := httptest.NewRecorder()

	// test get
	for _, file := range getFiles {
		if !file.valid {
			continue
		}

		// make the file
		os.Mkdir(dir+filepath.Dir(file.path), os.ModePerm)
		fileWriter, _ := os.Create(dir + file.path)
		fmt.Fprintf(fileWriter, file.contents)
		// test getting file contents through the server
		req, _ := http.NewRequest(http.MethodGet, "/files/"+file.path, nil)
		router.ServeHTTP(writer, req)
		if writer.Code != http.StatusOK {
			t.Errorf("Receiving error code on valid request")
		}
		if writer.Body.String() != file.contents {
			t.Errorf("Served file contents for %s are incorrect", file.path)
		}

		os.RemoveAll(dir + filepath.Dir(file.path))
	}

	// test create
	for _, file := range createFiles {
		if !file.valid {
			continue
		}

		// create url
		URL, _ := url.Parse("/files/" + file.path)
		params := url.Values{}
		params.Add("contents", file.contents)
		URL.RawQuery = params.Encode()
		// test if server is creating the file correctly
		req, _ := http.NewRequest(http.MethodPost, URL.String(), nil)
		router.ServeHTTP(writer, req)
		if writer.Code != http.StatusOK {
			t.Errorf("Receiving error code on valid request")
		}
		if dat, _ := ioutil.ReadFile(dir + file.path); string(dat) != file.contents {
			t.Errorf("Created file %s doesn't contain correct contents", file.path)
		}

		os.RemoveAll(dir + filepath.Dir(file.path))
	}

	// test update
	for _, file := range updateFiles {
		if !file.valid {
			continue
		}

		// create url
		URL, _ := url.Parse("/files/" + file.path)
		params := url.Values{}
		params.Add("contents", file.contents)
		URL.RawQuery = params.Encode()
		// make the file
		os.MkdirAll(filepath.Dir(file.path), os.ModePerm)
		os.Create(dir + file.path)
		// test if server is update files
		req, _ := http.NewRequest(http.MethodPut, URL.String(), nil)
		router.ServeHTTP(writer, req)
		if writer.Code != http.StatusOK {
			t.Errorf("Receiving error code on valid request")
		}
		if dat, _ := ioutil.ReadFile(dir + file.path); string(dat) != file.contents {
			t.Errorf("Updated file %s doesn't contain correct contents", file.path)
		}

		os.RemoveAll(dir + filepath.Dir(file.path))
	}

	// test delete
	for _, file := range deleteFiles {
		if !file.valid {
			continue
		}

		// make the file
		os.MkdirAll(filepath.Dir(file.path), os.ModePerm)
		os.Create(dir + file.path)
		// test if server deletes the file
		req, _ := http.NewRequest(http.MethodDelete, "/files/"+file.path, nil)
		router.ServeHTTP(writer, req)
		if writer.Code != http.StatusOK {
			t.Errorf("Receiving error code on valid request")
		}
		if _, err := os.Stat(dir + file.path); !os.IsNotExist(err) {
			t.Errorf("%s not deleted", file.path)
		}

		os.RemoveAll(dir + filepath.Dir(file.path))
	}
}
