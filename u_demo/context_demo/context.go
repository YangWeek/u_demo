package contextdemo

import (
	"context"
	"fmt"
	"sync"

	"time"
)

var wg3 sync.WaitGroup

func worker3(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func Init_Context() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}