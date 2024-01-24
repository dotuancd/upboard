package main

import (
	"log"
	"os"

	upboard "github.com/dotuancd/upboard/internal/upboard/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/slack-go/slack"
)

func main() {
	app := fiber.New()

	slackToken, ok := os.LookupEnv("SLACK_TOKEN")

	if !ok {
		panic("The SLACK_TOKEN is empty")
	}

	slackClient := slack.New(slackToken)
	// http := http.DefaultClient
	// token, _, error := slack.GetOAuthToken(http, "", "", "", "")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/slack/auth", upboard.SlackAuth(slackClient))

	app.Post("/slack/events", upboard.SlackEvent)

	app.Post("/upvote", upboard.Upvote)

	log.Fatal(app.Listen(":3000"))
}
