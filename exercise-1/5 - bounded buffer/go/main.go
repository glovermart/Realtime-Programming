package main

import (
	"fmt"
	"time"
)

func producer(c chan int /*TODO: parameters?*/) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		// TODO: push real value to buffer
		c <- i
	}

}

func consumer(c chan int /*TODO: parameters?*/) {

	time.Sleep(1 * time.Second)
	for {
		i := <-c //TODO: get real value from buffer
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	ch := make(chan int, 5)

	go consumer(ch)
	go producer(ch)

	select {}
}
