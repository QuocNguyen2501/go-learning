package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name     string
	Age      int64
	Birthday time.Time
}

func find(db *gorm.DB, name string) User {
	var user User
	db.Where("name = ? ", name).Find(&user)
	return user
}

func main() {
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=postgres dbname=demo password=postgres sslmode=disable")
	db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	user := User{}
	db.Model(&user).Update("name", "John")
	defer db.Close()
	fmt.Println("Successful, shutdown program ...")
}
