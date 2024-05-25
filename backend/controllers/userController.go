package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/Kishore-kumar-07/neighbourly/backend/config"
	"github.com/Kishore-kumar-07/neighbourly/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func SignUp (c *gin.Context) {
	client := config.Client;
	var user models.UserModel;
	user.SetDefaults();

	if err:= c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()});
		return;
	}

	if(user.Email == "" || user.Password == "" || user.Name == "" || user.Phone == "") {
		c.JSON(400, gin.H{"error": "Invalid Credentials"});
		return;
	}

	collection := client.Database("muruga").Collection("users");
	fmt.Println(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := collection.FindOne(ctx, bson.M{"name": user.Name});
	if(doc != nil){
		c.JSON(400, gin.H{"error": "User already exists"});
		return;
	}

	doc = collection.FindOne(ctx, bson.M{"name": user.Email});
	if(doc != nil){
		c.JSON(400, gin.H{"error": "Email already exists"});
		return;
	}

	doc = collection.FindOne(ctx, bson.M{"name": user.Phone});
	if(doc != nil){
		c.JSON(400, gin.H{"error": "Phone number already exists"});
		return;
	}

	res, err := collection.InsertOne(ctx, user);
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while inserting user"});
		return;
	}
	c.JSON(200, gin.H{"message": "User created successfully", "id": res.InsertedID});
	return;
}