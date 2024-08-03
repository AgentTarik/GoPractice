package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sumSlices[int](intSlice))

	var float32Slice = []float32{1, 2, 3, 4, 5}
	fmt.Println(sumSlices[float32](float32Slice))
}

func sumSlices[T int | float32 | float64](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}
