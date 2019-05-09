package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis连接失败", err)
		return
	}
	rep, err := conn.Do("get", "v1")
	result, err := redis.String(rep, err)
	fmt.Println("获得的值是：", result)
}
