package main

import "fmt"

func main() {
	var intArray [3]int32
	intArray[0] = 10
	intArray[1] = 20
	intArray[2] = 30
	var firstPosition *int32 = &intArray[0]
	fmt.Println(firstPosition)
	fmt.Println(*firstPosition)

	secondPosition := &intArray[1]
	fmt.Println(*secondPosition)

	stringArray := [...]string{"bom", "dia", "!"}
	for i := 0; i < len(stringArray); i++ {
		printer(stringArray[i])
	}

}

func printer(text string) {
	fmt.Println(text)
}
