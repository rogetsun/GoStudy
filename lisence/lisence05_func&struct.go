package main

import "fmt"

type user struct {
	name string //首字母小写只有类似方法一样，小写的只有同包能访问
	age  uint8
}

//给结构体user定义方法ToString()
func (u user) ToString() string {
	fmt.Println("user ToString exec")
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user  //引入了user的属性
	title string
}
type user2 struct {
	name string
	age  uint8
}

type manager2 struct {
	user  //引入了user2的属性
	title string
}

func (m manager2) ToString() string {
	fmt.Println("manager2 ToString exec")
	return fmt.Sprintf("%+v", m)
}

func main() {
	var m manager //03课中定义的结构体
	m.name = "litx"
	m.age = 13
	m.title = "pig"
	fmt.Println(m.ToString()) //调用的是user.ToString()

	var m2 manager2
	m2.name = "litx2"
	m2.age = 13
	m2.title = "pig2"
	fmt.Println(m2.ToString())
}
