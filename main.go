package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User is a struct
type User struct {
	ID   int
	Name string
}

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open error %s", err.Error())
	}

	db.AutoMigrate(&User{})

	// Create
	u := User{
		Name: "Will",
	}
	db.Create(&u)

	// Read
	db.Last(&u)
	fmt.Println(u)

	// Update
	u.Name = "Duotify"
	db.Save(&u)

	// Read
	db.Last(&u)
	fmt.Println(u)

	// Delete
	// db.Delete(u)
}
