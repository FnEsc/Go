/*
Go语言将数据类型分为四类：
1. 基础类型：数字、字符串和布尔型。
2. 复合数据类型：数组和结构体（通过组合简单类型）。
3. 引用类型包括指针、切片、字典、函数、通道
4. 接口类型
*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// 3.1 整型

	/*
		*      /      %      <<       >>     &       &^
		+      -      |      ^
		==     !=     <      <=       >      >=
		&&
		||
		五种优先级，同一优先级从左到右
	*/
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

	// 3.2 浮点数
	for x4 := 0; x4 < 8; x4++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x4, math.Exp(float64(x4)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //0 -0 +Inf -Inf NaN
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // false false false
	fmt.Println(math.IsNaN(nan))                  // true

	// 3.3 复数
	var x31 complex128 = complex(1, 2)                 // 1+2i
	var y31 complex128 = complex(3, 4)                 // 3+4i
	fmt.Println(x31*y31, real(x31*y31), imag(x31*y31)) // (-5+10i) -5 10

	// 3.4 布尔型
	// true和false

	// 3.5 字符串
	// 字符串是不可修改的
	s35 := "hello, world"
	fmt.Println(len(s35))       // 12
	fmt.Println(s35[0], s35[7]) // 104 119 (h w)
	// c35 := s35[len(s35)]		// panic: index out of range
	// fmt.Println(c35) 		// panic: index out of range
	fmt.Println(s35[0:5]) // "hello"

	// 字符串 vs 数字
	x35 := 123
	y35 := fmt.Sprintf("%d", x35)
	fmt.Println(y35, strconv.Itoa(x35), y35 == strconv.Itoa(x35)) // 123 123 true
	// fmt.Fprintln(x35 == y35) // 编译错误

	// x, err := strconv.Atoi("123")              // x is an int
	// y, err := strconv.ParseInt("123", 10, 64)  // base 10, up to 64 bits

	fmt.Println(strconv.FormatInt(int64(5), 2)) // "101"

	// 3.6 常量
	const pi = 3.14159

	const (
		e   = 2.71828182845904523536028747135266249775724709369995957496696763
		pi2 = 3.14159265358979323846264338327950288419716939937510582097494459
	)

	const (
		a36 = 1
		b36
		c36 = 2
		d36
	)
	fmt.Println(a36, b36, c36, d36) // 1 1 2 2

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	// 0 ~ 6
}
