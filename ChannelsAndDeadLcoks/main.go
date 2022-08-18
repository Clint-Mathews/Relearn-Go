package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and Deadlocks")
	// make(chan int) - Will case error if we pass multiple chan  data
	myChannel := make(chan int, 2) // To fix that we set a buffer ie now its a buffer channel
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)
	// Recive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myChannel) Causes an error since the close caused the channel to be not used
		val, isChannelOpen := <-myChannel
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myChannel)
		defer wg.Done()

	}(myChannel, wg)

	// Send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		close(myChannel)

		// myChannel <- 10

		defer wg.Done()

	}(myChannel, wg)

	wg.Wait()

}
