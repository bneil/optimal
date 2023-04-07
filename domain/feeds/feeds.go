package feeds

/*

Opml -> Feeds
     -> classify those endpoints of personal content
     -> find unique  things

*/

type Feed struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Type    string `json:"type"`
	HtmlUrl string `json:"html_url"`
	XMLUrl  string `json:"xml_url"`
}

type List struct {
	Feeds []Feed
}

func (f *Feed) Read() {}

func (f *Feed) ReadAll() {}

func (f *Feed) Update() {}

func (f *Feed) Delete() {}

/*
func (f *Feed) Create() error {
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

	f.ID = int(id)
}

/*
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

*/
