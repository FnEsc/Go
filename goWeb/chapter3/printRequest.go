package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func printRequest(w http.ResponseWriter, r *http.Request) {
	// 打印
	fmt.Println("Request 解析")
	// http 方法
	fmt.Println("method:", r.Method)
	// RequestURI
	fmt.Println("RequestURI:", r.RequestURI)
	// URL 类型，列出 URL 各成员
	fmt.Println("URL_Path:", r.URL.Path)
	fmt.Println("URL_RawQuery", r.URL.RawQuery)
	fmt.Println("URL_Fragment", r.URL.Fragment)
	// 协议版本
	fmt.Println("proto", r.Proto)
	fmt.Println("protomajor", r.ProtoMajor)
	fmt.Println("protominor", r.ProtoMinor)
	// http 请求头
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Println("header key: " + k + " value: " + vv)
		}
	}
	// 判断是否未 multipart 方式
	isMultipart := false
	for _, v := range r.Header["Content-Type"] {
		if strings.Index(v, "multipart/form-data") != -1 {
			isMultipart = true
		}
	}
	// 解析 Form 表单
	if isMultipart == true {
		r.ParseMultipartForm(128)
		fmt.Println("解析方式: ParseMultipartForm")
	} else {
		r.ParseForm()
		fmt.Println("解析方式: ParseForm")
	}
	// http body 内容长度
	fmt.Println("ContentLength", r.ContentLength)
	// 是否在回复请求后关闭链接
	fmt.Println("Close", r.Close)
	// Host
	fmt.Println("host", r.Host)
	// 该请求的来源地址
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Println(w, "hello, let's go!") // 输出至客户端
}

func main() {
	http.HandleFunc("/hello", printRequest)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error", err)
	}
}
