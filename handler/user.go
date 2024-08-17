package handler

import (
	"ecommerce-gofiber/database"
	"ecommerce-gofiber/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validUser(id string, p string) bool {
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}

	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// GetUser godoc
//	@Summary		Get a user by ID
//	@Description	Retrieve user details by user ID
//	@Tags			User
//	@Param			id	path	string	true	"User ID"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"User found"
//	@Failure		404	{object}	map[string]interface{}	"No user found with ID"
//	@Failure		500	{object}	map[string]interface{}	"Internal Server Error"
//	@Router			/api/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// CreateUser new user
//	@Summary		Create User
//	@Description	Create a new user with the provided username and email
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.UserSwagger		true	"User info"
//	@Success		200		{object}	map[string]interface{}	"status: success, message: Created user, data: {model.User.Username, model.User.Email}"
//	@Failure		500		{object}	map[string]interface{}	"status: error, message: Review your input or Couldn't create user"
//	@Router			/api/user [post]
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// UpdateUser update user
// UpdateUser godoc
//	@Summary		Update a user's information
//	@Description	Update a user's names based on the provided ID and input data
//	@Tags			User
//	@Param			id		path	string					true	"User ID"
//	@Param			input	body	model.UserUpdateSwagger	true	"Update user input"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"User successfully updated"
//	@Failure		400	{object}	map[string]interface{}	"Invalid input"
//	@Failure		500	{object}	map[string]interface{}	"Internal Server Error"
//	@Security		BearerAuth
//	@Router			/api/user/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Names string `json:"names"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	db := database.DB
	var user model.User

	db.First(&user, id)
	user.Names = uui.Names
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

// DeleteUser delete user
// DeleteUser godoc
//	@Summary		Delete a user by ID
//	@Description	Delete a user based on the provided ID and password for verification
//	@Tags			User
//	@Param			id		path	string					true	"User ID"
//	@Param			input	body	model.UserDeleteSwagger	true	"Password input for user verification"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"User successfully deleted"
//	@Failure		400	{object}	map[string]interface{}	"Invalid input"
//	@Failure		500	{object}	map[string]interface{}	"Internal Server Error"
//	@Security		BearerAuth
//	@Router			/api/user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})

	}

	if !validUser(id, pi.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})

	}

	db := database.DB
	var user model.User

	db.First(&user, id)

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}