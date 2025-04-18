package gchan

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestBatchProcessor(t *testing.T) {
	input := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second * 2)
			input <- i
		}

		time.Sleep(time.Second * 5)

		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second * 5)
			input <- i
		}
		close(input)
	}()

	for batch := range BatchProcessor(input, 3) {
		fmt.Println("BatchProcessor:", batch)
	}
}

func TestBatchProcessorWithTimeout(t *testing.T) {
	input := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}

		time.Sleep(time.Second * 5)

		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second * time.Duration(i))
			log.Println(i)
			input <- i
		}
		close(input)
	}()

	for batch := range BatchProcessorWithTimeout(input, 3, time.Second*3) {
		log.Println("BatchProcessorWithTimeout:", batch)
	}
}
