package main

import (
	"strconv"

	"edwingarcia.dev/github/jhunterscore/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func (core *Core) HandldeOffers(c *fiber.Ctx) error {
	isCompact := c.QueryBool("c", false)

	if isCompact {
		c.JSON(fiber.Map{
			"ok":   true,
			"data": getCompactOffers(),
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
	_ = isCompact
	return nil
}

func (core *Core) HandldeResourceById(c *fiber.Ctx) error {
	return nil
}
