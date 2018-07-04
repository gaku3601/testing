package main

import (
	"log"
)

func main() {
	r, err := newRemoteServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	d := newDatabase()
	if err = d.storeDatabase(r.ul, r.fl); err != nil {
		log.Fatal(err.Error())
	}
}
