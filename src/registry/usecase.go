package registry

import (
	"github.com/yuki-toida/go-clean/src/application/interactors/interactor_email"
	"github.com/yuki-toida/go-clean/src/application/interactors/interactor_user"
	"github.com/yuki-toida/go-clean/src/application/usecase/usecase_email"
	"github.com/yuki-toida/go-clean/src/application/usecase/usecase_user"
	"github.com/yuki-toida/go-clean/src/registry/interfaces"
)

type useCase struct {
	repository interfaces.Repository
}

func NewUseCase(r interfaces.Repository) interfaces.UseCase {
	return &useCase{repository: r}
}

func (u *useCase) NewUserUseCase() usecase_user.UseCase {
	ur := u.repository.NewUserRepository()
	pu := u.repository.NewProfileRepository()
	return interactor_user.New(ur, pu)
}

func (u *useCase) NewEmailUseCase() usecase_email.UseCase {
	er := u.repository.NewEmailRepository()
	return interactor_email.New(er)
}
