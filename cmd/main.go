package main

import (
	_ "Shortify/docs"
	"Shortify/internal/logger"
	"Shortify/internal/mongodb"
	"Shortify/internal/server"
	"github.com/gin-gonic/gin"
)

// Основной файл
func main() {
	gin.SetMode(gin.ReleaseMode)
	logger.InitLogger()
	mongodb.InitMongo()
	server.StartServer()
}
