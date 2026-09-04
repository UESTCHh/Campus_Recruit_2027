// package main

// // import "net/http"
// // 这是 Go 标准库提供的 HTTP 包。
// // 它帮助我们完成：
// // 启动HTTP Server
// // 接收HTTP Request
// // 路由请求
// // 调用Handler
// // 构造HTTP Response
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // HTTP Handler
// // 简单理解：
// // Handler 是负责处理一次 HTTP Request 的函数。
// // r
// // → Request
// // → 客户端发给我们的信息
// // 和：
// // w
// // → ResponseWriter
// // → 我们用它向客户端写Response
// func helloHandler(
// 	// ResponseWriter 是一个：
// 	// interface
// 	// 我们通过它：
// 	// 向HTTP Response写入数据
// 	w http.ResponseWriter,
// 	// http.Request 是一个 struct。
// 	// 里面包含一次 HTTP 请求的信息，例如：
// 	// 请求方法
// 	// URL
// 	// Header
// 	// Body
// 	// 其他请求信息
// 	r *http.Request,
// ) {
// 	// //不要把内容打印到终端，而是写到 w。
// 	// fmt.Fprintln(
// 	// 	w,
// 	// 	"Hello, Go HTTP!",
// 	// )

// 	//HTTP Method
// 	//浏览器或 curl 默认访问时通常：GET
// 	//GET : 获取 / 查询资源
// 	fmt.Println(
// 		"method:",
// 		r.Method,
// 	)

// 	//请求路径
// 	fmt.Println(
// 		"path:",
// 		r.URL.Path,
// 	)

// 	fmt.Fprintln(
// 		w,
// 		"Hello, Go HTTP!",
// 	)
// }

// func main() {
// 	//路由注册
// 	//如果请求路径是 /hello → 调用 helloHandler
// 	http.HandleFunc(
// 		"/hello",
// 		helloHandler,
// 	)

// 	// URL
// 	// http://
// 	// → 协议

// 	// localhost
// 	// → Host

// 	// 8080
// 	// → Port

// 	// /hello
// 	// → Path
// 	fmt.Println(
// 		"server listening on http://localhost:8080",
// 	)

// 	//在 8080 端口启动 HTTP Server，并开始等待客户端请求。
// 	// 服务器正常运行时会一直：
// 	// 监听
// 	// ↓
// 	// 等待Request
// 	// ↓
// 	// 处理Request
// 	// ↓
// 	// 继续监听
// 	// 所以它通常不会在服务器正常运行期间返回。
// 	err := http.ListenAndServe(
// 		":8080",
// 		nil, //这里第二个参数本质上需要一个 HTTP Handler,nil代表使用 net/http 的默认路由器 DefaultServeMux
// 	)

// 	//端口已经被其他程序占用
// 	if err != nil {
// 		//打印错误 + 结束程序
// 		log.Fatal(err)
// 	}
// }

package main

// net/http 是 Go 标准库提供的 HTTP 包。
// 今天主要使用它完成：
// 1. 注册路由
// 2. 编写Handler
// 3. 启动HTTP Server
// 4. 设置HTTP Status Code
import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler 负责处理：
// GET /hello
//
// Handler 有两个核心参数：
//
// w http.ResponseWriter
// → 用来向客户端写HTTP Response
//
// r *http.Request
// → 表示客户端发送过来的HTTP Request
// func helloHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	// r.Method 表示HTTP请求方法。
// 	// 浏览器直接访问时通常是 GET。
// 	fmt.Println(
// 		"method:",
// 		r.Method,
// 	)

// 	// r.URL.Path 表示请求的Path。
// 	// 访问：
// 	// http://localhost:8080/hello
// 	// 时，这里的值是：
// 	// /hello
// 	fmt.Println(
// 		"path:",
// 		r.URL.Path,
// 	)

//		// 这里没有主动调用：
//		// w.WriteHeader(http.StatusOK)
//		//
//		// 当我们直接开始向Response Body写数据时，
//		// 如果之前没有设置状态码，
//		// net/http会自动使用：
//		// 200 OK
//		fmt.Fprintln(
//			w,
//			"Hello, Go HTTP!",
//		)
//	}
//

// helloHandler 负责处理：
// GET /hello
//
// 当前这个Handler只允许GET请求。
// 如果客户端使用其他Method，例如POST，
// 则返回：
// 405 Method Not Allowed
func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// r.Method表示客户端本次HTTP Request使用的方法。
	//
	// http.MethodGet是net/http提供的常量，
	// 它表示字符串：
	// "GET"
	if r.Method != http.MethodGet {
		// Header必须尽量在WriteHeader之前设置。
		//
		// 这里告诉客户端：
		// 当前资源允许的HTTP Method是GET。
		//
		// 最终Response Header中可以看到：
		// Allow: GET
		w.Header().Set(
			"Allow",
			http.MethodGet,
		)

		// http.StatusMethodNotAllowed
		// → 405
		//
		// 表示：
		// Path本身存在，
		// 但是当前HTTP Method不被这个资源允许。
		w.WriteHeader(
			http.StatusMethodNotAllowed,
		)

		// Status设置完成以后，
		// 再写Response Body。
		fmt.Fprintln(
			w,
			"405 - method not allowed",
		)

		// return非常重要。
		//
		// 如果这里不return，
		// 函数还会继续执行下面正常GET请求的逻辑，
		// 从而继续向Response写入内容。
		return
	}

	// 能执行到这里，说明：
	// r.Method == http.MethodGet
	fmt.Println(
		"method:",
		r.Method,
	)

	fmt.Println(
		"path:",
		r.URL.Path,
	)

	// 没有显式WriteHeader，
	// 第一次写Response Body时会自动使用：
	// 200 OK
	fmt.Fprintln(
		w,
		"Hello, Go HTTP!",
	)
}

// healthHandler 负责处理：
// GET /health
//
// 以后真实后端经常会有类似健康检查接口，
// 用来判断Server是否正常运行。
func healthHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Println(
		"method:",
		r.Method,
	)

	fmt.Println(
		"path:",
		r.URL.Path,
	)

	// http.StatusOK 是net/http提供的状态码常量。
	// 它的值是：
	// 200
	//
	// WriteHeader 用来显式写入HTTP Response Status Code。
	w.WriteHeader(http.StatusOK)

	// 状态码设置完成后，
	// 再写Response Body。
	fmt.Fprintln(
		w,
		"server is healthy",
	)
}

// notFoundHandler 用来处理没有匹配到其他业务路由的请求。
//
// 例如：
// GET /abc
// GET /users
//
// 目前这些Path都没有对应的业务Handler，
// 因此返回404 Not Found。
func notFoundHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Println(
		"not found path:",
		r.URL.Path,
	)

	// 先设置404状态码。
	//
	// http.StatusNotFound
	// → 404
	w.WriteHeader(http.StatusNotFound)

	// 再写Response Body。
	fmt.Fprintln(
		w,
		"404 - resource not found",
	)
}

func main() {
	// 注册：
	//
	// /hello
	// ↓
	// helloHandler
	http.HandleFunc(
		"/hello",
		helloHandler,
	)

	// 注册第二个业务路由：
	//
	// /health
	// ↓
	// healthHandler
	http.HandleFunc(
		"/health",
		healthHandler,
	)

	// "/" 在当前DefaultServeMux的使用方式下，
	// 可以作为其他没有更具体路由匹配时的兜底Handler。
	//
	// 比如：
	// /abc
	//
	// 没有注册 /abc，
	// 就会进入这里。
	http.HandleFunc(
		"/",
		notFoundHandler,
	)

	fmt.Println(
		"server listening on http://localhost:8080",
	)

	// ListenAndServe：
	//
	// ":8080"
	// → 监听8080端口
	//
	// nil
	// → 使用DefaultServeMux
	//
	// Server正常运行时，
	// ListenAndServe会持续监听Request，
	// 所以通常不会正常返回。
	err := http.ListenAndServe(
		":8080",
		nil,
	)

	// 如果ListenAndServe真的返回，
	// 通常意味着Server发生了错误，
	// 比如8080端口已经被其他程序占用。
	if err != nil {
		// log.Fatal：
		// 1. 打印错误
		// 2. 结束程序
		log.Fatal(err)
	}
}
