package main

import "fmt"

/**
首先要明确的是：defer是在return之前执行的。这个在 官方文档中明确说明了的。

然后是了解defer的实现方式，大致就是在defer出现的地方，插入指令CALL runtime.deferproc，
然后在函数返回之前的地方，插入指令CALL runtime.deferreturn。再就是明确go返回值的方式跟C是不一样的，
为了支持多值返回，go是用栈返回值的，而C是用寄存器。

最最重要的一点就是： *return xxx这一条语句并不是一条原子指令!*

整个return过程，没有defer之前，先在栈中写一个值，这个值会被当作返回值，然后再调用RET指令返回。
return xxx语句汇编后是先给返回值赋值，再做一个空的return，( 赋值指令 ＋ RET指令)。
defer的执行是被插入到return指令之前的，有了defer之后，就变成了(赋值指令 + CALL defer指令 + RET指令)。
而在CALL defer函数中，有可能将最终的返回值改写了...也有可能没改写。总之，如果改写了，那么看上去就像defer是在return xxx之后执行的~
 */
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

