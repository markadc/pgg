package main

import (
	"context"
	"pgg/utils/loger"
	"time"
)

func download(delay time.Duration) {
	loger.Info("下载中...")
	time.Sleep(delay)
	loger.Success("下载完成！！！")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan bool)
	go func() {
		download(2 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		loger.Info("已完成任务")
	case <-ctx.Done():
		loger.Warning("任务已超时")
	}
}
