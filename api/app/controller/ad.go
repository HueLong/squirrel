package controller

import (
	"github.com/gin-gonic/gin"
	"hbbapi/app/service"
	"hbbapi/app/validator"
	"hbbapi/util/response"
	"hbbapi/util/validate"
)

// GetAdByShowPosition 根据广告展示位置标识获取广告信息
func GetAdByShowPosition(ctx *gin.Context) {
	var adParam validator.AdShowPositionGet
	if err := validate.Validate(ctx.ShouldBindQuery, &adParam); err != nil {
		response.Fail(ctx, err.Error(), gin.H{})
		return
	}
	data := service.AdService{}.GetAdByShowPosition(adParam.ShowPosition)
	response.Success(ctx, gin.H{"ad_info": data})
}

// GetAdList 获取广告列表
func GetAdList(ctx *gin.Context) {
	var adParam validator.AdListGet
	if err := validate.Validate(ctx.ShouldBindQuery, &adParam); err != nil {
		response.Fail(ctx, err.Error(), gin.H{})
		return
	}
	data := service.AdService{}.GetAdList(adParam.ShowPosition)
	response.Success(ctx, gin.H{"ad_list": data})
}
