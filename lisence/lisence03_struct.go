package main

import "fmt"

type user struct {
	name string
	age  uint8
}

type manager struct {
	user  //引入了user的属性
	title string
}

func main() {
	var m manager

	m.name = "litx"
	m.age = 12
	m.title = "pig"

	fmt.Println(m)
	fmt.Printf("%+v\n", m)
	fmt.Println(m.title)
	fmt.Println(m.user)

	m2 := manager{user: user{age: 11, name: "litxxxx"}, title: "title"}
	fmt.Printf("%+v\n", m2)

}
