package controllers

import (
	"context"
	"time"

	"github.com/Kishore-kumar-07/neighbourly/services/service_seeker_service/config"
	"github.com/Kishore-kumar-07/neighbourly/services/service_seeker_service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func BuyService(c* gin.Context){
	client := config.Client;


	var buyServiceStruct models.Service;
	buyServiceStruct.SetDefaults();
	var serviceProviderModel models.ServiceProviderModel

	if err := c.ShouldBindJSON(&buyServiceStruct); err != nil {
		c.JSON(400, gin.H{
			"error": true,
			"message": err.Error(),
		})
		return
	}
	
	collection := client.Database("muruga").Collection("serviceProviders");
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := collection.FindOne(ctx, bson.M{"email": buyServiceStruct.ProviderEmail});

	if err := doc.Decode(&serviceProviderModel); err != nil{
		c.JSON(500, gin.H{
			"error": true,
			"message": "No such provider found",
		})
		return
	}

	collection = client.Database("muruga").Collection("buyService");
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	buyServiceStruct.SeekerEmail = c.GetString("email");
	buyServiceStruct.SeekerName = c.GetString("name");
	buyServiceStruct.SeekerPhone = c.GetString("phone");

	_, err := collection.InsertOne(ctx, buyServiceStruct);
	if err != nil {
		c.JSON(500, gin.H{
			"error": true,
			"message": "Error while inserting buy service",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": false,
		"message": "Service bought successfully",
	})
	

}