package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo2 struct {
	Name   string
	Gender string
	Age    int
}

func tmplSample(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"/Users/linshangman/SimulinkSelf/Go/goWeb/src/chapter2/t.html",
		"/Users/linshangman/SimulinkSelf/Go/goWeb/src/chapter2/ul.html",
	)
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}

	user := UserInfo2{
		Name:   "李四",
		Gender: "男",
		Age:    28,
	}

	tmpl.Execute(w, user)

}

func main() {
	http.HandleFunc("/3", tmplSample)
	http.ListenAndServe(":8086", nil)
}
