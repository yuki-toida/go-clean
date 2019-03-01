package interactor_user

import (
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_profile"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_user"
)

type interactor struct {
	userRepository    entity_user.Repository
	profileRepository entity_profile.Repository
}

func New(ur entity_user.Repository, pr entity_profile.Repository) *interactor {
	return &interactor{userRepository: ur, profileRepository: pr}
}

func (i *interactor) Find() ([]entity_user.User, error) {
	user := entity_user.New(i.userRepository)
	return user.Find()
}

func (i *interactor) First(id uint64) (*entity_user.User, error) {
	user := entity_user.New(i.userRepository)
	return user.First(id)
}

func (i *interactor) Create(name string) (*entity_user.User, error) {
	profile := entity_profile.New(i.profileRepository)
	p, err := profile.Create(name)
	if err != nil {
		return nil, err
	}

	user := entity_user.New(i.userRepository)
	return user.Create(p)
}

func (i *interactor) Delete(userID uint64) error {
	user := entity_user.New(i.userRepository)
	return user.Delete(userID)
}
