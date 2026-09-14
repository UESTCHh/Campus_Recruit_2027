package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ==================================================
// API Response Types
// ==================================================

// ErrorResponse 表示所有API错误统一返回的JSON结构。
//
// 例如：
//
//	{
//	    "error": "user not found"
//	}
type ErrorResponse struct {
	Error string `json:"error"`
}

// UserResponse 表示查询User成功时返回的JSON。
type UserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

// ==================================================
// Business Errors
// ==================================================

// errUserNotFound：
//
// 表示Request本身没有问题，
// 但是请求的User Resource不存在。
//
// Handler之后会把它映射成：
//
// 404 Not Found
var errUserNotFound = errors.New("user not found")

// errStorageUnavailable：
//
// 模拟真实后端中：
//
// Database
// Storage
// External Service
//
// 等服务器内部依赖出现故障。
//
// Handler之后会把它映射成：
//
// 500 Internal Server Error
var errStorageUnavailable = errors.New("storage unavailable")

// ==================================================
// Response Helpers
// ==================================================

// writeJSONError：
//
// 统一写JSON格式的错误Response。
func writeJSONError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	// Response Body使用JSON。
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// Status必须在Body之前写出。
	w.WriteHeader(status)

	response :=
		ErrorResponse{
			Error: message,
		}

	err :=
		json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println(
			"encode error response:",
			err,
		)
	}
}

// writeJSON：
//
// 统一写成功JSON Response。
func writeJSON(
	w http.ResponseWriter,
	status int,
	value any,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	err :=
		json.NewEncoder(w).Encode(value)

	if err != nil {
		log.Println(
			"encode response:",
			err,
		)
	}
}

// ==================================================
// Simulated Business Logic
// ==================================================

// findUserByID：
//
// 今天暂时没有真正连接数据库。
//
// 为了学习HTTP Error Mapping，
// 我们使用几个特殊ID模拟不同业务结果：
//
// id == 404
// → 模拟User不存在
// → errUserNotFound
//
// id == 500
// → 模拟服务器内部存储系统失败
// → errStorageUnavailable
//
// 其他正整数：
// → 模拟查询成功。
//
// --------------------------------------------------
//
// 这里非常重要：
//
// findUserByID并不知道HTTP 404 / 500。
//
// 它只知道：
//
// “User不存在”
//
// 或：
//
// “Storage失败”
//
// HTTP Status Code应该由HTTP Handler决定。
func findUserByID(
	id int,
) (UserResponse, error) {
	// 模拟Resource不存在。
	if id == 404 {
		return UserResponse{},
			errUserNotFound
	}

	// 模拟Server内部依赖失败。
	if id == 500 {
		return UserResponse{},
			errStorageUnavailable
	}

	// 模拟查询成功。
	return UserResponse{
		ID:      id,
		Message: "user found",
	}, nil
}

// ==================================================
// HTTP Handler
// ==================================================

func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ------------------------------------------
	// 1. Extract Input
	// ------------------------------------------
	idText :=
		r.PathValue("id")

	// ------------------------------------------
	// 2. Parsing
	// ------------------------------------------
	id, err :=
		strconv.Atoi(idText)

	if err != nil {
		// Client发送的输入无法解析。
		//
		// Client Error
		// → 400 Bad Request
		writeJSONError(
			w,
			http.StatusBadRequest,
			"user id must be an integer",
		)

		return
	}

	// ------------------------------------------
	// 3. Validation
	// ------------------------------------------
	if id <= 0 {
		// 输入能够解析，
		// 但不满足接口要求。
		//
		// Client Error
		// → 400 Bad Request
		writeJSONError(
			w,
			http.StatusBadRequest,
			"user id must be greater than 0",
		)

		return
	}

	// ------------------------------------------
	// 4. Business Logic
	// ------------------------------------------
	user, err :=
		findUserByID(id)

	if err != nil {
		// --------------------------------------
		// Resource Not Found
		// --------------------------------------
		if errors.Is(
			err,
			errUserNotFound,
		) {
			writeJSONError(
				w,
				http.StatusNotFound,
				"user not found",
			)

			return
		}

		// --------------------------------------
		// Internal Server Error
		// --------------------------------------
		//
		// 服务器内部详细错误：
		//
		// errStorageUnavailable
		//
		// 不直接原样暴露给客户端。
		//
		// Server日志记录详细信息。
		log.Println(
			"find user:",
			err,
		)

		// Client只得到比较稳定、
		// 不泄露内部实现细节的错误。
		writeJSONError(
			w,
			http.StatusInternalServerError,
			"internal server error",
		)

		return
	}

	// ------------------------------------------
	// 5. Success Response
	// ------------------------------------------
	writeJSON(
		w,
		http.StatusOK,
		user,
	)
}

func main() {
	mux :=
		http.NewServeMux()

	mux.HandleFunc(
		"GET /users/{id}",
		userHandler,
	)

	fmt.Println(
		"server listening on http://localhost:8080",
	)

	err :=
		http.ListenAndServe(
			":8080",
			mux,
		)

	if err != nil {
		log.Fatal(err)
	}
}
