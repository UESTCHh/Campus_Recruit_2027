package main

import (
	"context"
	"fmt"
	"time"
)

// slowWork模拟一个原本需要5秒才能完成的慢任务。
//
// 它除了等待任务完成，也同时监听ctx.Done()。
// 如果上游提前取消Context，它就可以提前结束。
func slowWork(ctx context.Context) {
	fmt.Println("worker: start")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("worker: finished normally")

	case <-ctx.Done():
		fmt.Println("worker: canceled")
		fmt.Println("worker error:", ctx.Err())
	}
}

func main() {
	// 创建一个可以手动取消的Context。
	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	// 启动一个慢任务。
	go slowWork(ctx)

	// 模拟主流程运行2秒后，
	// 决定不再需要这个慢任务。
	time.Sleep(2 * time.Second)

	fmt.Println("main: call cancel()")

	cancel()

	// 暂时等一小会，让worker有机会打印取消结果。
	time.Sleep(500 * time.Millisecond)
}
