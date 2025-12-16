package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{};
// var m = sync.Mutex{}; 
var m = sync.RWMutex{}; // read write mutex

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7"};

var results = []string{};


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
	fmt.Printf("\nFinal results is: %v", results);
}

func dbCall(i int) {
	var delay float32 = 2000;

	time.Sleep(time.Duration(delay) * time.Millisecond);
	fmt.Println("The result from the databse is: ", dbData[i]);
	
	save(dbData[i]);
	log();
	wg.Done(); // specify that go routine is done
}


func save(res string) {
	// writing to the memory at the same time with multiple threads is not good 
	// and can lead to corrupt memory issues, so we'll use mutexes.
	m.Lock();
	results = append(results, res);
	m.Unlock();
}

func log() {
	m.RLock();
	fmt.Printf("\nThe current result are: %v\n", results);
	m.RUnlock();
}


// The go routines will use cpu cores to maximize the performance.
// but if our all processor are busy then the go routine will wait for processor to be free.


