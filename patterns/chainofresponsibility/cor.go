package main

import "fmt"

type MiddlewareInterface interface {
	Handle()
	Next()
	SetNext(middleware *MiddlewareInterface)
}

type MiddlewareBase struct {
	next *MiddlewareInterface
}

func (m *MiddlewareBase) SetNext(middleware *MiddlewareInterface) {
	m.next = middleware
}
func (m *MiddlewareBase) Next() {
	if m.next != nil {
		(*m.next).Handle()
	}
}

type UserExistMiddleware struct {
	MiddlewareBase
	user User
}

func (m *UserExistMiddleware) Handle() {
	if m.user.email == "" {
		fmt.Println("This email is not registered!")
		return
	}
	if m.user.password == "" {
		fmt.Println("Wrong password!")
		return
	}
	fmt.Println("User exist")
	m.Next()
}

type AdminMiddleware struct {
	MiddlewareBase
	user User
}

func (m *AdminMiddleware) Handle() {
	if m.user.email != "admin@mail.ru" {
		fmt.Println("Dont have admin rights")
		return
	}
	fmt.Println("Hello admin")
	m.Next()
}

type User struct {
	email, password string
}

func main() {
	user := User{
		email:    "admin@mail.ru",
		password: "12345",
	}
	userMiddleware := UserExistMiddleware{
		user: user,
	}
	adminMiddleware := AdminMiddleware{
		user: user,
	}

	var test MiddlewareInterface = &adminMiddleware
	userMiddleware.SetNext(&test)
	userMiddleware.Handle()
}
