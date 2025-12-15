package main

import "fmt"


func square(arr *[3]int32) {
	fmt.Printf("Memory address of array in the function: %p.\n", arr);
	
	// modify the array
	for i := range arr {
		arr[i] *= arr[i];
	}
	
}
func main() {
	myArray := [...]int32{12, 4, 2};

	fmt.Printf("Memory address of array in the main function: %p.\n", &myArray);
	fmt.Printf("Array before: %v\n", myArray);
	
	square(&myArray);
	
	fmt.Printf("Array after: %v", myArray);
	

}