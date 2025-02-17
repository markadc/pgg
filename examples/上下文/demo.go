package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个有取消功能的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个 goroutine
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine completed")
	}()

	// 在主 goroutine 中调用 cancel 来取消操作
	time.Sleep(1 * time.Second)
	cancel() // 取消 goroutine 中的任务

	select {
	case <-ctx.Done():
		fmt.Println("Operation canceled:", ctx.Err()) // 输出：Operation canceled: context canceled
	}
}
