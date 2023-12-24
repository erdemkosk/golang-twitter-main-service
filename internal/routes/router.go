package router

import (
	"github.com/erdemkosk/golang-twitter-main-service/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func Initalize(router *fiber.App) {

	tweetController := controllers.CreateTweetController()
	userController := controllers.CreateUserController()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	tweets := router.Group("/tweets")
	tweets.Post("/", tweetController.CreateTweet)
	tweets.Delete("/:id", tweetController.DeleteTweet)
	tweets.Get("/", tweetController.GetTweets)
	tweets.Get("/:id", tweetController.GetTweetById)

	users := router.Group("/users")
	users.Post("/", userController.CreateUser)
	users.Get("/", userController.GetUsers)
	users.Get("/:id", userController.GetUserById)
	users.Post("/:followerID/follow/:followedID", userController.FollowUser)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
