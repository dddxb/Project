package biz

import (
	"Project/service/model"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func Test_Login(t *testing.T) {

	//正常登陆
	user1 := model.User{
		Password: "qwert",
		Username: "18753113305",
	}
	//密码错误
	user2 := model.User{
		Password: "qwe",
		Username: "18753113305",
	}
	//账号未注册
	user3 := model.User{
		Password: "qwe",
		Username: "18954526289",
	}

	assert.Equal(t, 1, Login(user1))
	assert.Equal(t, 2, Login(user2))
	assert.Equal(t, 3, Login(user3))
}
