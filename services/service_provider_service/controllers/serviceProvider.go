package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/Kishore-kumar-07/neighbourly/services/service_provider_service/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	SeekerEmail string `json:"seekerEmail"`
	ProviderEmail string `json:"providerEmail"`
	Date string `json:"date"`
	Time string `json:"time"`
	Description string `json:"description"`
	Status string `json:"status"`
	Seekername string `json:"seekername"`
	Seekerphone string `json:"Seekerphone"`
}

func ViewServices (c * gin.Context){
	var service []Service;

	client := config.Client;
	collection := client.Database("muruga").Collection("buyService")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	email := c.GetString("email");
	status := c.Param("status");

	doc, err := collection.Find(ctx, bson.M{"provideremail": email, "status": status});
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

func UpdateServiceStatus (c * gin.Context){
	
	var id primitive.ObjectID;

	client := config.Client;

	collection := client.Database("muruga").Collection("buyService")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": true,
			"message": "Invalid ID",
		})
		return
	}

	status := c.Param("status");

	fmt.Println( id ,status);

	doc, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}});
	if err != nil {
		c.JSON(500, gin.H{
			"error": true,
			"message": "Error while updating service status",
		})
		return
	}

	fmt.Println(doc);

	c.JSON(200, gin.H{
		"error": false,
		"message": "Service status updated",
	})
}