package handler

import (
	"easychat/common/message"
	"easychat/common/utils"
	"fmt"
	"io"
	"net"
)

type DispatchHandler struct {
	Conn net.Conn
}

func (this *DispatchHandler) ReadMesHandler() (err error) {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，断开连接", err)
			} else {
				fmt.Println("readPkg err=", err)
			}
			return err
		}
		err = this.dispatchHandler(&mes)
	}
}

//根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *DispatchHandler) dispatchHandler(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LOGINMES:
		uh := &UserHandler{
			Conn: this.Conn,
		}
		err = uh.LoginHandler(mes)
	case message.LOGOUTMES:
		fmt.Println("logout")
	case message.REGISTERMES:
		uh := UserHandler{
			Conn: this.Conn,
		}
		err = uh.RegisterHandler(mes)
	default:
		fmt.Println("err message!")
	}
	return
}
