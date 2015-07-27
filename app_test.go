package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateResourceHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(createResourceHandler))
	defer ts.Close()

	request := "{\"id\":\"\",\"link\":\"http://www.google.com\"}"
	body := strings.NewReader(request)
	res, err := http.Post(ts.URL, "application/json", body)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	actualBytes, err := ioutil.ReadAll(res.Body)
	actual := string(actualBytes)
	res.Body.Close()
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	if actual != request {
		t.Errorf("Failed asserting %s to be %s", actual, request)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Failed asserting status code %d to be %d", res.StatusCode, http.StatusCreated)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Failed asserting %s to be %s", contentType, "application/json")
	}
}
