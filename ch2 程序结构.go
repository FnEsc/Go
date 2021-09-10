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

	// http://books.studygolang.com/gopl-zh/ch2/ch2-03.html
	// 2.3.2. 指针

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
