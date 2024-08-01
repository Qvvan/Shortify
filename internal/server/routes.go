package server

import (
	"Shortify/internal/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("api/v1/shortify", v1.CreateShortUrl)
	router.GET("/:shortUrl", v1.RedirectShortUrl)

	return router
}
