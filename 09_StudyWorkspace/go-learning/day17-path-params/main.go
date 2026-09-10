// 学习：
// Path Parameter
// ServeMux Path Pattern
// r.PathValue()
// string → int
// Required Check
// Parsing
// Validation
// JSON Response
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// UserResponse 表示根据Path Parameter查询用户后
// 返回给客户端的JSON数据。
//
// 今天还没有数据库，
// 所以这里只返回一个简单的演示结果。
//
// 真正重点是：
//
// /users/{id}
// ↓
// r.PathValue("id")
// ↓
// string
// ↓
// Parsing
// ↓
// Validation
// ↓
// JSON Response
type UserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

// userHandler 处理：
//
// GET /users/{id}
//
// 例如：
//
// GET /users/42
//
// 其中：
//
// 42
//
// 会被ServeMux匹配为Path Parameter：
//
// id
func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// --------------------------------------------------
	// 1. 读取Path Parameter
	// --------------------------------------------------
	//
	// 路由Pattern：
	//
	// GET /users/{id}
	//
	// 如果Request：
	//
	// GET /users/42
	//
	// 那么：
	//
	// r.PathValue("id")
	//
	// 得到：
	//
	// "42"
	//
	// 注意：
	// 仍然是string。
	idText :=
		r.PathValue("id")

	// --------------------------------------------------
	// 2. Required Check
	// --------------------------------------------------
	//
	// 当前路由正常匹配时，
	// id通常应该存在。
	//
	// 这里仍然保留检查，
	// 用于延续我们前几天形成的
	// External Input处理流程。
	if idText == "" {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - user id is required",
		)

		return
	}

	// --------------------------------------------------
	// 3. Parsing
	// --------------------------------------------------
	//
	// Path Parameter本质仍然是文本。
	//
	// "42"
	// ↓
	// strconv.Atoi
	// ↓
	// 42
	id, err :=
		strconv.Atoi(idText)

	if err != nil {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - user id must be an integer",
		)

		return
	}

	// --------------------------------------------------
	// 4. Validation
	// --------------------------------------------------
	//
	// Parsing成功
	// ≠
	// 业务数据一定合法。
	//
	// 当前简单规定：
	//
	// id必须 > 0。
	if id <= 0 {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - user id must be greater than 0",
		)

		return
	}

	// --------------------------------------------------
	// 5. 构造Go Response数据
	// --------------------------------------------------
	response := UserResponse{
		ID:      id,
		Message: "user path accepted",
	}

	// Response Body使用JSON格式。
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
			"encode user response:",
			err,
		)
	}
}

func main() {
	// Go ServeMux支持带Method和Wildcard的Pattern。
	//
	// GET：
	// 只匹配GET请求。
	//
	// /users/{id}：
	// {id}是Path Parameter。
	//
	// 例如：
	//
	// GET /users/42
	//
	// 会进入userHandler，
	// 并让：
	//
	// r.PathValue("id") == "42"
	http.HandleFunc(
		"GET /users/{id}",
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
