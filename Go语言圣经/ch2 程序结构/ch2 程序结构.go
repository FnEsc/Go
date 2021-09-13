package main

import (
	"fmt"
	"math/rand"
)

const boilingF = 212.0

func main() {
	// 2.2 声明
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	const freezinF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezinF, fToC(freezinF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	// 2.3 变量
	var s string
	fmt.Println(s)

	var i, j, k int                   // int, int, int
	var b, ff, ss = true, 2.3, "four" // bool, float64, string
	fmt.Println(i, j, k)
	fmt.Println(b, ff, ss)

	// var f, err = os.Open(name) // os.Open returns a file and an error

	// 简短变量声明
	freq := rand.Float64() * 3.0
	fmt.Println(freq)

	var boiling float64 = 100
	fmt.Println(boiling)

	// 简短变量声明语句中必须至少要声明一个新的变量
	i, z := 1, 2 // 如果是 i, j := 1, 2 则编译不通过
	fmt.Println(i, z)

	// 2.3.2. 指针
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var x2, y2 int                                      // &x2 指向零值，是个有效变量
	fmt.Println(&x2 == &x2, &x2 == &y2, &x2 == nil, x2) // true false false

	fmt.Println(f2() == f2()) // "false"
	// http://books.studygolang.com/gopl-zh/ch2/ch2-03.html
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func f2() *int {
	v := 1
	return &v
}
