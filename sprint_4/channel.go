package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64

func thread(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&counter, 1)
	}
	wg.Done()
}

type Kek struct {
	ch chan int
}

func main() {
	kek := &Kek{make(chan int, 10)}
	readCh(kek)
	fmt.Println("1")
	kek.ch <- 1
	fmt.Println("2")
	kek.ch <- 2
	fmt.Println("3")
	kek.ch <- 3
	fmt.Println("3")
	kek.ch <- 3
	fmt.Println("3")
	kek.ch <- 3
	fmt.Println("3")
	kek.ch <- 3
	fmt.Println("3")
	kek.ch <- 3
	fmt.Println("30")
	kek.ch <- 30
	time.Sleep(time.Second * 5)
	fmt.Println("Finish")
}

func readCh(kek *Kek) {
	go func(kek *Kek) {
		for {
			i := <-kek.ch
			fmt.Printf("Number %d\n", i)
			time.Sleep(time.Second)
		}
	}(kek)
}
