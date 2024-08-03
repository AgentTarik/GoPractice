package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var myMap map[string]uint32 = make(map[string]uint32)
	fmt.Println(myMap)

	var myMap2 map[string]uint32 = map[string]uint32{"Bruno": 18, "Jaum": 21}
	fmt.Println(myMap)
	fmt.Println(myMap2["Bruno"])
	myMap2["Pedro"] = 15
	delete(myMap2, "Jaum")
	var age, isOk = myMap2["Jaum"]
	if isOk {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	for key, value := range myMap2 {
		fmt.Printf("The value for %v is %v\n", key, value)
	}
}
