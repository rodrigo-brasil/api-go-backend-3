package domain

import (
	"github.com/jinzhu/gorm"
)

type Consulta struct {
	gorm.Model
	Descricao  string   `json:"descricao"`
	Data       string   `json:"data"`
	Paciente   Paciente `json:"paciente"`
	Dentista   Dentista `json:"dentista"`
	PacienteID uint     `json:"-"`
	DentistaID uint     `json:"-"`
}
