package service

import (
	"go-learn/code/go_di_example/config"
	"go-learn/code/go_di_example/infra"
)

type PersonService struct {
	config           *config.Config
	personRepository *infra.PersonRepository
}

func NewPersonService(config *config.Config, personRepository *infra.PersonRepository) *PersonService {
	return &PersonService{config: config, personRepository: personRepository}
}

func (s *PersonService) FindAll() []*infra.Person {
	if s.config.Enabled {
		return s.personRepository.FindAll()
	}
	return []*infra.Person{}
}
