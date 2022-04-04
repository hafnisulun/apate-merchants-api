package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-merchants-api/controllers"
)

func Merchants(router *gin.RouterGroup) {
	merchantsGroup := router.Group("/merchants")
	{
		merchant := new(controllers.MerchantController)
		merchantsGroup.GET("", merchant.FindAll)
		merchantsGroup.GET("/:merchant_uuid", merchant.FindOne)
		merchantsGroup.POST("", merchant.Create)
	}
}
