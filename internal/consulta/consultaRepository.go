package consulta

import (
	"v0/internal"
	"v0/internal/domain"

	"github.com/jinzhu/gorm"
)

type ConsultaRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) internal.RepositoryInterface[domain.Consulta] {
	return &ConsultaRepository{db: db}
}

func (s *ConsultaRepository) Delete(id int) error {
	result := s.db.Delete(&domain.Consulta{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ConsultaRepository) Create(consulta domain.Consulta) (domain.Consulta, error) {
	err := r.db.Create(&consulta).Error
	if err != nil {
		return domain.Consulta{}, err
	}
	result := r.db.Save(&consulta)
	if result.Error != nil {
		return domain.Consulta{}, result.Error
	}
	return consulta, nil
}

func (r *ConsultaRepository) GetByID(id int) (domain.Consulta, error) {
	var consulta domain.Consulta
	result := r.db.Preload("Paciente").Preload("Dentista").First(&consulta, id)
	if result.Error != nil {
		return domain.Consulta{}, result.Error
	}
	return consulta, nil
}

func (r *ConsultaRepository) Update(consulta domain.Consulta) (domain.Consulta, error) {
	result := r.db.Model(&consulta).Updates(&consulta)
	if result.Error != nil {
		return domain.Consulta{}, result.Error
	}
	return consulta, nil
}
