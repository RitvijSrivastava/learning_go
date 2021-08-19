package main

import (
	"fmt"
	"sync"
)

/**
- Channels are used to synchronize data transmission between multiple goroutines.
*/

var wg = sync.WaitGroup{}

func main() {
	// Create a channel.
	// Channel is strongly typed.
	// Here, we can send/receive messages of type "int"
	ch := make(chan int) // Unbuffered channel (one data at a time)

	wg.Add(2)
	go func() {
		// RECEIVING data from the channel.
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		// SENDING data to the channel.
		i := 42
		ch <- i // PASS BY VALUE
		i = 27
		wg.Done()
	}()
	wg.Wait()

	// BIDIRECTIONAL (SENDING + RECEIVING)
	wg.Add(2)
	go func() {
		// RECEIVING ONLY CHANNEL
		i := <-ch
		fmt.Println(i)
		ch <- 24
		wg.Done()
	}()
	go func() {
		// SENDING ONLY CHANNEL
		i := 42
		ch <- i
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()

	// ADD BIAS TO THE CHANNELS (Send only or Receive only)
	wg.Add(2)
	go func(ch <-chan int) {
		// RECEIVING ONLY CHANNEL
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		// SENDING ONLY CHANNEL
		i := 42
		ch <- i
		i = 27
		wg.Done()
	}(ch)
	wg.Wait()

	// BUFFERED CHANNELS
	// Used when the receiver and the sender operate at different frequencies, so,
	// Data won't be lost and the program woudln't shut down due to "no space" error (deadlock condition on sending)
	bufCh := make(chan int, 50) // Buffered channels (50 ints at a time)
	wg.Add(2)
	go func(bufCh <-chan int) {
		// Range over channel
		for i := range bufCh {
			fmt.Println(i)
		}

		// Alternative method
		for {
			if i, ok := <-bufCh; ok {
				fmt.Println(i)
			} else {
				//Channel is closed (ok = false)
				fmt.Println("Channel Closed")
				break
			}
		}
		wg.Done()
	}(bufCh)
	go func(bufCh chan<- int) {
		bufCh <- 50
		bufCh <- 70
		// Closing the channel, since if not done, receiver will deadlocked, since,
		// the receiver has no way of knowing how many data is coming through the channel.
		close(bufCh)

		wg.Done()
	}(bufCh)
	wg.Wait()
}
