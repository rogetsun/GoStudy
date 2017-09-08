package main

import (
	fmt "log"
	"reflect"
	"time"
)

/***
slice是一个指向底层数组的引用，它是一个引用类型
因此多个slice指向同一个底层数组，一个改变数组内容，全变

注意：
	slice的容量（假设10）初始化后，如果超出了，会再次重新分配一个连续的内存，容量翻倍，变成20；在超出40
*/
func main() {
	/**
	创建
	*/
	var s1 []int                      //直接定义slice
	ss1 := []int{1, 2, 3, 4, 5, 6, 7} //直接定义且赋值
	fmt.Println(s1, ss1)

	a1 := [...]int{9: 3}
	fmt.Println(a1)
	s2 := a1[5:8] //通过已经存在数组定义 s2的len()=8-5,cap()=len(a1)-5
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	s3 := make([]int, 3, 10) //make定义，[]int类型；初始化3个元素，len()=3；容量10,cap()=10
	fmt.Println(s3)

	/**
	共享底层数组测试
	*/
	fmt.Println("共享数组测试...")
	a2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a2)
	s4 := a2[1:5]
	s5 := a2[3:8]
	fmt.Println(s4, s5, a2)
	s4[3] = 22
	fmt.Println(s4, s5, a2)

	// **********************************
	// 这里s4 append的时候，如果一次性append的内容没有超出s4的cap 9，则会将底层数组覆盖掉。
	// 如果一次性append的内容超出了cap，则直接重新分配内存，不会覆盖掉原有底层数组。
	// **********************************
	s4 = append(s4, 11, 11, 11, 11, 11)
	s4 = append(s4, 23)
	fmt.Println(s4, s5, a2)

	/**
	slice的copy。俄罗斯方块覆盖原则
	*/
	fmt.Println("slice的copy...")
	sc1 := []int{1, 2, 3, 4, 5}
	sc2 := []int{7, 8, 9}
	copy(sc1, sc2)
	fmt.Println(sc1, sc2)

	sc3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sc4 := []int{7, 8, 9}
	copy(sc4[2:], sc3[1:2]) //指定位置copy
	fmt.Println(sc3, sc4)

	/**
	slice-->chan 注意slice要重新分配内存
	*/
	fmt.Println("slice -> chan...")
	var channel1 = make(chan []int, 3)
	var done = make(chan bool, 1)
	go func(channel1 chan<- []int) {
		var count = 2
		sc3 = make([]int, 0, 20)
		for i := 0; i < count; i++ {
			for j := 0; j < 20; j++ {
				sc3 = append(sc3, i*100+j+1)
			}
			fmt.Printf("append: %v\n", sc3)
			for m := 0; m < 20; m += 5 {
				tsc := sc3[m : m+5]
				fmt.Printf("in channel:%v\n", tsc)
				channel1 <- tsc
			}
		}
		time.AfterFunc(time.Second*2, func() {
			fmt.Println("AfterFunc")
			done <- true
		})
	}(channel1)
	go func(channel <-chan []int) {
		for {
			data, ok := <-channel
			if ok {
				fmt.Println(data)
			}
		}
	}(channel1)
	<-done
	//fmt.Println(<-channel1)
	//fmt.Println(<-channel1)
	/**
	遍历
	range类似于一个迭代器，但是拿到的v是一个copy。当然如果v是一个复杂的引用类型，而非基本类型则类似js的特性。再14课map有测试
	*/
	fmt.Println("遍历")
	for idx, v := range sc4 {
		fmt.Printf("idx=%v, v=%v\n", idx, v)
	}

	/**
	slice len cap 测试
	*/
	fmt.Println("slice len cap 测试....")
	sc5 := make([]int, 5, 5)
	sc6 := make([]int, 1, 5)
	arr1 := [5]int{9, 9, 9, 9, 9}
	sc7 := arr1[0:0]
	fmt.Printf("sc5:len=%v, cap=%v;\tsc6:len=%v,cap=%v;\tsc7:len=%v,cap=%v\n", len(sc5), cap(sc5), len(sc6), cap(sc6), len(sc7), cap(sc7))
	fmt.Println(sc5[0:5], sc6[0:5], sc7)
	fmt.Println(arr1)
	sc7 = append(sc7, 0, 1, 1, 1)
	copy(sc5, sc3)
	copy(sc6, sc3)
	copy(sc7, sc3)
	fmt.Printf("sc5:len=%v, cap=%v;\tsc6:len=%v,cap=%v;\tsc7:len=%v,cap=%v\n", len(sc5), cap(sc5), len(sc6), cap(sc6), len(sc7), cap(sc7))
	fmt.Println(sc5[0:5], sc6[0:5], sc7)
	fmt.Println(arr1)

	/**
	二维slice
	*/
	fmt.Println("二维slice...")
	sc8 := make([][5]int, 1, 10)
	sc9 := make([][]int, 1, 10)
	fmt.Println(sc8, sc9)
	fmt.Println(reflect.TypeOf(sc8).Kind(), reflect.TypeOf(sc8[0]).Kind())
	fmt.Println(reflect.TypeOf(sc9).Kind(), reflect.TypeOf(sc9[0]).Kind())

	/**
	地址测试
	*/
	fmt.Println("slice地址测试...")
	sc10 := make([]int, 10)
	fmt.Printf("堆内存地址：%p, 栈内存地址：%p, %v", sc10, &sc10, sc10)
	sc10 = make([]int, 10)
	fmt.Printf("堆内存地址：%p, 栈内存地址：%p, %v", sc10, &sc10, sc10)

}
