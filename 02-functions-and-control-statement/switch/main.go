package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No month provided.")
		os.Exit(1)
	}

	month := strings.ToLower(os.Args[1])

	switch month {
	case "january":
		fmt.Println(1)
	case "february":
		fmt.Println(2)
	case "march":
		fmt.Println(3)
	case "april":
		fmt.Println(4)
	case "may":
		fmt.Println(5)
	case "june":
		fmt.Println(6)
	case "july":
		fmt.Println(7)
	case "august":
		fmt.Println(8)
	case "september":
		fmt.Println(9)
	case "october":
		fmt.Println(10)
	case "november":
		fmt.Println(11)
	case "december":
		fmt.Println(12)
	default:
		fmt.Println("No valid month.")
	}

}
