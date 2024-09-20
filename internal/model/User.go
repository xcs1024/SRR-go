package model

type User struct {
	Id      int    `gorm:"primary_key;auto_increment;comment:自增主键"`
	Name    string `gorm:"type:varchar(32);not null;comment:用户名"`
	Age     int    `gorm:"type:int;size:3;comment:年龄"`
	Sex     int    `gorm:"type:int;size:2;comment:性别"`
	Address string `gorm:"type:varchar(32);comment:地址"`
}
