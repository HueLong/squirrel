package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"hbbapi/util/dingding"
	"os"
	"strconv"
	"time"
)

var redisPool *redis.Pool

// Init 初始化redis
func Init() {
	redisPool = initRedisClientPool()
}

func initRedisClientPool() *redis.Pool {
	maxIdle, _ := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
	maxActive, _ := strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE"))
	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲数
		MaxActive:   maxActive,   //最大活跃数
		IdleTimeout: time.Minute, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Wait:        false,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
			if err != nil {
				fmt.Println("Dial redis连接错误" + err.Error())
				return nil, err
			}
			if redisPass := os.Getenv("REDIS_PASS"); len(redisPass) >= 1 {
				_, err = conn.Do("AUTH", redisPass)
				if err != nil {
					fmt.Println("auth redis连接错误" + err.Error())
					_ = conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxConnLifetime: time.Minute,
	}
	return redisPool
}

// GetOneRedisClient 从连接池获取一个redis连接
func GetOneRedisClient() *Client {
	var oneConn redis.Conn
	maxRetryTimes := 5
	for i := 1; i <= maxRetryTimes; i++ {
		oneConn = redisPool.Get()
		if oneConn.Err() != nil {
			if i == maxRetryTimes {
				_ = dingding.Ding{}.
					SetTitle("retry redis连接出现错误").
					SetContent(oneConn.Err().Error()).
					AtAll().
					Send()
				os.Exit(0)
				return nil
			}
			//如果出现网络短暂的抖动，短暂休眠后，支持自动重连
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	return &Client{oneConn}
}

// Client 定义一个redis客户端结构体
type Client struct {
	Connect redis.Conn
}

// Execute 为redis-go 客户端封装统一操作函数入口
func (r *Client) Execute(cmd string, args ...interface{}) (interface{}, error) {
	defer r.ReleaseOneRedisClient()
	res, err := r.Connect.Do(cmd, args...)
	if err != nil {
		return nil, err
	}
	return res, err
}

// ReleaseOneRedisClient 释放连接到连接池
func (r *Client) ReleaseOneRedisClient() {
	_ = r.Connect.Close()
}

type ClientRes struct {
	Result interface{}
	Error  error
}
