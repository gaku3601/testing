package main

import (
	"github.com/jinzhu/gorm"
)

type database struct {
	db *gorm.DB
}

func newDatabase() *database {
	d := new(database)
	d.db, _ = gorm.Open("postgres", "user=postgres dbname=app sslmode=disable")
	return d
}

func (d *database) storeDatabase(ul []*User, fl []*Friend) error {
	defer d.db.Close()
	tx := d.db.Begin()
	if err := d.allDeleteTable(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := d.createData(ul, fl, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (d *database) allDeleteTable(tx *gorm.DB) error {
	if err := tx.Delete(&Friend{}).Error; err != nil {
		return err
	}
	if err := tx.Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (d *database) createData(ul []*User, fl []*Friend, tx *gorm.DB) error {
	for _, v := range ul {
		if err := tx.Create(v).Error; err != nil {
			return err
		}
	}
	for _, v := range fl {
		if err := tx.Create(v).Error; err != nil {
			return err
		}
	}
	return nil
}
