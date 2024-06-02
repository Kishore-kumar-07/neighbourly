package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/Kishore-kumar-07/neighbourly/backend/config"
	"github.com/Kishore-kumar-07/neighbourly/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func ViewServices (c * gin.Context){
	var service []models.Service;

	client := config.Client;
	collection := client.Database("muruga").Collection("buyService")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := c.GetString("email");

	doc, err := collection.Find(ctx, bson.M{"provideremail": email});
	if err != nil {
		c.JSON(500, gin.H{
			"error": true,
			"message": "Error while fetching services",
		})
		return
	}

	fmt.Println(doc);
	fmt.Println(email);

	if err := doc.All(ctx, &service); err != nil {
		c.JSON(500, gin.H{
			"error": true,
			"message": "Error while decoding services",
		})
		return
	}

	if(service == nil){
		c.JSON(200, gin.H{
			"error": false,
			"message": "No services found",
		})
		return
	
	}

	c.JSON(200, gin.H{
		"error": false,
		"message": service,
	})

}