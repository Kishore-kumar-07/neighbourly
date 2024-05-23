package main

import (
	"github.com/Kishore-kumar-07/neighbourly/backend/routes"
)

func main() {
	r := routes.SetupRouter();
	r.Run(":8000");
}