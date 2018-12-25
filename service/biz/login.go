package biz

import (
	"Project/service/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//调库，查询状态
func Login(user model.User) int {
	fmt.Println(user)
	//创建orm引擎
	engine, err := xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//查询
	has, err := engine.Get(&user)
	fmt.Println(err)
	if !has {
		user1 := model.User{Username: user.Username}
		if ok, _ := engine.Get(&user1); ok {
			return 2 //2代表密码错误
		} else {
			return 3 //3代表手机号未注册
		}
	} else {
		return 1 //1代表有匹配项，登陆成功
	}
}
