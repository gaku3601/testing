package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"sync"

	"github.com/tidwall/gjson"
)

type remoteServer struct {
	url string
}

func newRemoteServer() *remoteServer {
	r := new(remoteServer)
	r.url = "http://fg-69c8cbcd.herokuapp.com/user/"
	return r
}

type dataForMultipleProcess struct {
	err error
	ul  []*User
	fl  []*Friend
}

func newUserAndFriendList(r *remoteServer) ([]*User, []*Friend, error) {
	var ul []*User
	var fl []*Friend

	// multiple process
	dataChan := make(chan dataForMultipleProcess)
	wg := new(sync.WaitGroup)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			storeUserStructAndFriendStruct(i, r, dataChan)
		}(i)
	}
	go func() {
		wg.Wait()
		close(dataChan)
	}()
	for v := range dataChan {
		ul = append(ul, v.ul...)
		fl = append(fl, v.fl...)
		if v.err != nil {
			return nil, nil, v.err
		}
	}
	return ul, fl, nil
}

func storeUserStructAndFriendStruct(i int, r *remoteServer, dataChan chan<- dataForMultipleProcess) {
	var ul []*User
	var fl []*Friend

	s, err := fetchUserData(r.createUrl(i))
	if err != nil {
		dataChan <- dataForMultipleProcess{err: err, ul: nil, fl: nil}
		return
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
	dataChan <- dataForMultipleProcess{err: nil, ul: ul, fl: fl}
}

func (r *remoteServer) createUrl(id int) string {
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
