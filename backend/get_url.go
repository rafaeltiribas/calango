package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func get_url(ctx context.Context, c *gin.Context, collection *mongo.Collection, shortUrl string) {
	var result struct {
		LongUrl string `bson:"longUrl"`
	}

	err := collection.FindOne(ctx, bson.M{"_id": shortUrl}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, result.LongUrl)
}

func consult_url(ctx context.Context, collection *mongo.Collection, shortUrl string) bool {
	var result struct {
		LongUrl string `bson:"longUrl"`
	}

	err := collection.FindOne(ctx, bson.M{"_id": shortUrl}).Decode(&result)
	if err != nil {
		return false
	}
	return true
}
