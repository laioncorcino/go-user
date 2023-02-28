package repository

import (
	"github.com/laioncorcino/go-user/domain/model"
	"github.com/laioncorcino/go-user/rest_error"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MongodbUserDb = "MONGODB_USER_DB"
)

func NewUserRepository(database *mongo.Database) IUserRepository {
	return &UserRepository{database}
}

type UserRepository struct {
	databaseConnection *mongo.Database
}

type IUserRepository interface {
	CreateUser(user model.UserDomain) (model.UserDomain, *rest_error.StandardError)
	FindUserByEmail(email string) (model.UserDomain, *rest_error.StandardError)
	FindUserByID(id string) (model.UserDomain, *rest_error.StandardError)
}

func (u UserRepository) CreateUser(user model.UserDomain) (model.UserDomain, *rest_error.StandardError) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindUserByEmail(email string) (model.UserDomain, *rest_error.StandardError) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindUserByID(id string) (model.UserDomain, *rest_error.StandardError) {
	//TODO implement me
	panic("implement me")
}
