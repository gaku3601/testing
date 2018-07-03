package main

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func newTestDB() *database {
	d := new(database)
	d.db, _ = gorm.Open("postgres", "user=postgres dbname=app sslmode=disable")
	return d
}

func TestAllDeleteTable(t *testing.T) {
	d := newTestDB()
	// Initialization
	d.db.Delete(&User{})
	d.db.Delete(&Friend{})
	d.db.Create(&User{ID: 1, Name: "gaku"})
	d.db.Create(&Friend{From: 1, To: 2})
	// process
	d.allDeleteTable(d.db)

	// result
	count := 0
	d.db.Table("user").Count(&count)
	if count != 0 {
		t.Error("AllDeleteTable Function: delete error")
	}
	count = 0
	d.db.Table("friend").Count(&count)
	if count != 0 {
		t.Error("AllDeleteTable Function: delete error")
	}
}

func TestCreateData(t *testing.T) {
	d := newTestDB()
	// Initialization
	d.db.Delete(&User{})
	d.db.Delete(&Friend{})
	// process
	ul := []*User{
		&User{ID: 1, Name: "gaku1"},
		&User{ID: 2, Name: "gaku2"},
	}
	fl := []*Friend{
		&Friend{From: 1, To: 2},
		&Friend{From: 2, To: 2},
		&Friend{From: 2, To: 1},
	}
	d.createData(ul, fl, d.db)

	// result
	var u User
	d.db.First(&u, 2)
	if u.ID != 2 || u.Name != "gaku2" {
		t.Error("createData Function Error: store user data")
	}
	var f Friend
	d.db.First(&f)
	if f.From != 1 || f.To != 2 {
		t.Error("createData Function Error: store friend data")
	}
}
