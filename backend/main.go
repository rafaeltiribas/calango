package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func connectToMongoDB() (*mongo.Client, context.Context, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, nil
}

func main() {
	client, ctx, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("calango").Collection("url")

	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/calango", func(c *gin.Context) {
		longUrl := c.PostForm("url")

		reqCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shortUrl := urlShortener(longUrl, collection, reqCtx)

		c.JSON(http.StatusOK, gin.H{
			"shortUrl": shortUrl,
		})
	})

	r.GET("/:url", func(c *gin.Context) {
		shortUrl := c.Param("url")

		reqCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var result struct {
			LongUrl string `bson:"longUrl"`
		}
		err := collection.FindOne(reqCtx, bson.M{"_id": shortUrl}).Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Redirect(http.StatusMovedPermanently, result.LongUrl)
	})

	r.Run()
}
