package app

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Feed struct {
	ID          string
	Description string
	Complete    bool
}

var feeds = []Feed{
	{ID: "1", Description: "Testing1", Complete: false},
	{ID: "2", Description: "Testing2", Complete: false},
	{ID: "3", Description: "Testing3", Complete: false},
}

func AttachHandler(app *fiber.App) {
	app.Get("/", index)
	app.Get("/feeds/:id/edit", edit)
	app.Patch("/feeds/:id", patchFeed)
	app.Delete("/feeds/:id", deleteFeed)
}

func deleteFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	feed := getFeed(id)
	if feed == nil {
		return nil
	}

	var l []Feed
	for _, feed := range feeds {
		f := feed
		if strings.ToLower(f.ID) == strings.ToLower(id) {
			continue
		}
		l = append(l, f)
	}

	feeds = l

	return nil
}

func getFeed(id string) *Feed {
	for _, feed := range feeds {
		if strings.ToLower(feed.ID) == id {
			f := feed
			return &f
		}
	}
	return nil
}

func patchFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	feed := getFeed(id)
	if feed == nil {
		return nil
	}

	newDesc := c.FormValue("name")

	var l []Feed
	var updatedFeed Feed
	for _, feed := range feeds {
		f := feed
		if strings.ToLower(f.ID) == strings.ToLower(id) {
			f.Description = newDesc
			updatedFeed = f
		}
		l = append(l, f)
	}

	feeds = l

	return c.Render("view_feed", updatedFeed)
}

func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":            "Hi Folks",
		"Feeds":            feeds,
		"PopularFeeds":     feeds,
		"RecentFeeds":      feeds,
		"CategorizedFeeds": feeds,
	})
}

func edit(c *fiber.Ctx) error {
	//id, err := uuid.Parse(c.Params("id"))
	//if err != nil {
	//	return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	//}
	feed := getFeed(c.Params("id"))
	if feed == nil {
		return c.Status(fiber.StatusNotFound).SendString("not found")
	}

	return c.Render("edit_feed", feed)
}

/*
package main

import (
    "database/sql"
    "fmt"
    "github.com/gofiber/fiber/v2"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // Open a connection to the database
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create the users table if it doesn't exist
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT
    )`)
    if err != nil {
        fmt.Println(err)
        return
    }

    app := fiber.New()

    // Create a new user
    app.Post("/users", func(c *fiber.Ctx) error {
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            return err
        }

        stmt, err := db.Prepare("INSERT INTO users(name, email) values(?, ?)")
        if err != nil {
            return err
        }
        defer stmt.Close()

        result, err := stmt.Exec(user.Name, user.Email)
        if err != nil {
            return err
        }

        id, err := result.LastInsertId()
        if err != nil {
            return err
        }

        user.ID = int(id)

        return c.JSON(user)
    })

    // Read users
    app.Get("/users", func(c *fiber.Ctx) error {
        rows, err := db.Query("SELECT id, name, email FROM users")
        if err != nil {
            return err
        }
        defer rows.Close()

        var users []*User
        for rows.Next() {
            user := new(User)
            err := rows.Scan(&user.ID, &user.Name, &user.Email)
            if err != nil {
                return err
            }
            users = append(users, user)
        }

        return c.JSON(users)
    })

    // Update a user
    app.Put("/users/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")

        user := new(User)
        if err := c.BodyParser(user); err != nil {
            return err
        }

        stmt, err := db.Prepare("UPDATE users SET name=?, email=? WHERE id=?")
        if err != nil {
            return err
        }
        defer stmt.Close()

        _, err = stmt.Exec(user.Name, user.Email, id)
        if err != nil {
            return err
        }

        return c.SendString("User updated")
    })

    // Delete a user
    app.Delete("/users/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")

        stmt, err := db.Prepare("DELETE FROM users WHERE id=?")
        if err != nil {
            return err
        }
        defer stmt.Close()

        _, err = stmt.Exec(id)
        if err != nil {
            return err
        }

        return c.SendString("User deleted")
    })

    app.Listen(":3000")
}

*/
