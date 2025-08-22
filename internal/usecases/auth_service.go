package usecases

import (
	"errors"
	"time"

	"backend-service-api/internal/models"
	"backend-service-api/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	users      repositories.UserRepository
	jwtSecret  []byte
	tokenTTL   time.Duration
}

func NewAuthService(users repositories.UserRepository, jwtSecret string, tokenTTL time.Duration) *AuthService {
	return &AuthService{users: users, jwtSecret: []byte(jwtSecret), tokenTTL: tokenTTL}
}

func (s *AuthService) Register(email, password, name string) (models.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{ID: uuid.NewString(), Email: email, PasswordHash: string(passwordHash), Name: name}
	return s.users.Create(user)
}

func (s *AuthService) Login(email, password string) (string, models.User, error) {
	user, ok, err := s.users.GetByEmail(email)
	if err != nil {
		return "", models.User{}, err
	}
	if !ok {
		return "", models.User{}, ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", models.User{}, ErrInvalidCredentials
	}
	claims := jwt.MapClaims{
		"sub": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(s.tokenTTL).Unix(),
		"iat": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(s.jwtSecret)
	if err != nil {
		return "", models.User{}, err
	}
	return token, user, nil
}

func (s *AuthService) GetProfile(userID string) (models.User, error) {
	user, ok, err := s.users.GetByID(userID)
	if err != nil {
		return models.User{}, err
	}
	if !ok {
		return models.User{}, ErrTodoNotFound
	}
	return user, nil
}

func (s *AuthService) UpdateProfile(userID, name string) (models.User, error) {
	user, err := s.GetProfile(userID)
	if err != nil {
		return models.User{}, err
	}
	user.Name = name
	return s.users.Update(user)
}


