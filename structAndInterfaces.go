package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type eletricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e eletricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canIMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You can't make it")
	}
}

func main() {
	myEngine := gasEngine{25, 15}
	myEletricEngine := eletricEngine{17, 210}
	canIMakeIt(myEngine, 100)
	canIMakeIt(myEletricEngine, 130)

}
