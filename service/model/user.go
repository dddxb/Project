package model

type User struct {
	Password string `json:"password" xorm:"'PASSWORD'"`
	Username string `json:"username" xorm:"'USERNAME'"`
	Userid   string `json:"user_id"  xorm:"'USERID'"`
	Access   string `json:"access"   xorm:"'ACCESS'"`
	Token    string `json:"token"    xorm:"'TOKEN'"`
	Avator   string `json:"avator"   xorm:"'AVATOR'"`
}


type Errlogger struct {
	Error string
}

type LoginTest struct {
	Username int
}
