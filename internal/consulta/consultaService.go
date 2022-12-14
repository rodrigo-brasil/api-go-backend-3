package consulta

import (
	"v0/internal"
	"v0/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Consulta, error)
	Create(c domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	Update(id int, c domain.Consulta) (domain.Consulta, error)
}

type service struct {
	r internal.RepositoryInterface[domain.Consulta]
}

func NewService(repository internal.RepositoryInterface[domain.Consulta]) Service {
	return &service{repository}
}

func (s *service) GetByID(id int) (domain.Consulta, error) {
	c, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	return c, nil
}

func (s *service) Create(c domain.Consulta) (domain.Consulta, error) {
	c, err := s.r.Create(c)
	if err != nil {
		return domain.Consulta{}, err
	}
	return c, nil
}

func (s *service) Update(id int, u domain.Consulta) (domain.Consulta, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	if u.Descricao != "" {
		d.Descricao = u.Descricao
	}
	if u.Data != "" {
		d.Data = u.Data
	}

	newDestista, err := s.r.Update(d)
	if err != nil {
		return domain.Consulta{}, err
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
