package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/handler"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.GET("/user/:name", handler.GetUser)
	router.POST("/post/add", handler.AddPost)
	//router.GET("/posts", handler.AllPosts)
	router.GET("/post/:id", handler.GetPost)

	return router
}
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":4000", r)
}
