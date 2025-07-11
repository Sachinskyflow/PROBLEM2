package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, numGenerated int) {
	for i := 1; i <= numGenerated; i++ {
		ch <- i
		fmt.Printf("Produced number: %d\n", i)
	}
	close(ch)
}

func consumer(i int, ch chan int) {
	for number := range ch {
		ans := processItem(number)
		fmt.Printf("Consumer %d consumed number %d, processed result: %d\n", i, number, ans)
	}
}

func processItem(number int) int {
	return number * number
}

func main() {
	var (
		numConsumer  int
		numBuffChan  int
		numGenerated int
	)
	fmt.Print("Enter number of consumers: ")
	fmt.Scan(&numConsumer)
	fmt.Print("Enter number of buffered channels: ")
	fmt.Scan(&numBuffChan)
	fmt.Print("Enter number of generated integers: ")
	fmt.Scan(&numGenerated)
	if numConsumer <= 0 || numBuffChan <= 0 || numGenerated <= 0 {
		fmt.Println("All numbers must be greater than zero.")
		return
	}
	ch := make(chan int, numBuffChan)
	var wg sync.WaitGroup
	// producer(ch, numGenerated)
	// go func() {
	// 	producer(ch, numGenerated)
	// }()
	for i := 0; i < numConsumer; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumer(i+1, ch)
		}()
	}
	// producer(ch, numGenerated)
	go func() {
		producer(ch, numGenerated)
	}()
	wg.Wait()
}
