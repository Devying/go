package main

import (
	"fmt"
	"bytes"
)

func main() {
	var b bytes.Buffer
	fmt.Fprint(&b,"Hello World")
	fmt.Println(b.String())
}