// package main

// // import "encoding/json"
// // 这是 Go 标准库提供的 JSON 包。
// // 它主要负责：
// // 	Go value
// // 	↓
// // 	JSON
// // 以及以后：
// // 	JSON
// // 	↓
// // 	Go value
// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // User 是我们准备返回给客户端的数据结构。
// //
// // Go 中使用 struct 表达一组结构化数据。
// //
// // 字段：
// // Name
// // Age
// //
// // 首字母大写表示这些字段是 Exported，
// // 其他 package（包括 encoding/json）
// // 才能够访问它们。
// type User struct {
// 	// `json:"name"` 是 Struct Tag。
// 	//
// 	// 它告诉 encoding/json：
// 	// 当 Name 字段被编码成 JSON 时，
// 	// JSON 中字段名使用：
// 	//
// 	// "name"
// 	Name string `json:"name"`

// 	// Go字段：
// 	// Age
// 	//
// 	// JSON字段：
// 	// "age"
// 	Age int `json:"age"`
// }

// // userHandler 处理：
// //
// // GET /user
// //
// // 今天这个Handler会返回一个真正的JSON Response。
// func userHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	// 当前接口只允许GET。
// 	//
// 	// r.Method：
// 	// → 当前Request真正使用的方法
// 	//
// 	// http.MethodGet：
// 	// → 标准库提供的 "GET" 常量
// 	if r.Method != http.MethodGet {
// 		// 先设置Response Header。
// 		//
// 		// Allow告诉客户端：
// 		// 当前这个资源允许GET。
// 		w.Header().Set(
// 			"Allow",
// 			http.MethodGet,
// 		)

// 		// 再设置HTTP Status：
// 		// 405 Method Not Allowed
// 		w.WriteHeader(
// 			http.StatusMethodNotAllowed,
// 		)

// 		// 最后写Response Body。
// 		fmt.Fprintln(
// 			w,
// 			"405 - method not allowed",
// 		)

// 		return
// 	}

// 	// 创建一个Go struct实例。
// 	//
// 	// 目前数据直接写死，
// 	// 后面真正项目中会来自：
// 	//
// 	// Service
// 	// ↓
// 	// Repository
// 	// ↓
// 	// Database
// 	user := User{
// 		Name: "Alice",
// 		Age:  22,
// 	}

// 	// HTTP Response Header：
// 	//
// 	// Content-Type告诉客户端：
// 	// Response Body里面的数据是什么格式。
// 	//
// 	// application/json：
// 	// → 表示Body是JSON数据。
// 	//
// 	// Header应该在WriteHeader或写Body之前设置。
// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	// 显式设置：
// 	//
// 	// 200 OK
// 	//
// 	// 这里其实可以不写，
// 	// 因为第一次写Body时，
// 	// net/http默认也会发送200。
// 	//
// 	// 今天故意显式写出来，
// 	// 是为了把：
// 	//
// 	// Header
// 	// ↓
// 	// Status
// 	// ↓
// 	// Body
// 	//
// 	// 这个Response构造顺序再次固定下来。
// 	w.WriteHeader(
// 		http.StatusOK,
// 	)

// 	// json.NewEncoder(w)
// 	//
// 	// 创建一个JSON Encoder。
// 	//
// 	// Encoder的输出目标是：
// 	// w
// 	// 即HTTP ResponseWriter。
// 	//
// 	// Encode(user)
// 	//
// 	// 会把：
// 	// Go struct
// 	//
// 	// User{
// 	//     Name: "Alice",
// 	//     Age: 22,
// 	// }
// 	//
// 	// 编码为JSON，
// 	// 并直接写入HTTP Response Body。
// 	err := json.NewEncoder(w).Encode(user)

// 	// JSON编码也可能失败，
// 	// 所以Encode返回error。
// 	//
// 	// 目前User非常简单，
// 	// 正常情况下不会失败。
// 	//
// 	// 今天先保持最基础的错误处理：
// 	// 在Server终端记录错误。
// 	if err != nil {
// 		log.Println(
// 			"encode user:",
// 			err,
// 		)
// 	}
// }

// func main() {
// 	// 注册路由：
// 	//
// 	// /user
// 	// ↓
// 	// userHandler
// 	http.HandleFunc(
// 		"/user",
// 		userHandler,
// 	)

// 	fmt.Println(
// 		"server listening on http://localhost:8080",
// 	)

// 	// ":8080"
// 	// → 监听8080端口
// 	//
// 	// nil
// 	// → 使用DefaultServeMux
// 	err := http.ListenAndServe(
// 		":8080",
// 		nil,
// 	)

//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User 表示一个用户。
//
// Go中使用struct组织结构化数据。
// JSON Encoder会根据Struct Tag
// 将Go字段转换成对应的JSON字段。
type User struct {
	// Go字段：
	// Name
	//
	// JSON字段：
	// name
	Name string `json:"name"`

	// Go字段：
	// Age
	//
	// JSON字段：
	// age
	Age int `json:"age"`
}

// userHandler 处理：
//
// GET /user
//
// 返回一个User，
// 因此最终JSON是一个JSON Object。
func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 当前接口只允许GET。
	if r.Method != http.MethodGet {
		// 设置允许的HTTP Method。
		w.Header().Set(
			"Allow",
			http.MethodGet,
		)

		// 405：
		// Path存在，
		// 但是当前Method不允许。
		w.WriteHeader(
			http.StatusMethodNotAllowed,
		)

		fmt.Fprintln(
			w,
			"405 - method not allowed",
		)

		return
	}

	// 创建一个User struct。
	user := User{
		Name: "Alice",
		Age:  22,
	}

	// Response Body将是JSON。
	//
	// Header必须在WriteHeader
	// 或第一次写Body之前设置。
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// 显式返回：
	// 200 OK
	w.WriteHeader(
		http.StatusOK,
	)

	// Go struct：
	//
	// User
	//
	// ↓ JSON Encoder
	//
	// JSON Object：
	//
	// {
	//     "name": "Alice",
	//     "age": 22
	// }
	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		log.Println(
			"encode user:",
			err,
		)
	}
}

// usersHandler 处理：
//
// GET /users
//
// 和 /user 不同，
// 这里返回多个User。
//
// Go中使用：
// []User
//
// JSON中会编码成：
// JSON Array。
func usersHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 当前接口同样只允许GET。
	if r.Method != http.MethodGet {
		w.Header().Set(
			"Allow",
			http.MethodGet,
		)

		w.WriteHeader(
			http.StatusMethodNotAllowed,
		)

		fmt.Fprintln(
			w,
			"405 - method not allowed",
		)

		return
	}

	// []User：
	// 表示User类型的Slice。
	//
	// Slice中的每个元素
	// 都是一个User struct。
	users := []User{
		{
			Name: "Alice",
			Age:  22,
		},
		{
			Name: "Bob",
			Age:  25,
		},
		{
			Name: "Charlie",
			Age:  28,
		},
	}

	// 告诉客户端：
	// Body使用JSON格式。
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	// Go：
	//
	// []User
	//
	// ↓
	// encoding/json
	//
	// JSON：
	//
	// [
	//     {...},
	//     {...},
	//     {...}
	// ]
	//
	// Slice会被自动编码成JSON Array。
	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		log.Println(
			"encode users:",
			err,
		)
	}
}

func main() {
	// 单个User：
	//
	// GET /user
	// ↓
	// userHandler
	http.HandleFunc(
		"/user",
		userHandler,
	)

	// 多个User：
	//
	// GET /users
	// ↓
	// usersHandler
	http.HandleFunc(
		"/users",
		usersHandler,
	)

	fmt.Println(
		"server listening on http://localhost:8080",
	)

	// 启动HTTP Server。
	//
	// ":8080"
	// → 监听8080端口
	//
	// nil
	// → 使用DefaultServeMux
	err := http.ListenAndServe(
		":8080",
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
}
