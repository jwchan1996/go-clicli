package main

import (
	"github.com/132yse/acgzone-server/api/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/register", handler.Register)
	router.POST("/user/login", handler.Login)
	router.POST("/user/logout", handler.Logout)
	router.POST("/user/update/:id", handler.UpdateUser)
	router.POST("/user/delete/:id", handler.DeleteUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/user", handler.GetUser)

	router.POST("/post/add", handler.AddPost)
	router.POST("/post/delete/:id", handler.DeletePost)
	router.POST("/post/update/:id", handler.UpdatePost)
	router.GET("/post/:id", handler.GetPost)
	router.GET("/posts", handler.GetPosts)

	router.POST("/comment/add", handler.AddComment)
	router.DELETE("/comment/delete", handler.DeleteComment)
	router.GET("/comments", handler.GetComments)

	router.POST("/video/add", handler.AddVideo)
	router.POST("/video/update/:id", handler.UpdateVideo)
	router.DELETE("/video/delete", handler.DeleteVideo)
	router.GET("/video/:id", handler.GetVideo)
	router.GET("/videos", handler.GetVideos)

	router.GET("/search/posts", handler.SearchPosts)
	router.GET("/search/users", handler.SearchUsers)

	router.GET("/auth", handler.Auth)

	router.POST("/cookie/replace", handler.ReplaceCookie)
	router.GET("/cookie/:uid", handler.GetCookie)

	return router
}
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":4000", r)
}
