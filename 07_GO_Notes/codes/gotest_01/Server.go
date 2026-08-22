// 文件路径：server.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response 定义 JSON 响应的数据结构
type Response struct {
	Code    int    `json:"code"` // 结构体标签指定 JSON 字段名
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // omitempty：值为空时不输出该字段
}

// helloHandler 处理 /hello 路由的请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 仅允许 GET 请求
	if r.Method != http.MethodGet {
		http.Error(w, "仅支持 GET 请求", http.StatusMethodNotAllowed)
		return
	}

	// 从查询参数获取 name，默认为 "RUNOOB"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "RUNOOB"
	}

	// 构建响应数据
	resp := Response{
		Code:    200,
		Message: fmt.Sprintf("你好，%s！欢迎访问 Go HTTP 服务", name),
	}

	// 设置响应头并返回 JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// 注册路由处理函数
	http.HandleFunc("/hello", helloHandler)

	addr := ":8080"
	fmt.Printf("RUNOOB 服务器启动，监听地址: http://localhost%s\n", addr)
	fmt.Printf("访问示例: http://localhost%s/hello?name=Go开发者\n", addr)

	// 启动 HTTP 服务（阻塞，直到服务停止）
	log.Fatal(http.ListenAndServe(addr, nil))
}
