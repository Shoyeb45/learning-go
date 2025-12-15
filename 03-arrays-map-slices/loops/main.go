package main

import "fmt"

// there is no while loop in go

func main()  {
	// classic loop
	for i := 0; i < 12; i++ {
		fmt.Println(i);
	}

	// for loop converted into while loop
	i := 0;

	for i < 12 {
		// do something
		fmt.Println("Calling some function and doing work...");
		i++;
	}

	// without any arguments
	i = 0;
	for {
		if (i >= 10) {
			break;
		}
		fmt.Println("Doing heavy task, like calling api, db operation.");
		i += 2;
	}


	// iterating over arrays or slices
	myArr := []int32{10, 2, 40, -123};
	fmt.Println();
	for index, value := range myArr {
		fmt.Printf("Value at index %v is %v\n", index, value);
	}
	
	// iterating over map
	fmt.Println();

	myMap := map[int32]int32{ 12 : 34, 90 : -123, -123 : 23};

	for key, value := range myMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value);
	}
}