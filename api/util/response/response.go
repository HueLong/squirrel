package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success 返回正确结果
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errmsg": "成功", "data": data})
}

func Fail(c *gin.Context, message string, data interface{}) {
	var newData = gin.H{}
	newData["trace_id"], _ = c.Get("traceId")
	newData["errmsg"] = message
	newData["ip"] = c.ClientIP()
	newData["path"] = c.Request.RequestURI
	jsonStr, _ := json.Marshal(newData)
	fmt.Println(string(jsonStr))
	c.JSON(http.StatusOK, gin.H{"errcode": 1, "errmsg": message, "data": data})
	c.Abort()
}

func FailWithCode(c *gin.Context, code int, message string, data interface{}) {
	var newData = gin.H{}
	newData["trace_id"], _ = c.Get("traceId")
	newData["errmsg"] = message
	newData["ip"] = c.ClientIP()
	newData["path"] = c.Request.RequestURI
	jsonStr, _ := json.Marshal(newData)
	if code != 407 {
		fmt.Println(string(jsonStr))
	}
	c.JSON(http.StatusOK, gin.H{"errcode": code, "errmsg": message, "data": data})
	c.Abort()
}
