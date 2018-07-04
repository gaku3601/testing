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

	r := new(remoteServer)
	r.url = ts.URL
	_, err := r.getBody(1)
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

	r := new(remoteServer)
	r.url = ts.URL
	_, err := r.getBody(1)
	if err.Error() != "404 Not Found" {
		t.Error(err.Error())
	}
}

func TestRemoteServerDown(t *testing.T) {
	url := "http://aaiueo.aiadnuw.com/"
	r := new(remoteServer)
	r.url = url
	_, err := r.getBody(1)
	if err.Error() != "Get http://aaiueo.aiadnuw.com/1: dial tcp: lookup aaiueo.aiadnuw.com: no such host" {
		t.Error(err.Error())
	}
}

func TestNewUserDataInternalServerError(t *testing.T) {
	sampleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Hello HTTP Test")
	})
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	r := new(remoteServer)
	r.url = ts.URL
	err := r.fetchUserAndFriendList()
	if err.Error() != "500 Internal Server Error" {
		t.Error(err.Error())
	}
}

func TestNewUserData(t *testing.T) {
	sampleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"id": 1,
			"name": "Austin",
			"friends": [
				2,
				5
			]
		}`)
	})
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	r := new(remoteServer)
	r.url = ts.URL
	r.fetchUserAndFriendList()
	for _, v := range r.ul {
		if v.ID != 1 {
			t.Errorf("%#v", *v)
		}
		if v.Name != "Austin" {
			t.Errorf("%#v", *v)
		}
	}
	if r.fl[0].From != 1 {
		t.Errorf("%#v", r.fl[0].From)
	}
	if r.fl[0].To != 2 {
		t.Errorf("%#v", r.fl[0].To)
	}
	if r.fl[1].To != 5 {
		t.Errorf("%#v", r.fl[1].To)
	}
}

func TestCreateURL(t *testing.T) {
	r := new(remoteServer)
	r.url = "http://test/"
	u := r.createUrl(3)
	if u != "http://test/3" {
		t.Error(u)
	}
}
