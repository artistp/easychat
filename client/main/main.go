package main

import (
	"easychat/client/handler"
	"fmt"
)

var userID int
var userPSW string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续循环显示
	var loop=true

	for loop {
		fmt.Println("---------多人聊天------------")
		fmt.Println("\t\t\t1 登录")
		fmt.Println("\t\t\t2 注册")
		fmt.Println("\t\t\t3 退出")
		fmt.Println("\t\t\t请选择(1-3)：")
		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("欢迎登录")
			fmt.Print("请输入ID:")
			fmt.Scanf("%d\n",&userID)
			fmt.Print("请输入密码:")
			fmt.Scanf("%s\n",&userPSW)
			//创建一个userHandler的对象
			uh:=&handler.UserHandler{}
			uh.Login(userID,userPSW)
		case 2:
			fmt.Println("注册")

		case 3:
			fmt.Println("退出")

		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
