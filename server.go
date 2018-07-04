package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"sync"

	"github.com/tidwall/gjson"
)

type remoteServer struct {
	url string    // remote server url
	ul  []*User   // UserList
	fl  []*Friend // FriendList
}

func newRemoteServer() (*remoteServer, error) {
	r := new(remoteServer)
	r.url = "http://fg-69c8cbcd.herokuapp.com/user/"
	err := r.fetchUserAndFriendList()
	if err != nil {
		return nil, err
	}
	// sorting
	r.sortUserList()
	r.sortFriendList()
	return r, err
}

func (r *remoteServer) fetchUserAndFriendList() error {
	// multiple process
	errChan := make(chan error)
	wg := new(sync.WaitGroup)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			r.storeUserStructAndFriendStruct(i, errChan)
		}(i)
	}
	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *remoteServer) storeUserStructAndFriendStruct(i int, errChan chan<- error) {
	s, err := r.getBody(i)
	if err != nil {
		errChan <- err
		return
	}
	for _, v := range gjson.Get(s, "friends").Array() {
		r.fl = append(r.fl, &Friend{
			From: gjson.Get(s, "id").Int(),
			To:   v.Int(),
		})
	}

	r.ul = append(r.ul, &User{
		ID:   gjson.Get(s, "id").Int(),
		Name: gjson.Get(s, "name").String(),
	})
}

func (r *remoteServer) createUrl(id int) string {
	u, _ := url.Parse(r.url)
	u.Path = path.Join(u.Path, strconv.Itoa(id))
	return u.String()
}

func (r *remoteServer) getBody(i int) (string, error) {
	resp, err := http.Get(r.createUrl(i))
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

func (r *remoteServer) sortUserList() {
	sort.Slice(r.ul, func(i, j int) bool { return r.ul[i].ID < r.ul[j].ID })
}
func (r *remoteServer) sortFriendList() {
	sort.Slice(r.fl, func(i, j int) bool {
		if r.fl[i].From < r.fl[j].From {
			return true
		} else if r.fl[i].From < r.fl[j].From {
			if r.fl[i].To < r.fl[j].To {
				return true
			}
		}
		return false
	})
}
