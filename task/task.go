package task

import (
	"admin-api/models/redis"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

//定时任务
func init() {
	go tenMinute()
}

//十分钟一次
func tenMinute() {
	for range time.Tick(time.Second * 600) {
		deltoken()
	}
}

//删除过期token
func deltoken() {
	list, err := redis.Hgetall(`admin_token`)
	if err != nil {
		return
	}
	tokentime, err := beego.AppConfig.Int("tokenTIme")
	if err != nil {
		tokentime = 86400
	}
	time1 := int(time.Now().Unix()) - tokentime
	for _, v := range list {
		for key, val := range v {
			if time2, _ := strconv.Atoi(val[:10]); time1 > time2 {
				redis.Hdel(`admin_token`, key)
			}
		}
	}
}
