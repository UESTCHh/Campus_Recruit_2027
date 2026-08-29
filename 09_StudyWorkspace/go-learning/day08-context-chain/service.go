// 文件作用
// 模拟业务层调用 Repository，并继续向下传递 Context。
package main

import (
	"context"
	"fmt"
)

func GetUserProfile(
	ctx context.Context,
	id int,
) (string, error) {
	fmt.Println("service: get user profile")

	name, err := FindUser(ctx, id)

	if err != nil {
		return "", fmt.Errorf(
			"get user profile: %w",
			err,
		)
	}

	return "profile of " + name, nil
}
