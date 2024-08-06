package cron

import (
	"hbbapi/util/cache"
	"time"
)

func Lock(str string) error {
	cacheKey := str + "_" + time.Now().Format("2006-01-02 15:04")
	_, err := cache.GetInstance().Execute("SET", cacheKey, "1", "EX", "61", "NX")
	return err
}
