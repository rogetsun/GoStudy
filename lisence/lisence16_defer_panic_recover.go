package main

import (
	"fmt"
	"reflect"
)

/**
defer 放在func内部的某一条语句表达式或函数左边定义，在func执行完毕后执行。且一定会执行。类似finally
一个函数内部可以有多个defer，顺序先进后出
defer内可以调用recover()获取func的panic(信息)。

panic() 类似throw exception
recover() 类似catch
*/
func main() {
	defer fmt.Println("main ok")
	defer func(a ...interface{}) {
		fmt.Println(reflect.TypeOf(a[0]))
		f, ok := a[0].(func()) //断言机制，将a[0]断言强转为func()类型
		fmt.Println("assert ", ok)
		if ok {
			//断言强转成功
			f()
		}
		err := recover()
		fmt.Println(err)
		fmt.Println(err == nil)
	}(A1)
	A1()
	B1(A1)
	C1()
}

func A1() {
	fmt.Println("func A")
}

func B1(f func()) {
	f()
	fmt.Println("func B")
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	panic("panic in B")
}
func C1() {
	fmt.Println("func C")
}
