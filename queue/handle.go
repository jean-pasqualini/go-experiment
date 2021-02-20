package queue

import (
	"fmt"
	"time"
)

func Handle() {
	numbers := make(chan int, 100)
	waitEnd := make(chan bool, 1)

	go worker(numbers, waitEnd)

	for i := 1; i <= 200; i++ {
		fmt.Printf("sending %d\n", i)
		numbers <- i
	}

	close(numbers)
	<-waitEnd
}

func worker(ch <-chan int, notifyEnd chan<- bool ) {
	for value := range ch {
		fmt.Sprintf("receive %d\n", value)
		time.Sleep(time.Millisecond * 100)
	}
	notifyEnd <- true
}
