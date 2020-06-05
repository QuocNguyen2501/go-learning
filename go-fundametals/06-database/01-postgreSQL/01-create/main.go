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

var users = []User{
	{
		Name:     "Smith",
		Age:      10,
		Birthday: time.Now(),
	},
	{
		Name:     "Peter",
		Age:      20,
		Birthday: time.Now(),
	},
	{
		Name:     "Mary",
		Age:      30,
		Birthday: time.Now(),
	},
	{
		Name:     "Job",
		Age:      40,
		Birthday: time.Now(),
	},
	{
		Name:     "Will",
		Age:      40,
		Birthday: time.Now(),
	},
}

func create(db *gorm.DB, user User) {
	db.NewRecord(user)
	db.Create(&user)
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

	for _, v := range users {
		create(db, v)
	}

	var usrs []User
	db.Find(&usrs)
	for _, v := range usrs {
		fmt.Println(v)
	}
	defer db.Close()
	fmt.Println("Successful, shutdown program ...")
}
