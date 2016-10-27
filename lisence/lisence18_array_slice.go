package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	arr1 := &arr
	sc1 := arr1[0:5]
	fmt.Printf("栈内存地址:%p, %v\n", &arr, arr)
	fmt.Printf("arr1:堆内存地址:%p, 栈内存地址:%p, %v\n", arr1, &arr1, arr1)
	fmt.Printf("sc1:堆内存地址:%p, 栈内存地址:%p, %v\n", sc1, &sc1, sc1)
	fmt.Println("---------------重新定义---------------" +
		"\n- 栈内存地址都不变，堆都变" +
		"\n-------------------------------------")
	arr1 = new([10]int)
	arr1[1] = 1; arr1[2] = 2; arr1[8] = 8
	sc1 = arr1[0:5]
	fmt.Printf("arr1:堆内存地址:%p, 栈内存地址:%p, %v\n", arr1, &arr1, arr1)
	fmt.Printf("sc1:堆内存地址:%p, 栈内存地址:%p, %v\n", sc1, &sc1, sc1)

}
