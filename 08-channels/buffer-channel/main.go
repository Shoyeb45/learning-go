package main

import "time"
import "fmt"


func main()  {
	// this will create buffer channel with size
	var c = make(chan int32, 5);

	go process(c);

	for i := range c {
		fmt.Println(i);
		time.Sleep(time.Second + 1);
	}
}

func process(c chan int32) {
	// after defer keyword, whatever is there, execute it before exiting the function
	defer close(c); 

	for i := int32(1); i <= 6; i++ {
		c <- i;
	}
}