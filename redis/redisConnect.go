package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.dail err", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("Set", "annimal", "cat")
	if err != nil {
		fmt.Println("redis set err", err)
		return
	}

	_, err = conn.Do("HSet", "use1", "name", "chen", "age", 30)
	if err != nil {
		fmt.Println("redis hset err", err)
		return
	}
	info, err := redis.String(conn.Do("HGet", "use1", "age"))
	if err != nil {
		fmt.Println("redis get err", err)
		return
	}
	//info,err := redis.String(conn.Do("Get","annimal"))
	//if err != nil{
	//	fmt.Println("redis get err",err)
	//	return
	//}
	fmt.Println("getInfo", info)

}
