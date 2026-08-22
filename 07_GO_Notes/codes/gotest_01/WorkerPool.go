// 文件路径：workerpool.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// worker 从 jobs 通道读取任务，处理后将结果写入 results 通道
// 任务处理完毕时调用 wg.Done() 通知 WaitGroup
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 无论函数如何退出，都标记该 worker 完成
	for job := range jobs {
		fmt.Printf("[Worker %d] 处理任务 %d\n", id, job)
		time.Sleep(200 * time.Millisecond) // 模拟任务处理耗时
		results <- job * job               // 计算平方作为结果
	}
}

func main() {
	const numJobs = 10   // 总任务数
	const numWorkers = 3 // 并发 worker 数量

	jobs := make(chan int, numJobs)    // 带缓冲的任务通道
	results := make(chan int, numJobs) // 带缓冲的结果通道
	var wg sync.WaitGroup              // 用于等待所有 worker 完成

	fmt.Println("=== RUNOOB Worker Pool 示例 ===")

	// 启动固定数量的 worker goroutine
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // 每启动一个 worker，WaitGroup 计数加 1
		go worker(w, jobs, results, &wg)
	}

	// 向通道发送所有任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 jobs 通道，worker 的 range 循环会自行退出

	// 等待所有 worker 完成后，关闭结果通道
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集并打印所有结果
	fmt.Println("结果汇总:")
	for result := range results {
		fmt.Printf("%d ", result)
	}
	fmt.Println("\n所有任务处理完成")
}
