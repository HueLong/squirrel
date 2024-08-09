package controller

import (
	"github.com/gin-gonic/gin"
	"hbbapi/app/service"
	"hbbapi/app/validator"
	"hbbapi/util/response"
	"hbbapi/util/validate"
)

// GetPicGalleryList 获取图库列表
func GetPicGalleryList(ctx *gin.Context) {
	var param validator.PicGalleryList
	if err := validate.Validate(ctx.ShouldBindQuery, &param); err != nil {
		response.Fail(ctx, err.Error(), gin.H{})
		return
	}
	data := service.PicGalleryService{}.GetPicList(param.Page, param.Size)
	response.Success(ctx, gin.H{"list": data})
}
