package queue

import (
	"hbbapi/util/queue/drivers"
	"hbbapi/util/queue/router"
)

type Queue interface {
	Enqueue(scriptName string, data interface{}) bool
	EnqueueWithName(queueName, scriptName string, data interface{}) bool
	New() *router.Engine
	Run(engine *router.Engine) error
}

func GetInstance() Queue {
	var queue Queue
	driver := new(drivers.Redis)
	queue = driver
	return queue
}
