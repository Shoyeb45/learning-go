package main

import "fmt"

func main() {
	// we can create map using make function
	myMap := make(map[string]int8); // key is string, and value is int 8 bytes

	myMap["shoyeb"] = 21;

	fmt.Println(myMap);

	anotherMap := map[string]string{"shoyeb": "programmer"};

	val, ok := anotherMap["bob"];
	if !ok {
		fmt.Println("No 'bob' key is present in map.");
	} else {
		fmt.Println(val);
	}

	val, ok = anotherMap["bob"];
	
	fmt.Println(ok); // it does not insert the key
	fmt.Println(val);
}