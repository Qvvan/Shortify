package v1

import (
	"Shortify/internal/logger"
	"Shortify/internal/models"
	"Shortify/internal/mongodb"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

// RedirectShortUrl godoc
// @Summary Redirect to the long URL from the short URL
// @Description Redirects to the long URL using the short URL
// @Tags urls
// @Accept  json
// @Produce  json
// @Param shortUrl path string true "The short URL"
// @Success 302 {object} object{} "Successfully redirected"
// @Failure 404 {object} map[string]string "Short URL not found"
// @Router /api/v1/{shortUrl} [get]
func RedirectShortUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	logger := logger.GetLogger() // Получаем логгер

	logger.Info("Received short URL", zap.String("shortUrl", shortUrl))

	client := mongodb.GetClient()
	collection := client.Database("shortify").Collection("urls")

	var existingUrl models.Url
	filter := bson.M{"shorturl": shortUrl}

	logger.Debug("Filter for MongoDB query", zap.Any("filter", filter))

	err := collection.FindOne(context.Background(), filter).Decode(&existingUrl)
	if err != nil {
		logger.Error("Error finding short URL", zap.Error(err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	logger.Info("Found URL", zap.String("url", existingUrl.Url))

	_, err = collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": bson.M{"lastvisit": time.Now()}},
	)
	if err != nil {
		logger.Error("Failed to update last visit", zap.Error(err))
	}

	ctx.Redirect(http.StatusFound, existingUrl.Url)
}
