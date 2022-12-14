package domain

import "github.com/jinzhu/gorm"

type Paciente struct {
	gorm.Model
	Nome      string     `json:"nome"`
	Sobrenome string     `json:"sobrenome"`
	RG        string     `json:"rg"`
	Consultas []Consulta `gorm:"foreignKey:PacienteID " json:"-"`
}
