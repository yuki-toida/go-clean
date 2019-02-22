package usecase

import (
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type UserUseCase interface {
	Find() ([]entities.User, error)
	First(userID uint64) (*entities.User, error)
	Create(name string) (*entities.User, error)
	Delete(userID uint64) error
}
