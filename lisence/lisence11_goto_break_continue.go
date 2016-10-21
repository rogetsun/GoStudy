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
}
