package router

import (
	"hbbapi/app/controller"

	"github.com/gin-gonic/gin"
)

func InitApi(engine *gin.Engine) {
	api := engine.Group("api")
	//2.1.0版本

	ad := api.Group("ad")
	{
		ad.GET("show_position", controller.GetAdByShowPosition)
		ad.GET("list", controller.GetAdList)
	}
}
