package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Kishore-kumar-07/neighbourly/services/user_service/config"
	"github.com/Kishore-kumar-07/neighbourly/services/user_service/controllers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use()
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

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	return r
}