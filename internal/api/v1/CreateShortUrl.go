package v1

import (
	"Shortify/internal/models"
	"Shortify/internal/mongodb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "http://localhost:8080/"

func CreateShortUrl(ctx *gin.Context) {
	start := time.Now()
	var newUrl models.Url

	uri := ctx.Query("url")
	if uri == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is missing"})
		return
	}

	_, err := url.ParseRequestURI(uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	client := mongodb.GetClient()
	collection := client.Database("shortify").Collection("urls")

	var existingUrl models.Url
	filter := bson.M{"url": uri}
	err = collection.FindOne(context.Background(), filter).Decode(&existingUrl)
	if err == nil {
		fullShortUrl := baseURL + existingUrl.ShortUrl
		ctx.JSON(http.StatusOK, gin.H{"short_url": fullShortUrl})
		return
	}

	newUrl.Url = uri
	newUrl.ShortUrl = generateUniqueShortUrl(collection)
	newUrl.CreatedAt = start
	newUrl.LastVisit = start

	// Сохраняем новый URL в базе данных
	_, err = collection.InsertOne(context.Background(), newUrl)
	if err != nil {
		log.Fatal("Failed to insert new URL:", err)
	}

	fullShortUrl := baseURL + newUrl.ShortUrl
	ctx.JSON(http.StatusCreated, gin.H{"short_url": fullShortUrl})
}

func generateUniqueShortUrl(collection *mongo.Collection) string {
	for {
		shortUrl := generateShortUrl()
		filter := bson.M{"short_url": shortUrl}

		// Используем FindOne вместо CountDocuments
		var result bson.M
		err := collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil && err == mongo.ErrNoDocuments {
			// Если документ не найден, короткий URL уникален
			return shortUrl
		} else if err != nil {
			log.Fatal("Failed to check if short URL is unique:", err)
		}
	}
}

func generateShortUrl() string {
	b := make([]byte, 6) // 6 байт для генерации строки
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Failed to generate random bytes:", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
