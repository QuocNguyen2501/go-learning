package main

import "fmt"

func main() {
	dict := make(map[string]string)

	dict["Key1"] = "Value1"
	dict["Key2"] = "Value2"
	dict["Key3"] = "Value3"
	dict["Key4"] = "Value4"

	for i, v := range dict {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}

	fmt.Println("==================================")
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	for i, v := range dict2 {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}

	fmt.Println("==================================")
	// Delete an item in map
	delete(dict2, "Red")
	for i, v := range dict2 {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
}
