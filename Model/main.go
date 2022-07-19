package main

import "sync"

var wg sync.WaitGroup

func Consumer(ch <-chan int) []int {
	var num []int
	println("Receiving")
	for i := range ch {
		println(i)
		num = append(num, i)
	}
	return num
}

func Producer() <-chan int {
	ch := make(chan int, 255)
	wg.Add(1)
	go func() {
		println("Producing")
		for i := 1; i <= 10; i += 2 {
			ch <- i
			println(i)
		}
		close(ch)
		wg.Done()
	}()
	wg.Wait()
	return ch
}

func main() {
	ch1 := Producer()

	var num []int

	num = Consumer(ch1)

	println("Test")
	for i := range num {
		println(num[i])
	}
}
