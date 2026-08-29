// 文件作用
// 模拟：
// Repository正在执行一个很慢的数据库查询
package main

import (
	"context"
	"fmt"
	"time"
)

func FindUser(ctx context.Context, id int) (string, error) {
	fmt.Println("repository: start query")

	select {
	case <-ctx.Done():
		fmt.Println("repository: query canceled")
		return "", ctx.Err()
	case <-time.After(2 * time.Second):
		fmt.Println("repository: query finished")
		return fmt.Sprintf("user-%d", id), nil
	}
}
