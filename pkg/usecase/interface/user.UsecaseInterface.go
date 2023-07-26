package interfaces

import "github.com/athunlal/bookNow-auth-svc/pkg/domain"

type UserUseCase interface {
	Register(user domain.User) error
	RegisterValidate(user domain.User) (domain.User, error)
	Login(user domain.User) (domain.User, error)
	ValidateJwtUser(userid uint) (domain.User, error)
	ForgotPassword(user domain.User) error
	ChangePassword(user domain.User) error
}
