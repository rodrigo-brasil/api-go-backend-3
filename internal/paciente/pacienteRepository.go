package paciente

import (
	"v0/internal"
	"v0/internal/domain"

	"github.com/jinzhu/gorm"
)

type PacienteRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) internal.RepositoryInterface[domain.Paciente] {
	return &PacienteRepository{db: db}
}

func (s *PacienteRepository) Delete(id int) error {
	result := s.db.Delete(&domain.Paciente{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *PacienteRepository) Create(paciente domain.Paciente) (domain.Paciente, error) {
	result := r.db.Save(&paciente)
	if result.Error != nil {
		return domain.Paciente{}, result.Error
	}
	return paciente, nil
}

func (r *PacienteRepository) GetByID(id int) (domain.Paciente, error) {
	var paciente domain.Paciente
	result := r.db.First(&paciente, id)
	if result.Error != nil {
		return domain.Paciente{}, result.Error
	}
	return paciente, nil
}

func (r *PacienteRepository) Update(paciente domain.Paciente) (domain.Paciente, error) {
	result := r.db.Model(&paciente).Updates(paciente)
	if result.Error != nil {
		return domain.Paciente{}, result.Error
	}
	return paciente, nil
}
