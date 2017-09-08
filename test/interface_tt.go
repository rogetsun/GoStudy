package main

import (
	"fmt"
	"reflect"
)

type maker interface {
	make(int) (int, int)
}

type person struct {
	name string
	mk   func(int) (int, int)
}

func (p person) make(i int) (int, int) {
	fmt.Println("p.make ", i)
	if p.mk != nil {
		return p.mk(i)
	} else {
		fmt.Println("p.mk is nil")
	}
	return 0, 0
}
func main() {
	p := &person{name: "litx"}
	//p.mk = func(i int) (int, int) {
	//	fmt.Println("person make ", i)
	//	return i + 1, i + 2
	//}
	fmt.Println(p.mk)
	fmt.Println(p.make(1))
	i, ok := interface{}(p).(maker)
	fmt.Println(i, ok)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i.make(3))
}
