package main

import (
	"github.com/jinzhu/gorm"
)

func storeDatabase(ul []*User, fl []*Friend) error {
	db, _ := gorm.Open("postgres",
		"user=postgres dbname=app sslmode=disable")
	defer db.Close()
	tx := db.Begin()
	if err := allDeleteTable(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := createData(ul, fl, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func allDeleteTable(tx *gorm.DB) error {
	if err := tx.Delete(&Friend{}).Error; err != nil {
		return err
	}
	if err := tx.Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func createData(ul []*User, fl []*Friend, tx *gorm.DB) error {
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
