package main

import (
	"fmt"
	"math"
)

func main() {

	// 5.1 函数声明
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

	// 5.2 递归

	// 5.3 bare return
	// 在函数声明时候明确声明的多个返回值，并且在或函数内变量有明确的变量名，即可省略 return 的多个名字

	// 5.4 错误
	// http://books.studygolang.com/gopl-zh/ch5/ch5-04.html
}

func f(i, j, k int, s, t string) { /* 函数返回值列表可省略 */ }
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }
