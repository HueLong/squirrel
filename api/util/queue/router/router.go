package router

import (
	"errors"
	"hbbapi/util/queue/log"
)

type EngineInterface interface {
	Get(str string, f func(jsonByte []byte) error)
	Exec(funcName string, jsonByte []byte, log log.Log) error
}

type Engine struct {
	routerMap map[string]func(jsonByte []byte, log log.Log) error
}

func (r *Engine) Get(str string, f func(jsonByte []byte, log log.Log) error) {
	if nil == r.routerMap {
		r.routerMap = make(map[string]func(jsonByte []byte, log log.Log) error)
	}
	r.routerMap[str] = f
}

func (r *Engine) Exec(funcName string, jsonByte []byte, log log.Log) error {
	if funcName, ok := r.routerMap[funcName]; ok {
		return funcName(jsonByte, log)
	}
	return errors.New("函数不存在")
}
