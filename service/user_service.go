package service

import (
	"github.com/laioncorcino/go-user/config/logger"
	"github.com/laioncorcino/go-user/json"
	"github.com/laioncorcino/go-user/repository"
	"github.com/laioncorcino/go-user/rest_error"
	"go.uber.org/zap"
)

type IUserService interface {
	CreateUser(req json.UserRequest) (*json.UserResponse, *rest_error.StandardError)
	FindUserById(userId string) (json.UserResponse, *rest_error.StandardError)
	FindUserByEmail(email string) (json.UserResponse, *rest_error.StandardError)
}

type UserService struct {
	repo repository.IUserRepository
}

func (service UserService) CreateUser(req json.UserRequest) (*json.UserResponse, *rest_error.StandardError) {
	model := req.ToModel()
	model.EncryptPassword()

	user, err := service.repo.CreateUser(model)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", user.GetID()), zap.String("journey", "createUser"))
	response := json.ToResponse(user)
	return &response, nil
}
