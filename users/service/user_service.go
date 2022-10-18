package service

import (
	pb "bootcampProject/additional_information/proto"
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
	userRepository       domain.UserRepository
	tokenGenerator       domain.TokenGenerator
	additionalInfoClient domain.AdditionalInformationClient
}

func NewUserService(rep domain.UserRepository, tokenGen domain.TokenGenerator, addInfoClient domain.AdditionalInformationClient) domain.UserService {
	return &userService{
		userRepository:       rep,
		tokenGenerator:       tokenGen,
		additionalInfoClient: addInfoClient,
	}
}

// CreateUser Create User persistent
func (s *userService) CreateUser(ctx context.Context, user domain.Users) (int, error) {

	user.PwdHash = utils.HashSHA256(user.PwdHash)
	userTemp, err := s.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if err.Error() != utils.ErrRecordNotFound.Error() {
			return 0, err
		}
	}

	if userTemp.ID != 0 {
		return 0, ErrUserAlreadyExists
	}

	userID, err := s.userRepository.CreateUser(ctx, user)
	_, err = s.additionalInfoClient.CreateAdditionalInfo(ctx, &pb.CreateAdditionalInfoReq{
		UserId:         int32(userID),
		AdditionalInfo: user.AdditionalInfo,
	})
	return userID, err
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
		return "", utils.NewErrBadRequest()
	}

	// retrieve user with pwd from database
	user, err := s.GetUserByEmail(ctx, auth.Email)
	if err != nil {
		if err.Error() == utils.ErrRecordNotFound.Error() {
			return "", utils.NewErrInvalidCredentials()
		} else {
			return "", err
		}
	}

	// verify password
	if user.PwdHash != utils.HashSHA256(auth.Password) {
		return "", utils.NewErrInvalidCredentials()
	}
	//create TOKEN
	token, err := s.tokenGenerator.GenerateToken(auth.Email)
	if err != nil {
		return "", errors.New("GenerateToken: " + err.Error())
	}
	return token, nil
}
