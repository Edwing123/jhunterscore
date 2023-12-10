package main

import (
	"errors"
	"fmt"
	"strconv"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"github.com/gofiber/fiber/v2"
)

func (core *Core) HandldeOffers(c *fiber.Ctx) error {
	isCompact := c.QueryBool("c", false)

	offers, err := core.Database.OffersRepository.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err,
		})
	}

	if isCompact {
		return c.JSON(fiber.Map{
			"ok":   true,
			"data": getCompactOffers(offers),
		})
	}

	return c.JSON(fiber.Map{
		"ok":   true,
		"data": offers,
	})
}

func (core *Core) HandldeOfferById(c *fiber.Ctx) error {
	// Get offer id from params.
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": err,
		})
	}

	// Find offer by id.
	offer, err := core.Database.OffersRepository.GetById(id)
	if err != nil {
		if errors.Is(err, database.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("Oferta con id %d no existe", id),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"ok":   true,
		"data": offer,
	})
}

func (core *Core) HandldeResources(c *fiber.Ctx) error {
	isCompact := c.QueryBool("c", false)

	if isCompact {
		return c.JSON(fiber.Map{
			"ok":   true,
			"data": getCompactResources(mockResources),
		})
	}

	return c.JSON(fiber.Map{
		"ok":   true,
		"data": mockResources,
	})
}

func (core *Core) HandldeResourceById(c *fiber.Ctx) error {
	// Get resource id from params.
	id := c.Params("id")

	// Find resource by id.
	var resource models.Resource

	for _, r := range mockResources {
		if strconv.Itoa(r.Id) == id {
			resource = r
			break
		}
	}

	if resource.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":    false,
			"error": "Recurso no encontrado",
		})
	}

	return c.JSON(fiber.Map{
		"ok":   true,
		"data": resource,
	})
}

func (core *Core) HandldeCompanies(c *fiber.Ctx) error {
	companies, err := core.Database.CompaniesRepository.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
			"ok":    false,
		})
	}

	return c.JSON(fiber.Map{
		"ok":   true,
		"data": companies,
	})
}
