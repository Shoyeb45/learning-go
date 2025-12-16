package main

import "fmt"

// channels are:
// 1. Holds the data
// 2. Thread safe
// 3. Listen for data

func main()  {
	
	// creating the channel
	var intChan = make(chan int32);

	// we can add the value like:
	// intChan <- 123;
	// when we add the value, the channel will wait for some other routine to read from it.
	// so following code will panic and will give deadlock error
	// var i = <- intChan;

	// fmt.Println(i);

	go process(intChan, 123);

	// it will wait for process to finish adding the value in the channel

	for i := range intChan {
		fmt.Println(i);
	}
}


func process(c chan int32, val int32) {
	fmt.Printf("Adding %v in the intChan\n", val);
	for i := 1; i <= 5; i++ {
		c <- val;
	}
	close(c);
}