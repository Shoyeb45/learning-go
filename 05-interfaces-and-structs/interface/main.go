package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
	return  e.gallons * e.mpg;
}
func (e electricEngine) milesLeft() uint8 {
	return  e.mpkwh * e.kwh;
}


type engine interface {
	milesLeft() uint8
} 

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Yes you can make it.")
	} else {
		fmt.Println("Need to fuel up first.");
	}
}

func main()  {
	var myEngine gasEngine = gasEngine{25, 15};
	canMakeIt(myEngine, 123);
	
	electEngine := electricEngine{212, 3};
	canMakeIt(electEngine, 23);
}	