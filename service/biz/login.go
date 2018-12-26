package biz

import (
	"Project/service/model"
	"fmt"
	//_连库
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//Login 调库，查询状态
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
			//2代表密码错误
			return 2
		}
		//3代表手机号未注册
		return 3
	}
	//1代表登陆成功
	return 1
}
