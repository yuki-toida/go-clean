package web

import (
	"context"
	"encoding/json"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/yuki-toida/go-clean/src/adapter/transports/transport_email"
	"github.com/yuki-toida/go-clean/src/adapter/transports/transport_user"
	"github.com/yuki-toida/go-clean/src/registry/interfaces"
)

func Handle(u interfaces.UseCase) http.Handler {
	router := mux.NewRouter()

	uu := u.NewUserUseCase()
	router.Methods("GET").Path("/users").Handler(transport.NewServer(
		transport_user.MakeFindEndpoint(uu),
		defaultDecodeRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/users").Handler(transport.NewServer(
		transport_user.MakeCreateEndpoint(uu),
		transport_user.DecodeCreateRequest,
		encodeResponse,
	))
	router.Methods("GET").Path("/users/{uid}").Handler(transport.NewServer(
		transport_user.MakeFirstEndpoint(uu),
		transport_user.DecodeFirstRequest,
		encodeResponse,
	))
	router.Methods("DELETE").Path("/users/{uid}").Handler(transport.NewServer(
		transport_user.MakeDeleteEndpoint(uu),
		transport_user.DecodeDeleteRequest,
		encodeResponse,
	))

	eu := u.NewEmailUseCase()
	router.Methods("POST").Path("/emails").Handler(transport.NewServer(
		transport_email.MakeCreateEndpoint(eu),
		transport_email.DecodeCreateRequest,
		encodeResponse,
	))
	router.Methods("PATCH").Path("/emails").Handler(transport.NewServer(
		transport_email.MakeUpdateEndpoint(eu),
		transport_email.DecodeUpdateRequest,
		encodeResponse,
	))

	return router
}

func defaultDecodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return r, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
