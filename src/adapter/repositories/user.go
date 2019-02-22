package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type User struct {
	entities.Model
	Profile   Profile
	ProfileID uint64 `gorm:"not null"`
	Emails    []Email
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func toUserEntity(user *User) *entities.User {
	emails := make([]entities.Email, len(user.Emails))
	return &entities.User{
		Model: user.Model,
		Profile: &entities.Profile{
			Model: user.Profile.Model,
			Name:  user.Profile.Name,
		},
		Emails: emails,
	}
}

func (r *userRepository) find() ([]User, error) {
	var users []User
	if err := r.db.Preload("Profile").Preload("Emails").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) first(userID uint64) (*User, error) {
	user := User{Model: entities.Model{ID: userID}}
	if err := r.db.First(&user).Related(&user.Profile).Related(&user.Emails).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Find() ([]entities.User, error) {
	dbUsers, err := r.find()
	if err != nil {
		return nil, err
	}

	users := make([]entities.User, len(dbUsers))
	for i, u := range dbUsers {
		emails := make([]entities.Email, len(u.Emails))
		for j, e := range u.Emails {
			emails[j] = *toEmailEntity(e)
		}
		users[i] = *toUserEntity(&u)
		users[i].Profile = toProfileEntity(u.Profile)
		users[i].Emails = emails
	}
	return users, nil
}

func (r *userRepository) First(userID uint64) (*entities.User, error) {
	dbUser, err := r.first(userID)
	if err != nil {
		return nil, err
	}

	emails := make([]entities.Email, len(dbUser.Emails))
	for i, v := range dbUser.Emails {
		emails[i] = *toEmailEntity(v)
	}

	user := toUserEntity(dbUser)
	user.Profile = toProfileEntity(dbUser.Profile)
	user.Emails = emails
	return user, nil
}

func (r *userRepository) Create(profile *entities.Profile) (*entities.User, error) {
	dbUser := User{ProfileID: profile.ID}
	if err := r.db.Create(&dbUser).Error; err != nil {
		return nil, err
	}

	user := toUserEntity(&dbUser)
	user.Profile = profile
	return user, nil
}

func (r *userRepository) Delete(userID uint64) error {
	dbUser, err := r.first(userID)
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
