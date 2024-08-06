package middleware

import (
	"github.com/gin-gonic/gin"
	"hbbapi/app/enum/cache_enum"
	"hbbapi/util/cache"
	"hbbapi/util/response"
	"math/rand"
	"strconv"
)

func Close() gin.HandlerFunc {
	return func(c *gin.Context) {
		//服务器维护状态
		isMaintain, err := cache.GetInstance().Get(cache_enum.SysMaintain)
		if err == nil {
			response.Fail(c, "系统维护中，预计"+isMaintain+"结束", gin.H{})
			c.Abort()
			return
		}
		//服务器升级状态
		isUpgrade, err := cache.GetInstance().GetInt(cache_enum.SysUpgrade)
		if err == nil {
			response.Fail(c, "服务器被挤爆了，紧急升级中，预计"+strconv.Itoa(isUpgrade)+"分钟后恢复，请稍等", gin.H{})
			c.Abort()
			return
		}
		//服务器限流
		_, err = cache.GetInstance().GetInt(cache_enum.SysFilter)
		if err == nil {
			r := rand.Intn(2)
			if r == 1 {
				response.Fail(c, "服务器繁忙，请稍后再来", gin.H{})
				return
			}
		}
		c.Next()

	}
}
