package handler

import (
	"fmt"
)

var (
	UserMgr *UserManager
)

type UserManager struct {
	OnlineUserTable map[int]*UserHandler
}

func init() {
	UserMgr = &UserManager{
		OnlineUserTable: make(map[int]*UserHandler, 1024),
	}
}

func (this *UserManager) AddOnlineUser(up *UserHandler) {
	this.OnlineUserTable[up.UserId] = up
}

func (this *UserManager) DelOnlineUser(userId int) {
	delete(this.OnlineUserTable, userId)
}

func (this *UserManager) GetAllOnlineUsers() map[int]*UserHandler {
	return this.OnlineUserTable
}

func (this *UserManager) GetOnlineUserById(userId int) (up *UserHandler, err error) {
	up, ok := this.OnlineUserTable[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
