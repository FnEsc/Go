package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {
	//写入文件
	var filename = "./output1.json"
	var f *os.File
	var err1 error

	if checkFileIsExist(filename) {
		f, err1 = os.OpenFile(filename, os.O_RDWR, 0666)
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename)
		fmt.Println("文件不存在")
	}
	check(err1)

	// 写入json文件
	//enc := json.NewEncoder(f)
	//d := map[string]int{"apple": 5, "lettuce中文": 7}
	//enc.Encode(d)
	//check(err1)

	// 读取json文件
	d := map[string]int{}
	decoder := json.NewDecoder(f)
	err1 = decoder.Decode(&d)
	check(err1)
	fmt.Println(d)

}
