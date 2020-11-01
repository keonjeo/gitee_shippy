package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	Base
	Name     string    `gorm:"column:name;size:255;index:idx_name_add_id" json:"name"`
	Company  string    `gorm:"column:company" json:"company"`
	Email    string    `gorm:"column:email;type:varchar(100);unique_index" json:"email"`
	Password string    `gorm:"column:password;type:varchar(30);" json:"password"`
}


//设置表名，默认是结构体的名的复数形式
func (User) TableName() string {
	return "t_user"
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("Id", uuid.String())
}