package drivers

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	re "hbbapi/util/redis"
	"reflect"
)

type Redis struct {
	Client interface{}
}

func (r Redis) Get(key string) (string, error) {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("GET", key)
	resString, err1 := redis.String(res, err)
	return resString, err1
}

func (r Redis) GetInt(key string) (int, error) {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("GET", key)
	resString, err1 := redis.Int(res, err)
	return resString, err1
}

func (r Redis) GetJson(key string, obj interface{}) error {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("GET", key)
	resString, err := redis.String(res, err)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(resString), obj)
	if err != nil {
		return err
	}
	return nil
}

func (r Redis) MultiGetJson(key []string, obj interface{}) error {
	if len(key) > 50 {
		return errors.New("最多同时查询50个，过多请分批查询")
	}
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		return errors.New("obj必须为指针")
	}
	objType = reflect.ValueOf(obj).Elem().Type()
	if objType.Kind() != reflect.Slice && objType.Kind() != reflect.Array {
		return errors.New("obj必须为slice或array")
	}
	var args []interface{}
	for _, v := range key {
		args = append(args, v)
	}
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("MGET", args...)
	resString, err := redis.Strings(res, err)
	if err != nil {
		return err
	}
	objValue := reflect.ValueOf(obj).Elem()
	single := reflect.New(reflect.TypeOf(obj).Elem().Elem()).Interface()
	for _, value := range resString {
		if value == "" {
			continue
		}
		err = json.Unmarshal([]byte(value), single)
		if err != nil {
			return err
		}
		objValue.Set(reflect.Append(objValue, reflect.ValueOf(single).Elem()))
	}
	return nil
}

func (r Redis) MultiSetJson(keyValue map[string]interface{}, expire int) error {
	if len(keyValue) > 1000 {
		return errors.New("最多同时存储1000个，过多请分批查询")
	}
	var args []interface{}
	for k, v := range keyValue {
		jsonByte, jsonErr := json.Marshal(v)
		if jsonErr != nil {
			return jsonErr
		}
		args = append(args, k, jsonByte)
	}
	rc := r.Client.(*re.Client)
	defer func(Connect redis.Conn) {
		err := Connect.Close()
		if err != nil {
			return
		}
	}(rc.Connect)
	err := rc.Connect.Send("MSET", args...)
	if err != nil {
		return err
	}
	for key := range keyValue {
		_ = rc.Connect.Send("EXPIRE", key, expire)
	}
	_, _ = rc.Connect.Do("")
	return err
}

func (r Redis) Set(key string, data interface{}, expire int, others ...interface{}) error {
	var args []interface{}
	rc := r.Client.(*re.Client)
	args = append(args, key, data)
	if expire > 0 {
		args = append(args, "EX", expire)
	}
	if len(others) > 0 {
		args = append(args, others...)
	}
	res, err := rc.Execute("SET", args...)
	_, err1 := redis.String(res, err)
	return err1
}

func (r Redis) SetJson(key string, data interface{}, expire int, others ...interface{}) error {
	var args []interface{}
	jsonString, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		return jsonErr
	}
	args = append(args, key, jsonString)
	if expire > 0 {
		args = append(args, "EX", expire)
	}
	if len(others) > 0 {
		args = append(args, others...)
	}
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("SET", args...)
	_, err1 := redis.String(res, err)
	return err1
}

func (r Redis) Inc(key string) (int, error) {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("INCR", key)
	resInt, err1 := redis.Int(res, err)
	return resInt, err1
}

func (r Redis) IncBy(key string, step int) (int, error) {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("INCRBY", key, step)
	resInt, err1 := redis.Int(res, err)
	return resInt, err1
}

func (r Redis) Dec(key string) (int, error) {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("DECR", key)
	resInt, err1 := redis.Int(res, err)
	return resInt, err1
}

func (r Redis) DecBy(key string, step int) (int, error) {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("DECRBY", key, step)
	resInt, err1 := redis.Int(res, err)
	return resInt, err1
}

func (r Redis) Del(key string) error {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("DEL", key)
	_, err1 := redis.Int(res, err)
	return err1
}

// Exp 设置过期时间
func (r Redis) Exp(key string, expire int) error {
	rc := r.Client.(*re.Client)
	res, err := rc.Execute("EXPIRE", key, expire)
	_, err1 := redis.Int(res, err)
	return err1
}

func (r Redis) Execute(command string, args ...interface{}) (interface{}, error) {
	rc := r.Client.(*re.Client)
	return rc.Execute(command, args...)
}
