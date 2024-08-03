package main

import "fmt"

func main() {

	var intSlices []int32 = []int32{4, 5, 6}
	fmt.Println(intSlices)
	fmt.Println(cap(intSlices))
	position := &intSlices[0]
	fmt.Println(position)

	fmt.Println()

	intSlices = append(intSlices, 7)
	fmt.Println(intSlices)
	fmt.Println(cap(intSlices))
	position = &intSlices[0]
	fmt.Println(position)
}
