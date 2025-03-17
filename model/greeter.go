package model

import "gorm.io/gorm"

type Greeter struct {
	gorm.Model
	Hello string `gorm:"type:varchar(128);not null"`
}

func (g *Greeter) TableName() string {
	return "greeter"
}

func (g *Greeter) SayHello() string {
	return g.Hello
}

func (g *Greeter) SetHello(hello string) {
	g.Hello = hello
}
