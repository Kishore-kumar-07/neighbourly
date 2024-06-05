package main

import (
	"fmt"

	"github.com/Kishore-kumar-07/neighbourly/services/service_service/routes"
	"github.com/joho/godotenv"
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