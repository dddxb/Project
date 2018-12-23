package model



type UserLogin struct {
	Password string    //`json:"PASSWORD"`
	Username string    // `json:"USERNAME"`
}

type UserInfo struct {
	Name  string    `json:"username"`
	User_id   int   `json:"user_id"`
	Access   []string   `json:"access"`
	Token   string    `json:"token"`
	Avator   string  `json:"avator"`
}

type Errlogger struct {
	Error string
}