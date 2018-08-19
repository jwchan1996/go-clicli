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
	router.POST("/comment/add", handler.AddComment)
	router.POST("/comment/delete/:id", handler.DeleteComment)
	router.POST("/logout", handler.Logout)
	router.GET("/pv/add/:pid", handler.AddPageView)
	router.GET("/user", handler.GetUser)
	router.GET("/post/:id", handler.GetPost)
	router.GET("/comments", handler.GetComments)
	router.GET("/posts/type", handler.GetPostsOneOf)
	router.GET("/posts/both", handler.GetPostsBoth)
	router.GET("/users", handler.GetUsers)
	router.GET("/search/posts", handler.SearchPosts)
	router.GET("/search/users", handler.SearchUsers)
	router.GET("/auth", handler.Auth)
	router.GET("/count/:pid", handler.GetCount)

	return router
}
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":4000", r)
}
