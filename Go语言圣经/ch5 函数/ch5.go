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
	// http://books.studygolang.com/gopl-zh/ch5/ch5-02.html
}

func f(i, j, k int, s, t string) { /* 函数返回值列表可省略 */ }
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }
