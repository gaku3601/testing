package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	r, err := newRemoteServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = storeDatabase(r.ul, r.fl); err != nil {
		log.Fatal(err.Error())
	}
}
