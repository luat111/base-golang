package user_model

import "practice/auth/core/base"

type User struct {
	base.BaseModel
	FullName *string `gorm:"varchar(255);not null" json:"full_name,omitempty"`
	Age      *int    `gorm:"integer;default: 0" json:"age,omitempty"`
	Username *string `gorm:"varchar(255);uniqueIndex;not null" json:"username,omitempty"`
	Password *string `gorm:"varchar(255);not null" json:"password,omitempty"`
}
