package message

//确定一些消息类型
const (
	LOGINMES = "LoginMes"
	LOGOUTMES="LogoutMes"
	REGISTERMES="Register"
	LOGINRESMES = "LoginResMes"

)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserID int `json:"userId"`
	UserPSW string `json:"userPsw"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int //返回状态码 500表示用户未注册，200表示登录成功
	Error string
}

