package main

import "fmt"

func main() {

	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("nil")
	}
	//变种,a:=0;分号必须要
	switch a := 0; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough//接的执行后面的case，但是不做条件检查，通swift
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("nil")
	}

}
