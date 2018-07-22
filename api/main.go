package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/handler"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/register", handler.Register)
	//router.POST("/login", Login)
	router.GET("/user/:name", handler.GetUser)
	router.GET("/posts", handler.AllPosts)
	router.GET("/post", handler.GetPost)

	return router
}
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":4000", r)
}
