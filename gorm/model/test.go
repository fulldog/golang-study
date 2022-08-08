package model

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name string `gorm:"column:name;not null;type varchar(200);comment:'用户名'"`
	Desc string `gorm:"column:desc;not null;type varchar(200);comment:'描述'"`
}

func (receiver Test) TableName() string {
	return "test"
}

func (receiver *Test) Create() bool {
	Db.Model(receiver).Debug().AutoMigrate(receiver).Error()
	return true
}
