package handler

import "github.com/gofiber/fiber/v2"

// Hello godoc
//	@Summary		Hello endpoint
//	@Description	Simple hello endpoint to test the API
//	@Tags			Hello
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Hello World"
//	@Router			/api [get]
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}