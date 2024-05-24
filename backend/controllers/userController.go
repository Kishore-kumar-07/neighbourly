	package controllers

	import (
		"context"
		"time"
		"fmt"

		"github.com/gin-gonic/gin"
		"go.mongodb.org/mongo-driver/bson"
		"github.com/Kishore-kumar-07/neighbourly/backend/config"
	)


	func SignUp (c *gin.Context) {
		client := config.Client;
		fmt.Println("Inside signup controller", client);
		collection := client.Database("muruga").Collection("users");
		fmt.Println(collection)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := collection.InsertOne(ctx, bson.M{"name":"muruga"});
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while inserting user"});
			return;
		}
		c.JSON(200, gin.H{"message": "User created successfully", "id": res.InsertedID});
		return;
	}