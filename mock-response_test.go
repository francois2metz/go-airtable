// Copyright © 2020 Mike Berezin
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package airtable

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

func mockResponse(paths ...string) *httptest.Server {
	parts := []string{".", "testdata"}
	filename := filepath.Join(append(parts, paths...)...)

	mockData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		_, err := rw.Write(mockData)
		if err != nil {
			log.Fatal(err)
		}
	}))
}

func mockErrorResponse(code int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		http.Error(rw, "An error occurred", code)
	}))
}
