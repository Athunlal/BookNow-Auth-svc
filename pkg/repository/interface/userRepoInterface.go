package interfaces

import "github.com/athunlal/bookNow-auth-svc/pkg/domain"

type UserRepo interface {
	Create(user domain.User) error
	FindByUserName(user domain.User) (domain.User, error)
	FindByUserEmail(user domain.User) (domain.User, error)
	FindUserByOtp(user domain.User) (domain.User, error)
	NullTheOtp(user domain.User) int64
	FindUserById(userid uint) (domain.User, error)
	IsOtpVerified(username string) string
	DeleteUser(user domain.User) error
	UpdateOtp(user domain.User) error
	VerifyUser(user domain.User) (domain.User, error)
	ChangePassword(user domain.User) error
	
}
