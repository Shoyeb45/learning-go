package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5;

func main() {
	var chickenChannel = make(chan string);
	
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"};

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel);
	}
	sendMessage(chickenChannel);
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		fmt.Println("Fetching the price from the website.");
		time.Sleep(time.Second + 1);
		// fetch the price from the website (originally)
		var chickenPrice = rand.Float32() * 20;
		fmt.Printf("Fetched the price, price is : %v\n", chickenPrice);

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website;
			break;
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %v.", <- chickenChannel);
}