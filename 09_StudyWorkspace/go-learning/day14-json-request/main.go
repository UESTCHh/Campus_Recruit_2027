// 学习：
//
//	POST Request
//	JSON Request Body
//	json.Decoder
//	Decode(&user)
//	JSON → Go struct
//	错误JSON → 400
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User 表示客户端提交的用户数据。
//
// Struct Tag：
//
// `json:"name"`
// `json:"age"`
//
// 不仅用于：
// Go → JSON 的 Encoding，
//
// 同样也用于：
// JSON → Go 的 Decoding。
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// userHandler 处理：
//
// POST /user
//
// Client会把JSON放在HTTP Request Body中。
//
// 例如：
//
//	{
//	    "name": "Alice",
//	    "age": 22
//	}
//
// Handler需要完成：
//
// Request Body
// ↓
// JSON Decoder
// ↓
// User struct
func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 当前接口只允许POST。
	//
	// r.Method：
	// 当前客户端实际使用的HTTP Method。
	//
	// http.MethodPost：
	// Go标准库提供的"POST"常量。
	if r.Method != http.MethodPost {
		// 告诉客户端：
		// 当前资源允许POST。
		w.Header().Set(
			"Allow",
			http.MethodPost,
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

	// 先声明一个User变量。
	//
	// 此时字段还是Go的零值：
	//
	// user.Name == ""
	// user.Age  == 0
	//
	// 接下来Decoder会把Request Body中的JSON
	// 写进这个user变量。
	var user User

	// r.Body：
	//
	// 表示当前HTTP Request Body的数据流。
	//
	// json.NewDecoder(r.Body)：
	//
	// 创建一个JSON Decoder，
	// 并告诉它：
	//
	// JSON数据要从r.Body读取。
	decoder := json.NewDecoder(
		r.Body,
	)

	// Decode(&user)：
	//
	// 从Request Body中读取JSON，
	// 并把JSON字段解码到user中。
	//
	// 为什么传 &user？
	//
	// 因为Decoder需要修改原来的user变量。
	//
	// &user：
	// → user的地址
	//
	// Decoder拿到这个地址以后，
	// 才能把：
	//
	// "name"
	// ↓
	// user.Name
	//
	// "age"
	// ↓
	// user.Age
	//
	// 写进去。
	err := decoder.Decode(
		&user,
	)

	// 如果客户端发送的JSON格式不合法，
	// Decode会返回error。
	if err != nil {
		// 400 Bad Request：
		//
		// 表示客户端发送的请求数据有问题。
		//
		// 注意仍然遵循：
		//
		// Status
		// ↓
		// Body
		//
		// 这里没有额外设置Header，
		// 所以错误Body保持普通文本即可。
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - invalid JSON",
		)

		return
	}

	// --------------------------------------------------
	// Decoding成功 != 业务数据一定合法
	// --------------------------------------------------
	//
	// JSON Decoder负责：
	//
	// JSON
	// ↓
	// Go struct
	//
	// 但它并不知道我们的业务规则。
	//
	// 例如：
	//
	// {"name":"Alice"}
	//
	// 是合法JSON，
	// 也可以成功Decode。
	//
	// 只是Age没有出现在JSON中，
	// 所以user.Age会保持int的零值：0。
	//
	// 当前我们简单规定：
	//
	// Name不能为空
	// Age必须大于0
	//
	// 这一步叫：
	// Validation
	// 业务数据校验。
	if user.Name == "" || user.Age <= 0 {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - invalid user data",
		)

		return
	}

	// 到这里说明JSON已经成功解码。
	//
	// 可以在Server终端观察：
	//
	// Request Body中的JSON
	// 是否真的进入了Go struct。
	fmt.Println(
		"decoded user:",
		user.Name,
		user.Age,
	)

	// --------------------------------------------------
	// 下面复用昨天Day13的Encoding。
	// --------------------------------------------------
	//
	// 为了让客户端确认服务器确实解析成功，
	// 我们再把刚刚得到的user编码回JSON。
	//
	// 因此完整链路变成：
	//
	// Client JSON
	// ↓
	// Decode
	// ↓
	// Go User
	// ↓
	// Encode
	// ↓
	// Response JSON

	// Response Body将使用JSON格式。
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	// 把刚刚解码得到的user
	// 再编码成JSON写回客户端。
	err = json.NewEncoder(w).Encode(
		user,
	)

	if err != nil {
		// Response已经开始发送，
		// 所以这里当前只在Server日志中记录编码错误。
		log.Println(
			"encode user:",
			err,
		)
	}
}

func main() {
	// 注册：
	//
	// /user
	// ↓
	// userHandler
	http.HandleFunc(
		"/user",
		userHandler,
	)

	fmt.Println(
		"server listening on http://localhost:8080",
	)

	err := http.ListenAndServe(
		":8080",
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
}
