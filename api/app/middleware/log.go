package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/url"
	"time"
)

// LogOut 定义日志输出格式
type LogOut struct {
	Uid       int       `json:"uid"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Ip        string    `json:"ip"`
	Time      time.Time `json:"time"`
	Param     string    `json:"param"`
	UserAgent string    `json:"user_agent"`
	TraceId   string    `json:"trace_id"`
}

// Log 日志中间件，可以开启日志服务
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.NewString()
		c.Set("traceId", traceId)
		useragent := c.Request.Header.Get("user-agent")
		body := c.Request.Body
		param, _ := ioutil.ReadAll(body)
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(param))
		URI, _ := url.QueryUnescape(c.Request.RequestURI)
		log := LogOut{
			Uid:       c.GetInt("uid"),
			Path:      URI,
			Method:    c.Request.Method,
			Ip:        c.ClientIP(),
			Time:      time.Now(),
			Param:     string(param),
			UserAgent: useragent,
			TraceId:   traceId,
		}
		jsonStr, _ := json.Marshal(log)
		fmt.Println(string(jsonStr))
		c.Next()
	}
}
