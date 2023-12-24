package controllers

import (
	"github.com/erdemkosk/golang-twitter-main-service/internal/models"
	"github.com/erdemkosk/golang-twitter-main-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserService
}

func CreateUserController() *UserController {
	UserService := services.CreateUserService()
	return &UserController{userService: *UserService}
}

func (t UserController) CreateUser(c *fiber.Ctx) error {
	payload := models.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	err := t.userService.CreateUser(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": "success",
	})
}

func (t UserController) GetUsers(c *fiber.Ctx) error {
	users, err := t.userService.GetUsers(c.Context())

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": users,
	})
}

func (t UserController) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := t.userService.GetUserById(c.Context(), id)

	if err != nil {
		return c.JSON(fiber.Map{
			"code": fiber.ErrBadRequest.Code,
			"data": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": user,
	})
}

func (t UserController) FollowUser(c *fiber.Ctx) error {
	followerID := c.Params("followerID")
	followedID := c.Params("followedID")

	err := t.userService.FollowUser(c.Context(), followerID, followedID)

	if err != nil {
		return c.JSON(fiber.Map{
			"code": fiber.ErrBadRequest.Code,
			"data": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": "You followed with succesfully!",
	})
}
