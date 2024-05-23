package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Kishore-kumar-07/neighbourly/backend/config"
	"context"
)

func SetupRouter() *gin.Engine {
	
	r := gin.Default()
	client, err := config.ConnectDB();
	if(err != nil){
		panic(err)
	}
	defer client.Disconnect(context.Background())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	return r
}