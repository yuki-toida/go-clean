package services

import (
	"github.com/yuki-toida/go-clean/src/application/usecase"
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type userServiceImpl struct {
	userUseCase usecase.UserUseCase
}

func NewUserService(u usecase.UserUseCase) *userServiceImpl {
	return &userServiceImpl{userUseCase: u}
}

func (u *userServiceImpl) Find() ([]entities.User, error) {
	return u.userUseCase.Find()
}

func (u *userServiceImpl) First(userID uint64) (*entities.User, error) {
	return u.userUseCase.First(userID)
}

func (u *userServiceImpl) Create(name string) (*entities.User, error) {
	return u.userUseCase.Create(name)
}

func (u *userServiceImpl) Delete(userID uint64) error {
	return u.Delete(userID)
}
