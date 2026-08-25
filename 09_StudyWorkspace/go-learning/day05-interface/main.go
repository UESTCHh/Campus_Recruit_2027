package main

func main() {
	email := EmailNotifier{}
	service := NewUserService(email)

	service.Register("UESTCHh")
}
