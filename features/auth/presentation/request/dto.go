package request

import (
	"project3/eventapp/features/auth"
)

type User struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func ToCore(userReq User) auth.Core {
	userCore := auth.Core{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
