package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteServerInternalServerError(t *testing.T) {
	sampleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Hello HTTP Test")
	})
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	_, err := fetchUserData(ts.URL)
	if err.Error() != "500 Internal Server Error" {
		t.Error(err.Error())
	}
}

func TestRemoteServerStatusNotFound(t *testing.T) {
	sampleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Hello HTTP Test")
	})
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	_, err := fetchUserData(ts.URL)
	if err.Error() != "404 Not Found" {
		t.Error(err.Error())
	}
}

func TestRemoteServerDown(t *testing.T) {
	url := "http://aaiueo.aiadnuw.com/"
	_, err := fetchUserData(url)
	if err.Error() != "Get http://aaiueo.aiadnuw.com/: dial tcp: lookup aaiueo.aiadnuw.com: no such host" {
		t.Error(err.Error())
	}
}

func TestNewUserData(t *testing.T) {
	sampleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Hello HTTP Test")
	})
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	r := new(remoteUserDataServer)
	r.url = ts.URL
	_, err := newUserData(r)
	if err.Error() != "500 Internal Server Error" {
		t.Error(err.Error())
	}
}
