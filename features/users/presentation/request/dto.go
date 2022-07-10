package request

import "project3/eventapp/features/users"

type User struct {
	URL      string `json:"url" form:"url"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     users.Role
}

func ToCore(userReq User) users.Core {
	userCore := users.Core{
		URL:      userReq.URL,
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
		Role: users.Role{
			ID:   userReq.Role.ID,
			Name: userReq.Role.Name,
		},
	}
	return userCore
}
