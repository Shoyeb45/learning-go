package main

import (
	"errors"
	"fmt" 
)


func printMe(value string) {
	fmt.Println(value);
}


func main() {
	value := "Hello, world";

	printMe(value);

	val, reminder, err := divideInt(12, 0);

	if err != nil {
		fmt.Println("Error detected, error message: " + err.Error())
	}
	fmt.Printf("The reminder is %v and the division is %v", val, reminder);
}


func divideInt(numerator int, denominator int) (int, int, error) {
	var err error;

	if denominator == 0 {
		err = errors.New("can not divide by zero");
		return 0, 0, err;
	}

	val := numerator / denominator;
	reminder := numerator % denominator;

	return  val, reminder, err;
}