package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

const boilingF = 212.0

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "---", "seperator") // 默认分割符为 ---

var cwd string

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

	fmt.Println(f2() == f2()) // "false" 说明局部变量指针

	v := 1
	incr(&v)              // v = 2
	fmt.Println(incr(&v)) // v = 3

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}

	p2 := new(int)
	fmt.Println(*p2) // 0
	*p2 = 2
	fmt.Println(*p2) // 2

	// 2.4 赋值
	medals := []string{"glod", "silver", "bronze"} // 隐式赋值，类似 medals[1] = "silver"
	fmt.Println(medals)

	// nil可以赋值给任何指针或引用类型的变量

	// 2.5. 类型
	// tempconv

	// 2.6 包和文件
	// 如果一个p包导入了q包，那么在p包初始化的时候可以认为q包必然已经初始化过了

	// 2.7 作用域
	x3 := "hello!"
	for i := 0; i < len(x3); i++ {
		x3 := x3[i] // 循环内屏蔽外部得声明
		if x3 != '!' {
			x3 := x3 + 'A' - 'a'
			fmt.Printf("%c", x3) // HELLO
		}
	}
	fmt.Printf("%s", x3) // hello!

	fmt.Println("---")
	x4 := "hello"
	for _, x4 := range x4 {
		x4 := x4 + 'A' - 'a'
		fmt.Printf("%c", x4) // HELLO
	}
	fmt.Printf("%s", x4) // hello

	// 可见， for if switch 都会在条件部分创建隐式词法域

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func f2() *int {
	v := 1
	return &v
}

func incr(p *int) int { // 在这里必须做响应的返回类型
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	// *p就是变量v的别名
	return *p
}

func newInt() *int {
	return new(int) // ==> 等价于 var dummy int; return &dummy 返回新的变量地址
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println(cwd)
}
