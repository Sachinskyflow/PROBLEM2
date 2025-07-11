package consumer

import (
	"fmt"
)

func Consumer(i int, ch chan int) {
	for number := range ch {
		ans := processItem(number)
		fmt.Printf("Consumer %d consumed number %d, processed result: %d\n", i, number, ans)
	}
}

func processItem(number int) int {
	return number * number
}
