package business

import (
	"project3/eventapp/features/auth"
	"project3/eventapp/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userData auth.Data
}

func NewAuthBusiness(usrData auth.Data) auth.Business {
	return &authUseCase{
		userData: usrData,
	}
}

func (uc *authUseCase) Login(data auth.Core) (string, int, error) {
	response, errFind := uc.userData.FindUser(data.Email)
	if errFind != nil {
		return "", 0, errFind
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(data.Password))
	if errCompare != nil {
		return "", 0, errCompare
	}
	token, err := middlewares.CreateToken(int(response.ID))

	return token, response.ID, err
}
