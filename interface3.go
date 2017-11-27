package main
import (
    "fmt"
    "strconv" //for conversions to and from string
)

//任何数据类型，只要实现了Stringer接口，就能够传递给fmt.Print函数，然后打印出该类型String()函数的返回值。
//stringer 接口如下
//The Stringer interface found in fmt package
//type Stringer interface {
//     String() string
//}


type Human struct {
    name string
    age int
    phone string
}

//Returns a nice string representing a Human
//With this method, Human implements fmt.Stringer
func (h Human) String() string {
    //We called strconv.Itoa on h.age to make a string out of it.
    //Also, thank you, UNICODE!
    return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
    Bob := Human{"Bob", 39, "000-7777-XXX"}
    fmt.Println("This Human is : ", Bob)
}