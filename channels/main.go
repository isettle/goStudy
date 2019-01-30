package main

import (
	"fmt"
	_ "runtime"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, done func()) worker {
	w := worker{
		in:   make(chan int, 0),
		done: done,
	}

	go doWork(id, w.in, w.done)
	return w
}

func doWork(id int, c <-chan int, done func()) {
	for n := range c {
		fmt.Printf("worker: %d, data: %c\n", id, n)
		done()
	}
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, func() {
			wg.Done()
		})
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
