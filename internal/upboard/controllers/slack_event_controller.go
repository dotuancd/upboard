package controllers

import (
	"github.com/dotuancd/upboard/api"
	"github.com/gofiber/fiber/v2"
	"github.com/slack-go/slack"
)

func SlackEvent(c *fiber.Ctx) error {
	event := api.UrlVerification{}

	err := c.BodyParser(&event)
	if err != nil {
		return c.SendString("Invalid request")
	}

	return c.SendString(event.Challenge)
}

func SlackAuth(sc *slack.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sc.PublishView("", slack.HomeTabViewRequest{}, "")
		return c.JSON(fiber.Map{"ok": true})
	}
}
