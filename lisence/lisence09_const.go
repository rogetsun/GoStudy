package main

import (
	"fmt"
)

//iota 每定义一个常量，值+1,从0开始,每遇到一个const 关键字，重置为0,因此在常量组中才有实际意义
const (
	a = iota + 1
	b = string(iota + 65) + string(iota + 97)
	c = iota * 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
