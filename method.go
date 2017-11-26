package main

import "fmt"

//方法跟函数是有区别的 他们的区别是：方法在定义的时候，会在func和方法名之间增加一个参数，
//这个参数就是接收者，这样我们定义的这个方法就和接收者绑定在了一起，称之为这个接收者的方法。
func main() {
	p := person{name: "张三"}
	p.modify() //值接收者，修改无效
	//调用直接用`.`
	fmt.Println(p)
	//fmt.Println(p.String())

	print("1", "2", p)

}

type person struct {
	name string
}

//方法名是 String
func (p person) String() string {
	return "the person name is " + p.name
}

func (p person) modify() {
	p.name = "李四"
}
//有不确定个数的参数 都是interface类型
func print(b ...interface{}) {
	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println("------------------------")
}
