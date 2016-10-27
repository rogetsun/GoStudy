package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s == nil)
	var ss [][]byte
	fmt.Println(ss == nil)
	var arr [3]int
	fmt.Println(arr)
	//fmt.Println(arr == nil)
}
