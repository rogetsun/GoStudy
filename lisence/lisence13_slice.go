package main

import "fmt"
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
	var s1 []int //直接定义slice
	ss1 := []int{1, 2, 3, 4, 5, 6, 7}//直接定义且赋值
	fmt.Println(s1, ss1)

	a1 := [...]int{9:3}
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
	sc3 := []int{1, 2, 3, 4, 5}
	sc4 := []int{7, 8, 9}
	copy(sc3[3:], sc4[1:2])//指定位置copy
	fmt.Println(sc3, sc4)

}
