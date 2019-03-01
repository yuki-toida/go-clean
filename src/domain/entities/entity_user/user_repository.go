package entity_user

import (
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_profile"
)

type Repository interface {
	Find() ([]User, error)
	First(userID uint64) (*User, error)
	Create(profile *entity_profile.Profile) (*User, error)
	Delete(userID uint64) error
}
