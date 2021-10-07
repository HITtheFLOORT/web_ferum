package logics

import (
	"buble/dao/mysql"
	"buble/utils"
)

func Signup(){
	//1.判断用户存不存在
	mysql.QueryuserbyUsername()
	//2.生成uid
	utils.GetID()
	//3.
	mysql.InsertUser()
}
