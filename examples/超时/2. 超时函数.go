package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func FuncIsTimeout(maxRunTime time.Duration, runner func()) bool {
	ctx, cancel := context.WithTimeout(context.Background(), maxRunTime)
	defer cancel()

	done := make(chan bool)
	go func() {
		runner()
		close(done)
	}()

	select {
	case <-done:
		return true
	case <-ctx.Done():
		return false
	}
}

func task(delay time.Duration) {
	log.Printf("正在执行（%v）\n", delay)
	time.Sleep(delay)
	log.Printf("执行完成（%v）\n", delay)
}

var wg sync.WaitGroup

func job(i int) {
	defer wg.Done()
	var delay time.Duration
	success := FuncIsTimeout(1*time.Second, func() {
		if i == 1 {
			delay = 500 * time.Millisecond
		} else {
			delay = time.Duration(i) * time.Second
		}
		task(delay)
	})
	if success {
		log.Printf("任务 %v 完成\n", delay)
	} else {
		log.Printf("任务 %v 超时\n", delay)
	}
}

func main() {
	for i := 1; i < 6; i++ {
		wg.Add(1)
		go job(i)
	}
	wg.Wait()
}
