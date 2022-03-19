package handler

import (
	"easychat/common/message"
	"encoding/json"
	"fmt"
)

func showGroupSms(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("smsMes Unmarshal err=", err)
		return
	}

	info := fmt.Sprintf("用户id:\t%d对大家说:\t%s",
		smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)

}
