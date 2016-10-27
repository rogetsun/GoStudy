package main

import "fmt"
/***

go中数组非其他语言那样是一个引用类型，而是一个值类型，数组赋值是会copy一份的。
所以可以判等两个数组，但是不能比大小。如果需要引用的那种数组，可以参考slice
注意go中数组＋数组长度容量等是一个单独的类型，即[3]int [2]int 的变量不能赋值，不能判等
注意2：
	数组重新定义，不会重新分配内存空间，还是用旧的。
	数组指针重新定义，放指针的地址不会变，但是指针指向的底层数组会重新分配内存
	可以理解为数组是值类型，放在栈内存，数组指针的指针值放在栈，底层数组放在堆
	栈内存变量赋值是copy
	详看地址测试
 */
func main() {
	//只定义类型，
	var a [3]int
	fmt.Println(a)

	var b = [3]int{}
	fmt.Println(b)

	c := [3]int{}
	fmt.Println(c)

	d := [5]int{1, 2, 3, 4}
	fmt.Println(d)

	//只给e[8]赋值12
	e := [10]int{8:12}
	fmt.Println(e)

	//不指定长度
	f := [...]int{1, 3, 5, 7, 9}
	fmt.Println(f)
	fmt.Println(len(f))
	g := [...]int{9:9}
	fmt.Println(g)
	fmt.Println(len(g))

	//数组指针 注意＊号,数组指针只能用var方式创建，不能:=,或者下面的new方式
	var h *[10]int = &g
	fmt.Println(h)

	//指针数组 注意＊号
	v1, v2 := 1, 2
	i := [...]*int{&v1, &v2}
	fmt.Println(i)
	fmt.Println(*i[0])
	fmt.Println(*i[1])

	//数组判等
	j1 := [...]int{1, 2}
	j2 := [...]int{1, 2}
	fmt.Println(j1 == j2)

	//使用new创建数组，返回的是数组指针
	k := new([12]int)
	fmt.Println(k)

	//数组指针也可以通过下标直接操作元素
	//l := *[...]int{1,2}
	var l = [...]int{1, 2}
	fmt.Println(l)
	l[1] = 3
	fmt.Println(l)

	//多维数组
	m := [2][3]int{}
	fmt.Println(m)
	n := [...][4]int{{1, 2}, {2:3}}//只有顶级数组可以使用[...]
	fmt.Println(n)
	//地址测试
	fmt.Println("array address test ...")
	//地址测试 1 普通数组
	arr1 := [5]int{1, 2, 3, 4, 5}
	var arr2 *[5]int
	arr1[2] = 1
	arr2 = &arr1//数组地址复制到指针
	fmt.Printf("%p, %v\n", &arr2, arr2)
	fmt.Printf("%p, %v\n", &arr1, arr1)

	arr1 = [5]int{6, 7, 8, 9, 0}
	fmt.Printf("%p, %v\n", &arr1, arr1)

	fmt.Println(arr1, arr2)
	//地址测试 2 数组的指针
	arr3 := new([5]int)
	var arr4 *[5]int
	arr3[2] = 1
	arr4 = arr3//指针复制
	fmt.Printf("%p,%p, %v\n", arr4, &arr4, arr4)
	fmt.Printf("%p,%p, %v\n", arr3, &arr3, arr3)

	arr3 = new([5]int)
	fmt.Printf("%p,%p, %v\n", arr3, &arr3, arr3)

	fmt.Println(arr3, arr4)

}
