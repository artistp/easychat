package source

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var PoolInstance *redis.Pool

//address=10.16.92.63:6379
func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	PoolInstance = &redis.Pool{
		MaxIdle:     maxIdle,     //最大活跃连接池数量
		MaxActive:   maxActive,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
