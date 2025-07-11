package producer

import (
	"fmt"
)

func Producer(ch chan int, numGenerated int) {
	for i := 1; i <= numGenerated; i++ {
		ch <- i
		fmt.Printf("Produced number: %d\n", i)
	}
	close(ch)
}
