package repositories

import (
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type ProfileRepository interface {
	Create(name string) (*entities.Profile, error)
}
