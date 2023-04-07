package main

import (
	handler "github.com/bneil/gossr_tests/app"
	"github.com/bneil/gossr_tests/app/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	//quickly setup that db!
	go db.SetupSqlite()

	engine := html.New("./app/views/templates", ".gohtml")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		Prefork:     true,
	})
	app.Static("/static", "./app/views/static")

	// Attach the handlers
	handler.AttachHandler(app)

	log.Fatal(app.Listen(":3000"))
}
