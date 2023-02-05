package api

import (
	"anonichat-pgx-stdlib/middlewares"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	a := r.Group("/api")

	au := a.Group("/auth")
	au.POST("/register", Register)
	au.POST("/login", Login)
	au.GET("/genders", GetAllGenders)

	u := a.Group("/u")
	u.Use(middlewares.JwtAuthMiddleware())
	u.GET("/posts", GetAllPosts)
	u.GET("/posts/:user_id", GetAllPostsByUserID)
	u.POST("/post", CreatePost)
	u.DELETE("/post/:pid", DeletePost)

	u.GET("/likesdt/:pid", GetLikesByPostID)
	u.GET("/likes/:pid", CountLikesPostID)
	u.POST("/like/:pid", AddLikeToPostID)
	u.POST("/unlike/:pid", UnLikeFromPostID)

	u.POST("/comment/:pid", AddComment)
	u.GET("/comments/:pid", GetCommentsFromPostID)
	u.DELETE("/comment/:cid", DeleteCommentFromID)

	return r
}
