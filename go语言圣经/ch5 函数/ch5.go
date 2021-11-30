package main

import (
	"fmt"
	"math"
	"os"
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

	// 5.5 函数值
	// 函数可以直接北看作第一类值（类似别名）
	// 函数也可以直接先声明，零值是 nil。调用值为 nil 的函数值会引起 panic 错误
	// 但是函数值之间是不能比较的

	// 5.7 可变参数
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	// 5.8 defer 资源释放 类似 python - with

	// 5.9 panic 异常 类似 python - raise 表示主动抛出错误 如： panic(err)
	// 5.10 Recover 捕获异常 类似 try except 或者类似 defer - with 操作

}

func f(i, j, k int, s, t string) { /* 函数返回值列表可省略 */ }
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return nil, nil // ReadAll(f)
}
