package dentista

import (
	"fmt"
	"v0/internal"
	"v0/internal/domain"

	"github.com/jinzhu/gorm"
)

/* type Repository interface {
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
} */

type DentistaRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) internal.RepositoryInterface[domain.Dentista] {
	return &DentistaRepository{db: db}
}

func (s *DentistaRepository) Delete(id int) error {
	result := s.db.Delete(&domain.Dentista{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *DentistaRepository) Create(dentista domain.Dentista) (domain.Dentista, error) {
	result := r.db.Save(&dentista)
	if result.Error != nil {
		return domain.Dentista{}, result.Error
	}
	return dentista, nil
}

func (r *DentistaRepository) GetByID(id int) (domain.Dentista, error) {
	var dentista domain.Dentista
	result := r.db.First(&dentista, id)
	if result.Error != nil {
		return domain.Dentista{}, result.Error
	}
	return dentista, nil
}

func (r *DentistaRepository) Update(dentista domain.Dentista) (domain.Dentista, error) {
	fmt.Println("update", dentista)
	result := r.db.Model(&dentista).Updates(dentista)
	if result.Error != nil {
		return domain.Dentista{}, result.Error
	}
	return dentista, nil
}
