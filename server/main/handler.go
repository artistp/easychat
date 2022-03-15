package main

import (
	"easychat/common/message"
	"easychat/common/utils"
	handler2 "easychat/server/handler"
	"fmt"
	"io"
	"net"
)

type DispatchHandler struct {
	Conn net.Conn
}

func (this *DispatchHandler) connHandler()(error) {
	for {
		tf:=&utils.Transfer{
			Conn: this.Conn,
		}

		mes,err:= tf.ReadPkg()
		if err!=nil{
			if err!=io.EOF{
				fmt.Println("readPkg err=",err)
			}
			return err
		}
		err=this.dispatchHandler(&mes)
		return err
	}
}

//根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *DispatchHandler)dispatchHandler(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LOGINMES:
		uh:=&handler2.UserHandler{
		Conn: this.Conn,
		}
		err=uh.LoginHandler(mes)
	case message.LOGOUTMES:
		fmt.Println("logout")
	case message.REGISTERMES:
		fmt.Println("register")
	default:
		fmt.Println("err message!")
	}
	return
}