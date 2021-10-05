package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)
var DB *sqlx.DB
func Init() (err error){
	dbhost:=viper.GetString("app.host")
	dbport:=viper.GetInt("mysql.port")
	dbname:=viper.GetString("mysql.dbname")
	var dsn string =fmt.Sprintf("root:GLMlove19971212@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbhost,dbport,dbname)
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
		panic(nil)
	}
	DB.SetMaxOpenConns(200)
	DB.SetMaxIdleConns(100)
	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}else{
		fmt.Printf("%s is connected", "mysql")
	}
	return
}
func Close(){
	DB.Close()
}