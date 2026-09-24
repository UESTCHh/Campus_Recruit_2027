package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// CreateUserRequest表示Client提交的用户数据。
type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreateUserResponse表示Server成功处理后返回的数据。
type CreateUserResponse struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Message string `json:"message"`
}

func createUserHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 1. Decode：
	// JSON Request Body -> Go struct
	var req CreateUserRequest

	err :=
		json.NewDecoder(r.Body).
			Decode(&req)

	if err != nil {
		http.Error(
			w,
			"invalid json body",
			http.StatusBadRequest,
		)
		return
	}

	// 2. Validation：
	// Parse成功不代表业务数据一定合法。
	if req.Name == "" {
		http.Error(
			w,
			"name is required",
			http.StatusBadRequest,
		)
		return
	}

	if req.Age <= 0 {
		http.Error(
			w,
			"age must be greater than 0",
			http.StatusBadRequest,
		)
		return
	}

	// 3. Business Logic：
	// 当前还没有数据库，
	// 暂时模拟“创建用户成功”。
	resp := CreateUserResponse{
		Name:    req.Name,
		Age:     req.Age,
		Message: "user created",
	}

	// 4. Response Header
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	// 5. Response Status
	w.WriteHeader(
		http.StatusCreated,
	)

	// 6. Response Body：
	// Go struct -> JSON
	err =
		json.NewEncoder(w).
			Encode(resp)

	if err != nil {
		log.Println(
			"failed to encode response:",
			err,
		)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/users",
		createUserHandler,
	)

	log.Println(
		"starting server at http://localhost:8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			mux,
		),
	)
}
