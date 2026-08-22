package main

import "fmt"

// User 定义用户结构体（类似于其他语言中的类）
type User struct {
	Name  string // 首字母大写 = 公开字段（可被外部包访问）
	Email string
	age   int // 首字母小写 = 私有字段（仅包内可访问）
}

// Greet 是为 User 类型定义的方法
// (u User) 是值接收者，方法内对 u 的修改不会影响原始值
func (u User) Greet() string {
	u.age++ // 不影响实例属性
	return fmt.Sprintf("你好，我是 %s，邮箱是 %s", u.Name, u.Email)
}

// Birthday 使用指针接收者，方法内可以修改原始值
func (u *User) Birthday() {
	u.age++ // 通过指针直接修改原始结构体的 age 字段
}

func main() {
	// 创建结构体实例
	user := User{
		Name:  "runoob",
		Email: "runoob@example.com",
		age:   25,
	}

	// 调用方法
	fmt.Println(user.Greet())
	fmt.Printf("生日前年龄: %d\n", user.age)

	user.Birthday() // 使用指针接收者，age 被加 1
	fmt.Printf("生日后年龄: %d\n", user.age)
}
