// // // Status
// // // +
// // // Code
// // // +
// // // Message
// // // +
// // // Request ID
// // package main

// // import (
// // 	"context"
// // 	"encoding/json"
// // 	"errors"
// // 	"fmt"
// // 	"log"
// // 	"net/http"
// // 	"strconv"
// // 	"sync/atomic"
// // )

// // // ==================================================
// // // Context
// // // ==================================================

// // type contextKey string

// // const requestIDKey contextKey = "requestID"

// // // ==================================================
// // // Request ID
// // // ==================================================

// // var requestCounter uint64

// // func nextRequestID() string {
// // 	number :=
// // 		atomic.AddUint64(
// // 			&requestCounter,
// // 			1,
// // 		)

// // 	return "req-" +
// // 		strconv.FormatUint(
// // 			number,
// // 			10,
// // 		)
// // }

// // func requestIDFromContext(
// // 	ctx context.Context,
// // ) string {
// // 	requestID, ok :=
// // 		ctx.Value(
// // 			requestIDKey,
// // 		).(string)

// // 	if !ok {
// // 		return ""
// // 	}

// // 	return requestID
// // }

// // // ==================================================
// // // Stable API Error Codes
// // // ==================================================
// // //
// // // Error Code：
// // // 给程序判断错误类型。
// // //
// // // Message：
// // // 给人阅读。
// // //
// // // 所以Client不应该依赖Message文本判断错误类型。
// // const (
// // 	errorCodeInvalidUserID = "INVALID_USER_ID"

// // 	errorCodeUserNotFound = "USER_NOT_FOUND"

// // 	errorCodeInternal = "INTERNAL_ERROR"
// // )

// // // ==================================================
// // // API Response Types
// // // ==================================================

// // // ErrorResponse：
// // //
// // // 今天升级为：
// // //
// // // code
// // // → machine-readable
// // //
// // // message
// // // → human-readable
// // //
// // // request_id
// // // → request correlation / debugging
// // type ErrorResponse struct {
// // 	Code      string `json:"code"`
// // 	Message   string `json:"message"`
// // 	RequestID string `json:"request_id"`
// // }

// // type UserResponse struct {
// // 	ID        int    `json:"id"`
// // 	Message   string `json:"message"`
// // 	RequestID string `json:"request_id"`
// // }

// // // ==================================================
// // // Business Errors
// // // ==================================================

// // var errUserNotFound = errors.New("user not found")

// // var errStorageUnavailable = errors.New("storage unavailable")

// // // ==================================================
// // // Response Helpers
// // // ==================================================

// // func writeJSONError(
// // 	w http.ResponseWriter,
// // 	r *http.Request,
// // 	status int,
// // 	code string,
// // 	message string,
// // ) {
// // 	requestID :=
// // 		requestIDFromContext(
// // 			r.Context(),
// // 		)

// // 	w.Header().Set(
// // 		"Content-Type",
// // 		"application/json",
// // 	)

// // 	w.WriteHeader(status)

// // 	response :=
// // 		ErrorResponse{
// // 			Code:      code,
// // 			Message:   message,
// // 			RequestID: requestID,
// // 		}

// // 	err :=
// // 		json.NewEncoder(w).Encode(
// // 			response,
// // 		)

// // 	if err != nil {
// // 		log.Printf(
// // 			"[%s] encode error response: %v",
// // 			requestID,
// // 			err,
// // 		)
// // 	}
// // }

// // func writeJSON(
// // 	w http.ResponseWriter,
// // 	status int,
// // 	value any,
// // ) {
// // 	w.Header().Set(
// // 		"Content-Type",
// // 		"application/json",
// // 	)

// // 	w.WriteHeader(status)

// // 	err :=
// // 		json.NewEncoder(w).Encode(
// // 			value,
// // 		)

// // 	if err != nil {
// // 		log.Println(
// // 			"encode response:",
// // 			err,
// // 		)
// // 	}
// // }

// // // ==================================================
// // // Request ID Middleware
// // // ==================================================

// // func requestIDMiddleware(
// // 	next http.Handler,
// // ) http.Handler {
// // 	return http.HandlerFunc(
// // 		func(
// // 			w http.ResponseWriter,
// // 			r *http.Request,
// // 		) {
// // 			requestID :=
// // 				r.Header.Get(
// // 					"X-Request-ID",
// // 				)

// // 			if requestID == "" {
// // 				requestID =
// // 					nextRequestID()
// // 			}

// // 			w.Header().Set(
// // 				"X-Request-ID",
// // 				requestID,
// // 			)

// // 			ctx :=
// // 				context.WithValue(
// // 					r.Context(),
// // 					requestIDKey,
// // 					requestID,
// // 				)

// // 			r =
// // 				r.WithContext(ctx)

// // 			log.Printf(
// // 				"[%s] request started: %s %s",
// // 				requestID,
// // 				r.Method,
// // 				r.URL.Path,
// // 			)

// // 			next.ServeHTTP(
// // 				w,
// // 				r,
// // 			)

// // 			log.Printf(
// // 				"[%s] request finished",
// // 				requestID,
// // 			)
// // 		},
// // 	)
// // }

// // // ==================================================
// // // Business Logic
// // // ==================================================

// // func findUserByID(
// // 	id int,
// // ) (UserResponse, error) {
// // 	if id == 404 {
// // 		return UserResponse{},
// // 			errUserNotFound
// // 	}

// // 	if id == 500 {
// // 		return UserResponse{},
// // 			errStorageUnavailable
// // 	}

// // 	return UserResponse{
// // 		ID:      id,
// // 		Message: "user found",
// // 	}, nil
// // }

// // // ==================================================
// // // Handler
// // // ==================================================

// // func userHandler(
// // 	w http.ResponseWriter,
// // 	r *http.Request,
// // ) {
// // 	requestID :=
// // 		requestIDFromContext(
// // 			r.Context(),
// // 		)

// // 	idText :=
// // 		r.PathValue("id")

// // 	// ------------------------------------------
// // 	// Parsing Error
// // 	// ------------------------------------------
// // 	id, err :=
// // 		strconv.Atoi(idText)

// // 	if err != nil {
// // 		log.Printf(
// // 			"[%s] invalid user id: %q",
// // 			requestID,
// // 			idText,
// // 		)

// // 		writeJSONError(
// // 			w,
// // 			r,
// // 			http.StatusBadRequest,
// // 			errorCodeInvalidUserID,
// // 			"user id must be an integer",
// // 		)

// // 		return
// // 	}

// // 	// ------------------------------------------
// // 	// Validation Error
// // 	// ------------------------------------------
// // 	if id <= 0 {
// // 		log.Printf(
// // 			"[%s] invalid user id: %d",
// // 			requestID,
// // 			id,
// // 		)

// // 		writeJSONError(
// // 			w,
// // 			r,
// // 			http.StatusBadRequest,
// // 			errorCodeInvalidUserID,
// // 			"user id must be greater than 0",
// // 		)

// // 		return
// // 	}

// // 	// ------------------------------------------
// // 	// Business Logic
// // 	// ------------------------------------------
// // 	user, err :=
// // 		findUserByID(id)

// // 	if err != nil {
// // 		if errors.Is(
// // 			err,
// // 			errUserNotFound,
// // 		) {
// // 			log.Printf(
// // 				"[%s] user not found: %d",
// // 				requestID,
// // 				id,
// // 			)

// // 			writeJSONError(
// // 				w,
// // 				r,
// // 				http.StatusNotFound,
// // 				errorCodeUserNotFound,
// // 				"user not found",
// // 			)

// // 			return
// // 		}

// // 		log.Printf(
// // 			"[%s] find user failed: %v",
// // 			requestID,
// // 			err,
// // 		)

// // 		writeJSONError(
// // 			w,
// // 			r,
// // 			http.StatusInternalServerError,
// // 			errorCodeInternal,
// // 			"internal server error",
// // 		)

// // 		return
// // 	}

// // 	// ------------------------------------------
// // 	// Success
// // 	// ------------------------------------------
// // 	user.RequestID =
// // 		requestID

// // 	writeJSON(
// // 		w,
// // 		http.StatusOK,
// // 		user,
// // 	)
// // }

// // func main() {
// // 	mux :=
// // 		http.NewServeMux()

// // 	mux.HandleFunc(
// // 		"GET /users/{id}",
// // 		userHandler,
// // 	)

// // 	handler :=
// // 		requestIDMiddleware(mux)

// // 	fmt.Println(
// // 		"server listening on http://localhost:8080",
// // 	)

// // 	err :=
// // 		http.ListenAndServe(
// // 			":8080",
// // 			handler,
// // 		)

// //		if err != nil {
// //			log.Fatal(err)
// //		}
// //	}
// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"sync/atomic"
// )

// // ==================================================
// // Context
// // ==================================================

// type contextKey string

// const requestIDKey contextKey = "requestID"

// // ==================================================
// // Request ID
// // ==================================================

// var requestCounter uint64

// func nextRequestID() string {
// 	number :=
// 		atomic.AddUint64(
// 			&requestCounter,
// 			1,
// 		)

// 	return "req-" +
// 		strconv.FormatUint(
// 			number,
// 			10,
// 		)
// }

// func requestIDFromContext(
// 	ctx context.Context,
// ) string {
// 	requestID, ok :=
// 		ctx.Value(
// 			requestIDKey,
// 		).(string)

// 	if !ok {
// 		return ""
// 	}

// 	return requestID
// }

// // ==================================================
// // Stable Error Codes
// // ==================================================

// const (
// 	errorCodeInvalidUserID = "INVALID_USER_ID"

// 	errorCodeUserNotFound = "USER_NOT_FOUND"

// 	errorCodeInternal = "INTERNAL_ERROR"
// )

// // ==================================================
// // API Error Model
// // ==================================================
// //
// // APIError：
// //
// // Server内部用于描述：
// //
// // “这个错误应该怎样表现为HTTP API错误”
// //
// // Status：
// // HTTP层错误类型。
// //
// // Code：
// // 稳定Machine-readable错误语义。
// //
// // Message：
// // Human-readable错误说明.
// //
// // 注意：
// //
// // APIError本身不是直接发送给Client的JSON。
// type APIError struct {
// 	Status  int
// 	Code    string
// 	Message string
// }

// // ==================================================
// // Predefined API Errors
// // ==================================================
// //
// // 把Status / Code / Message的固定组合
// // 集中定义在一个位置。
// //
// // Handler不再到处重复填写这些值。

// var apiErrorInvalidUserIDFormat = APIError{
// 	Status: http.StatusBadRequest,

// 	Code: errorCodeInvalidUserID,

// 	Message: "user id must be an integer",
// }

// var apiErrorInvalidUserIDValue = APIError{
// 	Status: http.StatusBadRequest,

// 	Code: errorCodeInvalidUserID,

// 	Message: "user id must be greater than 0",
// }

// var apiErrorUserNotFound = APIError{
// 	Status: http.StatusNotFound,

// 	Code: errorCodeUserNotFound,

// 	Message: "user not found",
// }

// var apiErrorInternal = APIError{
// 	Status: http.StatusInternalServerError,

// 	Code: errorCodeInternal,

// 	Message: "internal server error",
// }

// // ==================================================
// // Response Types
// // ==================================================
// //
// // ErrorResponse：
// //
// // 真正发送给Client的JSON结构。
// //
// // 注意：
// // HTTP Status已经存在于HTTP Response Status Line，
// // 所以这里不需要再把Status重复放进JSON。
// type ErrorResponse struct {
// 	Code      string `json:"code"`
// 	Message   string `json:"message"`
// 	RequestID string `json:"request_id"`
// }

// type UserResponse struct {
// 	ID        int    `json:"id"`
// 	Message   string `json:"message"`
// 	RequestID string `json:"request_id"`
// }

// // ==================================================
// // Business Errors
// // ==================================================
// //
// // Business Layer表达业务错误语义，
// // 不知道HTTP Status，也不知道JSON Response。

// var errUserNotFound = errors.New("user not found")

// var errStorageUnavailable = errors.New("storage unavailable")

// // ==================================================
// // Business Logic
// // ==================================================

// func findUserByID(
// 	id int,
// ) (UserResponse, error) {
// 	if id == 404 {
// 		return UserResponse{},
// 			errUserNotFound
// 	}

// 	if id == 500 {
// 		return UserResponse{},
// 			errStorageUnavailable
// 	}

// 	return UserResponse{
// 		ID:      id,
// 		Message: "user found",
// 	}, nil
// }

// // ==================================================
// // Business Error → API Error Mapping
// // ==================================================
// //
// // 这个函数集中决定：
// //
// // 某一种Business Error
// //
// // 应该对应：
// //
// // 哪个HTTP Status
// // 哪个Error Code
// // 哪个Client Message
// //
// // --------------------------------------------------
// //
// // 以后即使多个Handler都会遇到：
// //
// // errUserNotFound
// //
// // 也不需要分别重新决定：
// //
// // 404
// // USER_NOT_FOUND
// // user not found
// func mapBusinessError(
// 	err error,
// ) APIError {
// 	if errors.Is(
// 		err,
// 		errUserNotFound,
// 	) {
// 		return apiErrorUserNotFound
// 	}

// 	// 未明确识别的内部Business Error，
// 	// 默认映射成安全的500。
// 	return apiErrorInternal
// }

// // ==================================================
// // Response Helpers
// // ==================================================

// // writeAPIError：
// //
// // 接收一个已经完成映射的APIError，
// //
// // 再负责把它转换成真正的HTTP Response。
// func writeAPIError(
// 	w http.ResponseWriter,
// 	r *http.Request,
// 	apiError APIError,
// ) {
// 	requestID :=
// 		requestIDFromContext(
// 			r.Context(),
// 		)

// 	// Headers
// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	// Status
// 	w.WriteHeader(
// 		apiError.Status,
// 	)

// 	// Body
// 	response :=
// 		ErrorResponse{
// 			Code: apiError.Code,

// 			Message: apiError.Message,

// 			RequestID: requestID,
// 		}

// 	err :=
// 		json.NewEncoder(w).Encode(
// 			response,
// 		)

// 	if err != nil {
// 		log.Printf(
// 			"[%s] encode error response: %v",
// 			requestID,
// 			err,
// 		)
// 	}
// }

// func writeJSON(
// 	w http.ResponseWriter,
// 	status int,
// 	value any,
// ) {
// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	w.WriteHeader(status)

// 	err :=
// 		json.NewEncoder(w).Encode(
// 			value,
// 		)

// 	if err != nil {
// 		log.Println(
// 			"encode response:",
// 			err,
// 		)
// 	}
// }

// // ==================================================
// // Request ID Middleware
// // ==================================================

// func requestIDMiddleware(
// 	next http.Handler,
// ) http.Handler {
// 	return http.HandlerFunc(
// 		func(
// 			w http.ResponseWriter,
// 			r *http.Request,
// 		) {
// 			requestID :=
// 				r.Header.Get(
// 					"X-Request-ID",
// 				)

// 			if requestID == "" {
// 				requestID =
// 					nextRequestID()
// 			}

// 			w.Header().Set(
// 				"X-Request-ID",
// 				requestID,
// 			)

// 			ctx :=
// 				context.WithValue(
// 					r.Context(),
// 					requestIDKey,
// 					requestID,
// 				)

// 			r =
// 				r.WithContext(ctx)

// 			log.Printf(
// 				"[%s] request started: %s %s",
// 				requestID,
// 				r.Method,
// 				r.URL.Path,
// 			)

// 			next.ServeHTTP(
// 				w,
// 				r,
// 			)

// 			log.Printf(
// 				"[%s] request finished",
// 				requestID,
// 			)
// 		},
// 	)
// }

// // ==================================================
// // Handler
// // ==================================================

// func userHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	requestID :=
// 		requestIDFromContext(
// 			r.Context(),
// 		)

// 	idText :=
// 		r.PathValue("id")

// 	// ------------------------------------------
// 	// Parsing
// 	// ------------------------------------------

// 	id, err :=
// 		strconv.Atoi(idText)

// 	if err != nil {
// 		log.Printf(
// 			"[%s] invalid user id format: %q",
// 			requestID,
// 			idText,
// 		)

// 		writeAPIError(
// 			w,
// 			r,
// 			apiErrorInvalidUserIDFormat,
// 		)

// 		return
// 	}

// 	// ------------------------------------------
// 	// Validation
// 	// ------------------------------------------

// 	if id <= 0 {
// 		log.Printf(
// 			"[%s] invalid user id value: %d",
// 			requestID,
// 			id,
// 		)

// 		writeAPIError(
// 			w,
// 			r,
// 			apiErrorInvalidUserIDValue,
// 		)

// 		return
// 	}

// 	// ------------------------------------------
// 	// Business Logic
// 	// ------------------------------------------

// 	user, err :=
// 		findUserByID(id)

// 	if err != nil {
// 		// Server保存详细内部错误。
// 		log.Printf(
// 			"[%s] find user failed: %v",
// 			requestID,
// 			err,
// 		)

// 		// Business Error
// 		// ↓
// 		// API Error Mapping
// 		apiError :=
// 			mapBusinessError(err)

// 		// API Error
// 		// ↓
// 		// HTTP Response
// 		writeAPIError(
// 			w,
// 			r,
// 			apiError,
// 		)

// 		return
// 	}

// 	// ------------------------------------------
// 	// Success
// 	// ------------------------------------------

// 	user.RequestID =
// 		requestID

// 	writeJSON(
// 		w,
// 		http.StatusOK,
// 		user,
// 	)
// }

// func main() {
// 	mux :=
// 		http.NewServeMux()

// 	mux.HandleFunc(
// 		"GET /users/{id}",
// 		userHandler,
// 	)

// 	handler :=
// 		requestIDMiddleware(mux)

// 	fmt.Println(
// 		"server listening on http://localhost:8080",
// 	)

// 	err :=
// 		http.ListenAndServe(
// 			":8080",
// 			handler,
// 		)

//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

// ==================================================
// Context
// ==================================================

type contextKey string

const requestIDKey contextKey = "requestID"

// ==================================================
// Request ID
// ==================================================

var requestCounter uint64

func nextRequestID() string {
	number :=
		atomic.AddUint64(
			&requestCounter,
			1,
		)

	return "req-" +
		strconv.FormatUint(
			number,
			10,
		)
}

func requestIDFromContext(
	ctx context.Context,
) string {
	requestID, ok :=
		ctx.Value(
			requestIDKey,
		).(string)

	if !ok {
		return ""
	}

	return requestID
}

// ==================================================
// Stable Error Codes
// ==================================================

const (
	errorCodeInvalidUserID = "INVALID_USER_ID"

	errorCodeUserNotFound = "USER_NOT_FOUND"

	errorCodeInternal = "INTERNAL_ERROR"
)

// ==================================================
// API Error Model
// ==================================================
//
// APIError 是 Server 内部的 API 错误映射模型。
//
// 它描述：
//
// 某一种错误应该以什么方式表现成 HTTP API Error。
//
// 它不是直接发送给 Client 的 JSON。
type APIError struct {
	Status  int
	Code    string
	Message string
}

// ==================================================
// Predefined API Errors
// ==================================================
//
// 这里是 Error Contract 的单一规则来源。
//
// Handler 不应该分别重复维护：
//
// Status
// Code
// Message

var apiErrorInvalidUserIDFormat = APIError{
	Status: http.StatusBadRequest,

	Code: errorCodeInvalidUserID,

	Message: "user id must be an integer",
}

var apiErrorInvalidUserIDValue = APIError{
	Status: http.StatusBadRequest,

	Code: errorCodeInvalidUserID,

	Message: "user id must be greater than 0",
}

var apiErrorUserNotFound = APIError{
	Status: http.StatusNotFound,

	Code: errorCodeUserNotFound,

	Message: "user not found",
}

var apiErrorInternal = APIError{
	Status: http.StatusInternalServerError,

	Code: errorCodeInternal,

	Message: "internal server error",
}

// ==================================================
// Response Types
// ==================================================
//
// ErrorResponse 才是真正的 Client JSON Contract。
//
// RequestID 属于某一次具体 Request，
// 所以不应该放进预定义 APIError 中。
type ErrorResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// ==================================================
// Business Errors
// ==================================================
//
// Business Error：
//
// 只表达服务内部业务语义。
//
// 它们不知道：
//
// HTTP Status
// API Error Code
// JSON Response

var errUserNotFound = errors.New("user not found")

var errStorageUnavailable = errors.New("storage unavailable")

// ==================================================
// Business Logic
// ==================================================

func findUserByID(
	id int,
) (UserResponse, error) {
	if id == 404 {
		return UserResponse{},
			errUserNotFound
	}

	if id == 500 {
		return UserResponse{},
			errStorageUnavailable
	}

	return UserResponse{
		ID:      id,
		Message: "user found",
	}, nil
}

// ==================================================
// Error Mapping
// ==================================================
//
// mapBusinessError：
//
// 只负责：
//
// Business Error
// ↓
// APIError
//
// 它负责 Error Policy，
// 不负责真正写 HTTP Response。
func mapBusinessError(
	err error,
) APIError {
	if errors.Is(
		err,
		errUserNotFound,
	) {
		return apiErrorUserNotFound
	}

	// 未明确识别的内部错误：
	//
	// 使用安全的500 fallback。
	return apiErrorInternal
}

// ==================================================
// Response Rendering
// ==================================================
//
// writeAPIError：
//
// 只负责：
//
// APIError
// ↓
// HTTP Response
//
// 它不负责判断Business Error是什么。
func writeAPIError(
	w http.ResponseWriter,
	r *http.Request,
	apiError APIError,
) {
	requestID :=
		requestIDFromContext(
			r.Context(),
		)

	// ------------------------------------------
	// Header
	// ------------------------------------------
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// ------------------------------------------
	// Status
	// ------------------------------------------
	w.WriteHeader(
		apiError.Status,
	)

	// ------------------------------------------
	// Body
	// ------------------------------------------
	response :=
		ErrorResponse{
			Code: apiError.Code,

			Message: apiError.Message,

			RequestID: requestID,
		}

	err :=
		json.NewEncoder(w).Encode(
			response,
		)

	if err != nil {
		log.Printf(
			"[%s] encode error response: %v",
			requestID,
			err,
		)
	}
}

// writeBusinessError：
//
// 这是一个很薄的组合Helper。
//
// 它没有定义新的Error Policy，
//
// 只是把已有的：
//
// mapBusinessError
// +
// writeAPIError
//
// 组合成：
//
// Business Error
// ↓
// API Error Mapping
// ↓
// HTTP Response
func writeBusinessError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	apiError :=
		mapBusinessError(err)

	writeAPIError(
		w,
		r,
		apiError,
	)
}

// writeJSON：
//
// 统一成功JSON Response。
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
		json.NewEncoder(w).Encode(
			value,
		)

	if err != nil {
		log.Println(
			"encode response:",
			err,
		)
	}
}

// ==================================================
// Request ID Middleware
// ==================================================

func requestIDMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			requestID :=
				r.Header.Get(
					"X-Request-ID",
				)

			if requestID == "" {
				requestID =
					nextRequestID()
			}

			w.Header().Set(
				"X-Request-ID",
				requestID,
			)

			ctx :=
				context.WithValue(
					r.Context(),
					requestIDKey,
					requestID,
				)

			r =
				r.WithContext(ctx)

			log.Printf(
				"[%s] request started: %s %s",
				requestID,
				r.Method,
				r.URL.Path,
			)

			next.ServeHTTP(
				w,
				r,
			)

			log.Printf(
				"[%s] request finished",
				requestID,
			)
		},
	)
}

// ==================================================
// Handler
// ==================================================

func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	requestID :=
		requestIDFromContext(
			r.Context(),
		)

	idText :=
		r.PathValue("id")

	// ------------------------------------------
	// 1. Parsing
	// ------------------------------------------

	id, err :=
		strconv.Atoi(idText)

	if err != nil {
		log.Printf(
			"[%s] invalid user id format: %q",
			requestID,
			idText,
		)

		writeAPIError(
			w,
			r,
			apiErrorInvalidUserIDFormat,
		)

		return
	}

	// ------------------------------------------
	// 2. Validation
	// ------------------------------------------

	if id <= 0 {
		log.Printf(
			"[%s] invalid user id value: %d",
			requestID,
			id,
		)

		writeAPIError(
			w,
			r,
			apiErrorInvalidUserIDValue,
		)

		return
	}

	// ------------------------------------------
	// 3. Business Logic
	// ------------------------------------------

	user, err :=
		findUserByID(id)

	if err != nil {
		// Server内部保留真实Error。
		log.Printf(
			"[%s] find user failed: %v",
			requestID,
			err,
		)

		// Handler不再关心：
		//
		// 这个Business Error具体应该：
		//
		// 404还是500？
		// 什么Code？
		// 什么Message？
		// 怎样Encode JSON？
		//
		// 统一交给Error Pipeline。
		writeBusinessError(
			w,
			r,
			err,
		)

		return
	}

	// ------------------------------------------
	// 4. Success
	// ------------------------------------------

	user.RequestID =
		requestID

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

	handler :=
		requestIDMiddleware(mux)

	fmt.Println(
		"server listening on http://localhost:8080",
	)

	err :=
		http.ListenAndServe(
			":8080",
			handler,
		)

	if err != nil {
		log.Fatal(err)
	}
}
