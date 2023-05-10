package controllers

import (
	"IdeaIntuition/app/models"
	"IdeaIntuition/services"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerRequest struct {
	request
	ConfirmPassword string `json:"confirm_password"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}

func Login(c *fiber.Ctx) error {
	var body request
	if e, err := validateUserParams(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": e,
		})
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	//search user in db
	user, err := models.GetUserByEmail(body.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	//verify password
	if err = user.ComparePassword(body.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}
	t, err := services.GenerateToken(user.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"user":  user,
		"token": t,
	})
}

func Register(c *fiber.Ctx) error {
	var body registerRequest

	if e, err := validateUserParams(c, true); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": e,
		})
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	//hash password
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot hash password",
		})
	}
	//create user
	user := models.User{
		Email:    body.Email,
		Password: string(password),
	}
	//save user
	if err = user.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create user",
		})
	}

	t, err := services.GenerateToken(user.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":  user,
		"token": t,
	})
}

func Restricted(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Welcome to the restricted area!",
	})
}

func validateUserParams(c *fiber.Ctx, register ...bool) (ErrorMap, error) {
	//get params of body request
	email := c.FormValue("email")
	password := c.FormValue("password")

	e := make(ErrorMap)

	// Valida el correo electrónico
	if !govalidator.IsEmail(email) {
		e["email"] = "Invalid email address"
	}

	// Valida la contraseña (ejemplo: mínimo 6 caracteres)
	if !govalidator.MinStringLength(password, "6") {
		e["password"] = "Password must be at least 6 characters long"
	}

	if len(register) > 0 && register[0] != false {
		// Valida la confirmación de la contraseña
		confirmPassword := c.FormValue("confirm_password")
		if password != confirmPassword {
			e["confirm_password"] = "Passwords do not match"
		}
		firstName := c.FormValue("first_name")
		lastName := c.FormValue("last_name")
		if !govalidator.MinStringLength(firstName, "2") {
			e["first_name"] = "First name must be at least 3 characters long"
		}
		if !govalidator.MinStringLength(lastName, "2") {
			e["last_name"] = "Last name must be at least 3 characters long"
		}
	}

	if len(e) > 0 {
		return e, errors.New("Invalid params")
	}

	return e, nil
}
