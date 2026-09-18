// // package main

// // // Go 标准库里的 log 可以先简单理解为：
// // // 让 Server 把运行信息记录到自己的终端里。
// // import (
// // 	"fmt"
// // 	"log"
// // 	"net/http"
// // )

// // // helloHandler负责处理 /hello 请求。
// // func helloHandler(
// // 	w http.ResponseWriter,
// // 	r *http.Request,
// // ) {
// // 	fmt.Fprintln(
// // 		w,
// // 		"hello",
// // 	)
// // }

// // // studyHandler负责处理 /study 请求。
// // func studyHandler(
// // 	w http.ResponseWriter,
// // 	r *http.Request,
// // ) {
// // 	fmt.Fprintln(
// // 		w,
// // 		"I am studying Go backend",
// // 	)
// // }

// // // nameHandler负责处理 /name 请求。
// // func nameHandler(
// // 	w http.ResponseWriter,
// // 	r *http.Request,
// // ) {
// // 	fmt.Fprintln(
// // 		w,
// // 		"my name is hui",
// // 	)
// // }

// // func main() {
// // 	// 创建一个属于我们自己的ServeMux。
// // 	//
// // 	// 可以暂时把它理解成：
// // 	// “保存Path -> Handler关系，并负责路由匹配的分发器”。
// // 	mux := http.NewServeMux()

// // 	// 在mux中注册路由：
// // 	//
// // 	// /hello
// // 	// ↓
// // 	// helloHandler
// // 	mux.HandleFunc(
// // 		"/hello",
// // 		helloHandler,
// // 	)

// // 	// 在mux中注册路由：
// // 	//
// // 	// /study
// // 	// ↓
// // 	// studyHandler
// // 	mux.HandleFunc(
// // 		"/study",
// // 		studyHandler,
// // 	)

// // 	// 在mux中注册路由：
// // 	//
// // 	// /name
// // 	// ↓
// // 	// nameHandler
// // 	mux.HandleFunc(
// // 		"/name",
// // 		nameHandler,
// // 	)

// // 	// 服务器运行记录
// // 	log.Println(
// // 		"server started at http://localhost:8080",
// // 	)

// // 	// Server监听8080端口。
// // 	//
// // 	// Request到达Server以后，
// // 	// Server会把Request交给mux。
// // 	//
// // 	// mux再根据Path找到对应Handler。

// //		// 启动 Server；如果 ListenAndServe 最终返回错误，就把这个错误打印出来，并结束程序。
// //		// 记录致命错误并退出
// //		log.Fatal(
// //			http.ListenAndServe(
// //				":8080",
// //				mux,
// //			),
// //		)
// //	}
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func helloHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	fmt.Fprintln(
// 		w,
// 		"hello",
// 	)
// }

// func studyHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	fmt.Fprintln(
// 		w,
// 		"I am studying Go backend",
// 	)
// }

// func main() {
// 	mux := http.NewServeMux()

// 	// 显式把普通函数helloHandler
// 	// 转换成http.HandlerFunc。
// 	//
// 	// HandlerFunc实现了Handler接口，
// 	// 因此可以传给mux.Handle。
// 	mux.Handle(
// 		"/hello",
// 		http.HandlerFunc(
// 			helloHandler,
// 		),
// 	)

// 	// HandleFunc是更方便的写法。
// 	//
// 	// 可以直接把普通处理函数传进来。
// 	mux.HandleFunc(
// 		"/study",
// 		studyHandler,
// 	)

// 	log.Println(
// 		"starting server at http://localhost:8080",
// 	)

//		log.Fatal(
//			http.ListenAndServe(
//				":8080",
//				mux,
//			),
//		)
//	}
package main

import (
	// 自己判断
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
func studyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am studying Go backend")
}
func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is running")
}

func main() {
	// 1. 创建mux
	mux := http.NewServeMux()
	// 2. 注册三个Route
	mux.HandleFunc(
		"/hello",
		helloHandler,
	)
	mux.HandleFunc(
		"/study",
		studyHandler,
	)
	mux.HandleFunc(
		"/status",
		statusHandler,
	)
	// 3. 输出Server启动日志
	log.Println("Server is listening at http:localhost::8080")
	// 4. 启动Server
	log.Fatal(
		http.ListenAndServe(":8080", mux),
	)
}
