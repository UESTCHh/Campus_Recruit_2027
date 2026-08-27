// 文件作用
//
//	定义：
//	ValidationError
//
// 让错误不仅有文本，还能携带结构化字段。
package main

import "fmt"

type ValidationError struct {
	Field   string
	Value   int
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(
		"validation failed: field = %s value = %d message = %s",
		e.Field,
		e.Value,
		e.Message,
	)
}

func ValidateUserID(id int) error {
	if id <= 0 {
		return &ValidationError{
			Field:   "user_id",
			Value:   id,
			Message: "must be positive",
		}
	}

	return nil
}
