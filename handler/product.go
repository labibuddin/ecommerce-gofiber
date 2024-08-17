package handler

import (
	"ecommerce-gofiber/database"
	"ecommerce-gofiber/model"

	"github.com/gofiber/fiber/v2"
)

// Get All Products
// GetAllProducts godoc
//	@Summary		Get all products
//	@Description	Retrieve a list of all products available in the store
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"List of products"
//	@Failure		500	{object}	map[string]interface{}	"Internal Server Error"
//	@Router			/api/product [get]
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []model.Product
	db.Find(&products)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// Get product with id
// GetProduct godoc
//	@Summary		Get a product by ID
//	@Description	Retrieve details of a specific product by its ID
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true	"Product ID"
//	@Success		200	{object}	map[string]interface{}	"Product details"
//	@Failure		404	{object}	map[string]interface{}	"Product not found"
//	@Failure		500	{object}	map[string]interface{}	"Internal Server Error"
//	@Router			/api/product/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var product model.Product
	db.Find(&product, id)
	if product.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success","message": "Product found",	"data": product})
}

// membuat product baru
// CreateProduct godoc
//	@Summary		Create a new product
//	@Description	Create a new product with the provided title, description, and amount
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		model.ProductSwagger	true	"Product data"	Example({"title": "Product A", "description": "A great product", "amount": 10})
//	@Success		200		{object}	map[string]interface{}	"Product created"
//	@Failure		400		{object}	map[string]interface{}	"Invalid input"
//	@Failure		500		{object}	map[string]interface{}	"Internal Server Error"
//	@Security		BearerAuth
//	@Router			/api/product [post]
func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	db.Create(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// delete product 
// Delete product with id
// Delete Product godoc
//	@Summary		Delete a product by ID
//	@Description	Delete a specific product by its ID
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true	"Product ID"
//	@Success		200	{object}	map[string]interface{}	"Product deleted"
//	@Failure		404	{object}	map[string]interface{}	"Product not found"
//	@Failure		500	{object}	map[string]interface{}	"Internal Server Error"
//	@Security		BearerAuth
//	@Router			/api/product/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.First(&product, id)

	if product.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	db.Delete(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}

