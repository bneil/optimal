package main

import (
	"github.com/bneil/gossr_tests/app/api"
	"github.com/bneil/gossr_tests/app/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	//quickly setup that db!
	d := db.GetInstance()
	err := d.SetupDb()
	if err != nil {
		panic("couldnt setup db, freak out")
	}

	engine := html.New("./app/views/templates", ".gohtml")
	config := fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		Prefork:     false,
	}
	app := api.SetupRoutes(config)
	app.Static("/static", "./app/views/static")

	log.Fatal(app.Listen(":3000"))
}
