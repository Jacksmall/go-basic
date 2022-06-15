package main

import "fmt"

type notifer interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 嵌套类型
// 外部类型 内部类型实现的方法会提升到外部类型
type admin struct {
	user  // 内部类型
	level string
}

// func (u *admin) notify() {
// 	fmt.Printf("Sending admin email to %s<%s>\n", u.name, u.email)
// }

func main() {
	ad := admin{user: user{
		name:  "adminUser",
		email: "admin@example.com",
	}, level: "adminLevel"}
	sendNotification(&ad)
}

func sendNotification(u notifer) {
	u.notify()
}
