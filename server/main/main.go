package main

import (
	"easychat/server/dao"
	"easychat/server/handler"
	"easychat/server/service"
	"easychat/server/source"
	"fmt"
	"net"
	"time"
)

func main() {
	//服务器启动时就初始化连接池
	source.InitPool("10.16.65.76:6379", 16, 0, 300*time.Second)
	//初始化userDao
	dao.InitUserDao(source.PoolInstance)
	service.InitUserService(dao.UserDaoInstance)
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("lister err=", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err=", err)
		} else {
			go Connect(conn)
		}
	}
}

func Connect(conn net.Conn) {
	defer conn.Close()
	ha := &handler.DispatchHandler{
		Conn: conn,
	}
	err := ha.ReadMesHandler()
	if err != nil {
		fmt.Println("客户端与服务器通信协程错误err=", err)
		return
	}
}
