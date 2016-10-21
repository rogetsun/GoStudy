package main

import (
	"fmt"
	"time"
)

var c int = 0

func task(id int) {
	for i := 0; i < 5; i++ {
		c ++
		fmt.Printf("%d : %d----%d\n", id, i, c)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("main begin")
	go task(1)
	go task(2)
	time.Sleep(time.Second * 8)
	fmt.Println("main end")
}
