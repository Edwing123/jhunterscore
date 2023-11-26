package main

import (
	"github.com/gofiber/fiber/v2"
)

func (core *Core) HandldeOffers(c *fiber.Ctx) error {
	isCompact := c.QueryBool("c", false)
	_ = isCompact
	return nil
}

func (core *Core) HandldeOfferById(c *fiber.Ctx) error {
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
