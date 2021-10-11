package logics

import (
	"buble/dao/mysql"
	"buble/models"
	"buble/utils"
	"errors"
)

func Signup(p *models.ParamSignUp)error{
	//1.判断用户存不存在
	exit,err:=mysql.QueryuserbyUsername(p.Username)
	if err!=nil{//数据库出错
		return err
	}
	if exit{//用户已存在
		return errors.New("用户已存在")
	}
	//2.生成uid
	userid,_:=utils.GetID()
	var u=models.User{
		userid,
		p.Username,
		utils.EncryptPassword(p.Password),
	}
	//3.插入数据
	_,err=mysql.InsertUser(&u)
	if err!=nil{
		return err
	}
	return err
}
func Login(p *models.ParamLogin)(token string,err error){
	//判断用户名和密码
	pass,err:=mysql.QueryuserbyPassword(p)
	if err!=nil{
		return "",err
	}
	if pass!=utils.EncryptPassword(p.Password){//用户名和密码不正确
		return "",errors.New("用户名或密码不正确")
	}
	//登陆成功获取token
	id,err:=mysql.QueryidbyPassword(p)
	if err!=nil{
		return "",err
	}
	return utils.GenToken(id,p.Username),nil
}