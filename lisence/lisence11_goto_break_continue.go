package main

import "fmt"

func main() {
	//break
	label1:
	for i := 0; i < 10; i++ {
		fmt.Println("out ", i)
		if i > 3 {
			break label1
		}
		for i := 0; i < 10; i++ {
			fmt.Println("in ", i)
			if i > 3 {
				continue label1
			}
		}
	}

	//goto
	for i := 1; i < 100; i++ {
		fmt.Println("goto ", i)
		if i > 10 {
			goto next
		}
	}
	next:
	fmt.Println("end")

	//test goto
	fmt.Println("test goto...")
	var x int
	X:
	fmt.Println(x)
	x ++
	if x < 5 {
		goto X
	}
}
