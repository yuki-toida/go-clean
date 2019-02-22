package repositories

import (
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type UserRepository interface {
	Find() ([]entities.User, error)
	First(userID uint64) (*entities.User, error)
	Create(profile *entities.Profile) (*entities.User, error)
	Delete(userID uint64) error
}
