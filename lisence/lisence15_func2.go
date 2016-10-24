package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(A(1, 2))
	fmt.Println(B(2, 3))
	C(1, 23, 3)
}

//这里返回值是2个int，但是没有给名称，所以return必须指定要返回的数据
func A(a, b int) (int, int) {
	return a, b
}

//这里返回值显性命名为m,n那么在进入函数体前，这两个已经被作为局部变量声明了，所以返回不需要指定要返回的数据，go也知道返回谁
func B(a, b int) (m, n int) {
	m = a + 1
	n = b + 1
	return
}

//这里入参是多个不定长的，s在调用是定义成一个slice。注意不定长参数必须作为入参的最后一个
func C(s ...int) {
	fmt.Println(reflect.TypeOf(s))
}