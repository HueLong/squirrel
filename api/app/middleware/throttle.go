package middleware

import (
	"github.com/gin-gonic/gin"
	"hbbapi/util/cache"
	"hbbapi/util/response"
)

// Throttle api访问限流
//GET请求，每个IP UID
//POST每个IP
func Throttle() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.FullPath() + c.ClientIP()
		var isFrequent bool
		isFrequent = false
		m := c.Request.Method
		//允许每秒钟访问reqTimes次
		var reqTimes int
		if m == "GET" || m == "OPTION" {
			reqTimes = 10
		} else {
			reqTimes = 5
		}
		initNum := 10
		times, _ := cache.GetInstance().Inc(key)
		if times < initNum {
			_ = cache.GetInstance().Set(key, initNum, 1)
		}
		if times >= initNum+reqTimes {
			isFrequent = true
		}
		if isFrequent {
			response.Fail(c, "哎呀，超速了，再试一次吧", gin.H{})
			c.Abort()
			return
		}
		c.Next()
	}
}
