package handler

import (
	"easychat/common/entity"
	"easychat/common/message"
	"encoding/json"
	"fmt"
)

var UserTableInstance map[int]*entity.User = make(map[int]*entity.User, 10)

func updateUserTable(mes *message.Message) {
	var notifyMes message.NotifyUserStatusMes
	err := json.Unmarshal([]byte(mes.Data), &notifyMes)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
		return
	}
	user, ok := UserTableInstance[notifyMes.UserId]
	if !ok { //UserTable中没有
		user = &entity.User{
			UserId:     notifyMes.UserId,
			UserStatus: notifyMes.UserStatus,
		}
	}
	UserTableInstance[notifyMes.UserId] = user
	showUserTable()
}

func showUserTable() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range UserTableInstance {
		fmt.Println("用户ID: \t", id)
	}
}
