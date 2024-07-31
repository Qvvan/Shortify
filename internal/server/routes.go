package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/qvvan/short_urls/short_urls/internal/api/v1"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/api/v1/shortify", v1.GetUrl)
	router.POST("api/v1/shortify", v1.CreateShortUrl)

	return router
}
