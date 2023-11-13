package model

type Tags struct {
	Id   int    `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"type:varchar(255)"`
}
