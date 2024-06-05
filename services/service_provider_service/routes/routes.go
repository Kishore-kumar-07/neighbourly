package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Kishore-kumar-07/neighbourly/services/service_provider_service/config"
	"github.com/Kishore-kumar-07/neighbourly/services/service_provider_service/controllers"
	"github.com/Kishore-kumar-07/neighbourly/services/service_provider_service/middlewares"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	err := config.ConnectDB();
	if(err != nil){
		fmt.Println(err)
	}
	// defer client.Disconnect(context.Background())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/viewServices/:status", controllers.ViewServices)
	r.POST("/updateService/:id/:status", controllers.UpdateServiceStatus)

	return r
}