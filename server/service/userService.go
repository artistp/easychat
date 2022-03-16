package service

import (
	"easychat/common/entity"
	error2 "easychat/common/error"
	"easychat/server/dao"
)

var (
	UserServiceInstance *UserService
)

type UserService struct {
	Ud *dao.UserDao
}

func (this *UserService) setUd(userDao *dao.UserDao) {
	this.Ud = userDao
}

func InitUserService(userDao *dao.UserDao) {
	UserServiceInstance = &UserService{}
	UserServiceInstance.setUd(userDao)
}

func (this *UserService) UserLoginService(userId int, userPwd string) (user *entity.User, err error) {
	user, err = this.Ud.GetUserDao(userId)
	if err != nil {
		return
	}

	//获取到了用户，密码不一定正确，需要验证密码
	if user.UserPwd != userPwd {
		err = error2.ERROR_USER_PWD
		return
	}
	return
}

func (this *UserService) UserRegisterService(userId int, userPwd, userName string) (user *entity.User, err error) {
	user, err = this.Ud.GetUserDao(userId)
	if err == nil {
		err = error2.ERROR_USER_EXIST
		return
	}
	user = &entity.User{
		userId,
		userPwd,
		userName,
	}
	err = this.Ud.SaveUserDao(user)
	//获取到了用户，密码不一定正确，需要验证密码

	return
}
