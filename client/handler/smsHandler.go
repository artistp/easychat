package handler

import (
	"easychat/client/model"
	"easychat/common/message"
	"easychat/common/utils"
	"encoding/json"
	"fmt"
)

var SmsHandlerInstance *SmsHandler = &SmsHandler{}

type SmsHandler struct {
}

func (this *SmsHandler) SendGroupSms(content string) {
	var mes message.Message
	mes.Type = message.SMSMES

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User = model.CurUserInstance.User

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("smsMes Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: model.CurUserInstance.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send sms err=", err)
	}
	return
}
