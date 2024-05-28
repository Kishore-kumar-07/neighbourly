package controllers

import (
	"context"
	"time"

	"github.com/Kishore-kumar-07/neighbourly/backend/config"
	"github.com/Kishore-kumar-07/neighbourly/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetProvider (c *gin.Context) {
	var service models.ServiceProviderModel;
	service.SetDefaults();

	client := config.Client;

	if err := c.BindJSON(&service); err != nil {
		c.JSON(400, gin.H{
			"error": true,
			"message": err.Error(),
		})
		return
	}

	if( service.Description == "" || service.Experience == "" || service.ServiceDescription == "" || service.Title == ""){
		c.JSON(400, gin.H{
			"error": true,
			"message": "All fields are required",
		})
		return
	}

	service.Email = c.GetString("email");
	if service.Email == "" {
		c.JSON(400, gin.H{
			"error": true,
			"message": "Email is missing",
		})
		return
	}

	collection := client.Database("muruga").Collection("serviceProviders");
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := collection.InsertOne(ctx, service);

	if err != nil {
		c.JSON(500, gin.H{
			"error": true,
			"message": "Error while inserting service provider",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": false,
		"message": "Service provider added successfully",
	})
}

func TopRatedProviders(c *gin.Context) {
    var serviceProviderModel models.ServiceProviderModel
    serviceProviderModel.SetDefaults()

    client := config.Client

    collection := client.Database("muruga").Collection("serviceProviders")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    findOptions := options.Find().SetSort(map[string]int{"rating": -1}).SetLimit(5)

    cursor, err := collection.Find(ctx, bson.M{}, findOptions)
    if err != nil {
        c.JSON(500, gin.H{"error": true, "message": "error finding providers"})
        return
    }
    defer cursor.Close(ctx)

    var serviceProviders []models.ServiceProviderModel
    if err := cursor.All(ctx, &serviceProviders); err != nil {
        c.JSON(500, gin.H{"error": true, "message": "error decoding providers"})
        return
    }

    c.JSON(200, gin.H{
        "error":   false,
        "message": serviceProviders,
    })
}

func SearchService(c *gin.Context){
	var serviceProviders []models.ServiceProviderModel;
	client := config.Client;
	collection := client.Database("muruga").Collection("serviceProviders")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	searchTitle := c.GetString("title");

	filter := bson.M{"title": bson.M{"$regex": searchTitle, "$options": "i"}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(500, gin.H{"error": true, "message": "error finding providers"})
		return
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &serviceProviders); err != nil {
        c.JSON(500, gin.H{"error": true, "message": "error decoding providers"})
        return
    }

	c.JSON(200, gin.H{
		"error":   false,
		"message": serviceProviders,
	})
}