package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // 密码
		DB:       0,        // 数据库
		PoolSize: 20,       // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return err
}

func redisExample() {
	if err := rdb.Set("score", 100, 0).Err(); err != nil {
		fmt.Printf("set score failed,err:%v\n", err)
		return
	}

	if score, err := rdb.Get("score").Result(); err != nil {
		fmt.Printf("get score failed,err:%v\n", err)
		return
	} else {
		fmt.Println("score=", score)
	}

	// 优先判断Key为空
	if name, err := rdb.Get("name").Result(); err == redis.Nil {
		fmt.Printf("name does not exist")
		return
	} else if err != nil {
		fmt.Printf("get name failed,err:%v\n", err)
		return
	} else {
		fmt.Println("name=", name)
	}
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("Connect init failed,err:%v\n", err)
		return
	}

	defer rdb.Close()

	redisExample()

}
