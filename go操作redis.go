package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	//连接Redis
	conn,err:=redis.Dial("tcp","10.16.92.63:6379")
	if err!=nil{
		fmt.Printf("Redis conn err=%v\n",err)
		return
	}

	defer conn.Close()
	//通过go向Redis写入数据
	//_, err=conn.Do("Set","name","xiaowu")
	//if err!=nil{
	//	fmt.Printf("Redis set err=%v\n",err)
	//	return
	//}
	//通过go获取Redis的 数据
	r, err:=redis.String(conn.Do("Get","name"))
	if err!=nil{
		fmt.Printf("Redis set err=%v\n",err)
		return
	}
	fmt.Println(r)
}


