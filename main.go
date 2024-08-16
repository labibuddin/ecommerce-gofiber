package main

import (
	"ecommerce-gofiber/database"
	_ "ecommerce-gofiber/docs"
	"ecommerce-gofiber/router"
	"log"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//	@title						Simple e-Commerce
//	@version					1.0
//	@description				This is a simple e-Commerce with Auth and postgres as database
//	@contact.name				Alfin Afifi
//	@contact.email				labibuddinalfin@gmail.com
//	@license.name				labibuddin 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:3000
//	@BasePath					/
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

