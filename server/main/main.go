package main

import (
	"fmt"
	"net"
)

func main() {
	listen,err:=net.Listen("tcp","127.0.0.1:8888")
	if err!=nil{
		fmt.Println("lister err=",err)
		return
	}
	defer listen.Close()
	for {
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("Accept err=",err)
		}

		go handler(conn)
	}
}
func handler(conn net.Conn)  {
	defer conn.Close()
	ha:=&DispatchHandler{
		Conn: conn,
	}
	err:=ha.connHandler()
	if err!=nil{
		fmt.Println("客户端与服务器通信协程错误err=",err)
		return
	}
}
