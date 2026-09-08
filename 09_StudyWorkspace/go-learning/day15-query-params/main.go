// 文件作用：
//
//	学习：
//	GET Query Parameter
//	r.URL.Query()
//	Query.Get()
//	strconv.Atoi
//	参数解析
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

// QueryResponse 表示 /users 接口返回给客户端的数据。
//
// 今天重点不是这个struct本身，
// 而是：
//
// URL Query Parameter
// ↓
// string
// ↓
// strconv.Atoi
// ↓
// int
// ↓
// Validation
type QueryResponse struct {
	// MinAge 是Server成功解析后的整数参数。
	MinAge int `json:"min_age"`

	// Message 用于让客户端确认请求已经成功处理。
	Message string `json:"message"`
}

// usersHandler 处理：
//
// GET /users?min_age=18
//
// 今天学习的是：
//
// # Query Parameter
//
// 它和昨天POST JSON Body不同。
func usersHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
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
	// 1. 读取URL Query Parameters
	// --------------------------------------------------
	//
	// 假设Request：
	//
	// GET /users?min_age=18
	//
	// r.URL.Query()
	//
	// 会解析URL中：
	//
	// ?min_age=18
	//
	// 这一部分Query String。
	query := r.URL.Query()

	// Query Parameter默认以string形式读取。
	//
	// 如果：
	//
	// /users?min_age=18
	//
	// 那么这里得到：
	//
	// "18"
	//
	// 注意：
	//
	// 这是string，
	// 还不是int 18。
	minAgeText :=
		query.Get("min_age")

	// --------------------------------------------------
	// 2. 检查参数是否缺失
	// --------------------------------------------------
	//
	// 如果URL中没有：
	//
	// ?min_age=...
	//
	// Get("min_age")
	// 会得到空字符串。
	if minAgeText == "" {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - min_age is required",
		)

		return
	}

	// --------------------------------------------------
	// 3. string → int
	// --------------------------------------------------
	//
	// strconv.Atoi：
	//
	// 将十进制数字字符串
	// 转换为int。
	//
	// 例如：
	//
	// "18"
	// ↓
	// 18
	//
	// 如果传入：
	//
	// "abc"
	//
	// 则转换失败，
	// err != nil。
	minAge, err :=
		strconv.Atoi(minAgeText)

	if err != nil {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - min_age must be an integer",
		)

		return
	}

	// --------------------------------------------------
	// 4. Validation
	// --------------------------------------------------
	//
	// 到这里说明：
	//
	// string → int
	//
	// 已经解析成功。
	//
	// 但是：
	//
	// 成功解析
	// ≠
	// 业务数据一定合法。
	//
	// 这和昨天：
	//
	// Decoding
	// ≠
	// Validation
	//
	// 是完全一样的思想。
	//
	// 当前简单规定：
	//
	// min_age必须 > 0。
	if minAge <= 0 {
		w.WriteHeader(
			http.StatusBadRequest,
		)

		fmt.Fprintln(
			w,
			"400 - min_age must be greater than 0",
		)

		return
	}

	// --------------------------------------------------
	// 5. 构造Go Response数据
	// --------------------------------------------------
	response := QueryResponse{
		MinAge:  minAge,
		Message: "query accepted",
	}

	// Response Body使用JSON。
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	// 复用Day13已经学习过的：
	//
	// Go
	// ↓
	// JSON Encoding
	// ↓
	// ResponseWriter
	err = json.NewEncoder(w).Encode(
		response,
	)

	if err != nil {
		log.Println(
			"encode response:",
			err,
		)
	}
}

func main() {
	// 注册：
	//
	// /users
	// ↓
	// usersHandler
	http.HandleFunc(
		"/users",
		usersHandler,
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
