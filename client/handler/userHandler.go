package handler

import (
	"easychat/common/message"
	"easychat/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserHandler struct {

}

func (this *UserHandler) Login(userID int, userPSW string) (err error)  {
	//1 连接到服务器
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err!=nil{
		fmt.Println("Dial err=",err)
		return
	}
	defer conn.Close()

	//2 准备通过conn发送的消息
	var sendMes message.Message
	sendMes.Type= message.LOGINMES

	//3 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserID=userID
	loginMes.UserPSW=userPSW

	//4 序列化loginMes
	data,err:=json.Marshal(loginMes)
	if err!=nil{
		fmt.Println("loginMes json marshal err=",err)
		return
	}
	sendMes.Data=string(data)

	//5 将mes进行序列化
	data,err=json.Marshal(sendMes)
	if err!=nil{
		fmt.Println("sendMes json marshal err=",err)
		return
	}

	tf:=&utils.Transfer{
		Conn: conn,
	}
	//6 data 为需要发送给服务器的消息
	//6.1 先把data的长度发送给服务器
	//先获取到data的长度，然后转为一个表示长度的切片
	err= tf.WritePkg(data)
	if err!=nil{
		fmt.Println("send mes err=",err)
		return
	}

	//处理服务器返回的消息
	resMes,err:= tf.ReadPkg()
	if err!=nil{
		fmt.Println("read mes err=",err)
		return
	}

	var loginResMes message.LoginResMes
	err=json.Unmarshal([]byte(resMes.Data),&loginResMes)
	if loginResMes.Code==200{
		fmt.Println("success login")

		//需要启动一个协程，
		//该协程保持和服务器的通信，如果服务器有数据推送
		//则接收并显示


		//1显示登录成功的菜单
		for{
			ShowMenu()
		}
	}else if loginResMes.Code==500{
		fmt.Println("err login")
	}
	return
}
