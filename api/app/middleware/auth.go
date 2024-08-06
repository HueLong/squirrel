package middleware

import (
	"github.com/gin-gonic/gin"
	"hbbapi/config"
	"hbbapi/util/jwt"
	"hbbapi/util/response"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var canSkip = false
		var notLogin = true
		for _, v := range config.Excepts {
			if v == c.FullPath() {
				canSkip = true
			}
		}
		auth := c.GetHeader("AUTHORIZATION")
		tokenArr := strings.Split(auth, "Bearer ")
		if len(tokenArr) == 2 && len(tokenArr[1]) > 5 {
			claim, err := jwt.SignObj.ParseToken(tokenArr[1])
			if err == nil {
				c.Set("uid", claim.Uid)
				c.Next()
				notLogin = false
			}
		}
		if !canSkip && notLogin {
			response.FailWithCode(c, 401, "未登录", gin.H{})
			c.Abort()
			return
		}
	}
}
