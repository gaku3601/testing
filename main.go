package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	r := newRemoteServer()
	ul, fl, err := newUserAndFriendList(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = storeDatabase(ul, fl); err != nil {
		log.Fatal(err.Error())
	}
}
