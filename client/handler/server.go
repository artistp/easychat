package handler

import (
	"easychat/common/message"
	"easychat/common/utils"
	"fmt"
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
	var content string
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		showUserTable()
	case 2:
		fmt.Println("输入聊天信息：")
		fmt.Scanf("%s\n", &content)
		SmsHandlerInstance.SendGroupSms(content)
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		//客户端不停地读取服务器发送的消息
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg err=", err)
			return
		}
		switch mes.Type {
		case message.NOTIFYUSERSTATUSMES:
			//1.取出mes中的data
			//2.把用户的信息状态保存在客户端的map[int]User中
			updateUserTable(&mes)
		case message.SMSMES:
			showGroupSms(&mes)
		default:
			fmt.Println("服务器返回了未知消息类型！")
		}
	}
}
