package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func urlShortener(longUrl string, collection *mongo.Collection, ctx context.Context) string {
	urlencoder := sha256.New()
	urlencoder.Write([]byte(longUrl))
	encodedurl := urlencoder.Sum(nil)
	encodedStr := hex.EncodeToString(encodedurl)

	shorturl := encodedStr[:5]

	repeatedUrl := isUrlRepeated(ctx, collection, shorturl)

	if !repeatedUrl {
		_, err := collection.InsertOne(ctx, bson.M{
			"_id":     shorturl,
			"longUrl": longUrl,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	return shorturl
}
