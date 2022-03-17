package dao

import (
	"easychat/common/entity"
	error2 "easychat/common/error"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	UserDaoInstance *UserDao
)

type UserDao struct {
	Pool *redis.Pool
}

//使用工厂模式创建一个userDao 的实例
func newUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

func InitUserDao(pool *redis.Pool) {
	UserDaoInstance = newUserDao(pool)
}

func (this *UserDao) GetUserDao(id int) (user *entity.User, err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	//通过id去Redis里面去查询
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//在users中没有用户
		if err == redis.ErrNil {
			err = error2.ERROR_USER_NOTEXIST
		}
		return
	}

	user = &entity.User{}

	//反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (this *UserDao) SaveUserDao(user *entity.User) (err error) {

	conn := this.Pool.Get()
	defer conn.Close()
	//存储到redis中
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//插入到Redis中
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户失败 err=", err)
	}
	return
}
