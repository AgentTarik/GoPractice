package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchasableInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("%+v\n", contacts)

	var purchasable []purchasableInfo = loadJSON[purchasableInfo]("./purchasbaleInfo.json")
	fmt.Printf("%+v\n", purchasable)
}

func loadJSON[T contactInfo | purchasableInfo](filePath string) []T {
	data, _ := ioutil.ReadFile(filePath)
	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}
