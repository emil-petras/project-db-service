package models

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	target := os.Getenv("DB_TARGET")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, password, target, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = autoMigrate(db)
	if err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}

	DB = db

	return nil
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("user auto migrate failed: %w", err)
	}

	err = db.AutoMigrate(&Transaction{})
	if err != nil {
		return fmt.Errorf("transaction auto migrate failed: %w", err)
	}

	err = db.AutoMigrate(&Token{})
	if err != nil {
		return fmt.Errorf("token auto migrate failed: %w", err)
	}

	// add one user
	user := User{
		Username: "mark",
		Password: "9b8769a4a742959a2d0298c36fb70623f2dfacda8436237df08d8dfd5b37374c", // hashed "pass123"
		Balance:  0,
		Updated:  time.Now().Local(),
		Created:  time.Now().Local(),
	}
	_ = db.FirstOrCreate(&user)

	return nil
}
