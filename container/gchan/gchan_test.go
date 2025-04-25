package gchan

import (
	"context"
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

func TestUnboundedChan(t *testing.T) {
	ctx := context.Background()
	ch := NewUnboundedChan[int](ctx, 1000)
	// or ch := NewUnboundedChanSize(10,200,1000)

	go func() {
		for i := 0; i < 100; i++ {
			ch.In <- i
			if i == 50 {
				break
			}
		}
		fmt.Println("Out 0 close(ch.In)")
		close(ch.In) // close In channel
	}()
	fmt.Println("Out 1")
	for v := range ch.Out { // read values
		fmt.Println(v)
	}
	fmt.Println("Out 2")

	time.Sleep(time.Second * 1115)

}
