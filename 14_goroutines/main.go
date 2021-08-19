package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
BEST PRACTICES:

- Don't create goroutines in libraries
- When creating a goroutine, know how it will end
	- Avoids memory leaks
- Check for race conditions at compile time. [  "go run -race src/main.go" ]
*/

var wg = sync.WaitGroup{} // Used to synchronize multiple goroutines.
var counter = 0
var m = sync.RWMutex{} // Only one can write at a time, multiple goroutines can read though, but writing cannot happen during reading.

func main() {

	// Set the number of threads that go can work with.
	// Default value is the number of cores available on the system.
	// Negative values do nothing.
	// Mostly used for tuning for performence.
	runtime.GOMAXPROCS(100)

	// Spawn a "green" thread.
	// "Green" thread is an abstraction of a thread, which is called goroutines.
	// Advantage is that these threads have lightweight abstraction over a thread. (can spawn thousands of go routine).
	go sayHello()

	msg := "Hello"
	go func() {
		// Msg is available inside this func, due a concept called "closures"
		fmt.Println("MSG: ", msg)
	}()
	// ! This will overwrite the previous [msg] inside the goroutine. (hence, why closures should be avoided -> Race conditions)
	// To Solve: pass arguments to the function.
	msg = "Bye"
	time.Sleep(100 * time.Millisecond) // Don't use this! Just for testing purposes.

	wg.Add(1) // Add number of goroutines
	go sayHello()
	wg.Wait() // Wait until all the goroutines are done executing (wg counter turns 0)

	for i := 0; i < 10; i++ {
		// This block will run BUT the order will be arbitrary. (Race conditions)
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		// This block will print in order BUT the functions are still not in sync, as in,
		// all the values might be printed (like, "Hello: 1, then 8 times "Hello: 2" ")
		wg.Add(2)
		go sayHello2()
		go increment2()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		// This block will print all he values in order and in sync.
		// Since, we used locking before the actual each goroutine executes.

		// The problem with this approach is that, this is basically syncronous problem.
		// No goroutines are not needed here.
		wg.Add(2)
		m.RLock()
		go sayHello2()
		m.Lock()
		go increment2()
	}
	wg.Wait()

}

func sayHello() {
	fmt.Println("Hello: ", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done() // Decrement the goroutine count
}

func sayHello2() {
	m.RLock() // Read lock
	fmt.Println("Hello: ", counter)
	m.RUnlock()
	wg.Done()
}

func increment2() {
	m.Lock() // lock for writing
	counter++
	m.Unlock()
	wg.Done()
}

func sayHello3() {
	fmt.Println("Hello: ", counter)
	m.RUnlock()
	wg.Done()
}

func increment3() {
	counter++
	m.Unlock()
	wg.Done()
}
