package redis

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var config redis.Options

func init() {
	addr := beego.AppConfig.String("redis::addr")
	password := beego.AppConfig.String("redis::password")
	config = redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	}
}
func Set(key, value string, times time.Duration) error {
	client := redis.NewClient(&config)
	defer client.Close()
	err := client.Set(key, value, times*time.Second).Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

func Get(key string) (string, error) {
	client := redis.NewClient(&config)
	defer client.Close()
	value, err := client.Get(key).Result()
	if err != nil {
		fmt.Printf(err.Error())
		return "", err
	}
	return value, err
}
func Ping() {
	// pong, err := client.Ping().Result()
	// if err != nil {
	// 	fmt.Printf("", err.Error())
	// 	return
	// }
	// fmt.Printf("ping result: %s\n", pong)
}

// 返回键值剩余时间
func TTL() {
	// duration, err = client.TTL("foo1").Result()
	// if err != nil {
	// 	fmt.Printf("try get ttl of key[foo1] error[%s]\n", err.Error())
	// 	return
	// }

	// fmt.Printf("key[foo1]' ttl is [%s] %ds\n",
	// 	duration.String(), int64(duration.Seconds()))
}
func Rpush() {
	// err := client.RPush("tqueue", "tmsg1").Err()
	// if err != nil {
	// 	fmt.Printf("rpush list[tqueue] error[%s]\n", err.Error())
	// 	return
	// }
}

func Llen() {
	// list_len, err := client.LLen("tqueue").Result()
	// if err != nil {
	// 	fmt.Printf("try get len of list[tqueue] error[%s]\n",
	// 		err.Error())
	// 	return
	// }
	// fmt.Printf("len of list[tqueue] is %d\n", list_len)
}

func Hset(name, key, value string) error {
	client := redis.NewClient(&config)
	defer client.Close()
	err := client.HSet(name, key, value).Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

func Hget(name, key string) (string, error) {
	client := redis.NewClient(&config)
	defer client.Close()
	value, err := client.HGet(name, key).Result()
	if err != nil {
		fmt.Printf(err.Error())
		return "", err
	}
	return value, err
}

func Hgetall(name string) ([]map[string]string, error) {
	client := redis.NewClient(&config)
	defer client.Close()
	var arr []map[string]string
	result_kv, err := client.HGetAll(name).Result()
	if err != nil {
		fmt.Printf(err.Error())
		return arr, err
	}
	for k, v := range result_kv {
		arr = append(arr, map[string]string{k: v})
	}
	return arr, err
}
func Hdel(name, key string) error {
	client := redis.NewClient(&config)
	defer client.Close()
	err := client.HDel(name, key).Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}
func aaa() {
	client := redis.NewClient(&config)
	defer client.Close()

}
func Del(name string) error {
	client := redis.NewClient(&config)
	defer client.Close()
	err := client.Del(name).Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}
func Test() {
	// Set(`qqq`, `eeee`, 10)
	// fmt.Printf("%+v\n", Del(`foo1`))
	// Hset(`name`, `key`, `value`)
	// val, err := Hget(`name`, `key`)
	// fmt.Printf("%+v\n", val, err)
	// Hgetall(`tmap`)
	// Hdel(`tmap`, `aa`)

	// result, err := client.BLPop(time.Second*1, "tqueue").Result()
	// if err != nil {
	// 	fmt.Printf("blpop list[tqueue] error[%s]\n", err.Error())
	// 	return
	// }
	// fmt.Printf("blpop list[tqueue], value[%s]\n", result[1])

	// fmt.Printf("----------------------------------------\n")

	// fmt.Printf("hmap test\n")

	// kv_map := make(map[string]interface{})
	// kv_map["3"] = "f3"
	// kv_map["4"] = "f4"

	// err = client.HMSet("tmap", kv_map).Err()
	// if err != nil {
	// 	fmt.Printf("try mset map[tmap] field[3] field[4] error[%s]\n",
	// 		err.Error())
	// 	return
	// }

	// map_len, err := client.HLen("tmap").Result()
	// if err != nil {
	// 	fmt.Printf("try get len of map[tmap] error[%s]\n", err.Error())
	// 	return
	// }
	// fmt.Printf("len of map[tmap] is %d\n", map_len)

	// fmt.Printf("----------------------------------------\n")

	// fmt.Printf("pubsub test\n")

	// pubsub := client.Subscribe("test_channel")

	// _, err = pubsub.Receive()
	// if err != nil {
	// 	fmt.Printf("try subscribe channel[test_channel] error[%s]\n",
	// 		err.Error())
	// 	return
	// }

	// // go channel to used to receive message
	// ch := pubsub.Channel()

	// // publish a message
	// err = client.Publish("test_channel", "hello").Err()
	// if err != nil {
	// 	fmt.Printf("try publish message to channel[test_channel] error[%s]\n",
	// 		err.Error())
	// 	return
	// }

	// time.AfterFunc(time.Second*2, func() {
	// 	_ = pubsub.Close()
	// })

	// // consume message
	// for {
	// 	msg, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}

	// 	fmt.Printf("recv message[%s] from channel[%s]\n",
	// 		msg.Payload, msg.Channel)
	// }

}
