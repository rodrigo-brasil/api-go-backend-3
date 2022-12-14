package dentista

import (
	"v0/internal"
	"v0/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	Update(id int, p domain.Dentista) (domain.Dentista, error)
}

type service struct {
	r internal.RepositoryInterface[domain.Dentista]
}

func NewService(repository internal.RepositoryInterface[domain.Dentista]) Service {
	return &service{repository}
}

func (s *service) GetByID(id int) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

func (s *service) Create(d domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if u.Nome != "" {
		d.Nome = u.Nome
	}
	if u.Sobrenome != "" {
		d.Sobrenome = u.Sobrenome
	}
	if u.Matricula > 0 {
		d.Matricula = u.Matricula
	}

	newDestista, err := s.r.Update(d)
	if err != nil {
		return domain.Dentista{}, err
	}

	return newDestista, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
