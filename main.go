package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	r := newRemoteUserDataServer()
	ul, fl, err := newUserData(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	db, _ := gorm.Open("postgres",
		"user=postgres dbname=app sslmode=disable")
	defer db.Close()
	for _, v := range ul {
		fmt.Printf("%#v", v)
		db.Create(v)
	}
	for _, v := range fl {
		fmt.Printf("%#v", v)
		db.Create(v)
	}
}
