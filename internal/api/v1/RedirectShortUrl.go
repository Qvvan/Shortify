package v1

import (
	"Shortify/internal/models"
	"Shortify/internal/mongodb"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func RedirectShortUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	log.Println("Received short URL:", shortUrl)

	client := mongodb.GetClient()
	collection := client.Database("shortify").Collection("urls")

	var existingUrl models.Url
	filter := bson.M{"shorturl": shortUrl} // Измените здесь
	log.Println("Filter for MongoDB query:", filter)

	err := collection.FindOne(context.Background(), filter).Decode(&existingUrl)
	if err != nil {
		log.Println("Error finding short URL:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	log.Println("Found URL:", existingUrl.Url)

	_, err = collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": bson.M{"lastvisit": time.Now()}}, // Используйте правильное имя поля
	)
	if err != nil {
		log.Println("Failed to update last visit:", err)
	}

	ctx.Redirect(http.StatusFound, existingUrl.Url)
}
