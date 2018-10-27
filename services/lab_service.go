package services

import "github.com/jinzhu/gorm"

type ILab interface {
	Create()
}

type lab struct {
	Conn *gorm.DB
}

type Person struct {
	gorm.Model
	//ID uint `json:”id”`
	FirstName string `json:”firstname”`
	LastName string `json:”lastname”`
}

func NewLab(conn *gorm.DB) ILab {
	return &lab{conn}
}

func (l *lab) Create() {
	l.Conn.AutoMigrate(&Person{})
	p1 := Person{FirstName: "Tom", LastName: "Me"}
	l.Conn.Create(&p1)
}
