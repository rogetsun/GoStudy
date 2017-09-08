package main

import (
	"fmt"
	"time"
)

func main() {
	cn1 := make(chan int, 10)
	cn2 := make(chan int, 10)
	done := make(chan int)

	fmt.Println(cn1, cn2)
	go func(cns []chan<- int) {
		for i := 0; i < 10; i++ {
			for _, c := range cns {
				c <- i
			}
			time.Sleep(time.Second)
		}
	}([]chan<- int{cn1, cn2})

	go func(c1 <-chan int, c2 <-chan int) {
		cnt := 0
		for {
			select {
			case i := <-c1:
				cnt = 0
				fmt.Println("get c1:", i)
			case i := <-c2:
				cnt = 0
				fmt.Println("get c2:", i)
			default:
				fmt.Println("no data sleep 0.5s")
				cnt++
				if cnt > 5 {
					done <- 1
				} else {
					time.Sleep(time.Millisecond * 500)
				}
			}
		}
	}(cn1, cn2)
	<-done
}
