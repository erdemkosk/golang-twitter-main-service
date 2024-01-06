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

type UserService struct {
	repository repositories.MongoDBRepository
}

func CreateUserService() *UserService {
	userRepository := repositories.NewMongoDBRepository(database.Client.Database("twitter"), "user", models.User{})
	return &UserService{repository: *userRepository}
}

func (userService UserService) CreateUser(ctx context.Context, user models.User) error {
	_, error := userService.repository.InsertOne(ctx, user)

	return error
}

func (userService UserService) GetUsers(ctx context.Context) ([]*models.User, error) {
	filterAll := bson.M{}

	resultAll, err := userService.repository.Find(ctx, filterAll)

	if err != nil {
		log.Fatal(err)
	}

	var users []*models.User

	fmt.Println(resultAll)

	for _, result := range resultAll {

		user, ok := result.(*models.User)
		if !ok {
			log.Fatal("cast error")
		}
		users = append(users, user)
	}

	return users, err
}

func (userService UserService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	result, err := userService.repository.FindOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	user, ok := result.(*models.User)
	if !ok {
		log.Fatal("Type Error")
	}

	return user, nil
}

func (userService UserService) UpdateUser(ctx context.Context, user models.User) error {
	err := userService.repository.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}

	return nil
}

func (userService UserService) FollowUser(ctx context.Context, followerID, followedID string) error {
	follower, err := userService.GetUserById(ctx, followerID)
	if err != nil {
		return err
	}

	followed, err := userService.GetUserById(ctx, followedID)
	if err != nil {
		return err
	}

	follower.Following = append(follower.Following, followed.ID)
	followed.Followers = append(followed.Followers, follower.ID)

	if err := userService.UpdateUser(ctx, *follower); err != nil {
		return err
	}

	if err := userService.UpdateUser(ctx, *followed); err != nil {
		return err
	}

	return nil

}
