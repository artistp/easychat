package entity

type User struct {

	//保证序列化成功，需要加入json的tag
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户的状态
}
