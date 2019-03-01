package repository_user

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/go-clean/src/adapter/repositories/repository_email"
	"github.com/yuki-toida/go-clean/src/adapter/repositories/repository_profile"
	"github.com/yuki-toida/go-clean/src/domain/entities"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_email"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_profile"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_user"
)

type User struct {
	entities.Model
	Profile   repository_profile.Profile
	ProfileID uint64 `gorm:"not null"`
	Emails    []repository_email.Email
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
}

func ToEntity(user *User) *entity_user.User {
	emails := make([]entity_email.Email, len(user.Emails))
	return &entity_user.User{
		Model: user.Model,
		Profile: &entity_profile.Profile{
			Model: user.Profile.Model,
			Name:  user.Profile.Name,
		},
		Emails: emails,
	}
}

func (r *repository) findDB() ([]User, error) {
	var users []User
	if err := r.db.Preload("Profile").Preload("Emails").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) firstDB(userID uint64) (*User, error) {
	user := User{Model: entities.Model{ID: userID}}
	if err := r.db.First(&user).Related(&user.Profile).Related(&user.Emails).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Find() ([]entity_user.User, error) {
	dbUsers, err := r.findDB()
	if err != nil {
		return nil, err
	}

	users := make([]entity_user.User, len(dbUsers))
	for i, u := range dbUsers {
		emails := make([]entity_email.Email, len(u.Emails))
		for j, e := range u.Emails {
			emails[j] = *repository_email.ToEntity(&e)
		}
		users[i] = *ToEntity(&u)
		users[i].Profile = repository_profile.ToEntity(&u.Profile)
		users[i].Emails = emails
	}
	return users, nil
}

func (r *repository) First(userID uint64) (*entity_user.User, error) {
	dbUser, err := r.firstDB(userID)
	if err != nil {
		return nil, err
	}

	emails := make([]entity_email.Email, len(dbUser.Emails))
	for i, v := range dbUser.Emails {
		emails[i] = *repository_email.ToEntity(&v)
	}

	user := ToEntity(dbUser)
	user.Profile = repository_profile.ToEntity(&dbUser.Profile)
	user.Emails = emails
	return user, nil
}

func (r *repository) Create(profile *entity_profile.Profile) (*entity_user.User, error) {
	dbUser := User{ProfileID: profile.ID}
	if err := r.db.Create(&dbUser).Error; err != nil {
		return nil, err
	}

	user := ToEntity(&dbUser)
	user.Profile = profile
	return user, nil
}

func (r *repository) Delete(userID uint64) error {
	dbUser, err := r.First(userID)
	if err != nil {
		return err
	}
	if err := r.db.Delete(&dbUser.Profile).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&dbUser).Error; err != nil {
		return err
	}
	for _, email := range dbUser.Emails {
		r.db.Delete(&email)
	}
	return nil
}
