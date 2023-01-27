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
	a.POST("/auth/register", Register)
	a.POST("/auth/login", Login)

	u := a.Group("/u")
	u.Use(middlewares.JwtAuthMiddleware())
	u.GET("/posts", GetAllPosts)
	u.GET("/posts/:user_id", GetAllPostsByUserID)
	u.POST("/post", CreatePost)
	u.DELETE("/post/:pid", DeletePost)
	u.GET("/post/:pid", GetLikesByPostID)
	u.GET("/likes/:pid", CountLikesPostID)
	u.POST("/like/:pid", AddLikeToPostID)
	u.POST("/unlike/:pid", UnLikeFromPostID)

	return r
}
