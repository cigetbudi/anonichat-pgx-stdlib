package api

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.POST("/auth/register", Register)
	r.POST("/auth/login", Login)

	r.GET("/getAllPosts", GetAllPosts)
	r.POST("/post", CreatePost)
	r.DELETE("/post/:id", DeletePost)

	return r
}
