package entity_user

import (
	"github.com/yuki-toida/go-clean/src/domain/entities"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_email"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_profile"
)

type User struct {
	entities.Model
	Profile *entity_profile.Profile
	Emails  []entity_email.Email
}

type user struct {
	repository Repository
}

func New(r Repository) *user {
	return &user{repository: r}
}

func (u *user) Find() ([]User, error) {
	return u.repository.Find()
}

func (u *user) First(userID uint64) (*User, error) {
	return u.repository.First(userID)
}

func (u *user) Create(profile *entity_profile.Profile) (*User, error) {
	return u.repository.Create(profile)
}

func (u *user) Delete(userID uint64) error {
	return u.repository.Delete(userID)
}
