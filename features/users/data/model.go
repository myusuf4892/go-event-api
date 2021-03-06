package data

import (
	"project3/eventapp/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	URL      string `json:"url" form:"url"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	RoleID   int
	Role     Role
}

type Role struct {
	gorm.Model
	Name string
}

//DTO

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		URL:       data.URL,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

}

// func toCoreList(data []User) []users.Core {
// 	result := []users.Core{}
// 	for key := range data {
// 		result = append(result, data[key].toCore())
// 	}
// 	return result
// }

func fromCore(core users.Core) User {
	return User{
		URL:      core.URL,
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
	}
}

func toCore(data User) users.Core {
	return data.toCore()
}
