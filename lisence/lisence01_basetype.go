package main

import (
	. "fmt"
	"time"
)

func main() {
	Println("fdsa hello world")
	var n = 1
	Println(n)
	var n1, n2 = 2, 3
	Println(n1)
	Println(n2)
	var t = time.Millisecond
	Println(t)
	Println(011)
	Println(0xff)
	Println(3.14e2)
	Println('这')
	var c rune = 'a'
	Println(c)
	var str1 string = "\\\""


	//这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
	Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。%d\n", str1, str1, c)
	var arr1 = [3]int{1, 2, 3}
	Println(len(arr1))
	Println(arr1)
	var arr2 = [...]int{2, 34, 5, 7}
	Println(arr2)
	var arr3 [5]int
	Println(arr3)


	/**切片,实现了动态数组的功能*/
	//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	Println("切片")
	Println(arr2[1:3])
	b := make([]int, 5, 10) // len(b)=5, cap(b)=10
	Printf("arr is %v, \tlen(b)=%d,cap(b)=%d\n", b, len(b), cap(b))
	//继续切片，注意len和cap的变化
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	Printf("arr is %v, \tlen(b)=%d,cap(b)=%d\n", b, len(b), cap(b))
	b = b[1:]      // len(b)=4, cap(b)=4
	Printf("arr is %v, \tlen(b)=%d,cap(b)=%d\n", b, len(b), cap(b))
	b = append(b, 2)
	b = append(b, 3)
	Println(b)//此时已经11个元素


	//访问底层数组
	var arr4 = arr2[1:3]
	Println(len(arr4))
	arr4 = arr4[:cap(arr4)]
	Println(arr4)


	/**字典类型*/
	Println("字典")
	var person = make(map[string]string)
	person["name"] = "litx"
	person["age"] = "16"
	Printf("%+v\n", person)
	Printf("gender is [%s]\n", person["gender"])
	delete(person, "name")
	Println(person)
	//ok表示person字典里是否存在gender这个key
	gender, ok := person["gender"]
	Printf("gender len is [%v]\n", len(gender))
	Printf("person dict exists gender?%t\n", ok)
	var ts int = 1
	Println(ts)
	rn, err := Println("123")
	Printf("%v, %v\n", rn, err)

}
