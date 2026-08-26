// 模拟一个后端业务函数：
//
//	根据id查询用户名
package main

// Go内置的error本质上是一个interface
import "errors"

var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrUserNotFound  = errors.New("user not found")
)

func GetUserName(id int) (string, error) {
	if id <= 0 {
		return "", ErrInvalidUserID
	}

	if id == 1 {
		return "UESTCHh", nil
	}

	return "", ErrUserNotFound
}
