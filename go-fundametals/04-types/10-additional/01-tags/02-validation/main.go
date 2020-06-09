package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	FirstName string `validation:"required,max:50"`
	LastName  string `validation:"required,max:50"`
	Email     string `validation:"required,max:50"`
}

func (u *User) validate() bool {
	val := reflect.ValueOf(u).Elem()
	for i := 0; i < val.NumField(); i++ {
		req, ok := val.Type().Field(i).Tag.Lookup("validation")
		if ok {
			valis := strings.Split(req, ",")
			if valis[0] == "required" && len(val.FieldByName(val.Type().Field(i).Name).String()) == 0 {
				return false
			}

			if valMaxs := strings.Split(valis[1], ":"); len(valMaxs[0]) > 0 {
				maxLength, _ := strconv.Atoi(valMaxs[1])
				if valMaxs[0] == "max" && maxLength <= len(val.FieldByName(val.Type().Field(i).Name).String()) {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	u := User{
		FirstName: "Will",
		LastName:  "Nguyen",
		Email:     "will.nguyen@gmail.com",
	}
	fmt.Println("User is valid: ", u.validate())
}
