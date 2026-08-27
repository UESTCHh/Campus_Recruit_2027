// 文件作用
//
//	第一次观察：
//	context timeout
//	↓
//	长任务主动发现取消
//	↓
//	提前停止
//
// 把原来的：
//
//	ctx.Err()
//	+
//	time.Sleep()
//	升级成：
//	ctx.Done()
//	+
//	select
//	使任务可以更及时响应 context 取消。
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func slowWork(ctx context.Context) error {
// 	for i := 1; i <= 10; i++ {
// 		if err := ctx.Err(); err != nil {
// 			return err
// 		}

// 		fmt.Println("working step:", i)

// 		time.Sleep(500 * time.Millisecond)
// 	}

// 	return nil
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		10*time.Second,
// 	)

// 	//让这一小段代码同时在旁边执行。
// 	go func() {
// 		time.Sleep(1200 * time.Millisecond)

// 		fmt.Println("calling cancel()")

// 		cancel()
// 	}()

// 	err := slowWork(ctx)

// 	if err != nil {
// 		fmt.Println("work stopped:", err)
// 		return
// 	}

//		fmt.Println("work finished")
//	}
package main

import (
	"context"
	"fmt"
	"time"
)

func slowWork(ctx context.Context) error {
	for i := 1; i <= 10; i++ {
		//同时等待多个 channel 操作，哪个先准备好，就执行哪个 case。
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(500 * time.Millisecond):
			fmt.Println("working step:", i)
		}
	}

	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1200*time.Millisecond,
	)
	//当前函数准备返回时，再执行这个调用。
	defer cancel()

	err := slowWork(ctx)

	if err != nil {
		fmt.Println("work stopped:", err)
		return
	}

	fmt.Println("work finished")
}
