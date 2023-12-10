package main

import (
	"strconv"

	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"github.com/gofiber/fiber/v2"
)

func (core *Core) HandldeOffers(c *fiber.Ctx) error {
	isCompact := c.QueryBool("c", false)

	if isCompact {
		c.JSON(fiber.Map{
			"ok":   true,
			"data": getCompactOffers(mockOffers),
		})
		return nil
	}

	c.JSON(fiber.Map{
		"ok":   true,
		"data": mockOffers,
	})

	return nil
}

func (core *Core) HandldeOfferById(c *fiber.Ctx) error {
	// Get offer id from params.
	id := c.Params("id")

	// Find offer by id.
	var offer models.Offer

	for _, o := range mockOffers {
		if strconv.Itoa(o.Id) == id {
			offer = o
			break
		}
	}

	if offer.Id == 0 {
		c.JSON(fiber.Map{
			"ok":    false,
			"error": "Oferta no encontrada",
		})
		return nil
	}

	c.JSON(fiber.Map{
		"ok":   true,
		"data": offer,
	})

	return nil
}

func (core *Core) HandldeResources(c *fiber.Ctx) error {
	isCompact := c.QueryBool("c", false)

	if isCompact {
		c.JSON(fiber.Map{
			"ok":   true,
			"data": getCompactResources(mockResources),
		})
		return nil
	}

	c.JSON(fiber.Map{
		"ok":   true,
		"data": mockResources,
	})

	return nil
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
		c.JSON(fiber.Map{
			"ok":    false,
			"error": "Recurso no encontrado",
		})
		return nil
	}

	c.JSON(fiber.Map{
		"ok":   true,
		"data": resource,
	})

	return nil
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
