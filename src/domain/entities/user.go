package entities

type User struct {
	Model
	Profile *Profile
	Emails  []Email
}
