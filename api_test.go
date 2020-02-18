package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_userHandler(t *testing.T) {
	DB = ConfigureMockedDB()
	req, err := http.NewRequest(
		"POST",
		"/users",
		ioutil.NopCloser(strings.NewReader(`{
			"name" : "Marco",
			"pass" : "1234",
			"email": "xmarcoied@gmail.com",
			"address" : "Maaadi",
			"age" : 125
		}
	`)))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
