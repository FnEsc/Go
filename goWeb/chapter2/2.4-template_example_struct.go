package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析置顶文件生成模板对象
	tmpl, err := template.ParseFiles("/Users/linshangman/SimulinkSelf/Go/goWeb/src/chapter2/template_example_struct.tmpl")
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}
	// 利用指定数据渲染模板，并将结果写入 w
	user := UserInfo{
		Name:   "李四",
		Gender: "男",
		Age:    28,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/2", sayHello)
	http.ListenAndServe(":8086", nil)
}
