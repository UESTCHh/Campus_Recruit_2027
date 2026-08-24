package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Print() {
	fmt.Println(
		"name:",
		u.Name,
		"age:",
		u.Age,
	)
}

// func (u User) ChangeName() {
// 	u.Name = "new name"
// }

func (u *User) ChangeName() {
	u.Name = "new name"
}

func main() {
	user := User{
		Name: "hui",
		Age:  22,
	}
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)

	user.Print()

	user.ChangeName()

	fmt.Println(user.Name)

	service := UserService{}

	service.Add()

	fmt.Println(service.count)

	service.AddPointer()

	fmt.Println(service.count)
}
