package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, index int) error {
	defer wg.Done()

	cnt := 1
	for {
		select {
		default:
			fmt.Printf("worker(%v): hello %v\n", index+1, cnt)
			cnt++
			time.Sleep(100 * time.Millisecond)
		case <-ctx.Done():
			fmt.Printf("worker(%v): %s\n", index+1, ctx.Err())
			return ctx.Err()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}
