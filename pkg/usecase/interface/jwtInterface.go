package interfaces

import (
	"github.com/athunlal/bookNow-auth-svc/pkg/domain"
	"github.com/golang-jwt/jwt"
)

type JwtUseCase interface {
	GenerateAccessToken(userid int, email string, role string) (string, error)
	GenerateRefreshToken(userid int, email string, role string) (string, error)
	VerifyToken(token string) (bool, *domain.JwtClaims)
	GetTokenFromString(signedToken string, claims *domain.JwtClaims) (*jwt.Token, error)
}
