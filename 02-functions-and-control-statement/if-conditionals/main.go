package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	age := 1;
	if len(os.Args) < 2 {
		fmt.Println("No age passed.");
		return;
	}

	ageInStr := os.Args[1]
	age, err := strconv.Atoi(ageInStr);

	if err != nil {
		fmt.Println(err.Error());
	}

	if (age >= 50) {
		fmt.Println("Old");
	} else if (age >= 25) {
		fmt.Println("Middle aged");
	} else if (age >= 18) {
		fmt.Println("Adult");
	} else if (age >= 12) {
		fmt.Println("Teen");
	} else if (age >= 5) {
		fmt.Println("Child");
	} else if (age >= 1) {
		fmt.Println("Toddler");
	} else {
		fmt.Println("Infant");
	}
} 
