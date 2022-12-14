package model

import "time"

type Permission struct {
	Id       int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`     // internal index, not used outside
	CardID   string `gorm:"size:512;not null;unique" json:"cardID"`   // Card ID
	Username string `gorm:"size:512;not null;unique" json:"username"` // unique Username, could be phone number
}

type Admin struct {
	Id       int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`     // internal index, also used in UserToken
	Username string `gorm:"size:512;not null;unique" json:"username"` // unique Username, could be phone number
	Password string `gorm:"size:128;not null" json:"password"`        // password
}

type UserToken struct {
	Id         int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`              // internal index, not used outside
	AdminId    int64     `gorm:"not null;unique" json:"adminId"`                    // from table Admin
	Token      string    `gorm:"size:32;unique;not null" json:"token" form:"token"` // token
	ExpiredAt  int64     `gorm:"not null" json:"expiredAt" form:"expiredAt"`        // expired at
	CreateTime time.Time `gorm:"not null" json:"createTime"`                        // create time
}
