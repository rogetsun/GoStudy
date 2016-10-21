package main

import "fmt"

func main() {
	//只定义类型，注意go中数组＋数组长度容量等是一个单独的类型，即[3]int [2]int 的变量不能赋值
	var a [3]int
	fmt.Println(a)

	var b = [3]int{}
	fmt.Println(b)

	c := [3]int{}
	fmt.Println(c)

	d := [5]int{1, 2, 3, 4}
	fmt.Println(d)

	//只给e[8]赋值12
	e := [10]int{8:12}
	fmt.Println(e)
}
