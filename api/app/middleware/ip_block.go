package middleware

import (
	"github.com/gin-gonic/gin"
	"hbbapi/util/cache"
	"hbbapi/util/response"
	"io"
)

func IpBlock() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		block, err := cache.GetInstance().Get(ip + "_block")
		if err != nil && err == io.ErrShortWrite {
			response.Fail(c, "系统繁忙，请重试", gin.H{})
			c.Abort()
			return
		}
		if len(block) > 0 {
			response.Fail(c, "检测到恶意攻击，关小黑屋24小时", gin.H{})
			c.Abort()
			return
		}
		inc, _ := cache.GetInstance().Inc(ip)
		if inc < 11 {
			_ = cache.GetInstance().Set(ip, 10, 60)
		}
		if inc > 1200 {
			_ = cache.GetInstance().Set(ip+"_block", 1, 86400)
			response.Fail(c, "检测到恶意攻击，关小黑屋24小时", gin.H{})
			c.Abort()
			//关闭缓存连接
			return
		}
		c.Next()
	}
}
