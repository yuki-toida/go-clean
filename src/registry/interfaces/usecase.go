package interfaces

import (
	"github.com/yuki-toida/go-clean/src/application/usecase"
)

type UseCase interface {
	NewUserUseCase() usecase.UserUseCase
	NewEmailUseCase() usecase.EmailUseCase
}
