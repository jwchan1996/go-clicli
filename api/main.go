package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/handler"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/register", handler.Register)
	router.POST("/api/login", handler.Login)
	router.POST("/user/update/:id", handler.UpdateUser)
	router.POST("/user/delete/:id", handler.DeleteUser)
	router.POST("/post/add", handler.AddPost)
	router.POST("/post/delete/:id", handler.DeletePost)
	router.POST("/post/update/:id", handler.UpdatePost)
	router.GET("/user/:name", handler.GetUser)
	router.GET("/post/:id", handler.GetPost)
	router.GET("/posts", handler.GetPosts)
	router.GET("/users", handler.GetUsers)
	router.GET("/search/posts", handler.SearchPosts)
	router.GET("/search/users", handler.SearchUsers)
	router.GET("/auth", handler.Auth)

	return router
}
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":4000", r)
}
