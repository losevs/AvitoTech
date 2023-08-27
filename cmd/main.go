package main

import (
	"avitoModRhino/database"
	"avitoModRhino/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":80") //localhost:80
}

func setupRoutes(app *fiber.App) {
	//SEGMENTS

	//New Segment
	app.Post("/add", handlers.NewSegment)
	//Show all segments
	app.Get("/show", handlers.ShowSegment)
	//Delete segment
	app.Delete("/del", handlers.DeleteSegment)

	//USER

	//Add new user / change user's segments
	app.Post("/user", handlers.UserSegment)
	//Show all users
	app.Get("/user/show", handlers.UserShow)
	//Delete user
	app.Delete("/user/delete/:id", handlers.UserDelete)
	//Show exact user
	app.Get("/user/show/:id", handlers.ShowExactUser)

}
