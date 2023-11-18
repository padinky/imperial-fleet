package services

import (
	"errors"

	"github.com/padinky/imperial-fleet/model"
	"github.com/padinky/imperial-fleet/repository"
)

type SpaceshipService struct {
	spaceshipRepository repository.ISpaceshipRepository
}

func NewSpaceshipService(repo repository.ISpaceshipRepository) *SpaceshipService {
	return &SpaceshipService{
		spaceshipRepository: repo,
	}
}

func (ps *SpaceshipService) CreateSpaceship(data model.Spaceship) error {
	_, err := ps.spaceshipRepository.Create(data)
	return err
}

func (ps *SpaceshipService) GetSpaceshipById(id uint) (model.Spaceship, error) {
	return ps.spaceshipRepository.FindByID(id)
}

func (ps *SpaceshipService) UpdateSpaceship(id uint, data model.Spaceship) error {
	return ps.spaceshipRepository.Update(id, data)
}

func (ps *SpaceshipService) GetAllSpaceship() (*[]model.Spaceship, error) {
	products, err := ps.spaceshipRepository.GetAll()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("no data found")
	}
	return &products, nil
}

func (ps *SpaceshipService) DeleteSpaceship(id uint) error {
	return ps.spaceshipRepository.Delete(id)
}
