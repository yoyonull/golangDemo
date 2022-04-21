package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	fmt.Println(1111)
	//程序启动，初始化链接
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	fmt.Println(222)
	conn := pool.Get()
	defer conn.Close()

	info, err := redis.String(conn.Do("HGet", "use1", "age"))
	if err != nil {
		fmt.Println("redis get err", err)
		return
	}
	fmt.Println("getInfo", info)

	info, err = redis.String(conn.Do("Get", "annimal"))
	if err != nil {
		fmt.Println("redis get err", err)
		return
	}
	fmt.Println("getInfo", info)

}
