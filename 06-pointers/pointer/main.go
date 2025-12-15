package main

import "fmt"

func main()  {
	// var p *int32;  // p is pointer to the int32 variable
	// default it will store nil
	var p *int32 = new(int32); // it will point to some memory location of 8 bytes long

	fmt.Printf("The value at p points to the %v.\n", *p);

	// we can change the value which pointer points
	*p = 123;
	fmt.Printf("*p = %v\n", *p);

	newVariable := int32(150);
	
	fmt.Printf("Before newVariable: %v\n",newVariable);

	// we can change the pointer 
	p = &newVariable;

	// now if we change the value of the p, then newVariable should change
	
	*p = 450;

	fmt.Printf("After newVariable = %v, *p = %v", newVariable, *p);
}