package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/yuki-toida/go-clean/src/application/usecase"
)

func MakeUserFindEndpoint(u usecase.UserUseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		users, err := u.Find()
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}
