package controller

import (
	"Project/service/biz"
	"Project/service/model"
	"fmt"
	"regexp"
	"strings"

	"github.com/kataras/iris"
)

//Login 确认登陆信息
func Login(ctx iris.Context) {
	var usr model.User

	err := ctx.ReadJSON(&usr)

	fmt.Println(usr.Username, usr.Password)

	//http.body正确解析
	if err == nil {
		//正则匹配手机号
		reg := `^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$`
		rgx := regexp.MustCompile(reg)

		fmt.Println(usr.Username)
		//手机号格式正确
		if rgx.MatchString(usr.Username) {
			//密码格式正确
			if strings.Count(usr.Password, "") > 1 && strings.Count(usr.Password, "") <= 8 {
				//查找数据库
				response := biz.Login(usr)
				switch response {
				case 1: //登陆成功
					switch usr.Username {
					case "18753113305":
						ctx.JSON(iris.Map{"token": "super_admin"})
					case "18813053143":
						ctx.JSON(iris.Map{"token": "admin"})
					}
				case 2: //密码错误，账号正确
					ctx.StatusCode(iris.StatusBadRequest)
				case 3: //账号未注册
					ctx.StatusCode(iris.StatusBadRequest)
				}
			} else { //密码格式错误
				ctx.StatusCode(iris.StatusBadRequest)
			}
		} else { //用户名格式错误
			ctx.StatusCode(iris.StatusBadRequest)
		}
	} else {
		ctx.StatusCode(iris.StatusBadRequest)
	}
}
