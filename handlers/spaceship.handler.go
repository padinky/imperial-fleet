package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/padinky/imperial-fleet/helper"
	"github.com/padinky/imperial-fleet/model"
	"github.com/padinky/imperial-fleet/repository"
	"github.com/padinky/imperial-fleet/services"
	"gorm.io/gorm"
)

// SpaceshipHandler struct for Spaceship Handler
type SpaceshipHandler struct {
	spaceshipService *services.SpaceshipService
}

// NewSpaceshipHandler returns SpaceshipHandler instance
func NewSpaceshipHandler(spaceshipRepository repository.ISpaceshipRepository) *SpaceshipHandler {
	ps := services.NewSpaceshipService(spaceshipRepository)
	return &SpaceshipHandler{
		spaceshipService: ps,
	}
}

func (s *SpaceshipHandler) Create(c *fiber.Ctx) error {
	json := new(model.Spaceship)
	if err := c.BodyParser(json); err != nil {
		return helper.ResponseBadRequest(c, err.Error())
	}

	//TODO: Validate param

	err := s.spaceshipService.CreateSpaceship(*json)
	if err != nil {
		return helper.ResponseError(c, err.Error())
	}

	return helper.ResponseSuccessOnly(c)
}

func (s *SpaceshipHandler) GetByID(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return helper.ResponseBadRequest(c, "invalid id format, should be a number")
	}

	// -- hit product service to get product by id
	spaceship, err := s.spaceshipService.GetSpaceshipById(uint(id))

	if err == gorm.ErrRecordNotFound {
		return helper.ResponseNotFound(c, "data not found")
	}
	return helper.ResponseDataOnly(c, spaceship)
}

func (s *SpaceshipHandler) GetAll(c *fiber.Ctx) error {
	spaceships, err := s.spaceshipService.GetAllSpaceship()
	if err == gorm.ErrRecordNotFound {
		return helper.ResponseNotFound(c, "data not found")
	}
	return helper.ResponseDataOnly(c, spaceships)
}

func (s *SpaceshipHandler) Update(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return helper.ResponseBadRequest(c, "invalid id format, should be a number")
	}

	json := new(model.Spaceship)
	if err := c.BodyParser(json); err != nil {
		return helper.ResponseBadRequest(c, err.Error())
	}

	//TODO: Validate param

	err = s.spaceshipService.UpdateSpaceship(uint(id), *json)
	if err != nil {
		return helper.ResponseError(c, err.Error())
	}

	return helper.ResponseSuccessOnly(c)
}

func (s *SpaceshipHandler) Delete(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return helper.ResponseBadRequest(c, "invalid id format, should be a number")
	}

	// -- hit product service to get delete the product
	err = s.spaceshipService.DeleteSpaceship(uint(id))
	if err != nil {
		return helper.ResponseError(c, err.Error())
	}

	return helper.ResponseSuccessOnly(c)
}
