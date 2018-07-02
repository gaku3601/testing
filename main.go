package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

func main() {
	r := newRemoteUserDataServer()
	ul, err := newUserData(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range *ul {
		fmt.Println(*v)
	}
}

type remoteUserDataServer struct {
	url string
}

type userData struct {
	data string
}
type userList []*userData

func newRemoteUserDataServer() *remoteUserDataServer {
	r := new(remoteUserDataServer)
	r.url = "http://fg-69c8cbcd.herokuapp.com/user/"
	return r
}

func newUserData(r *remoteUserDataServer) (*userList, error) {
	var ul userList
	for i := 1; i <= 10; i++ {
		s, err := fetchUserData(r.createUrl(i))
		if err != nil {
			return nil, err
		}
		ul = append(ul, &userData{data: s})
	}
	return &ul, nil
}

func (r *remoteUserDataServer) createUrl(id int) string {
	u, _ := url.Parse(r.url)
	u.Path = path.Join(u.Path, strconv.Itoa(id))
	return u.String()
}

func fetchUserData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return "", errors.New(resp.Status)
	}
	if resp.StatusCode == 500 {
		return "", errors.New(resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
