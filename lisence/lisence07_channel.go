package main

import (
	"fmt"
	"time"
)

func consumer(id int, data chan int, done chan bool) {
	fmt.Println("******consumer-%d", id)
	for {
		i, isOk := <-data //数据通道中死循环取数据，除非通道关闭isOk==false
		fmt.Printf("consumer %d i=%d, isOpen=%v\t", id, i, isOk)
		if !isOk {
			fmt.Println("\n")
			break
		}
		fmt.Printf("deal %d\n", i)
	}
	done <- true //放入结果到结果通道
}

func producer(data chan int) {
	for i := 0; i < 10; i++ {
		data <- 100 + i
		data <- 200 + i
		time.Sleep(time.Second)
	}
	fmt.Println("producer end")
	time.Sleep(time.Second * 5) //5秒后关闭数据通道
	fmt.Println("data close")
	close(data)
}

func main() {
	data := make(chan int)
	done := make(chan bool)

	go consumer(1, data, done)

	go producer(data)
	time.Sleep(time.Second * 2) //2秒后加入第二个消费者
	go consumer(2, data, done)
	fmt.Println(<-done) //阻塞，等待处理结果

}
