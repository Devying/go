package main

import "fmt"

type user struct {
	name  string
	email string
}
type admin struct {
	user
	level string
}

func main() {
	ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员"}
	ad.email = "huangby@imoo.com"
	fmt.Println("可以直接调用,名字为：", ad.name)
	fmt.Println("可以直接调用,邮箱为：", ad.email)
	fmt.Println("也可以通过内部类型调用,名字为：", ad.user.name)
	fmt.Println("但是新增加的属性只能直接调用，级别为：", ad.level)
	ad.sayHello()
}
func (u user) sayHello() {
	fmt.Println("Hello，i am a user")
}
func (a admin) sayHello() {
	fmt.Println("Hello，i am a admin")
}
