package repository

import (
	"errors"

	"github.com/padinky/imperial-fleet/database"
	"github.com/padinky/imperial-fleet/model"
	"gorm.io/gorm"
)

type ISpaceshipRepository interface {
	Create(model.Spaceship) (uint, error)
	FindByID(uint) (model.Spaceship, error)
	Update(uint, model.Spaceship) error
	GetAll() ([]model.Spaceship, error)
	Delete(uint) error
}

type SpaceshipRepository struct {
	db *gorm.DB
}

func NewSpaceshipRepository() *SpaceshipRepository {
	return &SpaceshipRepository{
		db: database.DB,
	}
}

func (r *SpaceshipRepository) Create(data model.Spaceship) (uint, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *SpaceshipRepository) FindByID(id uint) (model.Spaceship, error) {
	spaceship := model.Spaceship{}
	query := model.Spaceship{}
	query.ID = uint(id)

	// get armament first
	armament, err := r.getArmamentsBySpaceshipID(uint(id))
	if err != nil {
		return spaceship, err
	}

	err = r.db.Where("deleted_at IS NULL").First(&spaceship, &query).Error

	spaceship.Armaments = armament

	return spaceship, err
}

func (r *SpaceshipRepository) GetAll() ([]model.Spaceship, error) {
	spaceship := []model.Spaceship{}

	// get armament first
	armaments, err := r.getAllArmaments()
	if err != nil {
		return spaceship, err
	}

	err = r.db.Where("deleted_at IS NULL").Find(&spaceship).Error

	if err != nil {
		return spaceship, err
	}

	// match spaceship with armament
	for i, d := range spaceship {
		for _, arm := range armaments {
			if d.ID == arm.SpaceshipID {
				spaceship[i].Armaments = append(spaceship[i].Armaments, arm)
			}
		}
	}

	return spaceship, err
}

func (r *SpaceshipRepository) Update(id uint, data model.Spaceship) error {
	found := model.Spaceship{}
	query := model.Spaceship{}
	query.ID = uint(id)
	err := r.db.Where("deleted_at IS NULL").First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("data not found")
	}
	found.Name = data.Name
	found.Class = data.Class
	found.Crew = data.Crew
	found.Image = data.Image
	found.Status = data.Status
	found.ID = id
	// found.Armaments = data.Armaments

	err = r.db.Save(&found).Error
	if err != nil {
		return err
	}

	// failed to cascade update I don't know why
	r.db.Where("spaceship_id = ?", id).Delete(&model.Armament{})
	for _, d := range data.Armaments {
		// save new row
		newArmament := new(model.Armament)
		newArmament.Title = d.Title
		newArmament.Qty = d.Qty
		newArmament.SpaceshipID = id
		err := r.db.Create(&newArmament).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SpaceshipRepository) Delete(id uint) error {
	found := model.Spaceship{}
	query := model.Spaceship{}
	query.ID = uint(id)
	err := r.db.Where("deleted_at IS NULL").First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("data not found")
	}
	return r.db.Delete(&found).Error
}

func (r *SpaceshipRepository) getArmamentsBySpaceshipID(spaceshipID uint) ([]model.Armament, error) {
	armaments := []model.Armament{}
	query := model.Armament{}
	query.SpaceshipID = spaceshipID
	err := r.db.Where("deleted_at IS NULL").Find(&armaments, &query).Error
	return armaments, err
}

func (r *SpaceshipRepository) getAllArmaments() ([]model.Armament, error) {
	armaments := []model.Armament{}
	err := r.db.Where("deleted_at IS NULL").Find(&armaments).Error
	return armaments, err
}
