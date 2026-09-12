// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// )

// // ==============================
// // Middleware 类型定义
// // ==============================

// // Middleware本质:
// // 一个函数类型
// //
// // 输入:
// // http.Handler
// //
// // 输出:
// // http.Handler
// //
// // 作用:
// // 接收一个handler
// // 返回一个包装后的handler
// type Middleware func(http.Handler) http.Handler

// // ==============================
// // Context Key
// // ==============================

// // 不推荐直接使用string作为context key
// //
// // 例如:
// // context.WithValue(ctx,"userID",1001)
// //
// // 因为不同package可能出现相同key导致冲突
// //
// // 所以定义自己的类型
// type contextKey string

// const userIDKey contextKey = "userID"

// // ==============================
// // Middleware Chain
// // ==============================

// // Chain负责把多个middleware组合起来
// //
// // 输入:
// //
// // handler
// //
// // middleware:
// // logging
// // auth
// //
// // 输出:
// //
// // logging(
// //      auth(
// //          handler
// //      )
// // )
// func Chain(
// 	handler http.Handler,
// 	middlewares ...Middleware,
// ) http.Handler {

// 	// 为什么倒序？
// 	//
// 	// 因为middleware执行顺序希望:
// 	//
// 	// logging
// 	//    ↓
// 	// auth
// 	//    ↓
// 	// handler
// 	//
// 	// 所以包装过程必须从最后一个开始

// 	for i := len(middlewares)-1; i >= 0; i-- {

// 		handler = middlewares[i](handler)

// 	}

// 	return handler
// }

// // ==============================
// // Logging Middleware
// // ==============================

// func loggingMiddleware(
// 	next http.Handler,
// ) http.Handler {

// 	return http.HandlerFunc(
// 		func(
// 			w http.ResponseWriter,
// 			r *http.Request,
// 		){

// 			fmt.Println(
// 				"[logging before]",
// 			)

// 			// 调用下一个handler
// 			next.ServeHTTP(
// 				w,
// 				r,
// 			)

// 			fmt.Println(
// 				"[logging after]",
// 			)

// 		},
// 	)

// }

// // ==============================
// // Auth Middleware
// // ==============================

// func authMiddleware(
// 	next http.Handler,
// ) http.Handler {

// 	return http.HandlerFunc(
// 		func(
// 			w http.ResponseWriter,
// 			r *http.Request,
// 		){

// 			fmt.Println(
// 				"[auth before]",
// 			)

// 			// 模拟认证成功
// 			// 得到用户ID
// 			userID := 1001

// 			// 将userID写入context
// 			ctx := context.WithValue(
// 				r.Context(),
// 				userIDKey,
// 				userID,
// 			)

// 			// 创建新的request
// 			// 携带新的context
// 			r = r.WithContext(
// 				ctx,
// 			)

// 			// 继续执行handler
// 			next.ServeHTTP(
// 				w,
// 				r,
// 			)

// 			fmt.Println(
// 				"[auth after]",
// 			)

// 		},
// 	)

// }

// // ==============================
// // Business Handler
// // ==============================

// func helloHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ){

// 	// 从context读取middleware传递的数据

// 	userID := r.Context().Value(
// 		userIDKey,
// 	)

// 	fmt.Println(
// 		"handler userID:",
// 		userID,
// 	)

// 	fmt.Fprintln(
// 		w,
// 		"hello",
// 	)

// }

// // ==============================
// // main
// // ==============================

// func main(){

// 	mux := http.NewServeMux()

// 	mux.HandleFunc(
// 		"/hello",
// 		helloHandler,
// 	)

// 	// 将多个middleware包装mux

// 	handler := Chain(
// 		mux,
// 		loggingMiddleware,
// 		authMiddleware,
// 	)

// 	fmt.Println(
// 		"server listening on http://localhost:8080",
// 	)

// 	err := http.ListenAndServe(
// 		":8080",
// 		handler,
// 	)

// 	if err != nil {

// 		panic(err)

// 	}

// }
package main

import (
	"context"
	"fmt"
	"net/http"
)

// ==============================
// Middleware 类型定义
// ==============================

// Middleware本质:
// 一个函数类型
//
// 输入:
// http.Handler
//
// 输出:
// http.Handler
//
// 作用:
// 接收一个handler
// 返回一个包装后的handler
type Middleware func(http.Handler) http.Handler

// ==============================
// Context Key
// ==============================

// 不推荐直接使用string作为context key
//
// 例如:
// context.WithValue(ctx,"userID",1001)
//
// 因为不同package可能出现相同key导致冲突
//
// 所以定义自己的类型
type contextKey string

const userIDKey contextKey = "userID"

// ==============================
// Middleware Chain
// ==============================

// Chain负责把多个middleware组合起来
//
// 输入:
//
// handler
//
// middleware:
// logging
// auth
//
// 输出:
//
// logging(
//
//	auth(
//	    handler
//	)
//
// )
func Chain(
	handler http.Handler,
	middlewares ...Middleware,
) http.Handler {

	// 为什么倒序？
	//
	// 因为middleware执行顺序希望:
	//
	// logging
	//    ↓
	// auth
	//    ↓
	// handler
	//
	// 所以包装过程必须从最后一个开始

	for i := len(middlewares) - 1; i >= 0; i-- {

		handler = middlewares[i](handler)

	}

	return handler
}

// ==============================
// Logging Middleware
// ==============================

func loggingMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			fmt.Println(
				"[logging before]",
			)

			// 调用下一个handler
			next.ServeHTTP(
				w,
				r,
			)

			fmt.Println(
				"[logging after]",
			)

		},
	)

}

// ==============================
// Auth Middleware
// ==============================

func authMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			fmt.Println(
				"[auth before]",
			)

			// 模拟认证成功
			// 得到用户ID
			userID := 1001

			// 将userID写入context
			ctx := context.WithValue(
				r.Context(),
				userIDKey,
				userID,
			)

			// 创建新的request
			// 携带新的context
			r = r.WithContext(
				ctx,
			)

			// 继续执行handler
			next.ServeHTTP(
				w,
				r,
			)

			fmt.Println(
				"[auth after]",
			)

		},
	)

}

// ==============================
// Business Handler
// ==============================

func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// 从context读取middleware传递的数据

	userID := r.Context().Value(
		userIDKey,
	)

	fmt.Println(
		"handler userID:",
		userID,
	)

	fmt.Fprintln(
		w,
		"hello",
	)

}

// ==============================
// main
// ==============================

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/hello",
		helloHandler,
	)

	// 将多个middleware包装mux

	handler := Chain(
		mux,
		loggingMiddleware,
		authMiddleware,
	)

	fmt.Println(
		"server listening on http://localhost:8080",
	)

	err := http.ListenAndServe(
		":8080",
		handler,
	)

	if err != nil {

		panic(err)

	}

}
