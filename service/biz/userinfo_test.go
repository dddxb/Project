package biz

import (
	"Project/service/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserInfo(t *testing.T) {
	expectedResponseSuper := model.User{Password: "", Username: "18753113305", Userid: "1", Access: "super_admin", Token: "super_admin", Avator: "https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png"}
	assert.Equal(t, expectedResponseSuper, UserInfo("super_admin"))
	expectedResponseAdmin := model.User{Password: "", Username: "18813053143", Userid: "2", Access: "admin", Token: "admin", Avator: "https://avatars0.githubusercontent.com/u/20942571?s=460&v=4"}
	assert.Equal(t, expectedResponseAdmin, UserInfo("admin"))
}
