package usecase

import (
	"errors"

	"github.com/athunlal/bookNow-auth-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNow-auth-svc/pkg/repository/interface"
	userCase "github.com/athunlal/bookNow-auth-svc/pkg/usecase/interface"
	"github.com/athunlal/bookNow-auth-svc/pkg/utils"
)

type userUseCase struct {
	Repo interfaces.UserRepo
}

func (use *userUseCase) Register(user domain.User) error {
	validationError := utils.ValidateUser(user)
	if validationError != nil {
		return validationError
	}

	userData, err := use.Repo.FindByUserEmail(user)
	if err == nil {
		if userData.Isverified == false {
			err := use.Repo.DeleteUser(user)
			if err != nil {
				return errors.New("Could Not delete unethenticated user")
			}
		} else {
			return errors.New("Email Address already exists")
		}
	}
	_, err = use.Repo.FindByUserName(user)
	if err == nil {
		return errors.New("Username Already exists")
	}

	otp := utils.Otpgeneration(user.Email)
	user.Otp = otp

	user.Password = utils.HashPassword(user.Password)

	err = use.Repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (use *userUseCase) RegisterValidate(user domain.User) (domain.User, error) {
	user, err := use.Repo.FindUserByOtp(user)
	if err != nil {
		return user, errors.New("Enterd wrong OTP")
	}

	rows := use.Repo.NullTheOtp(user)
	if rows == 0 {
		return user, errors.New("Could not update the OTP")
	}

	user, err = use.Repo.VerifyUser(user)
	if err != nil {
		return user, errors.New("Could not verify the user")
	}
	return user, nil
}

func (use *userUseCase) Login(user domain.User) (domain.User, error) {
	var userDetatils domain.User
	var err error
	if user.Username != "" {
		userDetatils, err = use.Repo.FindByUserName(user)
		if err != nil {
			return userDetatils, errors.New("User not found")
		}
	} else if user.Email != "" {
		userDetatils, err = use.Repo.FindByUserEmail(user)
		if err != nil {
			return userDetatils, errors.New("User not found")
		}
	}

	if userDetatils.Isverified == false {
		err := use.Repo.DeleteUser(userDetatils)
		if err != nil {
			return userDetatils, errors.New("Could not delete unauthenticateduser")
		}
		return userDetatils, errors.New("User not Authenticated, Register again")

	}

	if !utils.VerifyPassword(user.Password, userDetatils.Password) {
		return userDetatils, errors.New("Password is not matched or worg")
	}
	return userDetatils, nil
}

func (use *userUseCase) ForgotPassword(user domain.User) error {
	user, err := use.Repo.FindByUserEmail(user)
	if err != nil {
		return errors.New("Email Address not found!")
	}

	otp := utils.Otpgeneration(user.Email)
	user.Otp = otp

	err = use.Repo.UpdateOtp(user)
	if err != nil {
		return errors.New("Could not update the OTP")
	}
	return nil
}

func (use *userUseCase) ChangePassword(user domain.User) error {
	user.Password = utils.HashPassword(user.Password)
	err := use.Repo.ChangePassword(user)
	if err != nil {
		return errors.New("Could not change the password")
	}
	return nil
}

func (use *userUseCase) ValidateJwtUser(userId uint) (domain.User, error) {
	user, err := use.Repo.FindUserById(userId)
	if err != nil {
		return user, errors.New("Unauthorized User")
	}
	return user, nil
}

func (use *userUseCase) AdminLogin(user domain.User) (domain.User, error) {
	if user.Username != "" {
		userDetatils, err := use.Repo.FindByUserName(user)
		if err != nil {
			return userDetatils, errors.New("User not found")
		}

		if userDetatils.Isadmin == false {
			return userDetatils, errors.New("User not found")
		}

		if !utils.VerifyPassword(user.Password, userDetatils.Password) {
			return userDetatils, errors.New("Password in worng or did not match ")
		}
		return userDetatils, nil
	} else if user.Email != "" {
		userDetails, err := use.Repo.FindByUserEmail(user)
		if err != nil {
			return userDetails, errors.New("User not found")
		}
		if err != nil {
			return userDetails, errors.New("User not found")
		}
		if !utils.VerifyPassword(user.Password, userDetails.Password) {
			return userDetails, errors.New("Password in wrong or did not match ")
		}
		return userDetails, nil
	}
	return user, nil
}

func NewUserUseCase(repo interfaces.UserRepo) userCase.UserUseCase {
	return &userUseCase{
		Repo: repo,
	}
}
