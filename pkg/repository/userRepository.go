package repository

import (
	"github.com/athunlal/bookNow-auth-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNow-auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func (r *userDatabase) FindByUserName(user domain.User) (domain.User, error) {
	result := r.DB.First(&user, "username LIKE ?", user.Username).Error
	return user, result
}
func (r *userDatabase) FindByUserEmail(user domain.User) (domain.User, error) {
	result := r.DB.First(&user, "email LIKE ?", user.Email).Error
	return user, result
}
func (r *userDatabase) Create(user domain.User) error {
	result := r.DB.Create(&user).Error
	return result
}
func (r *userDatabase) FindUserByOtp(user domain.User) (domain.User, error) {
	result := r.DB.Where("otp LIKE ?", user.Otp).First(&user)
	return user, result.Error
}
func (r *userDatabase) NullTheOtp(user domain.User) int64 {
	var userData domain.User
	result := r.DB.Model(&userData).Where("id = ?", user.Id).Update("otp", nil)
	return result.RowsAffected
}

func (r *userDatabase) FindUserById(userid uint) (domain.User, error) {
	user := domain.User{}
	result := r.DB.First(&user, "id = ?", userid).Error
	return user, result
}

func (r *userDatabase) IsOtpVerified(username string) string {
	var otp string
	r.DB.Raw("select otp from users where username LIKE ?", username).Scan(&otp)
	return otp
}

func (r *userDatabase) DeleteUser(user domain.User) error {
	result := r.DB.Exec("DELETE FROM users WHERE email LIKE ?", user.Email).Error
	return result
}

func (r *userDatabase) UpdateOtp(user domain.User) error {
	result := r.DB.Model(&user).Where("id = ?", user.Id).Update("otp", user.Otp)
	return result.Error
}

func (r *userDatabase) VerifyUser(user domain.User) (domain.User, error) {
	result := r.DB.Model(&user).Where("id = ?", user.Id).Update("isverified", true)
	return user, result.Error
}

func (r *userDatabase) ChangePassword(user domain.User) error {
	result := r.DB.Model(&user).Where("id = ?", user.Id).Update("password", user.Password)
	return result.Error
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &userDatabase{
		DB: db,
	}
}
