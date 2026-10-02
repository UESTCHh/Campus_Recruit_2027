// 文件职责：
// 验证：
// Context自动超时
// +
// Done通知
// +
// Err区分取消原因
package main

import (
	"context"
	"fmt"
	"time"
)

func slowWork(ctx context.Context) {
	fmt.Println("worker: start")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("worker: finished normally")

	case <-ctx.Done():
		fmt.Println("worker: stopped")
		fmt.Println("worker error:", ctx.Err())
	}
}

func main() {
	// 基于Background创建一个最长存活2秒的Context。
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	// 即使Context最终会自动超时，
	// 仍然建议调用cancel释放相关资源。
	defer cancel()

	slowWork(ctx)

	fmt.Println("main: finished")
}
