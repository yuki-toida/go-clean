package interactors

import (
	"github.com/yuki-toida/go-clean/src/application/repositories"
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type userInteractor struct {
	userRepository    repositories.UserRepository
	profileRepository repositories.ProfileRepository
}

func NewUserInteractor(ur repositories.UserRepository, pr repositories.ProfileRepository) *userInteractor {
	return &userInteractor{userRepository: ur, profileRepository: pr}
}

func (i *userInteractor) Find() ([]entities.User, error) {
	return i.userRepository.Find()
}

func (i *userInteractor) First(id uint64) (*entities.User, error) {
	return i.userRepository.First(id)
}

func (i *userInteractor) Create(name string) (*entities.User, error) {
	p, err := i.profileRepository.Create(name)
	if err != nil {
		return nil, err
	}
	return i.userRepository.Create(p)
}

func (i *userInteractor) Delete(userID uint64) error {
	return i.userRepository.Delete(userID)
}
