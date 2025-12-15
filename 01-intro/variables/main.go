package main

import "fmt"

func main() {
	// defining variable using var keyword
	
	// int data types: unsigned and signed:
	// (u)int - {8, 16, 32, 64}

	var myInt int32 = 123;


	// Real numbers - float32 and float64
	var myFloat float64 = 3.1415124154123123;

	// var newVal = myInt + myFloat;  not allowed

	var newVal = myInt + int32(myFloat);

	fmt.Println(newVal);

	// divide and modulo
	var int1 int32 = 123;
	var int2 int32 = 4;
	fmt.Println(int1 / int2);
	fmt.Println(int1 % int2);


	// defining values
	var thisIsint int64 = (1 << 55);
	// we can omit the type 
	var omittedType = 1 << 55;

	// we can even omit the var
	omittedVarAndType := 1 << 55;

	// let's see if this is giving the same value
	fmt.Println(thisIsint, omittedType, omittedVarAndType);

	// boolean data types
	maybeBoolean := thisIsint == int64(omittedType);

	fmt.Println(maybeBoolean); // true
}
