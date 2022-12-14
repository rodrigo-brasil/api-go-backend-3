package paciente

import (
	"v0/internal"
	"v0/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	Update(id int, p domain.Paciente) (domain.Paciente, error)
}

type service struct {
	r internal.RepositoryInterface[domain.Paciente]
}

func NewService(repository internal.RepositoryInterface[domain.Paciente]) Service {
	return &service{repository}
}

func (s *service) GetByID(id int) (domain.Paciente, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return d, nil
}

func (s *service) Create(p domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}

	if u.Nome != "" {
		p.Nome = u.Nome
	}
	if u.Sobrenome != "" {
		p.Sobrenome = u.Sobrenome
	}
	if u.RG != "" {
		p.RG = u.RG
	}

	newPaciente, err := s.r.Update(p)
	if err != nil {
		return domain.Paciente{}, err
	}

	return newPaciente, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
