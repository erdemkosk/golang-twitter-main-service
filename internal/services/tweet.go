package services

import (
	"context"
	"fmt"
	"log"

	"github.com/erdemkosk/golang-twitter-main-service/internal/models"
	"github.com/erdemkosk/golang-twitter-main-service/internal/repositories"
	"github.com/erdemkosk/golang-twitter-main-service/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetService struct {
	repository  repositories.MongoDBRepository
	userService UserService
}

func CreateTweetService(userService *UserService) *TweetService {
	tweetRepository := repositories.NewMongoDBRepository(database.Client.Database("twitter"), "tweet", models.Tweet{})

	return &TweetService{repository: *tweetRepository, userService: *userService}
}

func (this TweetService) CreateTweet(ctx context.Context, tweet models.Tweet) error {
	user, err := this.userService.GetUserById(ctx, tweet.UserId)
	if err != nil {
		return err
	}

	tweetId, error := this.repository.InsertOne(ctx, tweet)

	tweetObjectID, err := primitive.ObjectIDFromHex(tweetId)
	if err != nil {
		return err
	}

	user.Tweets = append(user.Tweets, tweetObjectID)

	err = this.userService.UpdateUser(ctx, *user)

	return error
}

func (this TweetService) GetTweets(ctx context.Context) ([]*models.Tweet, error) {
	filterAll := bson.M{}

	resultAll, err := this.repository.Find(ctx, filterAll)

	if err != nil {
		log.Fatal(err)
	}

	var tweets []*models.Tweet

	for _, result := range resultAll {

		tweet, ok := result.(*models.Tweet)
		if !ok {
			log.Fatal("cast error")
		}
		tweets = append(tweets, tweet)
	}

	return tweets, err
}

func (this TweetService) GetTweetById(ctx context.Context, id string) (*models.Tweet, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	result, err := this.repository.FindOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	tweet, ok := result.(*models.Tweet)
	if !ok {
		log.Fatal("Type Error")
	}

	return tweet, nil
}

func (this TweetService) DeleteTweet(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	this.repository.DeleteOne(ctx, filter)

	return err
}
