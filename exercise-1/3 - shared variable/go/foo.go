// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"time"
)

var i = 0

func incrementing(ch chan int) {
	//TODO: increment i 1000000 times
	// for x := 0; x < 1000000; x++ {
	// 	i++
	// }
	for x := 0; x < 1000000; x++ {
		i = <-ch
		i++
		ch <- i
	}

}

func decrementing(ch chan int) {
	//TODO: decrement i 1000000 times
	// for y := 0; y < 1000000; y++ {
	// 	i--
	// }
	for x := 0; x < 1000000; x++ {
		i = <-ch
		i--
		ch <- i
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	// GOMAXPROCS sets the maximum number of CPUs(threads) that can be executing simultaneously
	// There is a zero ouput when the number of CPUs are set to 1 because a single thread is used which operates on the go routines sequentially.
	//runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	ch1 := make(chan int, 1)
	ch1 <- i

	ch2 := make(chan int, 1)
	ch2 <- i
	go incrementing(ch1)
	go decrementing(ch2)

	select {
	case <-ch1:
		i = <-ch1
	case <-ch2:
		i = <-ch2
	}
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.

}
