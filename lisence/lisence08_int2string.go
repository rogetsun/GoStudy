package main

import (
	"fmt"
	"strconv"
)

const (
	UB int64 = 1 << (10 * iota)
	KB
	GB
	TB
)

func main() {

	a := 65
	fmt.Println(string(a))
	fmt.Println(strconv.Itoa(a))
	fmt.Println(strconv.Atoi("65"))
	fmt.Println(int('a'))
	fmt.Printf("B=%dbyte, KB=%dbyte, GB=%dbyte, TB=%dbyte", UB, KB, GB, TB)

}
