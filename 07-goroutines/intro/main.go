package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{};

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7"};



func main() {
	t0 := time.Now();
	
	for i := 0; i < len(dbData); i++ {
		// now for making a function go routine, we can use `go` keyword, 
		// it will run the function in background, but it will exit the program
		// to wait for it, we can create the wait groups
		wg.Add(1); // for spawning the go routine
		go dbCall(i);
	}
	wg.Wait(); // wait for the counter to get back down to 0
	fmt.Printf("\nTotal Time of executation is: %v", time.Since(t0));
}

func dbCall(i int) {
	var delay float32 = 2000;

	time.Sleep(time.Duration(delay) * time.Millisecond);
	fmt.Println("The result from the databse is: ", dbData[i]);
	
	wg.Done(); // specify that go routine is done
}
