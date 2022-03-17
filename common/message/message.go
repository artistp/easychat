package message

import "easychat/common/entity"

//确定一些消息类型
const (
	LOGINMES    = "LoginMes"
	LOGINRESMES = "LoginResMes"

	LOGOUTMES = "LogoutMes"

	REGISTERMES    = "Register"
	REGISTERRESMES = "RegisterResMes"

	NOTIFYUSERSTATUSMES = "NotifyUserStatusMes"
)

//用户状态
const (
	USERONLINE = iota
	USEROFFLINE
	USERBUSYSTATUS
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	User entity.User
}

type LoginResMes struct {
	Code          int //返回状态码 500表示用户未注册，200表示登录成功
	OnlineUserIds []int
	Error         string
}

type RegisterMes struct {
	User entity.User
}

type RegisterResMes struct {
	Code  int //返回状态码 500表示用户未注册，200表示登录成功
	Error string
}

//配合服务器端推送用户上线通知
type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`
	UserStatus int `json:"userStatus"` //用户的状态
}
