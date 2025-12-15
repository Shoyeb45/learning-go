package main

import "fmt"

func main()  {
	var arr [3]int32;
	fmt.Println("Default values of int32 arrays\n", arr); 

	arr[0] = 123;
	arr[1] = 2;
	arr[2] = 5;

	fmt.Println("Indexing at position 2(i.e, 1 in go): ", arr[1]);
	fmt.Println("Slicing: arr[1: 2]", arr[1: 2]);

	fmt.Println("Memory address: ");

	fmt.Println("Memory address of the first element: ", &arr[0]);
	fmt.Println("Memory address of the second element: ", &arr[1]);
	fmt.Println("Memory address of the third element: ", &arr[2]);


	// valid syntax
	intArr := [...]int32{1, 3, 4}; // compiler will automatically infer the length of the array.
	fmt.Println(intArr);
}