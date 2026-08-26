// 模拟后端service调用底层函数
// 并给底层错误增加业务上下文
package main

import "fmt"

func GetUserDisplayName(id int) (string, error) {
	name, err := GetUserName(id)

	if err != nil {
		return "", fmt.Errorf("get user display name: %w", err)
	}

	return "User:" + name, nil
}
