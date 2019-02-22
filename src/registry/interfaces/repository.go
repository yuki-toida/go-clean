package interfaces

import (
	"github.com/yuki-toida/go-clean/src/application/repositories"
)

type Repository interface {
	NewUserRepository() repositories.UserRepository
	NewProfileRepository() repositories.ProfileRepository
	NewEmailRepository() repositories.EmailRepository
}
