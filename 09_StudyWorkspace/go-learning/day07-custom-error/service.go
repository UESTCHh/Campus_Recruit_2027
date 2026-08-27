// 文件作用
// 模拟 Service 调用参数校验，并继续使用 %w 包装
package main

import "fmt"

func RegisterUser(id int) error {
	err := ValidateUserID(id)

	if err != nil {
		return fmt.Errorf(
			"register user: %w",
			err,
		)
	}

	return nil
}
