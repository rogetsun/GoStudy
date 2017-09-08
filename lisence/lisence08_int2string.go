package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 65
	fmt.Println(string(a))
	fmt.Println(strconv.Itoa(a))
	fmt.Println(strconv.Atoi("65"))
	fmt.Println(int('a'))
}
