package service

import (
	"bootcampProject/config"
	"bootcampProject/users/domain"
	"bootcampProject/utils"
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type tokenGenerator struct {
	jwtSecret string
}

const EXPIRATION = 30

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func NewTokenGenerator() domain.TokenGenerator {
	return &tokenGenerator{
		jwtSecret: config.GetJwtSecret()}
}

func (gen *tokenGenerator) GenerateToken(email string) (string, error) {
	uClaim := &domain.UserClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(EXPIRATION) * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaim)
	return token.SignedString([]byte(gen.jwtSecret))
}

type userService struct {
	userRepository domain.UserRepository
	tokenGenerator domain.TokenGenerator
}

func NewUserService(rep domain.UserRepository, tokenGen domain.TokenGenerator) domain.UserService {
	return &userService{
		userRepository: rep,
		tokenGenerator: tokenGen,
	}
	//return userServiceLogging{logger, userSvc}
}

// CreateUser Create User persistent
func (s *userService) CreateUser(ctx context.Context, user domain.Users) (int, error) {
	user.PwdHash = utils.HashSHA256(user.PwdHash)
	userTemp, err := s.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return 0, err
	}

	if userTemp.ID != 0 {
		return 0, ErrUserAlreadyExists
	}

	return s.userRepository.CreateUser(ctx, user)
}

func (s *userService) GetUsers(ctx context.Context, limit int, offset int) ([]domain.Users, error) {
	return s.userRepository.GetUsers(ctx, limit, offset)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (domain.Users, error) {
	return s.userRepository.GetUserByEmail(ctx, email)
}

func (s *userService) Authenticate(ctx context.Context, auth domain.Auth) (string, error) {
	// Check for any empty fields
	if utils.CheckEmptyField(auth) {
		return "", errors.New("invalid fields")
	}

	// retrieve user with pwd from database
	user, err := s.GetUserByEmail(ctx, auth.Email)
	if err != nil {
		if err.Error() == utils.ErrRecordNotFound.Error() {
			return "", utils.ErrInvalidCredentials
		} else {
			return "", err
		}
	}

	// verify password
	if user.PwdHash != utils.HashSHA256(auth.Password) {
		return "", utils.ErrInvalidCredentials
	}
	//create TOKEN
	token, err := s.tokenGenerator.GenerateToken(auth.Email)
	if err != nil {
		return "", errors.New("GenerateToken: " + err.Error())
	}
	return token, nil
}
