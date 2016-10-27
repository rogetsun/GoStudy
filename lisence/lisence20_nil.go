package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s == nil)
	var ss [][]byte
	fmt.Println(ss == nil)
}
