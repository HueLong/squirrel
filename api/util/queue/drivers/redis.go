package drivers

import (
	"encoding/json"
	"fmt"
	redis2 "github.com/gomodule/redigo/redis"
	"hbbapi/util/cache"
	"hbbapi/util/queue/log"
	"hbbapi/util/queue/router"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type MqStatus struct {
	mu    sync.Mutex
	isRun bool
}

func (r *MqStatus) isRunning() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.isRun
}

func (r *MqStatus) ChangeRunning(run bool) {
	r.mu.Lock()
	r.isRun = run
	r.mu.Unlock()
}

type Redis struct {
}

var status *MqStatus

// Enqueue 消息入队
func (Redis) Enqueue(scriptName string, data interface{}) bool {
	//定义队列的名称
	key := "queue_streams"
	dataJson, err := json.Marshal(data)
	if err != nil {
		return false
	}
	_, err = cache.GetInstance().Execute("XADD", key, "MAXLEN", 500000, "*", scriptName, dataJson)
	if err != nil {
		return false
	}
	return true
}

func (Redis) EnqueueWithName(queueName, scriptName string, data interface{}) bool {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return false
	}
	_, err = cache.GetInstance().Execute("XADD", queueName, "MAXLEN", 500000, "*", scriptName, dataJson)
	if err != nil {
		return false
	}
	return true
}

// Run 开始消费
func (r Redis) Run(engine *router.Engine) error {
	status = &MqStatus{isRun: true}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-quit
		status.ChangeRunning(false)
	}()
	//最多支持100个协程同时执行
	max := make(chan int, 100)
	for {
		// 如果收到k8s的term信号，则停止运行
		if !status.isRunning() {
			break
		}
		max <- 1
		consumeName := fmt.Sprintf("consume_%d_%d", time.Now().UnixNano(), rand.Int())
		res, readErr := cache.GetInstance().Execute("XREADGROUP", "GROUP", "group1", consumeName,
			"COUNT", 1, "BLOCK", 10000, "STREAMS", "queue_streams", ">")
		messages, msgErr := r.getMessage(res, readErr)
		if msgErr != nil {
			fmt.Println(msgErr.Error())
			continue
		}
		value := messages[0]
		go func(funcName string, params []byte) {
			defer func() {
				if pErr := recover(); pErr != nil {

				}
			}()
			startTime := time.Now().UnixNano()
			var log2 = log.Log{}
			log2.Info("start_time", fmt.Sprintf("%d", startTime))
			defer func(map[string]string) {
				jsonString, _ := json.Marshal(log2.Messages)
				fmt.Println(jsonString)
			}(log2.Messages)
			err := engine.Exec(funcName, params, log2)
			endTime := time.Now().UnixNano()
			log2.Info("end_time", fmt.Sprintf("%d", endTime))
			if err != nil {
				log2.Error(err.Error())
			}
			//执行完释放
			<-max
		}(value.Key, []byte(value.Value))

	}
	time.Sleep(time.Minute * 5)
	return nil
}

func (r Redis) New() *router.Engine {
	eng := new(router.Engine)
	return eng
}

func (r Redis) Ack(messageId string) {
	cache.GetInstance().Execute("XACK", "GROUP", "group1", messageId)
}

type Message struct {
	QueueName string
	MessageId string
	Key       string
	Value     string
}

func (r Redis) getMessage(res interface{}, err error) ([]Message, error) {
	var msg []Message
	valueRes, err1 := redis2.Values(res, err)
	for kIndex := 0; kIndex < len(valueRes); kIndex++ {
		var keyInfo = valueRes[kIndex].([]interface{})
		var key = string(keyInfo[0].([]byte))
		var idList = keyInfo[1].([]interface{})
		for idIndex := 0; idIndex < len(idList); idIndex++ {
			var idInfo = idList[idIndex].([]interface{})
			var id = string(idInfo[0].([]byte))
			var fieldList = idInfo[1].([]interface{})
			var field = string(fieldList[0].([]byte))
			var value = string(fieldList[1].([]byte))
			msg = append(msg, Message{key, id, field, value})
		}
	}
	return msg, err1
}

// 主要队列
func (r Redis) primary() {

}

// 重试队列
func (r Redis) retry() {
	tick := time.NewTicker(time.Second)
	for {
		<-tick.C

	}
}

// 死信队列
func (r Redis) dead() {
	tick := time.NewTicker(time.Second * 10)
	for {
		<-tick.C

	}
}
