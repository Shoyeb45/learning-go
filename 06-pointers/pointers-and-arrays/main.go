package main

import "fmt"

func main() {
	myArray := []int32{12, 4, 4};

	// here only pointer is assigned to newArray, entire array is not copied
	// so both myArray and newArray points to the same array in memory
	newArray := myArray; 

	newArray[0] = 123;
	fmt.Printf("New Array: %v\n", newArray);
	fmt.Printf("My Array: %v\n", myArray);

}