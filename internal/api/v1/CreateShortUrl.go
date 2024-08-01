package v1

import (
	"Shortify/internal/logger"
	"Shortify/internal/models"
	"Shortify/internal/mongodb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "http://localhost:8080/"

// CreateShortUrl godoc
// @Summary Create a short URL
// @Description Create a new short URL from a long URL
// @Tags urls
// @Accept  json
// @Produce  json
// @Param url query string true "The long URL to shorten"
// @Success 201 {object} map[string]string "Successfully created short URL"
// @Failure 400 {object} map[string]string "Invalid URL or missing parameter"
// @Router /api/v1/shortify [post]
func CreateShortUrl(ctx *gin.Context) {
	start := time.Now()
	var newUrl models.Url
	logger := logger.GetLogger()

	uri := ctx.Query("url")
	if uri == "" {
		logger.Warn("URL parameter is missing", zap.String("method", "CreateShortUrl"))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is missing"})
		return
	}

	_, err := url.ParseRequestURI(uri)
	if err != nil {
		logger.Warn("Invalid URL", zap.String("url", uri), zap.String("method", "CreateShortUrl"))
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
		logger.Info("Short URL already exists", zap.String("short_url", fullShortUrl), zap.String("method", "CreateShortUrl"))
		ctx.JSON(http.StatusOK, gin.H{"short_url": fullShortUrl})
		return
	}

	newUrl.Url = uri
	newUrl.ShortUrl = generateUniqueShortUrl(collection)
	newUrl.CreatedAt = start
	newUrl.LastVisit = start

	_, err = collection.InsertOne(context.Background(), newUrl)
	if err != nil {
		logger.Fatal("Failed to insert new URL", zap.Error(err))
	}

	fullShortUrl := baseURL + newUrl.ShortUrl
	logger.Info("Short URL created", zap.String("short_url", fullShortUrl), zap.String("method", "CreateShortUrl"), zap.Duration("duration", time.Since(start)))
	ctx.JSON(http.StatusCreated, gin.H{"short_url": fullShortUrl})
}

func generateUniqueShortUrl(collection *mongo.Collection) string {
	logger := logger.GetLogger()
	for {
		shortUrl := generateShortUrl()
		filter := bson.M{"shorturl": shortUrl}

		var result bson.M
		err := collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil && err == mongo.ErrNoDocuments {
			return shortUrl
		} else if err != nil {
			logger.Fatal("Failed to check if short URL is unique", zap.Error(err))
		}
	}
}

func generateShortUrl() string {
	logger := logger.GetLogger()
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		logger.Fatal("Failed to generate random bytes", zap.Error(err))
	}
	return base64.URLEncoding.EncodeToString(b)
}
