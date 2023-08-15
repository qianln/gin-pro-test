package models

import (
	"fmt"
	"gin-pro/app/core/system"
)

func NewUsersModel() *UsersModel {
	return &UsersModel{
		BaseModel: BaseModel{DB: UseDbConn()},
	}
}

type UsersModel struct {
	BaseModel
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
	Address  string `gorm:"column:address" json:"address"`
	Status   int    `gorm:"column:status" json:"status"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Gender   string `gorm:"column:gender" json:"gender"`
}

func (u *UsersModel) TableName() string {
	return fmt.Sprintf("%susers", system.Config.GetString("Mysql.Prefix"))
}
