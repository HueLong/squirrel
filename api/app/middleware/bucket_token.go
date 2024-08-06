package middleware

import (
	"github.com/gin-gonic/gin"
	"hbbapi/util/response"
	"sync"
)

type TokenStack struct {
	Mu    sync.Mutex
	Count int
	Max   int
}

const max = 500

var TokenStackIns = TokenStack{
	Count: max,
	Max:   max,
}

func BucketToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		TokenStackIns.Mu.Lock()
		if TokenStackIns.Count <= 0 {
			response.Fail(c, "排队中，请稍后……", gin.H{})
			c.Abort()
		}
		TokenStackIns.Count--
		TokenStackIns.Mu.Unlock()
		c.Next()
	}
}
