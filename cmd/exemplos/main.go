package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	qtdWorkers := 10
	for i := range qtdWorkers {
		go worker(i, ch)
	}

	for i := range 15 {
		ch <- i
	}
}
