package main

import (
	"fmt"
	"sync"
)

// The channel is a write-only type with the <- arrow pointing into the channel.
func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("sending %d to channel\n", idx)
		ch <- idx
	}
}

// PrintNumbers only needs to be able to read numbers from the channel,
// so it’s a read-only type with the <- arrow pointing away from the channel.
func printNumbers(idx int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("%d: read %d from channel\n", idx, num)
	}
}

func main() {
	var wg sync.WaitGroup //Declaring WaitGroup
	numberChan := make(chan int)

	for idx := 1; idx <= 3; idx++ {
		wg.Add(1) //How many things to wait for, before starting the goroutines,
		// wait for two Done calls before considering the group finished.
		go printNumbers(idx, numberChan, &wg)
	}

	generateNumbers(5, numberChan, &wg)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
