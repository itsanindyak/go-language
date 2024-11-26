package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannels := make(chan int, 3)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// go routine for only write the value
	go func(c chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		c <- 5
		c <- 6

		close(c) // close the channel after putting all value

	}(myChannels, wg)
	// go routine for only read the value
	go func(c chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		for value := range c {
			fmt.Println(value)
		}

	}(myChannels, wg)

	wg.Wait()

}
