//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/athunlal/bookNow-auth-svc/pkg/api"
	handler "github.com/athunlal/bookNow-auth-svc/pkg/api/handler"
	config "github.com/athunlal/bookNow-auth-svc/pkg/config"
	db "github.com/athunlal/bookNow-auth-svc/pkg/db"
	repo "github.com/athunlal/bookNow-auth-svc/pkg/repository"
	useCase "github.com/athunlal/bookNow-auth-svc/pkg/usecase"
	"github.com/google/wire"
)

func InitApi(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(
		db.ConnectToDb,
		repo.NewUserRepo,
		useCase.NewUserUseCase,
		useCase.NewJWTUseCase,
		handler.NewUserHandler,
		http.NewServerHttp)

	return &http.ServerHttp{}, nil
}
