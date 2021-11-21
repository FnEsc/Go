package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/default", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.Write([]byte("default get"))
	})
	router.POST("/default", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.Write([]byte("default post"))
	})
	// 精准匹配
	router.GET("/user/name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Print(params.ByName("name"))
		writer.Write([]byte("user name 1: " + params.ByName("name")))
	})
	// 匹配所有
	router.GET("/user/name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("user name 2: " + params.ByName("name")))
	})

	http.ListenAndServe(":8086", router)

	// 未理解该逻辑，参考 GoWeb 实战派 p143
}
