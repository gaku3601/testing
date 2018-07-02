package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	r := newRemoteUserDataServer()
	ul, fl, err := newUserData(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = storeDatabase(ul, fl); err != nil {
		log.Fatal(err.Error())
	}
}
