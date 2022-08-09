package model

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name string `gorm:"column:name;not null;size:255;type varchar;comment:用户名"`
	Desc string `gorm:"column:desc;not null;size:255;type varchar;comment:描述"`
}

func (receiver Test) TableName() string {
	return "test"
}

func (receiver *Test) Create() error {
	var err error
	if !Db.Migrator().HasTable(receiver) {
		err = Db.Migrator().CreateTable(receiver)
	}
	return err
}
