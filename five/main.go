package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		// write to a channel
		c <- "The Gadget Show"
	}()

	// read from a channel
	val := <-c

	//Print Out Result
	fmt.Printf("Watching %q\n", val)

}
