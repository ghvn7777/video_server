package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login) // 会把冒号之后的东西作为参数传到 handler 里，具体看 github

	return router
}

// handler -> validation{1.request, 2.user} -> business logic -> response
// validation:
//   1. data model
//   2. response
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8070", r)
}
