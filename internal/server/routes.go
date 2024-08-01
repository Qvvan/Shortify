package server

import (
	_ "Shortify/docs"
	"Shortify/internal/api/v1"
	"Shortify/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// Инициализация маршрутов
func InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLoggerMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("api/v1/shortify", v1.CreateShortUrl)
	router.GET("/:shortUrl", v1.RedirectShortUrl)

	return router
}
