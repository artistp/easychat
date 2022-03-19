package handler

import (
	"easychat/common/message"
	"easychat/common/utils"
	"encoding/json"
	"fmt"
)

type SmsHandler struct {
}

func (this *SmsHandler) SendGroupSms(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("smsMes Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes Unmarshal err=", err)
		return
	}

	for id, up := range UserMgrInstance.OnlineUserTable {
		//过滤掉自己
		if id == smsMes.User.UserId {
			continue
		}
		tf := &utils.Transfer{
			Conn: up.Conn,
		}
		err = tf.WritePkg(data)
		if err != nil {
			fmt.Println("转发消息失败err=", err)
		}
	}
}
