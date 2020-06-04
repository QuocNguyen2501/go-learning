package main

import (
	"fmt"
)

type Customer struct {
	Id        int
	Name      string
	Genre     string
	Address   string
	IsMarried bool
}

func main() {
	customer := Customer{
		Id:        1,
		Name:      "Peter",
		Genre:     "Male",
		Address:   "Ajeltake Road, Ajeltake Island, Majuro, Marshall Islands",
		IsMarried: true,
	}

	fmt.Println(customer)
	fmt.Printf("%T\n", customer)
	fmt.Println("Id: ", customer.Id)
	fmt.Println("Name: ", customer.Name)
	fmt.Println("Genre: ", customer.Genre)
	fmt.Println("Address: ", customer.Address)
	fmt.Println("IsMarried: ", customer.IsMarried)
}
