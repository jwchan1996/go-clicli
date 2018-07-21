package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/register", RegisterUser)
	router.POST("/login", Login)
	router.GET("/posts", AllPosts)
	router.GET("/post", GetPost)

	return router
}
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":4000", r)
}
