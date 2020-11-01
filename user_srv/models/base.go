package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        string     `gorm:"type:varchar(36) comment '表自增ID'; primary_key;" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at; type:datetime not null default current_timestamp comment '创建时间'" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at; type:datetime not null default current_timestamp on update current_timestamp comment '更新时间'" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at; type:datetime comment '删除时间';index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	return nil
}
