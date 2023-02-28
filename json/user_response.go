package json

import "github.com/laioncorcino/go-user/domain/model"

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
}

func ToResponse(userDomain model.UserDomain) UserResponse {
	return UserResponse{
		userDomain.GetID(),
		userDomain.GetName(),
		userDomain.GetAge(),
		userDomain.GetEmail(),
	}
}
