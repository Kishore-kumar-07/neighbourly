package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Kishore-kumar-07/neighbourly/services/service_seeker_service/config"
	"github.com/Kishore-kumar-07/neighbourly/services/service_seeker_service/controllers"
	"github.com/Kishore-kumar-07/neighbourly/services/service_seeker_service/middlewares"
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

	r.POST("/buyService", controllers.BuyService)

	return r
}