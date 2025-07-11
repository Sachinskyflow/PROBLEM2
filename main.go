package main

import (
	"fmt"
	"problem2/consumer"
	"problem2/producer"
	"sync"
)

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
	for i := 0; i < numConsumer; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			consumer.Consumer(i, ch)
		}(i + 1)
	}
	go func() {
		producer.Producer(ch, numGenerated)
	}()
	wg.Wait()
}
