package mysql

import (
	"buble/models"
	"database/sql"
	"errors"
)
// QueryuserbyUsername 检查用户名是否存在
func QueryuserbyUsername(username string)(bool,error)  {
	sqlStr:=`select count(user_id) from user where username = ?`
	var count int
	if err:=DB.Get(&count,sqlStr,username);err!=nil{
		return false,err
	}
	return count>0,nil
}
// InsertUser 插入一条记录
func InsertUser(user *models.User)(sql.Result,error){
	sqlStr:=`insert into user(user_id,username,password) values(?,?,?)`
	re,err:=DB.Exec(sqlStr,user.UserID,user.Username,user.Password)
	return re,err
}
func QueryuserbyPassword(login *models.ParamLogin)(string,error){
	sqlStr:=`select password from user where username = ?`
	var name string
	if err:=DB.Get(&name,sqlStr,login.Username);err!=nil{
		return name,errors.New("用户名或密码不真确")
	}
	return name,nil
}
