package controllers

import "github.com/gofiber/fiber/v2"

func Upvote(c *fiber.Ctx) error {
	return c.SendString("Success")
}
