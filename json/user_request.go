package json

import (
	"github.com/laioncorcino/go-user/domain/model"
)

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Age      int8   `json:"age" binding:"required,min=2,max=140"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

func (request *UserRequest) ToModel() model.UserDomain {
	var domain model.UserDomain
	domain.SetName(request.Name)
	domain.SetEmail(request.Email)
	domain.SetAge(request.Age)
	domain.SetPassword(request.Password)
	return domain
}
