package domain

import "github.com/jinzhu/gorm"

type Dentista struct {
	gorm.Model
	Matricula uint       `json:"matricula" binding:"required" gorm:"unique;"`
	Nome      string     `json:"nome" binding:"required"`
	Sobrenome string     `json:"sobrenome" binding:"required"`
	Consultas []Consulta `gorm:"foreignKey:DentistaID " json:"-"`
}
