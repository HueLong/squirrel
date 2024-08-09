package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hbbapi/app/middleware"
	"hbbapi/router"
	"hbbapi/util/db"
	"hbbapi/util/redis"
	"hbbapi/util/validate"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if ginEnv := os.Getenv("GIN_ENV"); ginEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	go func() {
		for {
			middleware.TokenStackIns.Mu.Lock()
			middleware.TokenStackIns.Count = middleware.TokenStackIns.Max
			middleware.TokenStackIns.Mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	validate.Init()
	//初始化数据库
	db.Init()
	redis.Init()
	engine := gin.New()
	engine.Use(middleware.Cors(), middleware.IpBlock(), middleware.Auth(), middleware.Log(), middleware.BucketToken(), middleware.Recovery(), middleware.Throttle())
	router.InitApi(engine)
	err1 := engine.Run("0.0.0.0:8080")
	if err1 != nil {
		fmt.Print("启动出错")
	}
}
