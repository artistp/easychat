package tools

import (
	"easychat/server/handler"
	"fmt"
)

var (
	UserMgr *UserManager
)

type UserManager struct {
	OnlineUserTable map[int]*handler.UserHandler
}

func init() {
	UserMgr = &UserManager{
		OnlineUserTable: make(map[int]*handler.UserHandler, 1024),
	}
}

func (this *UserManager) AddOnlineUser(up *handler.UserHandler) {
	this.OnlineUserTable[up.UserId] = up
}

func (this *UserManager) DelOnlineUser(userId int) {
	delete(this.OnlineUserTable, userId)
}

func (this *UserManager) GetAllOnlineUsers() map[int]*handler.UserHandler {
	return this.OnlineUserTable
}

func (this *UserManager) GetOnlineUserById(userId int) (up *handler.UserHandler, err error) {
	up, ok := this.OnlineUserTable[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
