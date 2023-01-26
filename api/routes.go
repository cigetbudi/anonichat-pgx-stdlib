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
	u.GET("/getAllPosts", GetAllPosts)
	u.GET("/getAllPostsByUserID/:user_id", GetAllPostsByUserID)
	u.POST("/post", CreatePost)
	u.DELETE("/post/:pid", DeletePost)
	u.GET("/getLikesByPostID/:pid", GetLikesByPostID)
	u.GET("/countLikesPostID/:pid", CountLikesPostID)
	u.POST("/addLikeToPostID/:pid", AddLikeToPostID)
	u.POST("/unLikePostFromID/:pid", UnLikeFromPostID)

	return r
}
