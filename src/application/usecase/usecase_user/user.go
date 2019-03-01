package usecase_user

import (
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_user"
)

type UseCase interface {
	Find() ([]entity_user.User, error)
	First(userID uint64) (*entity_user.User, error)
	Create(name string) (*entity_user.User, error)
	Delete(userID uint64) error
}
