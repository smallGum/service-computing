package entities

import (
	"time"
)

//UserInfo .
type UserInfo struct {
	UID        int64      `xorm:"autoincr notnull pk 'uid'"`
	UserName   string     `xorm:"username"`
	DepartName string     `xorm:"departname"`
	CreateAt   *time.Time `xorm:"created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not be null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}
