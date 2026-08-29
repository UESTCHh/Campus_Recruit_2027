// 文件作用
//
//	暂时不用真实 HTTP，而是模拟：
//	Handler收到请求
//	↓
//	给请求设置700ms timeout
//	↓
//	调用Service
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		700*time.Millisecond,
	)
	defer cancel()

	fmt.Println("handler: request started")

	profile, err := GetUserProfile(
		ctx,
		1001,
	)

	if err != nil {
		if errors.Is(
			err,
			context.DeadlineExceeded,
		) {
			fmt.Println(
				"handler: request timeout",
			)
			return
		}

		fmt.Println(
			"handler: error:",
			err,
		)
		return
	}
	fmt.Println(
		"handler:",
		profile,
	)
}
