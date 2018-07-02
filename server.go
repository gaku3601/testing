package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/tidwall/gjson"
)

type remoteUserDataServer struct {
	url string
}

func newRemoteUserDataServer() *remoteUserDataServer {
	r := new(remoteUserDataServer)
	r.url = "http://fg-69c8cbcd.herokuapp.com/user/"
	return r
}

func newUserData(r *remoteUserDataServer) ([]*User, []*Friend, error) {
	var ul []*User
	var fl []*Friend
	for i := 1; i <= 10; i++ {
		s, err := fetchUserData(r.createUrl(i))
		if err != nil {
			return nil, nil, err
		}
		for _, v := range gjson.Get(s, "friends").Array() {
			fl = append(fl, &Friend{
				From: gjson.Get(s, "id").Int(),
				To:   v.Int(),
			})
		}

		ul = append(ul, &User{
			ID:   gjson.Get(s, "id").Int(),
			Name: gjson.Get(s, "name").String(),
		})
	}
	return ul, fl, nil
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
