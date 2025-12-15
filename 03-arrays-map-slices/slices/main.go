package main

import "fmt"

func main() {
	var intSlice []int32 = []int32{12, 32, 12};
	fmt.Printf("Length of the slice: %v, capacity of the slice: %v\n", len(intSlice), cap(intSlice));
	
	intSlice = append(intSlice, 123);
	fmt.Println(intSlice);
	
	fmt.Printf("Length of the slice: %v, capacity of the slice: %v\n", len(intSlice), cap(intSlice));

	// we can also append like this:
	intSlice = append(intSlice, []int32{-12, 23}...);  // spread operator
	fmt.Println(intSlice);

	// another way to make slice:
	intSlice2 := make([]float64, 4, 10); // type, length, capcity
	fmt.Println(intSlice2);
}