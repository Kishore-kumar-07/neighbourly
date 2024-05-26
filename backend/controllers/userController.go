package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Kishore-kumar-07/neighbourly/backend/config"
	"github.com/Kishore-kumar-07/neighbourly/backend/models"
	"github.com/dgrijalva/jwt-go"
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
		c.JSON(400, gin.H{"error":true, "message": "Invalid Credentials"});
		return;
	}

	collection := client.Database("muruga").Collection("users");
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var res models.UserModel;

	doc := collection.FindOne(ctx, bson.M{"email": user.Email});
	if err := doc.Decode(&res); err ==nil {
		fmt.Println(res.Email);
		c.JSON(400, gin.H{"error":true, "message": "Email already exists"});
		return;
	}

	doc = collection.FindOne(ctx, bson.M{"email": user.Phone});
	if err := doc.Decode(&res); err==nil {
		c.JSON(400, gin.H{"error":true, "message": "Phone number already exists"});
		return;
	}

	_, err := collection.InsertOne(ctx, user);
	if err != nil {
		c.JSON(500, gin.H{"error":true, "message": "Error while inserting user"});
		return;
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role" : user.Role,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo": "bar",
	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })
	
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{"error":false, "message": "User created successfully", "token": tokenString})
}

func Login (c *gin.Context) {
	fmt.Println("check");
	client := config.Client;
	var user models.UserModel;
	user.SetDefaults();

	type LoginReq struct {
		Email string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}

	var loginReq LoginReq;

	if err:= c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error":true, "message": err.Error()});
		return;
	}

	if(loginReq.Email == "" || loginReq.Password == "") {
		c.JSON(400, gin.H{"error":true, "message": "Invalid Credentials"});
		return;
	}

	collection := client.Database("muruga").Collection("users");
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := collection.FindOne(ctx, bson.M{"email": loginReq.Email});
	if(doc == nil){
		c.JSON(400, gin.H{"error":true, "message": "User does not exist"});
		return;
	}


	if err:= doc.Decode(&user); err != nil {
		c.JSON(500, gin.H{"error":true, "message": "user not found"});
		return
	}

	if(user.Password != loginReq.Password){
		c.JSON(400, gin.H{"error":true, "message": "Invalid Password"});
		return;
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role" : user.Role,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{"error":false, "message": "User logged in successfully", "token": tokenString});
}