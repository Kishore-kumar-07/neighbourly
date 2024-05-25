package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Kishore-kumar-07/neighbourly/backend/models"
)

func setProvider (c *gin.Context) {
	var serviceProviderModel models.ServiceProviderModel;
	serviceProviderModel.SetDefaults();

	if err := c.BindJSON(&serviceProviderModel); err != nil {
		c.JSON(400, gin.H{
			"error": true,
			"message": err.Error(),
		})
		return
	}

	if( serviceProviderModel.Description == "" || serviceProviderModel.Email == "" || serviceProviderModel.Experience == "" || serviceProviderModel.ServiceDescription == ""){
		c.JSON(400, gin.H{
			"error": true,
			"message": "All fields are required",
		})
		return
	}

}