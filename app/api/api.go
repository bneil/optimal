package api

import (
	"github.com/bneil/optimal/app/db"
	"github.com/bneil/optimal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func SetupRoutes(config fiber.Config) *fiber.App {
	app := *fiber.New(config)
	app.Get("/", index)
	app.Post("/feed", createFeed)
	app.Get("/feed/:id", readFeed)
	app.Get("/feeds", readFeeds)
	app.Put("/feed/:id", updateFeed)
	return &app
}

func index(c *fiber.Ctx) error {
	feeds, err := db.GetFeeds()
	if err != nil {
		slog.Warn("unable to find feed", err)
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "no feeds found",
		})
	}

	return c.Render("index", fiber.Map{
		"Title":            "Hi Folks",
		"Feeds":            feeds,
		"PopularFeeds":     feeds,
		"RecentFeeds":      feeds,
		"CategorizedFeeds": feeds,
	})
}

func createFeed(c *fiber.Ctx) error {
	feed := new(model.Feed)
	err := c.BodyParser(&feed)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	feed.ID = uuid.New().String()

	err = db.CreateFeed(feed)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.Render("view_feed", feed)
}

func readFeed(c *fiber.Ctx) error {
	id := c.Params("id")
	feed, err := db.GetFeedById(id)
	if err != nil {
		slog.Warn("unable to find feed", err)
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "no feed found",
		})
	}

	return c.Render("view_feed", feed)
}

func readFeeds(c *fiber.Ctx) error {
	feeds, err := db.GetFeeds()
	if err != nil {
		slog.Warn("unable to find feed", err)
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "no feeds found",
		})
	}

	return c.Render("feeds", feeds)
}

func updateFeed(c *fiber.Ctx) error {
	updateFeed := new(model.Feed)

	err := c.BodyParser(updateFeed)
	if err != nil {
		err := c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		if err != nil {
			return err
		}
		return err
	}
	id := c.Params("id")
	feed, err := db.GetFeedById(id)
	if err != nil {
		err := c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		if err != nil {
			return err
		}
		return err
	}

	if feed != nil {
		feed.Title = updateFeed.Title
		feed.Text = updateFeed.Text
		feed.Type = updateFeed.Type
		feed.HtmlUrl = updateFeed.HtmlUrl
		feed.XMLUrl = updateFeed.XMLUrl
	}

	return c.Render("view_feed", feed)
}
