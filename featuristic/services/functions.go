package services

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
}

type FunctionPackage struct{}

func (fp FunctionPackage) HelloWorld(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func (fp FunctionPackage) Add(a, b int) int {
	return a + b
}

func (fp FunctionPackage) UserGreeting(user User) string {
	return fmt.Sprintf("Hello, %s %s, age %d!", user.FirstName, user.LastName, user.Age)
}

func (fp FunctionPackage) Multiply(a, b int) int {
	return a * b
}
