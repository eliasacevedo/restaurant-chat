package services

import (
	"fastfoodrestaurant.com/api/models"
)

type IUserService interface {
	GetUser() (*models.User, error)
	GetBotInfo() (*models.User, error)
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUser() (*models.User, error) {
	user := &models.User{
		Name:      "Elias",
		PhotoPath: "https://previews.123rf.com/images/yutthaphan/yutthaphan1702/yutthaphan170200003/71426489-comida-r%C3%A1pida-conjunto-de-comida-r%C3%A1pida-comida-chatarra-.jpg",
	}
	return user, nil
}

func (us *UserService) GetBotInfo() (*models.User, error) {
	user := &models.User{
		Name:      "Fast food bot",
		PhotoPath: "https://previews.123rf.com/images/yutthaphan/yutthaphan1702/yutthaphan170200003/71426489-comida-r%C3%A1pida-conjunto-de-comida-r%C3%A1pida-comida-chatarra-.jpg",
	}
	return user, nil
}
