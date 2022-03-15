package handler

import (
	"easychat/common/message"
	"easychat/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserHandler struct {
	Conn net.Conn
}

func (this *UserHandler) LoginHandler(mes *message.Message) (err error) {
	//1 从message中取出data
	var loginMes message.LoginMes
	err=json.Unmarshal([]byte(mes.Data),&loginMes)
	if err!=nil{
		fmt.Println("login mes unmarshal err=",err)
		return
	}
	//返回的消息
	var resMes message.Message
	resMes.Type= message.LOGINRESMES

	var loginResMes message.LoginResMes

	if loginMes.UserID==100&&loginMes.UserPSW=="123"{
		//合法
		loginResMes.Code=200
	}else{
		//不合法
		loginResMes.Code=500
		loginResMes.Error="illegal user"
	}

	//序列化登录信息
	data,err:=json.Marshal(loginResMes)
	if err!=nil{
		fmt.Println("loginResMes marshal err=",err)
		return
	}

	//序列化要发送的消息
	resMes.Data=string(data)
	data,err=json.Marshal(resMes)
	if err!=nil{
		fmt.Println("resMes marshal err=",err)
		return
	}

	tf:=&utils.Transfer{
		Conn: this.Conn,
	}

	err= tf.WritePkg(data)

	return
}
