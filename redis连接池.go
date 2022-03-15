package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool=&redis.Pool{
		MaxIdle: 8, //最大活跃连接池数量
		MaxActive: 0,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 300, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","10.16.92.63:6379")
		},
	}
}

func main() {
	//从连接池取数据
	conn:=pool.Get()
	defer conn.Close()
	r,err:=redis.String(conn.Do("Get","name"))
	if err!=nil{
		fmt.Printf("Redis Get err=",err)
		return
	}
	fmt.Println(r)
}

