package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chikenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chikenChannel)
		go checkTofuPrices(websites[i], tofuChannel)

	}
	sendMessage(chikenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrices = rand.Float32() * 20
		if chickenPrices <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}

	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrices = rand.Float32() * 20
		if chickenPrices <= MAX_CHICKEN_PRICE {
			tofuChannel <- website
			break
		}

	}
}

func sendMessage(chikenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chikenChannel:
		fmt.Printf("\nText sent: Found deal on chiken at %v.", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail sent: Found deal on tofu at %v.", website)
	}
}
