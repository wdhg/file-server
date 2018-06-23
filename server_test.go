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

func makeFile(file File, writeContents bool) {
	os.Mkdir(dir+filepath.Dir(file.path), os.ModePerm)
	fileWriter, _ := os.Create(dir + file.path)
	if writeContents {
		fmt.Fprintf(fileWriter, file.contents)
	}
}

func testServerMethod(method string, files []File) func(*testing.T) {
	return func(t *testing.T) {
		os.Mkdir(dir, os.ModePerm)
		defer os.RemoveAll(dir)
		gin.SetMode(gin.TestMode)
		router := CreateRouter()
		writer := httptest.NewRecorder()

		for _, file := range files {
			if !file.valid {
				continue
			}

			// if the server isn't going to make the file itself
			if method != http.MethodPost {
				makeFile(file, method != http.MethodPut)
			}

			// create the url for the request
			requestURL := "/files/" + file.path
			if method == http.MethodPost || method == http.MethodPut {
				// add the url contents parameter
				URL, _ := url.Parse(requestURL)
				params := url.Values{}
				params.Add("contents", file.contents)
				URL.RawQuery = params.Encode()
				requestURL = URL.String()
			}

			req, _ := http.NewRequest(method, requestURL, nil)
			router.ServeHTTP(writer, req)
			// test status code is correct
			if writer.Code != http.StatusOK {
				t.Errorf("Receiving error code on valid request")
			}
			// run test depending on method
			switch method {
			case http.MethodGet:
				if writer.Body.String() != file.contents {
					t.Errorf("Served file contents for %s are incorrect", file.path)
				}
			case http.MethodPost:
				if dat, _ := ioutil.ReadFile(dir + file.path); string(dat) != file.contents {
					t.Errorf("Created file %s doesn't contain correct contents", file.path)
				}
			case http.MethodPut:
				if dat, _ := ioutil.ReadFile(dir + file.path); string(dat) != file.contents {
					t.Errorf("Updated file %s doesn't contain correct contents", file.path)
				}
			case http.MethodDelete:
				if _, err := os.Stat(dir + file.path); !os.IsNotExist(err) {
					t.Errorf("%s not deleted", file.path)
				}
			}
		}
	}
}

func TestServer(t *testing.T) {
	t.Run(http.MethodGet, testServerMethod(http.MethodGet, getFiles))
	t.Run(http.MethodPost, testServerMethod(http.MethodPost, createFiles))
	t.Run(http.MethodPut, testServerMethod(http.MethodPut, updateFiles))
	t.Run(http.MethodDelete, testServerMethod(http.MethodDelete, deleteFiles))
}
