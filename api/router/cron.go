package router

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func day() string {
	return "@daily"
}

func hour() string {
	return "@hourly"
}

func minute() string {
	return "* * * * *"
}

func every(str string) string {
	return "@every " + str
}

func dayAt(hour int, minute int) string {
	return fmt.Sprintf("%d %d * * *", minute, hour)
}

func hourAt(minute int) string {
	return fmt.Sprintf("%d * * * *", minute)
}

//InitCronJob 初始化任务
func InitCronJob(eng *cron.Cron) {

}
