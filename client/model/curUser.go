package model

import (
	"easychat/common/entity"
	"net"
)

var CurUserInstance *CurUser = &CurUser{}

type CurUser struct {
	Conn net.Conn //维护自身的链接
	User entity.User
}
