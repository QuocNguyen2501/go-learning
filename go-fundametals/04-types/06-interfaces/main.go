package main

import "fmt"

type user struct {
	Id        int
	Firstname string
	Lastname  string
	Email     string
}

type IUserValidation interface {
	validateInfo() bool
	validateId() bool
}

func (u user) validateInfo() bool {
	if u.Firstname == "" {
		return false
	}
	if u.Lastname == "" {
		return false
	}
	if u.Email == "" {
		return false
	}
	return true
}

func (u user) validateId() bool {
	if u.Id == 0 {
		return false
	}
	return true
}

func InputUser(userValidation IUserValidation) {
	isValidInfo := userValidation.validateInfo()
	fmt.Println("Validate user info: ", isValidInfo)
	isValidID := userValidation.validateId()
	fmt.Println("Validate user info: ", isValidID)
}

func main() {
	u := user{
		Id:        0,
		Firstname: "Peter",
		Lastname:  "Nguyen",
		Email:     "peternguyen@gmail.com",
	}
	InputUser(u)
}
