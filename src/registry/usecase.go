package registry

import (
	"github.com/yuki-toida/go-clean/src/application/interactors"
	"github.com/yuki-toida/go-clean/src/application/usecase"
	"github.com/yuki-toida/go-clean/src/registry/interfaces"
)

type useCase struct {
	repository interfaces.Repository
}

func NewUseCase(r interfaces.Repository) interfaces.UseCase {
	return &useCase{repository: r}
}

func (u *useCase) NewUserUseCase() usecase.UserUseCase {
	ur := u.repository.NewUserRepository()
	pu := u.repository.NewProfileRepository()
	return interactors.NewUserInteractor(ur, pu)
}

func (u *useCase) NewEmailUseCase() usecase.EmailUseCase {
	er := u.repository.NewEmailRepository()
	return interactors.NewEmailInteractor(er)
}
