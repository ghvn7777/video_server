package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)
	router.POST("/", homeHandler)

	router.GET("/userhome", userhomeHandler)
	router.POST("/userhome", userhomeHandler)

	router.POST("/api", apiHandler)

	// 把 template 下的文件自动挂载到 statics 目录下，后面访问方便
	// 127.0.0.1:8080/statics/*... 其实访问的是 template 文件夹下的文件
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
