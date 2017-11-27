package main

import (
	"fmt"
)
//空接口类型interface{}一个方法签名也不包含，所以所有的数据类型都实现了该方法。
//空接口类型在描述一个对象实例的行为上力不从心，但是当我们需要存储任意数据类型的实例的时候，
//空接口类型的使用使得我们得心应手。

func main() {
	var a interface{}
	var i int = 5
	s := "Hello world"
	var r float32 = 3.14159
	// These are legal statements
	a = i
	fmt.Println(a)
	a = s
	fmt.Println(a)
	a = r
	fmt.Println(a)

}

//如果一个函数的参数包括空接口类型interface{}，实际上函数是在说“兄弟，我接受任何数据”。
//如果一个函数返回一个空接口类型，那么函数再说“我也不确定返回什么，你只要知道我一定返回一个值就好了”。