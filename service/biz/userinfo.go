package biz

import (
	"Project/service/model"
	"fmt"
	//_调库
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//UserInfo 调库，查询状态
func UserInfo(token string) model.User {
	fmt.Println(token)
	//创建orm引擎
	engine, err := xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//查询，省略查询password
	users := make([]model.User, 0)
	engine.Omit("PASSWORD").Where("TOKEN = ?", token).Find(&users)
	fmt.Println(users)
	//[{ 18753113305 1 super_admin super_admin https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png}]
	//需要去掉外边的[ ]符号,然后返回
	if len(users) > 0 {
		return users[0]
	}
	return users[1]

}
