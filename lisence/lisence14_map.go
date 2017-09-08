package main

import "fmt"

func main() {

	/**
	定义
	*/
	m := map[int]string{1: "1", 2: "2"}
	fmt.Println(m)

	/**
	测试map的value还是一个map
	*/
	mm := make(map[int]map[int]string)
	fmt.Println(mm)

	mm[1] = make(map[int]string)
	mm[1][1] = "ok"
	/**
	mm的value是另一个map。因此range拿到的value是一个引用copy。
	直接value ＝ newvalue是不会改变原有mm这个map的值。
	但是value[1]是可以改变原有map的value（map）的内部值
	这个特性类似javascript
	*/
	for key, value := range mm {
		fmt.Printf("key=%d,value=%v\n", key, value)
		value[1] = "false"
	}
	fmt.Printf("%+v\n", mm)
	/**
	测试map的value为string
	*/
	fmt.Println("test 2")
	mm2 := make(map[int]string)
	mm2[1] = "1111"
	for k, v := range mm2 {
		v = "123" //直接对v重新赋值，不会改变原有map的value
		fmt.Printf("key=%v,value=%v \n", k, v)
	}
	fmt.Println(mm2)

}
