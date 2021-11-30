package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
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
	//router.GET("/user/name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//	fmt.Print(params.ByName("name"))
	//	writer.Write([]byte("user name 1: " + params.ByName("name")))
	//})
	//匹配所有
	router.GET("/user/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("user name 2: " + params.ByName("name")))
	})
	router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name 3:" + p.ByName("name")))
	})

	// 捕获服务器异常
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "error %s", i)
	}

	log.Fatal(http.ListenAndServe(":8086", router))

}
