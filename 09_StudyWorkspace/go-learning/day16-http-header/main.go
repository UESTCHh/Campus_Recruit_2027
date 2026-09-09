// 文件作用：
//
//	学习：
//	HTTP Request Header
//	r.Header
//	Header.Get()
//	自定义Header
//	Required Check
//	strconv.Atoi
//	Validation
//	JSON Response
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ProfileResponse 表示 /profile 接口成功时
// 返回给客户端的JSON数据。
//
// 今天真正的重点是：
//
// HTTP Request Header
// ↓
// string
// ↓
// Parsing
// ↓
// Validation
//
// 这个struct只是为了继续复用之前学过的
// JSON Encoding。
type ProfileResponse struct {
	// ClientID 是从 Request Header 中读取、
	// 解析并通过Validation后的客户端ID。
	ClientID int `json:"client_id"`

	// Message 用于说明Header已经成功处理。
	Message string `json:"message"`
}

// profileHandler 处理：
//
// GET /profile
//
// 客户端还必须在Request Header中提供：
//
// X-Client-ID: 1001
//
// 完整的数据处理链路：
//
// Request Header
// ↓
// r.Header.Get(...)
// ↓
// string
// ↓
// Required Check
// ↓
// strconv.Atoi
// ↓
// int
// ↓
// Validation
// ↓
// JSON Response
func profileHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// --------------------------------------------------
	// 1. Method Check
	// --------------------------------------------------
	//
	// 当前接口只允许GET。
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

	// --------------------------------------------------
	// 2. 从Request Header读取X-Client-ID
	// --------------------------------------------------
	//
	// r.Header：
	//
	// 表示当前HTTP Request携带的Headers。
	//
	// Get("X-Client-ID")：
	//
	// 根据Header Name取得对应的Header Value。
	//
	// 例如客户端发送：
	//
	// X-Client-ID: 1001
	//
	// 那么这里得到：
	//
	// "1001"
	//
	// 注意仍然是string。
	clientIDText :=
		r.Header.Get("X-Client-ID")

	// --------------------------------------------------
	// 3. Required Header Check
	// --------------------------------------------------
	//
	// 如果客户端没有提供：
	//
	// X-Client-ID
	//
	// Get会得到空字符串。
	//
	// 这和昨天Query Parameter缺失时
	// query.Get(...)得到""很像。
	if clientIDText == "" {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - X-Client-ID is required",
		)

		return
	}

	// --------------------------------------------------
	// 4. Parsing
	// --------------------------------------------------
	//
	// Header Value也是文本。
	//
	// 当前业务希望Client ID是int，
	// 因此使用昨天刚学过的：
	//
	// strconv.Atoi
	//
	// "1001"
	// ↓
	// 1001
	clientID, err :=
		strconv.Atoi(clientIDText)

	if err != nil {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - X-Client-ID must be an integer",
		)

		return
	}

	// --------------------------------------------------
	// 5. Validation
	// --------------------------------------------------
	//
	// Parsing成功：
	//
	// 不代表业务数据一定合法。
	//
	// 当前简单规定：
	//
	// Client ID必须 > 0。
	if clientID <= 0 {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - X-Client-ID must be greater than 0",
		)

		return
	}

	// --------------------------------------------------
	// 6. 构造成功响应
	// --------------------------------------------------
	response := ProfileResponse{
		ClientID: clientID,
		Message:  "header accepted",
	}

	// Response Header：
	//
	// 这里的w.Header()
	// 与上面的r.Header完全不是同一个方向。
	//
	// r.Header
	// → Client发给Server的Request Headers
	//
	// w.Header()
	// → Server发给Client的Response Headers
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	// Go struct
	// ↓
	// JSON
	// ↓
	// Response Body
	err = json.NewEncoder(w).Encode(
		response,
	)

	if err != nil {
		log.Println(
			"encode profile response:",
			err,
		)
	}
}

func main() {
	// 注册：
	//
	// /profile
	// ↓
	// profileHandler
	http.HandleFunc(
		"/profile",
		profileHandler,
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
