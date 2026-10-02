package main

import (
	"context"
	"fmt"
)

func main() {
	parentCtx, parentCancel :=
		context.WithCancel(
			context.Background(),
		)

	childCtx, childCancel :=
		context.WithCancel(
			parentCtx,
		)

	// 创建出来的cancel函数都应该最终调用，
	// 便于释放对应Context相关资源。
	defer parentCancel()
	defer childCancel()

	fmt.Println(
		"before cancel:",
	)

	fmt.Println(
		"parent error:",
		parentCtx.Err(),
	)

	fmt.Println(
		"child error:",
		childCtx.Err(),
	)

	fmt.Println(
		"call parentCancel()",
	)

	parentCancel()

	fmt.Println(
		"after parent cancel:",
	)

	fmt.Println(
		"parent error:",
		parentCtx.Err(),
	)

	fmt.Println(
		"child error:",
		childCtx.Err(),
	)
}
