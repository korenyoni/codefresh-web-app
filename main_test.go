package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var path404 string = "/askdasjldka"

func TestIndexHandler(t *testing.T) {
	router := BuildRouter()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fail()
	}
}

func TestHealthCheckHandler(t *testing.T) {
	router := BuildRouter()

	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fail()
	}
}

func TestDefaultHandler(t *testing.T) {
	router := BuildRouter()

	req, err := http.NewRequest("GET", path404, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fail()
	}
}
