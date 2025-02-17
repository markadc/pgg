package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 5; i++ {
		go func(id int) {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine", id, "canceled")
			case <-time.After(2 * time.Second):
				fmt.Println("Goroutine", id, "completed")
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	cancel() // 取消所有 goroutine

	time.Sleep(3 * time.Second)

}
