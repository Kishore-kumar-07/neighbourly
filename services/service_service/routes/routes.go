package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Kishore-kumar-07/neighbourly/services/service_service/config"
	"github.com/Kishore-kumar-07/neighbourly/services/service_service/controllers"
	"github.com/Kishore-kumar-07/neighbourly/services/service_service/middlewares"
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
	
	r.POST("/setProvider", controllers.SetProvider)

	r.GET("/topRatedProviders", controllers.TopRatedProviders)
	r.POST("/searchService", controllers.SearchService)

	return r
}