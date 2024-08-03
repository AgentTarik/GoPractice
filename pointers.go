package main

import "fmt"

func main() {
	fmt.Println()

	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value that p points to is %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)

	p = &i
	*p = 10

	fmt.Printf("The value that p points to is %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)

	var k int32 = 2
	i = k

	fmt.Printf("The value that p points to is %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)

}
