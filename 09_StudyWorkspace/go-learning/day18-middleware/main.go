// 文件作用：
// 学习：
//
//	HTTP Middleware
//	http.Handler
//	http.HandlerFunc
//	next.ServeHTTP()
//	Middleware包装Handler
//	Logging Middleware
//	Middleware执行前后顺序
//	Middleware读取Request Header
//	Handler / Middleware职责区别
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HelloResponse 表示 /hello 接口
// 成功时返回给客户端的JSON数据。
//
// 这个struct不是今天的新重点，
// 只是继续复用前面已经学习过的：
//
// Go struct
// ↓
// JSON Encoding
// ↓
// HTTP Response
type HelloResponse struct {
	Message string `json:"message"`
}

// helloHandler 是真正负责 /hello 业务的Handler。
//
// Handler的职责：
//
// 处理当前具体业务。
//
// 今天这里的业务很简单：
// 返回一个JSON Response。
func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Println(
		"[handler] helloHandler is running",
	)

	response := HelloResponse{
		Message: "hello from handler",
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	err := json.NewEncoder(w).Encode(
		response,
	)

	if err != nil {
		log.Println(
			"encode response:",
			err,
		)
	}
}

// loggingMiddleware 是今天学习的核心。
//
// Middleware常见形式：
//
// func(next http.Handler) http.Handler
//
// 也就是：
//
// Handler
// ↓
// Middleware包装
// ↓
// 新Handler
//
// next 表示：
//
// 当前Middleware后面
// 将要继续执行的那个Handler。
func loggingMiddleware(
	next http.Handler,
) http.Handler {
	// http.HandlerFunc 可以把：
	//
	// func(
	//     http.ResponseWriter,
	//     *http.Request,
	// )
	//
	// 这种普通函数适配成：
	//
	// http.Handler
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			// ------------------------------------------
			// 1. Middleware前置逻辑
			// ------------------------------------------
			//
			// 这里发生在真正业务Handler执行之前。
			fmt.Println(
				"[middleware before]",
				r.Method,
				r.URL.Path,
			)

			// Middleware也能够读取Request中的信息。
			//
			// 这里复习昨天学过的：
			//
			// r.Header.Get(...)
			//
			// X-Request-ID只是今天用来演示的
			// 一个Request Header。
			requestID :=
				r.Header.Get(
					"X-Request-ID",
				)

			if requestID == "" {
				fmt.Println(
					"[middleware] X-Request-ID: (none)",
				)
			} else {
				fmt.Println(
					"[middleware] X-Request-ID:",
					requestID,
				)
			}

			// ------------------------------------------
			// 2. 把Request交给下一个Handler
			// ------------------------------------------
			//
			// 这是Middleware最关键的一句：
			//
			// next.ServeHTTP(w, r)
			//
			// 表示：
			//
			// 当前Middleware自己的前置工作完成
			// ↓
			// 继续执行后面的Handler
			//
			// 如果没有这一句，
			// 请求就会停在Middleware，
			// helloHandler不会执行。
			next.ServeHTTP(
				w,
				r,
			)

			// ------------------------------------------
			// 3. Middleware后置逻辑
			// ------------------------------------------
			//
			// next.ServeHTTP返回以后，
			// 说明后面的Handler已经执行结束。
			//
			// 所以这里发生在Handler之后。
			fmt.Println(
				"[middleware after]",
				r.Method,
				r.URL.Path,
			)
		},
	)
}

func main() {
	// --------------------------------------------------
	// 1. 创建一个独立ServeMux
	// --------------------------------------------------
	//
	// 以前我们经常使用：
	//
	// http.HandleFunc(...)
	//
	// 配合：
	//
	// ListenAndServe(..., nil)
	//
	// 也就是使用DefaultServeMux。
	//
	// 今天为了更清楚地看到：
	//
	// ServeMux
	// ↓
	// Middleware
	// ↓
	// ListenAndServe
	//
	// 我们自己创建一个mux。
	mux :=
		http.NewServeMux()

	// --------------------------------------------------
	// 2. 注册业务Route
	// --------------------------------------------------
	mux.HandleFunc(
		"GET /hello",
		helloHandler,
	)

	// mux 本身实现了：
	//
	// http.Handler
	//
	// 所以可以传给Middleware。
	//
	// 原来：
	//
	// mux
	//
	// 被：
	//
	// loggingMiddleware
	//
	// 包装以后得到新的http.Handler。
	handler :=
		loggingMiddleware(mux)

	fmt.Println(
		"server listening on http://localhost:8080",
	)

	// --------------------------------------------------
	// 3. Server真正使用的是包装后的handler
	// --------------------------------------------------
	//
	// Request进来以后：
	//
	// handler
	// ↓
	// loggingMiddleware
	// ↓
	// next.ServeHTTP
	// ↓
	// mux
	// ↓
	// helloHandler
	err :=
		http.ListenAndServe(
			":8080",
			handler,
		)

	if err != nil {
		log.Fatal(err)
	}
}
