package controllers

import (
	"github.com/erdemkosk/golang-twitter-main-service/internal/models"
	"github.com/erdemkosk/golang-twitter-main-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type TweetController struct {
	tweetService services.TweetService
	userService  services.UserService
}

func CreateTweetController() *TweetController {
	userService := services.CreateUserService()
	TweetService := services.CreateTweetService(userService)
	return &TweetController{tweetService: *TweetService}
}

func (t TweetController) CreateTweet(c *fiber.Ctx) error {
	payload := models.Tweet{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	err := t.tweetService.CreateTweet(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": "success",
	})
}

func (t TweetController) GetTweets(c *fiber.Ctx) error {
	tweets, err := t.tweetService.GetTweets(c.Context())

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": tweets,
	})
}

func (t TweetController) GetTweetById(c *fiber.Ctx) error {
	id := c.Params("id")

	tweet, err := t.tweetService.GetTweetById(c.Context(), id)

	if err != nil {
		return c.JSON(fiber.Map{
			"code": fiber.ErrBadRequest.Code,
			"data": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": tweet,
	})
}

func (t TweetController) DeleteTweet(c *fiber.Ctx) error {
	id := c.Params("id")

	err := t.tweetService.DeleteTweet(c.Context(), id)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": "success",
	})

}
