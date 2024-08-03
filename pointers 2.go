package main

import "fmt"

func main() {
	fmt.Println()

	var list = []int32{1, 2, 3}
	var listCopy = list
	listCopy[1] = 6
	fmt.Println(list)
	fmt.Println(listCopy)
}
