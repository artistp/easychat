package handler

import (
	"easychat/common/utils"
	"fmt"
	"io"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("-------登录成功------------")
	fmt.Println("-------1.显示在线用户列表------------")
	fmt.Println("-------2.发送消息-----------")
	fmt.Println("-------3.信息列表------------")
	fmt.Println("-------4.退出系统------------")
	fmt.Println("请选择：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
		os.Exit(0)
	default:
		fmt.Println("error")
	}
}

//和服务器端保持通信
func serverHandlerMes(conn net.Conn) {
	//创建一个transfer实例，不停地读取服务器的消息
	tf := &utils.Transfer{Conn: conn}
	for {
		//客户端不停地读取服务器发送的消息
		_, err := tf.ReadPkg()
		if err != nil && err != io.EOF {
			fmt.Println("服务器错误=", err)
			return
		}

		//如果读取到了消息
		//fmt.Println("mes from server", mes)
	}
}
