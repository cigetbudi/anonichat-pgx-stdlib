package api

import (
	"anonichat-pgx-stdlib/middlewares"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.POST("/auth/register", Register)
	r.POST("/auth/login", Login)

	u := r.Group("/u")
	u.Use(middlewares.JwtAuthMiddleware())
	u.GET("/getAllPosts", GetAllPosts)
	u.GET("/getAllPostsByUserID/:user_id", GetAllPostsByUserID)
	u.POST("/post", CreatePost)
	u.DELETE("/post/:pid", DeletePost)
	u.GET("/getLikesByPostID/:pid", GetLikesByPostID)
	u.POST("/addLikeToPostID/:pid", AddLikeToPostID)
	u.POST("/unLikePostFromID/:pid", UnLikeFromPostID)

	return r
}
