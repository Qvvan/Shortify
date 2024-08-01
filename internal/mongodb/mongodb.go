package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURI = "mongodb://root:123qwe@mongodb:27017"

// InitMongo инициализирует подключение к MongoDB
func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Проверяем подключение
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("MongoDB connected successfully")
}

// GetClient возвращает экземпляр клиента MongoDB
func GetClient() *mongo.Client {
	if client == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return client
}
