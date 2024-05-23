package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://murugaperumalr2004:murugaperumal@muruga.35stdjd.mongodb.net/?retryWrites=true&w=majority&appName=muruga")

	client, err := mongo.Connect(context.Background(), clientOptions);
	if(err != nil){
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}