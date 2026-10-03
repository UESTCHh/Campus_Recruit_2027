package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type contextKey string

const requestIDKey contextKey = "requestID"

// 程序级共享计数器。
// 多个HTTP请求会并发访问，因此使用atomic操作。
var requestCounter uint64

// requestIDFromContext统一负责：
// 从Context中安全读取Request ID。
//
// 调用方不需要关心具体Context Key和类型断言细节。
func requestIDFromContext(
	ctx context.Context,
) (string, bool) {
	requestID, ok :=
		ctx.Value(
			requestIDKey,
		).(string)

	return requestID, ok
}

func requestIDMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			id := atomic.AddUint64(
				&requestCounter,
				1,
			)

			requestID :=
				fmt.Sprintf(
					"req-%d",
					id,
				)

			log.Printf(
				"[%s] middleware: start",
				requestID,
			)

			// 将Request ID放入当前请求的Context。
			ctx := context.WithValue(
				r.Context(),
				requestIDKey,
				requestID,
			)

			r = r.WithContext(ctx)

			// 将Request ID同时返回给客户端。
			//
			// 必须尽量在Handler写响应之前设置Header。
			w.Header().Set(
				"X-Request-ID",
				requestID,
			)

			next.ServeHTTP(
				w,
				r,
			)

			log.Printf(
				"[%s] middleware: finish",
				requestID,
			)
		},
	)
}

func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	requestID, ok :=
		requestIDFromContext(
			r.Context(),
		)

	if !ok {
		http.Error(
			w,
			"request id not found",
			http.StatusInternalServerError,
		)
		return
	}

	log.Printf(
		"[%s] handler: start",
		requestID,
	)

	time.Sleep(
		2 * time.Second,
	)

	log.Printf(
		"[%s] handler: finish",
		requestID,
	)

	fmt.Fprintf(
		w,
		"hello\n",
	)
}

func main() {
	mux := http.NewServeMux()

	businessHandler :=
		http.HandlerFunc(
			helloHandler,
		)

	wrappedHandler :=
		requestIDMiddleware(
			businessHandler,
		)

	mux.Handle(
		"/hello",
		wrappedHandler,
	)

	log.Println(
		"server started at http://localhost:8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			mux,
		),
	)
}
