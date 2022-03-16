package handler

import (
	"easychat/common/message"
	"easychat/common/utils"
	"easychat/server/service"
	"encoding/json"
	"fmt"
	"net"
)

type UserHandler struct {
	Conn   net.Conn
	UserId int
}

func (this *UserHandler) LoginHandler(mes *message.Message) (err error) {
	//1 从message中取出data
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("login mes unmarshal err=", err)
		return
	}
	//返回的消息
	var resMes message.Message
	resMes.Type = message.LOGINRESMES

	var loginResMes message.LoginResMes

	user, err := service.UserServiceInstance.UserLoginService(loginMes.User.UserId, loginMes.User.UserPwd)
	if err != nil {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = err.Error()
	} else {
		//合法
		loginResMes.Code = 200
		this.UserId = user.UserId
		UserMgr.AddOnlineUser(this)
		for id, _ := range UserMgr.OnlineUserTable {
			loginResMes.OnlineUserIds = append(loginResMes.OnlineUserIds, id)
		}
		fmt.Println(user, "登录成功")
	}

	//序列化登录信息
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes marshal err=", err)
		return
	}

	//序列化要发送的消息
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)

	return
}

func (this *UserHandler) RegisterHandler(mes *message.Message) (err error) {
	//1 从message中取出data
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("register mes unmarshal err=", err)
		return
	}
	//返回的消息
	var resMes message.Message
	resMes.Type = message.LOGINRESMES

	var registerResMes message.RegisterResMes

	user, err := service.UserServiceInstance.UserRegisterService(registerMes.User.UserId, registerMes.User.UserPwd, registerMes.User.UserName)
	if err != nil {
		//不合法
		registerResMes.Code = 500
		registerResMes.Error = err.Error()
	} else {
		//合法
		registerResMes.Code = 200
		fmt.Println(user, "注册成功")
	}

	//序列化登录信息
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("registerResMes marshal err=", err)
		return
	}

	//序列化要发送的消息
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)

	return
}
