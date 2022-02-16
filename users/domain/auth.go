package domain

import (
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenGenerator interface {
	GenerateToken(string) (string, error)
}

/*
func (uClaim *UserClaim) GenerateToken(email string) (string, error) {
	signingKey := []byte(config.GetJwtSecret())
	uClaim = &UserClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(EXPIRATION) * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaim)
	return token.SignedString(signingKey)
}*/
