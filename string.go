package main

import "fmt"

func main() {
	var s string = "huangby"
	fmt.Println(s)
	//字符串长度
	fmt.Println("字符串长",len(s))

	//字符串中的字符 打印出来为ASCII码值
	fmt.Println("字符串中的字符 0 和 3 分别是",s[0],s[3])

	//字符串切割 从0开始到5结束不包含5
	fmt.Println("切割前5位",s[0:5])
	fmt.Println("结尾",s[5:])

	//字符串连接
	var t string = "blog"
	fmt.Println(s+t)
	var plus string = " of huangby"
	t += plus
	fmt.Println(t)


	//字符串里面改一个字符
	var s1 string = "hello"
	c:=[]byte(s1)
	c[0]='H'
	s2:=string(c)
	fmt.Println(s2)
	//字符串原样输出
	raw:=`huang
	bao
	ying
	`
	fmt.Println(raw)

	//转义
	// \a      响铃
	// \b      退格
	// \f      换页
	// \n      换行
	// \r      回车
	// \t      制表符
	// \v      垂直制表符
	// \'      单引号 (只用在 '\'' 形式的rune符号面值中)
	// \"      双引号 (只用在 "..." 形式的字符串面值中)
	// \\      反斜杠
	fmt.Println("响铃","\a");
	
}