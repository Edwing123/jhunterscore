package main

import (
	"strconv"

	"edwingarcia.dev/github/jhunterscore/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func (core *Core) HandldeOffers(c *fiber.Ctx) error {
	compact := c.QueryBool("compact")

	var offersResult []any

	for _, o := range offers {
		if compact {
			offersResult = append(offersResult, models.CompactOffer{
				Id:          o.Id,
				Title:       o.Title,
				Contract:    o.Contract,
				Salary:      o.Salary,
				PublishedAt: o.PublishedAt,
			})
		} else {
			offersResult = append(offersResult, o)
		}
	}

	return c.JSON(fiber.Map{
		"offers": offersResult,
	})
}

func (core *Core) HandldeOfferById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	if strId == "" || err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is required.")
	}

	var offerIdx int
	var found bool

	for idx, o := range offers {
		if o.Id == id {
			offerIdx = idx
			found = true
			break
		}
	}

	if !found {
		return fiber.NewError(fiber.StatusNotFound, "offer not found.")
	}

	return c.JSON(fiber.Map{
		"offer": offers[offerIdx],
	})
}

func (core *Core) HandldeResources(c *fiber.Ctx) error {
	compact := c.QueryBool("compact")

	var resourcesResult []any

	for _, r := range resources {
		if compact {
			resourcesResult = append(resourcesResult, models.CompactResource{
				Id:          r.Id,
				Title:       r.Title,
				Author:      r.Author,
				PublishedAt: r.PublishedAt,
			})
		} else {
			resourcesResult = append(resourcesResult, r)
		}
	}

	return c.JSON(fiber.Map{
		"resources": resourcesResult,
	})
}

func (core *Core) HandldeResourceById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.Atoi(strId)

	if strId == "" || err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is required.")
	}

	var resourceIdx int
	var found bool

	for idx, r := range resources {
		if r.Id == id {
			resourceIdx = idx
			found = true
			break
		}
	}

	if !found {
		return fiber.NewError(fiber.StatusNotFound, "resource not found.")
	}

	return c.JSON(fiber.Map{
		"offer": resources[resourceIdx],
	})
}
