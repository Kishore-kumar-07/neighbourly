package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/Kishore-kumar-07/neighbourly/services/user_service/routes"
)

func init() {
	err := godotenv.Load();
	if err != nil {
		panic("Error loading .env file");
	}
	fmt.Println("Environment variables loaded successfully");
}

func main() {
	r := routes.SetupRouter();
	r.Run();
}