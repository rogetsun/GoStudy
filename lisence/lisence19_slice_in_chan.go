package main

import (
	"reflect"
	fmt "log"
)

const (
	ARR_LEN = 10
)
/**
测试结论：
	放入通道的操作为值copy
	放入slice，因为slice是一个指针，所以是放入了一个指针copy，两个指针任然指向同一个内存地址
 */
func main() {

	chanLen := 10
	ch1 := make(chan []int, chanLen)
	var arr1 = new([ARR_LEN]int)
	fmt.Printf("arr1:%p, %p\n", arr1, &arr1)
	makeArr(arr1, 1)
	for i := 0; i < ARR_LEN; i += ARR_LEN / 2 {
		tmpSc := arr1[i:i + ARR_LEN / 2]
		fmt.Printf("%p, %p,\t %v,%v, \t%v\n", tmpSc, &tmpSc, reflect.TypeOf(tmpSc), reflect.TypeOf(tmpSc).Kind(), tmpSc)
		ch1 <- tmpSc
	}
	close(ch1)
	makeArr(arr1, 20)//修改底层数组值
	fmt.Println("begin get from channel")
	for {
		tmpSc, ok := <-ch1
		if ok {
			fmt.Printf("%p, %p,\t %v,%v, \t%v\n", tmpSc, &tmpSc, reflect.TypeOf(tmpSc), reflect.TypeOf(tmpSc).Kind(), tmpSc)
		} else {
			break
		}
	}

}

func makeArr(arr *[ARR_LEN]int, beginV int) {
	for i := 0; i < ARR_LEN; i++ {
		arr[i] = beginV + i
	}
	fmt.Printf("makeArr:%p, %p, %v\n", arr, &arr, arr)
}
