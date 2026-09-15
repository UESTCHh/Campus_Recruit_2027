// // // 文件职责：
// // // Request
// // // ↓
// // // Request ID Middleware
// // // ↓
// // // 读取X-Request-ID
// // // ↓
// // // 必要时生成Request ID
// // // ↓
// // // 写入Response Header
// // // ↓
// // // Logging
// // // ↓
// // // next.ServeHTTP
// // // ↓
// // // Handler
// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"log"
// // 	"net/http"
// // 	"strconv"
// // 	"sync/atomic"
// // )

// // // requestCounter：
// // //
// // // 用来给本次学习实验生成简单的Request ID。
// // //
// // // 例如：
// // //
// // // req-1
// // // req-2
// // // req-3
// // //
// // // HTTP Server可能同时处理多个Request，
// // // 所以不能简单地：
// // //
// // // requestCounter++
// // //
// // // 因为多个goroutine可能同时修改同一个变量。
// // //
// // // atomic.AddUint64可以安全地完成：
// // //
// // // 读取旧值
// // // +
// // // 加1
// // // +
// // // 写回
// // //
// // // 这里先把它理解为：
// // //
// // // “并发安全地得到下一个编号”。
// // //
// // // atomic我们以前学习Data Race时已经接触过并发问题，
// // // 今天不展开atomic内部实现。
// // var requestCounter uint64

// // // HelloResponse：
// // //
// // // Handler成功时返回的JSON。
// // type HelloResponse struct {
// // 	Message string `json:"message"`
// // }

// // // nextRequestID：
// // //
// // // 生成服务器自己的Request ID。
// // //
// // // 当前只是教学版本：
// // //
// // // req-1
// // // req-2
// // // req-3
// // //
// // // 真实生产系统可能使用：
// // //
// // // UUID
// // // Trace ID
// // // 分布式Tracing系统生成的ID
// // //
// // // 今天先不引入额外库。
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

// // // requestIDMiddleware：
// // //
// // // Middleware签名仍然是我们之前学习过的：
// // //
// // // func(http.Handler) http.Handler
// // //
// // // 它负责：
// // //
// // // 1. 读取Request Header中的X-Request-ID
// // //
// // // 2. 如果Client没有提供，Server自己生成
// // //
// // // 3. 把Request ID写回Response Header
// // //
// // // 4. 记录当前Request ID
// // //
// // // 5. 调用next.ServeHTTP继续处理当前Request
// // //
// // // --------------------------------------------------
// // //
// // // 今天第一阶段还没有把Request ID写入Context。
// // //
// // // 所以当前重点只是：
// // //
// // // Request ID
// // // +
// // // Header
// // // +
// // // Middleware
// // func requestIDMiddleware(
// // 	next http.Handler,
// // ) http.Handler {
// // 	return http.HandlerFunc(
// // 		func(
// // 			w http.ResponseWriter,
// // 			r *http.Request,
// // 		) {
// // 			// ------------------------------------------
// // 			// 1. 尝试读取Client提供的Request ID
// // 			// ------------------------------------------
// // 			//
// // 			// Header读取方式继续复习：
// // 			//
// // 			// r.Header.Get(...)
// // 			requestID :=
// // 				r.Header.Get(
// // 					"X-Request-ID",
// // 				)

// // 			// ------------------------------------------
// // 			// 2. Client没有提供
// // 			// ------------------------------------------
// // 			//
// // 			// Server自己生成一个。
// // 			if requestID == "" {
// // 				requestID =
// // 					nextRequestID()
// // 			}

// // 			// ------------------------------------------
// // 			// 3. Response也返回Request ID
// // 			// ------------------------------------------
// // 			//
// // 			// 注意：
// // 			//
// // 			// Header必须在Response Body开始写出之前设置。
// // 			//
// // 			// 这正好连接昨天HTTP Error Handling学习的：
// // 			//
// // 			// Headers
// // 			// ↓
// // 			// Status
// // 			// ↓
// // 			// Body
// // 			w.Header().Set(
// // 				"X-Request-ID",
// // 				requestID,
// // 			)

// // 			// ------------------------------------------
// // 			// 4. Server Log
// // 			// ------------------------------------------
// // 			log.Printf(
// // 				"[%s] request started: %s %s",
// // 				requestID,
// // 				r.Method,
// // 				r.URL.Path,
// // 			)

// // 			// ------------------------------------------
// // 			// 5. 继续交给下一层Handler
// // 			// ------------------------------------------
// // 			//
// // 			// 如果next是mux：
// // 			//
// // 			// next.ServeHTTP
// // 			// ↓
// // 			// mux.ServeHTTP
// // 			// ↓
// // 			// Route Matching
// // 			// ↓
// // 			// helloHandler
// // 			next.ServeHTTP(
// // 				w,
// // 				r,
// // 			)

// // 			// ------------------------------------------
// // 			// 6. 下游处理完成后
// // 			// ------------------------------------------
// // 			//
// // 			// 这里再次体现Middleware Onion Model：
// // 			//
// // 			// before
// // 			// ↓
// // 			// next
// // 			// ↓
// // 			// after
// // 			log.Printf(
// // 				"[%s] request finished",
// // 				requestID,
// // 			)
// // 		},
// // 	)
// // }

// // // helloHandler：
// // //
// // // 今天第一阶段故意保持简单。
// // //
// // // Handler暂时还不知道requestID。
// // //
// // // 第二阶段再通过Context把requestID传进来。
// // func helloHandler(
// // 	w http.ResponseWriter,
// // 	r *http.Request,
// // ) {
// // 	response :=
// // 		HelloResponse{
// // 			Message: "hello",
// // 		}

// // 	w.Header().Set(
// // 		"Content-Type",
// // 		"application/json",
// // 	)

// // 	err :=
// // 		json.NewEncoder(w).Encode(
// // 			response,
// // 		)

// // 	if err != nil {
// // 		log.Println(
// // 			"encode response:",
// // 			err,
// // 		)
// // 	}
// // }

// // func main() {
// // 	// 创建Router。
// // 	mux :=
// // 		http.NewServeMux()

// // 	mux.HandleFunc(
// // 		"GET /hello",
// // 		helloHandler,
// // 	)

// // 	// --------------------------------------------------
// // 	// Middleware包装mux
// // 	// --------------------------------------------------
// // 	//
// // 	// 实际结构：
// // 	//
// // 	// requestIDMiddleware(mux)
// // 	//
// // 	// 所以Request进入：
// // 	//
// // 	// Client
// // 	// ↓
// // 	// requestIDMiddleware
// // 	// ↓
// // 	// mux
// // 	// ↓
// // 	// Route
// // 	// ↓
// // 	// helloHandler
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
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"sync/atomic"
// )

// // contextKey：
// //
// // 自定义Context Key类型。
// //
// // 不直接使用普通string作为Context Key，
// // 可以降低不同模块使用相同字符串Key时发生冲突的风险。
// type contextKey string

// // requestIDKey：
// //
// // 用于在Request Context中保存Request ID。
// const requestIDKey contextKey = "requestID"

// // requestCounter：
// //
// // 用于生成当前教学程序中的简单递增Request ID。
// //
// // 多个HTTP Request可能由不同goroutine并发处理，
// // 因此通过atomic安全递增。
// var requestCounter uint64

// // HelloResponse：
// //
// // 成功Response。
// //
// // 今天第二阶段新增RequestID字段，
// // 用于直接验证Handler已经能从Context中读取Request ID。
// type HelloResponse struct {
// 	Message   string `json:"message"`
// 	RequestID string `json:"request_id"`
// }

// // nextRequestID：
// //
// // 生成：
// //
// // req-1
// // req-2
// // req-3
// //
// // 这样的简单教学Request ID。
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

// // requestIDFromContext：
// //
// // 从Context中读取Request ID。
// //
// // Context.Value返回any，
// // 因此需要Type Assertion转换为string。
// //
// // 使用：
// //
// // requestID, ok := value.(string)
// //
// // 而不是直接：
// //
// // value.(string)
// //
// // 是为了在值不存在或类型不正确时避免panic。
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

// // requestIDMiddleware：
// //
// // 当前完整职责：
// //
// // 1. 从Request Header读取X-Request-ID
// //
// // 2. Client没有提供时Server生成
// //
// // 3. 写入Response Header
// //
// // 4. 将Request ID写入Context
// //
// // 5. 创建携带新Context的Request
// //
// // 6. 将Request继续交给下一层
// //
// // 7. 使用同一个Request ID记录Middleware日志
// func requestIDMiddleware(
// 	next http.Handler,
// ) http.Handler {
// 	return http.HandlerFunc(
// 		func(
// 			w http.ResponseWriter,
// 			r *http.Request,
// 		) {
// 			// ------------------------------------------
// 			// 1. 获取Request ID
// 			// ------------------------------------------
// 			requestID :=
// 				r.Header.Get(
// 					"X-Request-ID",
// 				)

// 			if requestID == "" {
// 				requestID =
// 					nextRequestID()
// 			}

// 			// ------------------------------------------
// 			// 2. 设置Response Header
// 			// ------------------------------------------
// 			//
// 			// 需要在下游Handler开始写Response之前设置。
// 			w.Header().Set(
// 				"X-Request-ID",
// 				requestID,
// 			)

// 			// ------------------------------------------
// 			// 3. 将Request ID放入Context
// 			// ------------------------------------------
// 			//
// 			// r.Context()
// 			// → 当前Request原Context
// 			//
// 			// context.WithValue(...)
// 			// → 创建Child Context
// 			//
// 			// Child Context额外保存：
// 			//
// 			// requestIDKey -> requestID
// 			ctx :=
// 				context.WithValue(
// 					r.Context(),
// 					requestIDKey,
// 					requestID,
// 				)

// 			// ------------------------------------------
// 			// 4. 创建携带新Context的Request
// 			// ------------------------------------------
// 			//
// 			// WithContext不会直接修改原Request本身，
// 			// 而是返回一个携带新Context的Request副本。
// 			r =
// 				r.WithContext(ctx)

// 			// ------------------------------------------
// 			// 5. Middleware Before
// 			// ------------------------------------------
// 			log.Printf(
// 				"[%s] request started: %s %s",
// 				requestID,
// 				r.Method,
// 				r.URL.Path,
// 			)

// 			// ------------------------------------------
// 			// 6. 继续交给下一层
// 			// ------------------------------------------
// 			//
// 			// 注意：
// 			//
// 			// 此时传下去的r，
// 			// 已经携带包含Request ID的Context。
// 			next.ServeHTTP(
// 				w,
// 				r,
// 			)

// 			// ------------------------------------------
// 			// 7. Middleware After
// 			// ------------------------------------------
// 			log.Printf(
// 				"[%s] request finished",
// 				requestID,
// 			)
// 		},
// 	)
// }

// // helloHandler：
// //
// // 第二阶段的核心变化：
// //
// // Handler现在不依赖Middleware中的局部变量。
// //
// // 它只从：
// //
// // r.Context()
// //
// // 中读取当前Request自己的Request ID。
// func helloHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	// ------------------------------------------
// 	// 1. 从Context读取Request ID
// 	// ------------------------------------------
// 	requestID :=
// 		requestIDFromContext(
// 			r.Context(),
// 		)

// 	// ------------------------------------------
// 	// 2. Handler自己的日志
// 	// ------------------------------------------
// 	//
// 	// 现在Middleware和Handler都能使用
// 	// 同一个Request ID。
// 	log.Printf(
// 		"[%s] hello handler",
// 		requestID,
// 	)

// 	// ------------------------------------------
// 	// 3. 构造Response
// 	// ------------------------------------------
// 	response :=
// 		HelloResponse{
// 			Message:   "hello",
// 			RequestID: requestID,
// 		}

// 	// ------------------------------------------
// 	// 4. 设置Response Content-Type
// 	// ------------------------------------------
// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	// ------------------------------------------
// 	// 5. 写JSON Body
// 	// ------------------------------------------
// 	err :=
// 		json.NewEncoder(w).Encode(
// 			response,
// 		)

// 	if err != nil {
// 		log.Printf(
// 			"[%s] encode response: %v",
// 			requestID,
// 			err,
// 		)
// 	}
// }

// func main() {
// 	mux :=
// 		http.NewServeMux()

// 	mux.HandleFunc(
// 		"GET /hello",
// 		helloHandler,
// 	)

// 	// 最终调用结构：
// 	//
// 	// requestIDMiddleware(mux)
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

// contextKey 是自定义的 Context Key 类型。
//
// 避免直接使用普通 string 作为 Key，
// 降低不同模块之间发生 Key Collision 的风险。
type contextKey string

const requestIDKey contextKey = "requestID"

// requestCounter 用于生成教学阶段的简单递增 Request ID。
//
// HTTP Server 会并发处理多个 Request，
// 所以使用 atomic 保证计数器递增是并发安全的。
var requestCounter uint64

// ==================================================
// API Response Types
// ==================================================

// ErrorResponse：
//
// 所有 API Error 使用统一的 JSON 结构。
//
// 今天新增 RequestID：
//
//	{
//	    "error": "...",
//	    "request_id": "req-1"
//	}
//
// 这样客户端拿到 Request ID 后，
// 可以和 Server Log 中的同一 Request 关联起来。
type ErrorResponse struct {
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}

// UserResponse：
//
// 查询 User 成功时返回。
type UserResponse struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// ==================================================
// Business Errors
// ==================================================

// User 不存在。
var errUserNotFound = errors.New("user not found")

// 模拟 Storage / Database 等服务器内部依赖失败。
var errStorageUnavailable = errors.New("storage unavailable")

// ==================================================
// Request ID Helpers
// ==================================================

// nextRequestID：
//
// 生成：
//
// req-1
// req-2
// req-3
//
// 这样的简单 Request ID。
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

// requestIDFromContext：
//
// 从 Context 中读取当前 Request 自己的 Request ID。
//
// Context.Value 返回 any，
// 所以通过 Type Assertion 转换成 string。
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
// Response Helpers
// ==================================================

// writeJSONError：
//
// 统一处理 Error Response。
//
// 今天它比 Day20 多了一项职责：
//
// 从当前 Request Context 中读取 Request ID，
// 并把它放进 Error JSON。
func writeJSONError(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
) {
	requestID :=
		requestIDFromContext(
			r.Context(),
		)

	// Response Headers 必须在 Response Body
	// 开始写出之前设置。
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// 显式写 HTTP Status Code。
	w.WriteHeader(status)

	response :=
		ErrorResponse{
			Error:     message,
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

// writeJSON：
//
// 成功 Response 统一 JSON 输出。
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

// requestIDMiddleware：
//
// 1. 读取 Client X-Request-ID
// 2. 没有则 Server 生成
// 3. 写回 Response Header
// 4. 保存进 Context
// 5. 创建携带新 Context 的 Request
// 6. 继续交给下游
// 7. 使用同一个 ID 串联日志
func requestIDMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			// ------------------------------------------
			// 1. 获取 Request ID
			// ------------------------------------------
			requestID :=
				r.Header.Get(
					"X-Request-ID",
				)

			if requestID == "" {
				requestID =
					nextRequestID()
			}

			// ------------------------------------------
			// 2. 写回 Response Header
			// ------------------------------------------
			w.Header().Set(
				"X-Request-ID",
				requestID,
			)

			// ------------------------------------------
			// 3. Request ID → Context
			// ------------------------------------------
			ctx :=
				context.WithValue(
					r.Context(),
					requestIDKey,
					requestID,
				)

			// ------------------------------------------
			// 4. Context → Request
			// ------------------------------------------
			r =
				r.WithContext(ctx)

			// ------------------------------------------
			// 5. Middleware Before
			// ------------------------------------------
			log.Printf(
				"[%s] request started: %s %s",
				requestID,
				r.Method,
				r.URL.Path,
			)

			// ------------------------------------------
			// 6. 进入下游
			// ------------------------------------------
			next.ServeHTTP(
				w,
				r,
			)

			// ------------------------------------------
			// 7. Middleware After
			// ------------------------------------------
			log.Printf(
				"[%s] request finished",
				requestID,
			)
		},
	)
}

// ==================================================
// Simulated Business Logic
// ==================================================

// findUserByID：
//
// 当前仍然使用特殊 ID 模拟不同结果：
//
// 404
// → Resource不存在
//
// 500
// → Server内部Storage故障
//
// 其他正整数
// → Success
//
// 注意：
//
// Business Function 不知道 HTTP 404 / 500。
// 它只返回业务结果或者 error。
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
// HTTP Handler
// ==================================================

func userHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 当前 Handler 从 Context 中得到 Request ID。
	requestID :=
		requestIDFromContext(
			r.Context(),
		)

	log.Printf(
		"[%s] user handler started",
		requestID,
	)

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
		log.Printf(
			"[%s] invalid user id: %q",
			requestID,
			idText,
		)

		writeJSONError(
			w,
			r,
			http.StatusBadRequest,
			"user id must be an integer",
		)

		return
	}

	// ------------------------------------------
	// 3. Validation
	// ------------------------------------------
	if id <= 0 {
		log.Printf(
			"[%s] user id validation failed: %d",
			requestID,
			id,
		)

		writeJSONError(
			w,
			r,
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
			log.Printf(
				"[%s] user not found: %d",
				requestID,
				id,
			)

			writeJSONError(
				w,
				r,
				http.StatusNotFound,
				"user not found",
			)

			return
		}

		// --------------------------------------
		// Internal Server Error
		// --------------------------------------
		//
		// Server Log：
		//
		// 保存详细内部错误。
		log.Printf(
			"[%s] find user failed: %v",
			requestID,
			err,
		)

		// Client：
		//
		// 只得到安全、稳定的错误信息。
		writeJSONError(
			w,
			r,
			http.StatusInternalServerError,
			"internal server error",
		)

		return
	}

	// ------------------------------------------
	// 5. Success Response
	// ------------------------------------------
	//
	// 把同一个 Request ID 也放入成功 Response，
	// 方便观察整个 Request 生命周期。
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

	// 实际结构：
	//
	// requestIDMiddleware(mux)
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
