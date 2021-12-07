package main

import (
	"fmt"
	"regexp"
)

func main() {
	//func Compile(expr string) (*Regexp, error)
	r, _ := regexp.Compile(`f([a-z]+)`)

	//func (re *Regexp) Match(b []byte) bool
	fmt.Println(r.Match([]byte("foo"))) //true

	//func (re *Regexp) MatchString(s string) bool
	fmt.Println(r.MatchString("foo")) //true

	//func (re *Regexp) FindString(s string) string
	//只匹配一次
	fmt.Println(r.FindString("foo func")) //foo

	//func (re *Regexp) FindStringIndex(s string) (loc []int)
	fmt.Println(r.FindStringIndex("demo foo func")) //[5 8]

	//func (re *Regexp) FindStringSubmatch(s string) []string
	//只匹配一次，返回的结果中，索引为0的值是整个匹配串的值，第二个值是子表达式的值
	fmt.Println(r.FindStringSubmatch("this foo func fan")) //[foo oo]

	//对于FindStringSubmatch，如果表达式中没有子表达式，则不检测子表达式
	demo, _ := regexp.Compile(`foo`)
	fmt.Println(demo.FindStringSubmatch("foo")) //[foo]

	//func (re *Regexp) FindStringSubmatchIndex(s string) []int
	fmt.Println(r.FindStringSubmatchIndex("foo func")) //[0 3 1 3]

	//func (re *Regexp) FindAllString(s string, n int) []string
	//n为-1时，匹配所有符合条件的字符串，n不为-1时，表示只匹配n次
	fmt.Println(r.FindAllString("foo func fan", -1)) //[foo func fan]
	fmt.Println(r.FindAllString("foo func fan", 2))  //[foo func]
	fmt.Println(r.FindAllString("fo fa", 5))         //[fo fa]
	fmt.Println(r.FindAllString("f f", 5))           //[]
	fmt.Println(r.FindAllString("foo func fan", 0))  //[]

	//func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
	//n同样是表示匹配的次数，-1表示匹配所有
	fmt.Println(r.FindAllStringSubmatchIndex("foo func demo fan", -1))
	//[[0 3 1 3] [4 8 5 8] [14 17 15 17]]

	//替换

	//func (re *Regexp) ReplaceAll(src []byte, repl []byte) []byte
	fmt.Println(string(r.ReplaceAll([]byte("this is foo, that is func, they are fan"), []byte("x"))))
	//this is x, that is x, they are x

	//func (re *Regexp) ReplaceAllString(src string, repl string) string
	fmt.Println(r.ReplaceAllString("this is foo, that is func, they are fan", "xx"))
	//this is xx, that is xx, they are xx

}
