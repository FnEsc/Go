/*
1. 数组
2. slice
3. map
4. 结构体
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. 数组
	// 固定长度，特定类型
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	a[2] = 100
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	var q [3]int = [3]int{3, 2}
	fmt.Println(q[0], q[1], q[2])

	r := [...]int{11, 21, 31}
	fmt.Println(r[0], r[1], r[2])

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // 3 ￥

	r1 := [...]int{99: -1, -2}                          // 99+2=101个元素，最后2个元素分别被初始化为-1和-2，其它元素都是用0初始化
	fmt.Println(r1[0], r1[1], r1[99], r1[100], len(r1)) // 0 0 -1 -2 101

	// 2. Slice
	// slice底层引用了数据对象，一个slice由三个部分构成：指针、长度和容量。
	// 要注意的是slice的第一个元素并不一定就是数组的第一个元素
	// 内置的len和cap函数分别返回slice的长度和容量
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	// 这里 months 是字符串的slice，而前面只是为序号
	fmt.Println(months)

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2, summer)

	// slice之间不能比较，能使用==操作符来判断两个slice是否含有全部相等元素
	// []byte 能使用标准库高度优化的 bytes.Equal 函数来比较
	// slice的元素是间接引用的

	// 自己展开比较每个元素--->但是安全的做法是直接禁止slice之间的比较操作
	// func equal(x, y []string) bool{
	// 	if len(x)!=len(y){
	// 		return false
	// 	}
	// 	for i := range x{
	// 		if x[i] !=y[i]{
	// 			return false
	// 		}
	// 	}
	// 	return true
	// }

	// 4.2.1 内置的append函数用于向slice追加元素
	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r) // 常规写法，append
	}
	fmt.Printf("%q\n", runes) // ['h' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']

	// 每次调用appendInt函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素。如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice
	// 如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组
	// 通过在每次扩展数组时直接将长度翻倍从而避免了多次内存分配，也确保了添加单个元素操的平均时间是一个常数时间。

	// 3. Map
	ages := make(map[string]int)
	ages["alice"] = 31
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages2["alice"] = 31
	delete(ages, "alice")
	ages2["charlie"] += 1
	ages2["charlie"]++

	// 为了避免可能存在 bob 0岁
	ages["bob"] = 0
	age, ok := ages["bob"]
	if !ok {
		fmt.Println("bob is not a key in this map; age=", age)
	} else {
		fmt.Println("bob is a key in this map; age=", age)
	}

	// 和slice一样，map之间也不能进行相等比较，唯一的特例又是可以和nil比较
	// 手写循环对比 equal2
	fmt.Println(equal2(map[string]int{"A": 0}, map[string]int{"B": 100}))

	// go 没有set类型，但是map中的key也是不相同的，可以用map实现类似set的功能。

	// 4.4. 结构体
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &dilbert              // Employee 类型的指针
	employeeOfTheMonth.Position += "(proactive team player)" // 等价语句
	(*employeeOfTheMonth).Position += "(add)"                // 等价语句
	fmt.Println(dilbert.Position)                            // Senior (proactive team player)(add)

	// gopl.io/ch4/treesort
	// http://books.studygolang.com/gopl-zh/ch4/ch4-04.html
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { // 长度自带
		z = x[:zlen]
	} else { // 需要扩展长度
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func equal2(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv { // 这里是个技巧，写在了同一行
			return false
		}
	}
	return true
}
