package main

import "fmt"

type Int int //可以理解为指向int的对象

func (i *Int) increment() Int {
	//给Int对象增加increment方法，入参是Int实体对象的指针
	*i++ //取指针指向的地址内容然后＋＋操作
	return *i
}

func main() {
	var x Int = 11
	fmt.Println(x.increment())
	x.increment() //对象直接调用方式
	x.increment()
	fmt.Println(x)
	(*Int).increment(&x) //第二种调用方法，像js的call
	fmt.Println(x + 10)
}
