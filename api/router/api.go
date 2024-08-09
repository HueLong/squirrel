package router

import (
	"hbbapi/app/controller"

	"github.com/gin-gonic/gin"
)

func InitApi(engine *gin.Engine) {
	api := engine.Group("api")
	//2.1.0版本

	picGallery := api.Group("pic_gallery")
	{
		picGallery.GET("list", controller.GetPicGalleryList)
	}
}
