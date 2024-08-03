package main

import "fmt"

func main() {
	var thing1 [5]float64 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memeory location of thing1 array is: %p", &thing1)
	square(&thing1)
	fmt.Printf("\nThe thing1 is: %v", thing1)
}

func square(thing2 *[5]float64) {
	fmt.Printf("\nThe memeory location of thing2 array is: %p", thing2)

	fmt.Printf("\nThe thing2 is of type: %T", thing2)
	fmt.Printf("\nThe thing2 is: %v", *thing2)

	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
}
