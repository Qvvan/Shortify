package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var client *mongo.Client
var mongoURI string

// InitMongo инициализирует подключение к MongoDB
func InitMongo() {
	logger := zap.L().Sugar()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file", err)
	}

	mongoURI = os.Getenv("MONGO_URI")
	if mongoURI == "" {
		logger.Fatal("MONGO_URI is not set in environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal("Failed to ping MongoDB:", err)
	}

	logger.Info("MongoDB connected successfully")
}

// GetClient возвращает экземпляр клиента MongoDB
func GetClient() *mongo.Client {
	if client == nil {
		zap.L().Fatal("MongoDB client is not initialized")
	}
	return client
}
